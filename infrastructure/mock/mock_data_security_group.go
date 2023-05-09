package mock

import (
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/arturiamu/lplms-public_cloud/utils/pointerutil"
	"github.com/arturiamu/lplms-public_cloud/utils/uuid"
)

var mockSecurityGroupRuleMap map[string]entity.SecurityGroupRule
var mockSecurityGroupMap map[string]entity.SecurityGroup
var mockSecurityGroupList = []entity.SecurityGroup{
	{
		SecurityGroupID:   uuid.GenerateUUID(),
		SecurityGroupName: "测试安全组1",
		Description:       "这是一个测试安全组",
		SecurityGroupRules: []entity.SecurityGroupRule{
			{
				SecurityGroupRuleID:   uuid.GenerateUUID(),
				Direction:             "ingress",
				Type:                  1,                         // 安全组授权类型，1 表示安全组，2 表示 ip 段
				RemoteSecurityGroupID: nil,                       // 远端安全组 ID
				CIDR:                  nil,                       // 远端 IP 段
				Protocol:              "tcp",                     // enum tcp | udp | icmp | gre | all
				FromPort:              pointerutil.Pointer(3306), // 最小端口号
				ToPort:                pointerutil.Pointer(3306), // 最大端口号
				Description:           "mysql",
				CreatedAt:             1683043199,
			},
			{
				SecurityGroupRuleID:   uuid.GenerateUUID(),
				Direction:             "egress",
				Type:                  1,                         // 安全组授权类型，1 表示安全组，2 表示 ip 段
				RemoteSecurityGroupID: nil,                       // 远端安全组 ID
				CIDR:                  nil,                       // 远端 IP 段
				Protocol:              "tcp",                     // enum tcp | udp | icmp | gre | all
				FromPort:              pointerutil.Pointer(3306), // 最小端口号
				ToPort:                pointerutil.Pointer(3306), // 最大端口号
				Description:           "mysql",
				CreatedAt:             1683043199,
			},
			{
				SecurityGroupRuleID:   uuid.GenerateUUID(),
				Direction:             "all",
				Type:                  1,                         // 安全组授权类型，1 表示安全组，2 表示 ip 段
				RemoteSecurityGroupID: nil,                       // 远端安全组 ID
				CIDR:                  nil,                       // 远端 IP 段
				Protocol:              "all",                     // enum tcp | udp | icmp | gre | all
				FromPort:              pointerutil.Pointer(8888), // 最小端口号
				ToPort:                pointerutil.Pointer(9999), // 最大端口号
				Description:           "自定义端口范围",
				CreatedAt:             1683043199,
			},
		},
	},
	{
		SecurityGroupID:   uuid.GenerateUUID(),
		SecurityGroupName: "测试安全组2",
		Description:       "这是一个测试安全组",
		SecurityGroupRules: []entity.SecurityGroupRule{
			{
				SecurityGroupRuleID:   uuid.GenerateUUID(),
				Direction:             "ingress",
				Type:                  1,                         // 安全组授权类型，1 表示安全组，2 表示 ip 段
				RemoteSecurityGroupID: nil,                       // 远端安全组 ID
				CIDR:                  nil,                       // 远端 IP 段
				Protocol:              "tcp",                     // enum tcp | udp | icmp | gre | all
				FromPort:              pointerutil.Pointer(3306), // 最小端口号
				ToPort:                pointerutil.Pointer(3306), // 最大端口号
				Description:           "mysql",
				CreatedAt:             1683043199,
			},
			{
				SecurityGroupRuleID:   uuid.GenerateUUID(),
				Direction:             "egress",
				Type:                  1,                         // 安全组授权类型，1 表示安全组，2 表示 ip 段
				RemoteSecurityGroupID: nil,                       // 远端安全组 ID
				CIDR:                  nil,                       // 远端 IP 段
				Protocol:              "tcp",                     // enum tcp | udp | icmp | gre | all
				FromPort:              pointerutil.Pointer(3306), // 最小端口号
				ToPort:                pointerutil.Pointer(3306), // 最大端口号
				Description:           "mysql",
				CreatedAt:             1683043199,
			},
			{
				SecurityGroupRuleID:   uuid.GenerateUUID(),
				Direction:             "egress",
				Type:                  1,                         // 安全组授权类型，1 表示安全组，2 表示 ip 段
				RemoteSecurityGroupID: nil,                       // 远端安全组 ID
				CIDR:                  nil,                       // 远端 IP 段
				Protocol:              "tcp",                     // enum tcp | udp | icmp | gre | all
				FromPort:              pointerutil.Pointer(6379), // 最小端口号
				ToPort:                pointerutil.Pointer(6379), // 最大端口号
				Description:           "redis",
				CreatedAt:             1683043199,
			},
			{
				SecurityGroupRuleID:   uuid.GenerateUUID(),
				Direction:             "all",
				Type:                  1,                         // 安全组授权类型，1 表示安全组，2 表示 ip 段
				RemoteSecurityGroupID: nil,                       // 远端安全组 ID
				CIDR:                  nil,                       // 远端 IP 段
				Protocol:              "all",                     // enum tcp | udp | icmp | gre | all
				FromPort:              pointerutil.Pointer(8000), // 最小端口号
				ToPort:                pointerutil.Pointer(9000), // 最大端口号
				Description:           "自定义端口范围",
				CreatedAt:             1683043199,
			},
		},
	},
}

func initSecurityGroup() {
	mockSecurityGroupRuleMap = make(map[string]entity.SecurityGroupRule)
	mockSecurityGroupMap = make(map[string]entity.SecurityGroup)

	for _, group := range mockSecurityGroupList {
		mockSecurityGroupMap[group.SecurityGroupID] = group
		for _, rule := range group.SecurityGroupRules {
			mockSecurityGroupRuleMap[rule.SecurityGroupRuleID] = rule
		}
	}
}

func getSecurityGroup(id string) entity.SecurityGroup {
	return mockSecurityGroupMap[id]
}

func getSecurityGroupList() []entity.SecurityGroup {
	return mockSecurityGroupList
}

func getSecurityGroupRules(id string) []entity.SecurityGroupRule {
	return mockSecurityGroupMap[id].SecurityGroupRules
}
