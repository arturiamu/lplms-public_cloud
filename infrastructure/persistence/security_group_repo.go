package persistence

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/arturiamu/lplms-public_cloud/common"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/arturiamu/lplms-public_cloud/utils/uuid"
	ovnv1 "github.com/kubeovn/kube-ovn/pkg/apis/kubeovn/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sort"
	"strings"
	"time"
)

const ruleIdPrefix = "ruleId_"

func (c *ComputeRepo) CreateSecurityGroup(args *entity.SecurityGroupCreateArg) (*entity.SecurityGroupCreateResp, error) {
	var (
		k    = stk.N.ovn.KubeovnV1().SecurityGroups()
		id   = uuid.GenerateUUID()
		desc = ""
		obj  *ovnv1.SecurityGroup
	)

	if args.Description != nil {
		desc = *args.Description
	}

	obj, err := k.Create(context.Background(), &ovnv1.SecurityGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name: id,
			Labels: map[string]string{
				common.LabelProject: args.ProjectID,
			},
			Annotations: map[string]string{
				common.AnnotationName:        *args.SecurityGroupName,
				common.AnnotationDescription: desc,
			},
		},
		Spec: ovnv1.SecurityGroupSpec{
			IngressRules:          []*ovnv1.SgRule{},
			EgressRules:           []*ovnv1.SgRule{},
			AllowSameGroupTraffic: false,
		},
	}, metav1.CreateOptions{})

	if err != nil {
		return nil, err
	}

	return &entity.SecurityGroupCreateResp{
		SecurityGroup: &entity.SecurityGroup{
			SecurityGroupID:    id,
			SecurityGroupName:  common.GetBizName(&obj.ObjectMeta),
			Description:        common.GetDesc(&obj.ObjectMeta),
			SecurityGroupRules: []entity.SecurityGroupRule{},
		},
	}, nil
}

type IndexSecurityGroupRule struct {
	Idx int `json:"index"`
	entity.SecurityGroupRule
}

func (c *ComputeRepo) DeleteSecurityGroup(args *entity.SecurityGroupDeleteArg) (*entity.SecurityGroupDeleteResp, error) {
	var (
		k = stk.N.ovn.KubeovnV1().SecurityGroups()
	)
	err := k.Delete(context.Background(), args.SecurityGroupID, metav1.DeleteOptions{})
	return nil, err
}

func (c *ComputeRepo) UpdateSecurityGroup(args *entity.SecurityGroupUpdateArg) (*entity.SecurityGroupUpdateResp, error) {
	var (
		k = stk.N.ovn.KubeovnV1().SecurityGroups()
	)
	group, err := k.Get(context.Background(), args.SecurityGroupID, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	toUpdate := group.DeepCopy()
	if args.SecurityGroupName != nil {
		toUpdate.Annotations[common.AnnotationName] = *args.SecurityGroupName
	}
	if args.Description != nil {
		toUpdate.Annotations[common.AnnotationDescription] = *args.Description
	}

	if common.GetBizName(&toUpdate.ObjectMeta) != common.GetBizName(&group.ObjectMeta) ||
		common.GetDesc(&toUpdate.ObjectMeta) != common.GetDesc(&group.ObjectMeta) {
		_, err = k.Update(context.Background(), toUpdate, metav1.UpdateOptions{})
		if err != nil {
			return nil, err
		}
	}
	return nil, err
}

func (c *ComputeRepo) GetSecurityGroup(args *entity.SecurityGroupGetArg) (*entity.SecurityGroupGetResp, error) {
	var (
		k = stk.N.ovn.KubeovnV1().SecurityGroups()
	)
	group, err := k.Get(context.Background(), args.SecurityGroupID, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return &entity.SecurityGroupGetResp{
		SecurityGroup: *makeStkSecurityGroup(group),
	}, nil
}

func (c *ComputeRepo) ListSecurityGroup(args *entity.SecurityGroupListArg) (*entity.SecurityGroupListResp, error) {
	var (
		k         = stk.N.ovn.KubeovnV1().SecurityGroups()
		wants     = map[string]bool{}
		stkGroups []entity.SecurityGroup
	)

	if len(args.SecurityGroupIDs) > 0 {
		for i := range args.SecurityGroupIDs {
			wants[args.SecurityGroupIDs[i]] = true
		}
	}

	projectSelector, _ := metav1.LabelSelectorAsSelector(&metav1.LabelSelector{MatchLabels: map[string]string{common.LabelProject: args.ProjectID}})
	groups, err := k.List(context.Background(), metav1.ListOptions{
		LabelSelector: projectSelector.String(),
	})
	if err != nil {
		return nil, err
	}

	for _, group := range groups.Items {
		stkGroup := makeStkSecurityGroup(&group)
		if len(wants) > 0 && !wants[stkGroup.SecurityGroupID] {
			continue
		}
		stkGroups = append(stkGroups, *makeStkSecurityGroup(&group))
	}

	return &entity.SecurityGroupListResp{
		SecurityGroups: stkGroups,
	}, nil
}

func (c *ComputeRepo) CreateSecurityGroupRule(args *entity.SecurityGroupRuleCreateArg) (*entity.SecurityGroupRuleCreateResp, error) {
	var (
		k     = stk.N.ovn.KubeovnV1().SecurityGroups()
		sgobj *ovnv1.SecurityGroup
	)

	if args.Protocol == common.GRENetworkProtocol {
		return nil, fmt.Errorf("protocol/%s is not supported", args.Protocol)
	}
	if args.Direction == common.AllDirection {
		return nil, fmt.Errorf("direction/%s is not supported", args.Direction)
	}

	sgobj, err := k.Get(context.Background(), args.SecurityGroupID, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	sgobj = sgobj.DeepCopy()

	var (
		ingressIdx = len(sgobj.Spec.IngressRules)
		egressIdx  = len(sgobj.Spec.EgressRules)
	)

	rule := entity.SecurityGroupRule{
		SecurityGroupRuleID:   uuid.GenerateUUID(),
		Direction:             args.Direction,
		Type:                  args.Type,
		RemoteSecurityGroupID: args.RemoteSecurityGroupID,
		CIDR:                  args.CIDR,
		Protocol:              args.Protocol,
		FromPort:              &args.FromPort,
		ToPort:                &args.ToPort,
		CreatedAt:             time.Now().UnixMilli(),
	}

	if args.Description != nil {
		rule.Description = *args.Description
	}

	ovnRule := makeOVNRule(&rule)

	switch rule.Direction {
	case common.IngressDirection:
		sgobj.Spec.IngressRules = append(sgobj.Spec.IngressRules, ovnRule)
		b, _ := json.Marshal(IndexSecurityGroupRule{Idx: ingressIdx, SecurityGroupRule: rule})
		sgobj.Annotations[ruleIdPrefix+rule.SecurityGroupRuleID] = string(b)
	case common.EgressDirection:
		sgobj.Spec.EgressRules = append(sgobj.Spec.EgressRules, ovnRule)
		b, _ := json.Marshal(IndexSecurityGroupRule{Idx: egressIdx, SecurityGroupRule: rule})
		sgobj.Annotations[ruleIdPrefix+rule.SecurityGroupRuleID] = string(b)
	}

	_, err = k.Update(context.Background(), sgobj, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}

	return &entity.SecurityGroupRuleCreateResp{
		SecurityGroupRule: &rule,
	}, nil
}

func (c *ComputeRepo) DeleteSecurityGroupRule(args *entity.SecurityGroupRuleDeleteArg) (*entity.SecurityGroupRuleDeleteResp, error) {
	var (
		k = stk.N.ovn.KubeovnV1().SecurityGroups()
	)
	group, err := k.Get(context.Background(), args.SecurityGroupID, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	group = group.DeepCopy()
	var rule IndexSecurityGroupRule
	rawRule := group.Annotations[ruleIdPrefix+args.SecurityGroupRuleID]
	if rawRule == "" {
		return nil, fmt.Errorf("SecurityGroupRule/%s not found", args.SecurityGroupRuleID)
	}
	err = json.Unmarshal([]byte(rawRule), &rule)
	if err != nil {
		return nil, fmt.Errorf("unexpected SecurityGroupRule %s", err.Error())
	}
	delete(group.Annotations, ruleIdPrefix+args.SecurityGroupRuleID)
	switch rule.Direction {
	case common.IngressDirection:
		group.Spec.IngressRules = append(group.Spec.IngressRules[0:rule.Idx], group.Spec.IngressRules[rule.Idx+1:]...)
	case common.EgressDirection:
		group.Spec.EgressRules = append(group.Spec.EgressRules[0:rule.Idx], group.Spec.EgressRules[rule.Idx+1:]...)
	}

	// dynamic update role`s index
	var r IndexSecurityGroupRule
	for k, v := range group.Annotations {
		if strings.HasPrefix(k, ruleIdPrefix) {
			err = json.Unmarshal([]byte(v), &r)
			if err != nil {
				continue
			}
			if r.Idx > rule.Idx {
				r.Idx--
				b, _ := json.Marshal(r)
				group.Annotations[k] = string(b)
			}
		}
	}
	_, err = k.Update(context.Background(), group, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *ComputeRepo) GetSecurityGroupRule(args *entity.SecurityGroupRuleGetArg) (*entity.SecurityGroupRuleGetResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) ListSecurityGroupRule(args *entity.SecurityGroupRuleListArg) (*entity.SecurityGroupRuleListResp, error) {
	var (
		k = stk.N.ovn.KubeovnV1().SecurityGroups()
	)
	group, err := k.Get(context.Background(), args.SecurityGroupID, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	stkGroup := makeStkSecurityGroup(group)
	return &entity.SecurityGroupRuleListResp{SecurityGroupRoles: stkGroup.SecurityGroupRules}, nil
}

///////////////////////////// help functions /////////////

func makeStkSecurityGroup(group *ovnv1.SecurityGroup) *entity.SecurityGroup {
	var inRules []entity.SecurityGroupRule
	var outRules []entity.SecurityGroupRule

	for k, v := range group.Annotations {
		if !strings.HasPrefix(k, ruleIdPrefix) {
			continue
		}
		var rule IndexSecurityGroupRule
		err := json.Unmarshal([]byte(v), &rule)
		if err != nil {
			continue
		}
		if rule.Direction == common.IngressDirection {
			inRules = append(inRules, rule.SecurityGroupRule)
		} else if rule.Direction == common.EgressDirection {
			outRules = append(outRules, rule.SecurityGroupRule)
		}
	}

	sort.Slice(inRules, func(i, j int) bool { return inRules[i].CreatedAt < inRules[j].CreatedAt })
	sort.Slice(outRules, func(i, j int) bool { return outRules[i].CreatedAt < outRules[j].CreatedAt })

	return &entity.SecurityGroup{
		SecurityGroupID:    group.Name,
		SecurityGroupName:  common.GetBizName(&group.ObjectMeta),
		Description:        common.GetDesc(&group.ObjectMeta),
		SecurityGroupRules: append(inRules, outRules...),
	}
}

func makeOVNRule(ori *entity.SecurityGroupRule) *ovnv1.SgRule {
	dst := ovnv1.SgRule{
		IPVersion:           "ipv4",
		Protocol:            ovnv1.SgProtocol(ori.Protocol),
		Priority:            1,
		RemoteAddress:       "",
		RemoteSecurityGroup: "",
		PortRangeMin:        0,
		PortRangeMax:        0,
		Policy:              "",
	}
	switch ori.Type {
	case common.SecurityGroupType:
		dst.RemoteType = ovnv1.SgRemoteTypeSg
		dst.RemoteSecurityGroup = *ori.RemoteSecurityGroupID
	case common.CIDRType:
		dst.RemoteType = ovnv1.SgRemoteTypeAddress
		dst.RemoteAddress = *ori.CIDR
		dst.PortRangeMin = *ori.FromPort
		dst.PortRangeMax = *ori.ToPort
	default:
	}

	dst.Policy = ovnv1.PolicyAllow

	return &dst
}
