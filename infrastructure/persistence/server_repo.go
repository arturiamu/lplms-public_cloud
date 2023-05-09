package persistence

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/arturiamu/lplms-public_cloud/common"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/arturiamu/lplms-public_cloud/utils/crypt/sha512_crypt"
	"github.com/arturiamu/lplms-public_cloud/utils/pointerutil"
	"github.com/arturiamu/lplms-public_cloud/utils/uuid"
	k8sv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "kubevirt.io/api/core/v1"
	"log"
	"math/rand"
	"strings"
	"sync"
	"text/template"
	"time"
)

const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZs"

const userDataTpl = `#cloud-config
users:
  - name: root
    lock_passwd: false
    {{if ne .Password ""}}hashed_passwd: "{{ .Password }}"{{end}}
    ssh_pwauth: true
    {{if ne .Keypair ""}}ssh_authorized_keys:
      - '{{ .Keypair }}'{{end}}
runcmd:
  - [ bash, -c, "sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config"]
  - [ bash, -c, "sed -i 's/PasswordAuthentication no/PasswordAuthentication yes/' /etc/ssh/sshd_config"]
  - [ bash, -c, "systemctl restart ssh"]
  - [ bash, -c, "systemctl restart sshd"]
`

type userData struct {
	Password string
	Keypair  string
}

type createSecretArgs struct {
	//ZoneId    enums.Zone
	Namespace string
	Name      string
	Data      string
}

func (c *ComputeRepo) createSecretForServer(args *createSecretArgs) (secret *k8sv1.Secret, err error) {
	secret, err = c.k8Virt.CoreV1().Secrets(args.Namespace).Create(context.Background(), &k8sv1.Secret{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      args.Name,
			Namespace: args.Namespace,
		},
		Data: map[string][]byte{
			"root": []byte(args.Data),
		},
	}, metav1.CreateOptions{})
	return
}

func (c *ComputeRepo) secretDelete(pid string, sid string) (err error) {
	for retry := 0; retry < 40; retry++ {
		server, err := c.GetServer(&entity.ServerGetArg{
			ProjectID: pid,
			ServerID:  sid,
		})
		if err != nil {
			return err
		}
		if server.Server.Status == common.RunningServerStatus {
			err = c.k8Virt.CoreV1().Secrets(pid).Delete(context.Background(), sid, metav1.DeleteOptions{})
			return nil
		}
		if retry >= 40 {
			return errors.NewBadRequest("err run virtual machine")
		}
		time.Sleep(3 * time.Second)
	}
	return nil
}

func mkPassword(pwd string) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := make([]byte, 8)
	for i := range s {
		s[i] = charset[seededRand.Intn(len(charset))]
	}
	salt := []byte(fmt.Sprintf("$6$%s", s))
	// use salt to hash user-supplied password
	c := sha512_crypt.New()
	hash, err := c.Generate([]byte(pwd), salt)
	if err != nil {
		fmt.Printf("error hashing user's supplied password: %s\n", err)
		return ""
	}
	return string(hash)
}

func (c *ComputeRepo) generateUserData(args *entity.ServerCreateArg) *string {
	if args.Password == nil && args.KeyPairName == nil {
		return nil
	}

	u := userData{}
	if args.Password != nil {
		u.Password = mkPassword(*args.Password)
	}

	if args.KeyPairName != nil {
		k, err := c.getPubKey(*args.KeyPairName, args.ProjectID)
		if err != nil {
			return nil
		}
		u.Keypair = k
	}

	t, err := template.New("").Parse(userDataTpl)
	if err != nil {
		return nil
	}

	buf := bytes.NewBuffer(nil)
	t.Execute(buf, u)
	s := base64.URLEncoding.EncodeToString(buf.Bytes())
	return &s
}

// CreateServer
// 调用 CreateServer 创建一台 Server 实例。
func (c *ComputeRepo) CreateServer(args *entity.ServerCreateArg) (*entity.ServerCreateResp, error) {
	var (
		ns = args.ProjectID
		k  = c.k8Virt.VirtualMachine(ns)
		id = uuid.GenerateUUID()
	)
	f, err := c.GetFlavor(&entity.FlavorGetArg{FlavorID: args.FlavorID})
	flavor := f.Flavor
	if err != nil {
		return nil, err
	}

	annotationMap := map[string]string{
		common.AnnotationName:      *args.ServerName,
		common.AnnotationVPCID:     *args.VPCID,
		common.AnnotationVSwitchID: *args.VSwitchID,
		common.AnnotationFlavor:    args.FlavorID,
	}

	if args.Description != nil {
		annotationMap[common.AnnotationDescription] = *args.Description
	}

	disks, volumes, err := c.generateCreateServerDisks(args, id)
	if err != nil {
		return nil, err
	}

	userData := c.generateUserData(args)
	if userData != nil {
		volumes = append(volumes, v1.Volume{
			Name: "cloudinitdisk",
			VolumeSource: v1.VolumeSource{
				CloudInitNoCloud: &v1.CloudInitNoCloudSource{
					UserDataBase64: *userData,
				},
			},
		})
	}
	cpuQuantity, _ := resource.ParseQuantity(fmt.Sprint(flavor.CPU))
	memQuantity, _ := resource.ParseQuantity(fmt.Sprintf("%dMi", flavor.Memory))

	secret, err := c.createSecretForServer(&createSecretArgs{
		Namespace: ns,
		Name:      id,
		Data:      *args.Password,
	})
	if err != nil {
		return nil, err
	}

	createArgs := &v1.VirtualMachine{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:        id,
			Namespace:   ns,
			Labels:      map[string]string{},
			Annotations: annotationMap,
		},
		Spec: v1.VirtualMachineSpec{
			Running: pointerutil.Pointer(true),
			Template: &v1.VirtualMachineInstanceTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						common.AnnotationOvnLogicalRouter: *args.VPCID,
						common.AnnotationOvnLogicalSwitch: *args.VSwitchID,
						common.AnnotationOvnSecurityGroup: args.SecurityGroupID,
					},
				},
				Spec: v1.VirtualMachineInstanceSpec{
					// Hostname: hostname,
					DNSConfig: &k8sv1.PodDNSConfig{
						Nameservers: []string{"119.29.29.29", "223.5.5.5", "223.6.6.6"},
					},
					Domain: v1.DomainSpec{
						Resources: v1.ResourceRequirements{
							Limits: k8sv1.ResourceList{
								"cpu":    cpuQuantity,
								"memory": memQuantity,
							},
						},
						Devices: v1.Devices{
							Disks: disks,
							Interfaces: []v1.Interface{
								{Name: "default", InterfaceBindingMethod: v1.InterfaceBindingMethod{
									Bridge: &v1.InterfaceBridge{},
								}},
							},
						},
					},
					AccessCredentials: []v1.AccessCredential{
						{
							UserPassword: &v1.UserPasswordAccessCredential{
								Source: v1.UserPasswordAccessCredentialSource{
									Secret: &v1.AccessCredentialSecretSource{
										SecretName: secret.Name,
									},
								},
								PropagationMethod: v1.UserPasswordAccessCredentialPropagationMethod{
									QemuGuestAgent: &v1.QemuGuestAgentUserPasswordAccessCredentialPropagation{},
								},
							},
						},
					},

					Networks: []v1.Network{*v1.DefaultPodNetwork().DeepCopy()},
					Volumes:  volumes,
				},
			},
		},
	}

	if args.HostName != nil {
		createArgs.Spec.Template.Spec.Hostname = *args.HostName
	}

	if flavor.GPUSpec != "" {
		var gpus []v1.GPU
		cnt := int(flavor.GPUAmount)
		driver := GetDriver(flavor.GPUSpec)
		if driver != "" {
			for i := 0; i < cnt; i++ {
				gpus = append(gpus, v1.GPU{
					Name:       fmt.Sprintf("GPU-%s_%d", driver, i+1),
					DeviceName: driver,
				})
			}
		}
		createArgs.Spec.Template.Spec.Domain.Devices.GPUs = gpus
	}

	_, err = k.Create(createArgs)

	go func() {
		if err = c.secretDelete(ns, id); err != nil {
			log.Println("err delete secret,", err)
		}
	}()
	return &entity.ServerCreateResp{ServerID: id}, nil
}

func (c *ComputeRepo) StartServer(args *entity.ServerStartArgs) (*entity.ServerStartResp, error) {
	var (
		ns = args.ProjectID
		k  = c.k8Virt.VirtualMachine(ns)
	)
	err := k.Start(args.ServerID, &v1.StartOptions{})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *ComputeRepo) StopServer(args *entity.ServerStopArgs) (*entity.ServerStopResp, error) {
	var (
		ns  = args.ProjectID
		k   = c.k8Virt.VirtualMachine(ns)
		err error
	)

	// 由于关机后 vmi 会被释放，从而导致拿不到 private_ip_address,
	// 所以在关机前先把 private_ip 写到 annotation 里面
	c.setPrivateIPAddrAndMacToAnnotation(args.ProjectID, args.ServerID)

	if args.ForceStop != nil && *args.ForceStop {
		err = k.ForceStop(args.ServerID, &v1.StopOptions{})
	} else {
		err = k.Stop(args.ServerID, &v1.StopOptions{})
	}
	return nil, err
}

func (c *ComputeRepo) RestartServer(args *entity.ServerRestartArgs) (*entity.ServerRestartResp, error) {
	var (
		ns  = args.ProjectID
		k   = c.k8Virt.VirtualMachine(ns)
		err error
	)

	if args.ForceRestart != nil && *args.ForceRestart {
		err = k.ForceRestart(args.ServerID, &v1.RestartOptions{})
	} else {
		err = k.Restart(args.ServerID, &v1.RestartOptions{})
	}
	return nil, err
}

func (c *ComputeRepo) DeleteServer(args *entity.ServerDeleteArg) (*entity.ServerDeleteResp, error) {
	var (
		ns = args.ProjectID
		k  = c.k8Virt.VirtualMachine(ns)
	)
	//serverResp, err := c.GetServer(&entity.ServerGetArg{
	//	ProjectID: args.ProjectID,
	//	ServerID:  args.ServerID,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//for _, vpc := range serverResp.Server.VPCAttributes {
	//	if vpc.EipAddress == nil {
	//		continue
	//	}
	//	deleteEIPArgs := entity.EipDeleteArg{
	//		EIPID:     vpc.EipAddress.AllocationID,
	//		ProjectID: args.ProjectID,
	//	}
	//	_, err := stk.N.DeleteEip(&deleteEIPArgs)
	//	if err != nil {
	//		log.Println("DeleteEIP failed ", err)
	//		err = nil
	//	}
	//}
	err := k.Delete(args.ServerID, &metav1.DeleteOptions{})
	if err != nil {
		return nil, err
	}

	// 删除系统盘
	disks, err := c.describeServerDisks(args.ProjectID, args.ServerID)
	if err != nil {
		log.Println("describeServerDisks failed ", err)
		err = nil // 删除失败不处理
	}

	for _, disk := range disks {
		if disk.DiskType != nil && *disk.DiskType == common.SystemDiskType {
			stk.S.DeleteDisk(&entity.DiskDeleteArg{
				ProjectID: args.ProjectID,
				DiskID:    disk.DiskID,
			})
		}
	}
	return nil, nil
}

// UpdateServer
// 调用 UpdateServer 调整一台 Server 实例的实例规格
func (c *ComputeRepo) UpdateServer(args *entity.ServerUpdateArg) (*entity.ServerUpdateResp, error) {
	var (
		ns = args.ProjectID
		k  = c.k8Virt.VirtualMachine(ns)
	)
	f, err := c.GetFlavor(&entity.FlavorGetArg{FlavorID: args.FlavorID})
	if err != nil {
		return nil, err
	}
	flavor := f.Flavor
	oldServer, err := k.Get(args.ServerID, &metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	cpuQuantity, _ := resource.ParseQuantity(fmt.Sprint(flavor.CPU))
	memQuantity, _ := resource.ParseQuantity(fmt.Sprintf("%dMi", flavor.Memory))
	newServer := oldServer.DeepCopy()
	newServer.ObjectMeta.Annotations[common.AnnotationFlavor] = args.FlavorID
	newServer.Spec.Template.Spec.Domain.Resources.Limits = k8sv1.ResourceList{
		k8sv1.ResourceCPU:    cpuQuantity,
		k8sv1.ResourceMemory: memQuantity,
	}
	_, err = k.Update(newServer)
	if err != nil {
		return nil, err
	}

	err = k.ForceRestart(args.ServerID, &v1.RestartOptions{})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// GetServer
// 调用 GetServer 获取一台 Server 实例的状态信息。
func (c *ComputeRepo) GetServer(args *entity.ServerGetArg) (*entity.ServerGetResp, error) {
	var (
		ns = args.ProjectID
		k  = c.k8Virt.VirtualMachine(ns)
	)
	resp, err := k.Get(args.ServerID, &metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	server, err := c.kubevirtServerToStackServer(*resp)
	if err != nil {
		return nil, err
	}
	return &entity.ServerGetResp{Server: *server}, nil
}

func (c *ComputeRepo) GetServerDisks(args *entity.ServerDisksGetArg) (*entity.ServerDisksGetResp, error) {
	disks, err := c.describeServerDisks(args.ProjectID, args.ServerID)
	if err != nil {
		return nil, err
	}

	res := &entity.ServerDisksGetResp{
		Disks: make([]*entity.ServerDiskInfo, 0, len(disks)),
	}

	for _, v := range disks {
		res.Disks = append(res.Disks, &entity.ServerDiskInfo{
			ID:           v.DiskID,
			Size:         int(v.Size),
			IsBoot:       v.Bootable,
			DiskCategory: v.Category.String(),
			DiskType:     v.DiskType,
		})
	}
	return res, nil
}

// ListServer
// 调用 DescribeServers 查询一台或多台 Server 实例的详细信息。
func (c *ComputeRepo) ListServer(args *entity.ServerListArg) (*entity.ServerListResp, error) {
	var (
		ns = args.ProjectID
		k  = c.k8Virt.VirtualMachine(ns)
	)

	resp, err := k.List(&metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	res := &entity.ServerListResp{
		Servers: make([]entity.Server, 0, len(resp.Items)),
	}
	for _, v := range resp.Items {
		info, er := c.kubevirtServerToStackServer(v)
		if er != nil {
			err = er
			return nil, err
		}

		res.Servers = append(res.Servers, *info)
	}

	return res, nil
}

func (c *ComputeRepo) generateCreateServerDisks(args *entity.ServerCreateArg, serverID string) (disks []v1.Disk, volumes []v1.Volume, err error) {
	var (
		systemDisk string
		dataDisks  []string
		wg         = sync.WaitGroup{}
		sysErr     error
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		systemDisk, sysErr = c.createDisk(&entity.DiskCreateArg{
			//ZoneID:      args.ZoneID,
			ProjectID:   args.ProjectID,
			Size:        pointerutil.Pointer(int64(args.SystemDisk.Size)),
			DiskName:    args.SystemDisk.DiskName,
			Description: args.SystemDisk.Description,
			ImageID:     &args.ImageID,
		})
		return
	}()

	for _, v := range args.DataDisks {
		wg.Add(1)
		go func(disk *entity.DiskArgs) {
			defer wg.Done()
			diskID, err := c.createDisk(&entity.DiskCreateArg{
				//ZoneID:      args.ZoneID,
				ProjectID:   args.ProjectID,
				Size:        pointerutil.Pointer(int64(disk.Size)),
				DiskName:    disk.DiskName,
				Description: disk.Description,
			})
			if err != nil {
				return
			}
			dataDisks = append(dataDisks, diskID)
		}(v)
	}

	wg.Wait()

	if sysErr != nil {
		err = fmt.Errorf("create system disk failed %w", sysErr)
		return
	}

	d := stk.S

	d.AfterAttachDisk(AfterAttachDiskArgs{
		PublicInfo: PublicInfo{
			//ZoneID:    args.ZoneID,
			DiskID:    systemDisk,
			ProjectID: args.ProjectID,
			ServerID:  serverID,
		},
		ServerName: *args.ServerName,
		ImageID:    args.ImageID,
	})

	disks = append(disks, v1.Disk{
		Name: systemDisk,
		DiskDevice: v1.DiskDevice{
			Disk: &v1.DiskTarget{
				Bus: v1.DiskBusSCSI,
			},
		},
	})

	volumes = append(volumes, v1.Volume{
		Name: systemDisk,
		VolumeSource: v1.VolumeSource{
			DataVolume: &v1.DataVolumeSource{
				Name: systemDisk,
			},
		},
	})

	for _, v := range dataDisks {
		disks = append(disks, v1.Disk{
			Name: v,
			DiskDevice: v1.DiskDevice{
				Disk: &v1.DiskTarget{
					Bus: v1.DiskBusSCSI,
				},
			},
		})

		volumes = append(volumes, v1.Volume{
			Name: v,
			VolumeSource: v1.VolumeSource{
				DataVolume: &v1.DataVolumeSource{
					Name: v,
				},
			},
		})

		d.AfterAttachDisk(AfterAttachDiskArgs{
			PublicInfo: PublicInfo{
				//ZoneID:    args.ZoneID,
				DiskID:    v,
				ProjectID: args.ProjectID,
				ServerID:  serverID,
			},
			ServerName: *args.ServerName,
		})
	}

	return
}

func (c *ComputeRepo) createDisk(args *entity.DiskCreateArg) (diskID string, err error) {
	d := stk.S

	diskInfo, err := d.CreateDisk(args)
	if err != nil {
		return
	}
	diskID = diskInfo.DiskID

	for {
		info, err := d.GetDisk(
			&entity.DiskGetArg{ProjectID: args.ProjectID, DiskID: diskID},
		)
		if err != nil {
			break
		}

		if info.Disk.Status == common.DiskAvailable {
			break
		}

		if info.Disk.Status == common.DiskError {
			err = fmt.Errorf("disk create failed %s", diskID)
			break
		}
		time.Sleep(time.Millisecond * 500)
	}

	return
}

func (c *ComputeRepo) setPrivateIPAddrAndMacToAnnotation(projectID, serverID string) (err error) {
	k := c.k8Virt.VirtualMachine(projectID)
	oldServer, err := k.Get(serverID, &metav1.GetOptions{})
	if err != nil {
		return
	}

	// newServer := oldServer.DeepCopy()
	annotation := oldServer.ObjectMeta.Annotations
	if ip, ok := annotation[common.AnnotationPrivateIPAddress]; ok && len(ip) > 0 {
		return
	}

	privateIP, mac := c.getPrivateIPAddrAndMac(projectID, serverID)
	if len(privateIP) > 0 {
		annotation[common.AnnotationPrivateIPAddress] = privateIP
		annotation[common.AnnotationMacAddress] = mac
		newServer := oldServer.DeepCopy()
		newServer.ObjectMeta.Annotations = annotation
		_, err = k.Update(newServer)
	}

	return
}

func (c *ComputeRepo) getPrivateIPAddrAndMac(projectID, serverID string) (privateIPAddr, mac string) {
	vmi, err := c.k8Virt.VirtualMachineInstance(projectID).Get(serverID, &metav1.GetOptions{})
	if err == nil && vmi != nil && len(vmi.Status.Interfaces) > 0 {
		privateIPAddr = vmi.Status.Interfaces[0].IP
		mac = vmi.Status.Interfaces[0].MAC
	}
	return
}

///////////////////////////// help functions /////////////

func (c *ComputeRepo) kubevirtServerToStackServer(kubevirtInfo v1.VirtualMachine) (info *entity.Server, err error) {
	var (
		meta       = kubevirtInfo.ObjectMeta
		annotation = kubevirtInfo.ObjectMeta.Annotations
		systemDisk entity.Disk
	)

	networkInfo, err := c.getNetworkInfo(getNetworkfInfoArgs{
		VPCID:     annotation[common.AnnotationVPCID],
		VswitchID: annotation[common.AnnotationVSwitchID],
		ServerID:  meta.Name,
		ProjectID: meta.Namespace,
	})
	if err != nil {
		return
	}

	if len(networkInfo.PrivateIPAddress) == 0 {
		networkInfo.PrivateIPAddress = kubevirtInfo.ObjectMeta.Annotations[common.AnnotationPrivateIPAddress]
		networkInfo.Mac = kubevirtInfo.ObjectMeta.Annotations[common.AnnotationMacAddress]
	}

	disks, err := c.describeServerDisks(meta.Namespace, meta.Name)
	if err != nil {
		return
	}

	diskInfos := make([]*entity.ServerDiskInfo, 0, len(disks))
	for _, v := range disks {
		if *v.DiskType == common.SystemDiskType {
			systemDisk = v
		}
		diskInfos = append(diskInfos, &entity.ServerDiskInfo{
			ID:           v.DiskID,
			Size:         int(v.Size),
			IsBoot:       v.Bootable,
			DiskCategory: v.Category.String(),
			DiskType:     v.DiskType,
		})
	}

	f, err := c.GetFlavor(&entity.FlavorGetArg{FlavorID: annotation[common.AnnotationFlavor]})
	flavor := f.Flavor

	securityGroupInfos, err := c.describeSecurityGroup(
		kubevirtInfo.Spec.Template.ObjectMeta.Annotations[common.AnnotationOvnSecurityGroup], meta.Namespace)
	if err != nil {
		return
	}

	info = &entity.Server{
		ServerID:           meta.Name,
		ServerName:         annotation[common.AnnotationName],
		Description:        annotation[common.AnnotationDescription],
		CreatedAt:          meta.CreationTimestamp.Unix() * 1000,
		HostName:           kubevirtInfo.Spec.Template.Spec.Hostname,
		Status:             kubevirtStatusToStack(kubevirtInfo.Status.PrintableStatus),
		VPCAttributes:      []*entity.VPCAttribute{networkInfo},
		SecurityGroupInfos: securityGroupInfos,
		Disks:              diskInfos,
		ImageID:            systemDisk.ImageID,
		ImageName:          systemDisk.ImageName,
		OSType:             systemDisk.OSType,
		OSName:             systemDisk.OSName,
		CPU:                int64(flavor.CPU),
		Memory:             int64(flavor.Memory),
		FlavorID:           flavor.FlavorID,
		ServerType:         flavor.Name,
	}

	return
}

func kubevirtStatusToStack(kubeStatus v1.VirtualMachinePrintableStatus) common.ServerStatus {
	switch kubeStatus {
	case v1.VirtualMachineStatusProvisioning:
		return common.CreatingServerStatus
	case v1.VirtualMachineStatusStopped:
		return common.StoppedServerStatus
	case v1.VirtualMachineStatusStopping:
		return common.StoppingServerStatus
	case v1.VirtualMachineStatusRunning:
		return common.RunningServerStatus
	case v1.VirtualMachineStatusStarting:
		return common.StartingServerStatus
	default:
		return common.UnknownServerStatus
	}
}

func (c *ComputeRepo) describeSecurityGroup(securityGroupIDs, projectID string) (sgInfos []*entity.SecurityGroupInfo, err error) {
	sgIDs := strings.Split(securityGroupIDs, ",")
	sgResp, err := c.ListSecurityGroup(&entity.SecurityGroupListArg{
		ProjectID:        projectID,
		SecurityGroupIDs: sgIDs,
	})

	for _, v := range sgResp.SecurityGroups {
		sgInfos = append(sgInfos, &entity.SecurityGroupInfo{
			ID:   v.SecurityGroupID,
			Name: v.SecurityGroupName,
		})
	}

	return
}

func (c *ComputeRepo) describeServerDisks(projectID, serverID string) (disks []entity.Disk, err error) {
	sto := stk.S
	resp, err := sto.ListDisk(&entity.DiskListArg{
		ProjectID: projectID,
		ServerID:  &serverID,
	})
	if err != nil {
		return
	}

	for _, v := range resp.Disks {
		if v.ServerID == serverID {
			disks = append(disks, v)
		}
	}

	return
}

type getNetworkfInfoArgs struct {
	VPCID     string
	VswitchID string
	ProjectID string
	ServerID  string
}

func (c *ComputeRepo) getNetworkInfo(args getNetworkfInfoArgs) (networkInfo *entity.VPCAttribute, err error) {
	net := stk.N

	networkInfo = &entity.VPCAttribute{}
	vswitch, err := net.ListVSwitch(&entity.VSwitchListArg{
		VPCID:      &args.VPCID,
		VSwitchIDs: []string{args.VswitchID},
	})
	if err == nil && len(vswitch.VSwitches) > 0 {
		networkInfo.VSwitchId = vswitch.VSwitches[0].VSwitchID
		networkInfo.VSwitchName = vswitch.VSwitches[0].VSwitchName
	}

	vpc, err := net.ListVpc(&entity.VpcListArg{
		VPCIDs:    []string{args.VPCID},
		ProjectID: args.ProjectID,
	})
	if err == nil && len(vpc.VPCs) > 0 {
		networkInfo.VPCId = args.VPCID
		networkInfo.VPCName = vpc.VPCs[0].VPCName
	}

	eip, err := net.ListEip(
		&entity.EipListArg{
			ProjectID: args.ProjectID,
			ServerID:  &args.ServerID,
		},
	)
	if err == nil && len(eip.EIPs) > 0 {
		networkInfo.EipAddress = &entity.EipAddress{
			Name:         eip.EIPs[0].EIPName,
			AllocationID: eip.EIPs[0].EIPID,
			IPAddress:    eip.EIPs[0].EIPAddress,
			Bandwidth:    int64(eip.EIPs[0].Bandwidth),
		}
	}

	vmi, err := c.k8Virt.VirtualMachineInstance(args.ProjectID).Get(args.ServerID, &metav1.GetOptions{})
	if err != nil && errors.IsNotFound(err) {
		err = nil
		return
	}

	if err == nil && vmi != nil && len(vmi.Status.Interfaces) > 0 {
		networkInfo.PrivateIPAddress = vmi.Status.Interfaces[0].IP
		networkInfo.Mac = vmi.Status.Interfaces[0].MAC
	}

	return
}
