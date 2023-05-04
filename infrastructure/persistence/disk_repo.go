package persistence

import (
	"context"
	"fmt"
	"github.com/arturiamu/lplms-public_cloud/common"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/arturiamu/lplms-public_cloud/utils/pointerutil"
	"github.com/arturiamu/lplms-public_cloud/utils/uuid"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "kubevirt.io/api/core/v1"
	"kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
)

const (
	diskAnnotationServerID   = "disk/server_id"
	diskAnnotationServerName = "disk/server_name"
	diskAnnotationImageID    = "disk/image_id"
	diskAnnotationStatus     = "disk/status"
)

type PublicInfo struct {
	//ZoneID    enums.Zone `json:"zone_id"`
	DiskID    string `json:"disk_id"`
	ProjectID string `json:"project_id"`
	ServerID  string `json:"server_id"`
}

type AfterAttachDiskArgs struct {
	PublicInfo
	ServerName string `json:"server_name"`
	ImageID    string `json:"image_id"`
}

func (s *StorageRepo) CreateDisk(arg *entity.DiskCreateArg) (*entity.DiskCreateResp, error) {
	var (
		ns = arg.ProjectID
		k  = s.k8Virt.CdiClient().CdiV1beta1().DataVolumes(ns)
		id = uuid.GenerateUUID()
		// current only support rw once
		accessMode = []corev1.PersistentVolumeAccessMode{corev1.ReadWriteMany}
	)
	quantity, err := resource.ParseQuantity(fmt.Sprintf("%dG", *arg.Size))
	if err != nil {
		return nil, err
	}
	annotations := map[string]string{
		common.AnnotationName: id,
	}
	source := &v1beta1.DataVolumeSource{
		Blank: &v1beta1.DataVolumeBlankImage{},
	}

	if arg.DiskName != nil {
		annotations[common.AnnotationName] = *arg.DiskName
	}
	if arg.Description != nil {
		annotations[common.AnnotationDescription] = *arg.Description
	}
	annotations[common.AnnotationDiskType] = common.DataDiskType.String()

	if arg.ImageID != nil {
		annotations[common.AnnotationDiskType] = common.SystemDiskType.String()
		annotations[diskAnnotationImageID] = *arg.ImageID
		imageInfo := GetImageInfo(*arg.ImageID)
		source = &v1beta1.DataVolumeSource{
			Registry: &v1beta1.DataVolumeSourceRegistry{
				PullMethod: pointerutil.Pointer(v1beta1.RegistryPullNode),
				URL:        pointerutil.Pointer(imageInfo.Addr),
			},
		}
	}

	createDVArgs := &v1beta1.DataVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name:        id,
			Annotations: annotations,
		},
		Spec: v1beta1.DataVolumeSpec{
			PVC: &corev1.PersistentVolumeClaimSpec{
				AccessModes: accessMode,
				Resources: corev1.ResourceRequirements{
					Requests: corev1.ResourceList{
						corev1.ResourceStorage: quantity,
					},
				},
			},
			Source: source,
		},
	}

	_, err = k.Create(context.Background(), createDVArgs, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return &entity.DiskCreateResp{DiskID: id}, err
}

func (s *StorageRepo) DeleteDisk(arg *entity.DiskDeleteArg) (*entity.DiskDeleteResp, error) {
	var (
		ns = arg.ProjectID
		k  = s.k8Virt.CdiClient().CdiV1beta1().DataVolumes(ns)
	)
	err := k.Delete(context.Background(), arg.DiskID, metav1.DeleteOptions{})
	return nil, err
}

func (s *StorageRepo) UpdateDisk(arg *entity.DiskUpdateArg) (*entity.DiskUpdateResp, error) {
	var (
		ns = arg.ProjectID
		k  = s.k8Virt.CdiClient().CdiV1beta1().DataVolumes(ns)
	)
	oldDisk, err := k.Get(context.Background(), arg.DiskID, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	newDisk := oldDisk.DeepCopy()
	annotation := newDisk.ObjectMeta.Annotations
	modified := false
	if arg.DiskName != nil {
		modified = true
		annotation[common.AnnotationName] = *arg.DiskName
	}

	if arg.Description != nil {
		modified = true
		annotation[common.AnnotationDescription] = *arg.Description
	}

	if modified {
		newDisk.ObjectMeta.Annotations = annotation
		_, err = k.Update(context.Background(), newDisk, metav1.UpdateOptions{})
		if err != nil {
			return nil, err
		}
	}
	return nil, err
}

func (s *StorageRepo) GetDisk(arg *entity.DiskGetArg) (*entity.DiskGetResp, error) {
	var (
		ns = arg.ProjectID
		k  = s.k8Virt.CdiClient().CdiV1beta1().DataVolumes(ns)
	)
	resp, err := k.Get(context.Background(), arg.DiskID, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	quantity := pointerutil.Pointer(resp.Spec.PVC.Resources.Requests[corev1.ResourceStorage])
	size, _ := quantity.AsInt64()
	status := common.DiskStatusType(resp.GetObjectMeta().GetAnnotations()[diskAnnotationStatus])
	if resp.Status.Phase != v1beta1.Succeeded {
		status = kubevirtStatusToStatckStatus(resp.Status.Phase)
	}

	if len(status) == 0 {
		status = common.DiskAvailable
	}

	diskInfo := entity.Disk{
		DiskID:     resp.GetObjectMeta().GetName(),
		DiskType:   pointerutil.Pointer(common.DiskType(resp.GetObjectMeta().GetAnnotations()[common.AnnotationDiskType])),
		ServerID:   resp.GetObjectMeta().GetAnnotations()[diskAnnotationServerID],
		ServerName: resp.GetObjectMeta().GetAnnotations()[diskAnnotationServerName],
		//ZoneID:     arg.ZoneID,
		Size:      size / 1000 / 1000 / 1000,
		DiskName:  resp.GetObjectMeta().GetAnnotations()[common.AnnotationName],
		CreatedAt: resp.GetObjectMeta().GetCreationTimestamp().Unix(),
		Status:    status,
		Category:  common.CloudSSDiskCategory,
	}

	if len(resp.GetObjectMeta().GetAnnotations()[diskAnnotationImageID]) > 0 {
		imageInfo := GetImageInfo(resp.GetObjectMeta().GetAnnotations()[diskAnnotationImageID])
		diskInfo.ImageID = imageInfo.ImageID
		diskInfo.ImageName = imageInfo.ImageName
		diskInfo.OSType = imageInfo.OSType
		diskInfo.OSName = imageInfo.OSName
	}

	return &entity.DiskGetResp{Disk: diskInfo}, nil
}

func (s *StorageRepo) ListDisk(arg *entity.DiskListArg) (*entity.DiskListResp, error) {
	var (
		ns = arg.ProjectID
		k  = s.k8Virt.CdiClient().CdiV1beta1().DataVolumes(ns)
	)
	resp, err := k.List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var res entity.DiskListResp
	diskFilter := make(map[string]struct{})
	for _, v := range arg.DiskIDs {
		diskFilter[v] = struct{}{}
	}

	for _, v := range resp.Items {
		if arg.ServerID != nil && *arg.ServerID != v.GetObjectMeta().GetAnnotations()[diskAnnotationServerID] {
			continue
		}
		if len(diskFilter) != 0 {
			if _, ok := diskFilter[v.GetName()]; !ok {
				continue
			}
		}
		quantity := pointerutil.Pointer(v.Spec.PVC.Resources.Requests[corev1.ResourceStorage])
		size, _ := quantity.AsInt64()
		status := common.DiskStatusType(v.GetObjectMeta().GetAnnotations()[diskAnnotationStatus])
		if v.Status.Phase != v1beta1.Succeeded {
			status = kubevirtStatusToStatckStatus(v.Status.Phase)
		}

		if len(status) == 0 {
			status = common.DiskAvailable
		}

		disk := entity.Disk{
			DiskID:     v.GetObjectMeta().GetName(),
			DiskType:   pointerutil.Pointer(common.DiskType(v.GetObjectMeta().GetAnnotations()[common.AnnotationDiskType])),
			ServerID:   v.GetObjectMeta().GetAnnotations()[diskAnnotationServerID],
			ServerName: v.GetObjectMeta().GetAnnotations()[diskAnnotationServerName],
			//ZoneID:     args.ZoneID,
			Size:      size / 1000 / 1000 / 1000,
			DiskName:  v.GetObjectMeta().GetAnnotations()[common.AnnotationName],
			CreatedAt: v.GetObjectMeta().GetCreationTimestamp().Unix() * 1000,
			Status:    status,
			Category:  common.CloudSSDiskCategory,
		}

		if len(v.GetObjectMeta().GetAnnotations()[diskAnnotationImageID]) > 0 {
			imageInfo := GetImageInfo(v.GetObjectMeta().GetAnnotations()[diskAnnotationImageID])
			disk.ImageID = imageInfo.ImageID
			disk.ImageName = imageInfo.ImageName
			disk.OSType = imageInfo.OSType
			disk.OSName = imageInfo.OSName
		}

		disk.Bootable = *disk.DiskType == common.SystemDiskType
		res.Disks = append(res.Disks, disk)
	}
	return &res, nil
}

func (s *StorageRepo) AttachDisk(args *entity.DiskAttachArg) ([]*entity.DiskAttachResp, error) {
	c := s.k8Virt.VirtualMachine(args.ProjectID)
	err := c.AddVolume(args.ServerID, &v1.AddVolumeOptions{
		Disk: &v1.Disk{
			Name: args.DiskID,
			DiskDevice: v1.DiskDevice{
				Disk: &v1.DiskTarget{
					Bus: v1.DiskBusSCSI,
				},
			},
		},
		VolumeSource: &v1.HotplugVolumeSource{
			DataVolume: &v1.DataVolumeSource{
				Name:         args.DiskID,
				Hotpluggable: true,
			},
		},
		Name: args.DiskID,
	})
	if err != nil {
		return nil, err
	}

	err = s.AfterAttachDisk(AfterAttachDiskArgs{
		PublicInfo: PublicInfo{
			ProjectID: args.ProjectID,
			DiskID:    args.DiskID,
			ServerID:  args.ServerID,
		},
	})

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *StorageRepo) DetachDisk(arg *entity.DiskDetachArg) ([]*entity.DiskDetachResp, error) {
	c := s.k8Virt.VirtualMachine(arg.ProjectID)
	err := c.RemoveVolume(arg.ServerID, &v1.RemoveVolumeOptions{
		Name: arg.DiskID,
	})
	if err != nil {
		return nil, err
	}
	err = s.afterDetachDisk(
		PublicInfo{ProjectID: arg.ProjectID, DiskID: arg.DiskID, ServerID: arg.ServerID},
	)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *StorageRepo) ResizeDisk(arg *entity.DiskResizeArg) ([]*entity.DiskResizeResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *StorageRepo) ResetDisk(arg *entity.DiskResetArg) ([]*entity.DiskResetResp, error) {
	var (
		d = s.k8Virt.CdiClient().CdiV1beta1().DataVolumes(arg.ProjectID)
	)
	oldDisk, err := d.Get(context.Background(), arg.DiskID, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	newDisk := oldDisk.DeepCopy()
	apigroup := "snapshot.storage.k8s.io"
	newDisk.Spec.PVC.DataSource = &corev1.TypedLocalObjectReference{
		Name:     arg.SnapshotID,
		Kind:     "VolumeSnapshot",
		APIGroup: &apigroup,
	}
	_, err = d.Update(context.Background(), newDisk, metav1.UpdateOptions{})
	return nil, err
}

func (s *StorageRepo) AfterAttachDisk(args AfterAttachDiskArgs) (err error) {
	d := s.k8Virt.CdiClient().CdiV1beta1().DataVolumes(args.ProjectID)
	oldDisk, err := d.Get(context.Background(), args.DiskID, metav1.GetOptions{})
	if err != nil {
		return
	}
	newDisk := oldDisk.DeepCopy()
	annotation := newDisk.ObjectMeta.Annotations

	annotation[diskAnnotationServerID] = args.ServerID
	server, err := s.k8Virt.VirtualMachine(args.ProjectID).Get(args.ServerID, &metav1.GetOptions{})
	if err == nil {
		annotation[diskAnnotationServerName] = server.ObjectMeta.Annotations[common.AnnotationName]
	}

	if len(args.ImageID) > 0 {
		annotation[diskAnnotationImageID] = args.ImageID
	}

	annotation[diskAnnotationStatus] = common.DiskInUse.String()
	newDisk.ObjectMeta.Annotations = annotation
	_, err = d.Update(context.Background(), newDisk, metav1.UpdateOptions{})

	return
}

func (s *StorageRepo) afterDetachDisk(args PublicInfo) (err error) {
	d := s.k8Virt.CdiClient().CdiV1beta1().DataVolumes(args.ProjectID)
	oldDisk, err := d.Get(context.Background(), args.DiskID, metav1.GetOptions{})
	if err != nil {
		return
	}
	newDisk := oldDisk.DeepCopy()
	annotation := newDisk.ObjectMeta.Annotations

	annotation[diskAnnotationStatus] = common.DiskAvailable.String()
	delete(annotation, diskAnnotationServerID)
	delete(annotation, diskAnnotationServerName)
	newDisk.ObjectMeta.Annotations = annotation
	_, err = d.Update(context.Background(), newDisk, metav1.UpdateOptions{})

	return
}

///////////////////////////// help functions /////////////

func kubevirtStatusToStatckStatus(s v1beta1.DataVolumePhase) common.DiskStatusType {
	switch s {
	case v1beta1.Pending,
		v1beta1.ImportInProgress,
		v1beta1.CloneInProgress,
		v1beta1.CSICloneInProgress,
		v1beta1.ImportScheduled:
		return common.DiskCreating
	default:
		return common.DiskError
	}
}
