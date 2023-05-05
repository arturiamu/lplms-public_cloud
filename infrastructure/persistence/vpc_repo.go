package persistence

import (
	"context"
	"fmt"
	"github.com/arturiamu/lplms-public_cloud/common"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/arturiamu/lplms-public_cloud/utils/uuid"
	ovnv1 "github.com/kubeovn/kube-ovn/pkg/apis/kubeovn/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	netutils "k8s.io/utils/net"
	"net"
	"strings"
	"time"
)

const (
	reconcileTimeout    = 5 * time.Second
	annotationBandwidth = "net.virt.lplms.com/bandwidth" // Mbps
	gatewaySubnetPrefix = "qvirt-gateway-default-"
)

func (n *NetworkRepo) CreateVpc(args *entity.VpcCreateArg) (*entity.VpcCreateResp, error) {
	var (
		vpc    = n.ovn.KubeovnV1().Vpcs()
		subnet = n.ovn.KubeovnV1().Subnets()
		id     = uuid.GenerateUUID()
	)

	desc := ""
	if args.Description != nil {
		desc = *args.Description
	}

	created, err := vpc.Create(context.Background(), &ovnv1.Vpc{
		ObjectMeta: metav1.ObjectMeta{
			Name: id,
			Annotations: map[string]string{
				common.AnnotationName:        *args.VPCName,
				common.AnnotationDescription: desc,
			},
			Labels: map[string]string{
				common.LabelProject: args.ProjectID,
			},
		},
		Spec: ovnv1.VpcSpec{
			Namespaces: []string{
				args.ProjectID,
			},
			StaticRoutes: []*ovnv1.StaticRoute{},
		},
	}, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	_, err = subnet.Create(context.Background(), &ovnv1.Subnet{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName:    gatewaySubnetPrefix,
			OwnerReferences: []metav1.OwnerReference{*metav1.NewControllerRef(created, ovnv1.SchemeGroupVersion.WithKind("Vpc"))},
		},
		Spec: ovnv1.SubnetSpec{
			CIDRBlock: "100.111.0.0/16",
			Vpc:       created.Name,
		},
	}, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	return &entity.VpcCreateResp{
		VPCID: id,
	}, nil
}

func (n *NetworkRepo) DeleteVpc(args *entity.VpcDeleteArg) (*entity.VpcDeleteResp, error) {
	var (
		k = n.ovn.KubeovnV1().Vpcs()
	)
	err := k.Delete(context.Background(), args.VPCID, metav1.DeleteOptions{})
	return nil, err
}

func (n *NetworkRepo) UpdateVpc(args *entity.VpcUpdateArg) (*entity.VpcUpdateResp, error) {
	var (
		k = n.ovn.KubeovnV1().Vpcs()
	)
	vpc, err := k.Get(context.Background(), args.VPCID, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	vpc = vpc.DeepCopy()
	if args.VPCName != nil {
		vpc.Annotations[common.AnnotationName] = *args.VPCName
	}
	if args.Description != nil {
		vpc.Annotations[common.AnnotationDescription] = *args.Description
	}

	_, err = k.Update(context.Background(), vpc, metav1.UpdateOptions{})
	return nil, err
}

func (n *NetworkRepo) GetVpc(args *entity.VpcGetArg) (*entity.VpcGetResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) ListVpc(args *entity.VpcListArg) (*entity.VpcListResp, error) {
	var (
		k    = n.ovn.KubeovnV1().Vpcs()
		vpcs []ovnv1.Vpc
		err  error
	)

	if len(args.VPCIDs) == 0 {
		projectSelector, _ := metav1.LabelSelectorAsSelector(
			&metav1.LabelSelector{MatchLabels: map[string]string{common.LabelProject: args.ProjectID}},
		)
		res, err := k.List(context.Background(), metav1.ListOptions{LabelSelector: projectSelector.String()})
		if err != nil {
			return nil, err
		}
		vpcs = append(vpcs, res.Items...)
	} else {
		for _, id := range args.VPCIDs {
			vpc, err := k.Get(context.Background(), id, metav1.GetOptions{})
			if err != nil {
				if !k8serrors.IsNotFound(err) {
					return nil, err
				} else {
					continue
				}
			}
			vpcs = append(vpcs, *vpc)
		}
	}

	r := entity.VpcListResp{VPCs: []entity.Vpc{}}
	for _, v := range vpcs {
		subnets := make([]string, 0)
		for i := range v.Status.Subnets {
			if !strings.HasPrefix(v.Status.Subnets[i], gatewaySubnetPrefix) {
				subnets = append(subnets, v.Status.Subnets[i])
			}
		}
		r.VPCs = append(r.VPCs, entity.Vpc{
			VPCID:       v.Name,
			Status:      common.AvailableVPCStatus,
			VPCName:     common.GetBizName(&v.ObjectMeta),
			VSwitchIDs:  subnets,
			CIDR:        "0.0.0.0/0", // CIDR 在 Subnet, 默认返回 0.0.0.0/0
			Description: common.GetDesc(&v.ObjectMeta),
			CreatedAt:   v.CreationTimestamp.UTC().UnixMilli(),
		})
	}
	return &r, err
}

func (n *NetworkRepo) CreateVSwitch(args *entity.VSwitchCreateArg) (*entity.VSwitchCreateResp, error) {
	var (
		k  = n.ovn.KubeovnV1().Subnets()
		id = uuid.GenerateUUID()
	)
	desc := ""
	if args.Description != nil {
		desc = *args.Description
	}
	ip, _, er := netutils.ParseCIDRSloppy(args.CIDR)
	if er != nil {
		return nil, er
	}
	if !ip.IsPrivate() {
		return nil, fmt.Errorf("invalid CIDR, please choose CIDR in 192.168.0.0/16, 172.16.0.0/12, 10.0.0.0/8")
	}
	vSwitch, err := DetectSubnetOverlay(n, args)
	if err != nil {
		return nil, err
	}
	if vSwitch != nil {
		return nil, fmt.Errorf("cidr %s overlay with %s", args.CIDR, vSwitch.CIDR)
	}

	_, err = k.Create(context.Background(), &ovnv1.Subnet{
		ObjectMeta: metav1.ObjectMeta{
			Name: id,
			Labels: map[string]string{
				common.LabelProject: args.ProjectID,
			},
			Annotations: map[string]string{
				common.AnnotationName:        *args.VSwitchName,
				common.AnnotationDescription: desc,
			},
		},
		Spec: ovnv1.SubnetSpec{
			Vpc:        args.VPCID,
			Namespaces: []string{args.ProjectID},
			CIDRBlock:  args.CIDR,
		},
		Status: ovnv1.SubnetStatus{},
	}, metav1.CreateOptions{})
	return &entity.VSwitchCreateResp{
		VSwitchID: id,
	}, err
}

func (n *NetworkRepo) DeleteVSwitch(args *entity.VSwitchDeleteArg) (*entity.VSwitchDeleteResp, error) {
	var (
		k = n.ovn.KubeovnV1().Subnets()
	)
	err := k.Delete(context.Background(), args.VSwitchID, metav1.DeleteOptions{})
	return nil, err
}

func (n *NetworkRepo) UpdateVSwitch(args *entity.VSwitchUpdateArg) (*entity.VSwitchUpdateResp, error) {
	var (
		k = n.ovn.KubeovnV1().Subnets()
	)
	subnet, err := k.Get(context.Background(), args.VSwitchID, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	subnet = subnet.DeepCopy()
	if args.VSwitchName != nil {
		subnet.Annotations[common.AnnotationName] = *args.VSwitchName
	}
	if args.Description != nil {
		subnet.Annotations[common.AnnotationDescription] = *args.Description
	}

	_, err = k.Update(context.Background(), subnet, metav1.UpdateOptions{})
	return nil, err
}

func (n *NetworkRepo) GetVSwitch(args *entity.VSwitchGetArg) (*entity.VSwitchGetResp, error) {
	var (
		k = n.ovn.KubeovnV1().Subnets()
	)
	subnet, err := k.Get(context.Background(), args.VSwitchID, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return &entity.VSwitchGetResp{
		Resources: []entity.VSwitchResourceInfo{
			{
				ID:   subnet.Name,
				Name: common.GetBizName(&subnet.ObjectMeta),
				Type: common.NetworkResourceType,
			},
		},
	}, err
}

func (n *NetworkRepo) ListVSwitch(args *entity.VSwitchListArg) (*entity.VSwitchListResp, error) {
	var (
		k        = n.ovn.KubeovnV1().Subnets()
		subnets  []*ovnv1.Subnet
		subnets2 []*ovnv1.Subnet
	)

	if len(args.VSwitchIDs) > 0 {
		for _, id := range args.VSwitchIDs {
			subnet, err := k.Get(context.Background(), id, metav1.GetOptions{})
			if err != nil && k8serrors.IsNotFound(err) {
				return nil, err
			}
			subnets = append(subnets, subnet)
		}
	} else {
		projectSelector, _ := metav1.LabelSelectorAsSelector(
			&metav1.LabelSelector{MatchLabels: map[string]string{common.LabelProject: args.ProjectID}},
		)

		subnets, err := k.List(context.Background(), metav1.ListOptions{LabelSelector: projectSelector.String()})
		if err != nil {
			return nil, err
		}

		if args.VPCID != nil {
			for _, v := range subnets.Items {
				if v.Spec.Vpc == *args.VPCID {
					subnets2 = append(subnets2, &v)
				}
			}
		}
	}

	res := &entity.VSwitchListResp{
		VSwitches: []entity.VSwitch{},
	}

	for i := range subnets2 {
		res.VSwitches = append(res.VSwitches, entity.VSwitch{
			VSwitchID:        subnets[i].Name,
			VPCID:            subnets[i].Spec.Vpc,
			Status:           common.AvailableVSwitchStatus,
			CIDR:             subnets[i].Spec.CIDRBlock,
			AvailableIPCount: int(subnets[i].Status.V4AvailableIPs),
			ResourceCount:    0,
			Description:      common.GetDesc(&subnets[i].ObjectMeta),
			VSwitchName:      common.GetBizName(&subnets[i].ObjectMeta),
			CreatedAt:        subnets[i].CreationTimestamp.UnixMilli(),
		})
	}

	return res, nil
}

///////////////////////////// help functions /////////////

func DetectSubnetOverlay(n *NetworkRepo, args *entity.VSwitchCreateArg) (vSwitch *entity.VSwitch, err error) {
	var vpcIds = make(map[string]struct{})
	var swts []*entity.VSwitch
	vpCsRes, err := n.ListVpc(&entity.VpcListArg{
		ProjectID: args.ProjectID,
	})
	if err != nil {
		return
	}
	for _, vpc := range vpCsRes.VPCs {
		vpcIds[vpc.VPCID] = struct{}{}
	}

	vSwitches, err := n.ListVSwitch(&entity.VSwitchListArg{
		ProjectID: args.ProjectID,
	})
	if err != nil {
		return
	}

	for _, st := range vSwitches.VSwitches {
		if _, ok := vpcIds[st.VPCID]; ok {
			swts = append(swts, &st)
		}
	}

	for _, sw := range swts {
		overlap, err := CIDROverlap(sw.CIDR, args.CIDR)
		if err != nil {
			return nil, err
		}
		if overlap {
			return sw, nil
		}
	}
	return nil, nil
}

func CIDROverlap(c1, c2 string) (bool, error) {
	_, net1, err := net.ParseCIDR(c1)
	if err != nil {
		return false, err
	}

	_, net2, err := net.ParseCIDR(c2)
	if err != nil {
		return false, err
	}

	return net1.Contains(net2.IP) || net2.Contains(net1.IP), nil
}
