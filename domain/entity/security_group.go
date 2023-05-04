package entity

import "github.com/arturiamu/lplms-public_cloud/common"

type SecurityGroup struct {
	SecurityGroupID    string
	SecurityGroupName  string
	Description        string
	SecurityGroupRules []SecurityGroupRule
}

type SecurityGroupRule struct {
	SecurityGroupRuleID   string
	Direction             common.Direction
	Type                  common.SecurityRuleGrantType // 安全组授权类型，1 表示安全组，2 表示 ip 段
	RemoteSecurityGroupID *string                      // 远端安全组 ID
	CIDR                  *string                      // 远端 IP 段

	// 传输层协议。取值大小写敏感。取值范围：
	// - icmp
	// - gre
	// - tcp
	// - udp
	// - all：支持所有协议
	Protocol common.NetworkProtocol // enum tcp | udp | icmp | gre | all

	// 目的端安全组开放的传输层协议相关的端口范围
	FromPort *int // 最小端口号
	ToPort   *int // 最大端口号

	// 安全组规则的描述信息。长度为1~512个字符。
	Description string

	CreatedAt int64
}

type SecurityGroupCreateArg struct {
	ProjectID string

	// 安全组名称
	SecurityGroupName *string

	// 安全组描述信息。
	Description *string
}

type SecurityGroupDeleteArg struct {
	SecurityGroupID string // 安全组ID
}

type SecurityGroupUpdateArg struct {
	SecurityGroupID   string  // 安全组ID。
	SecurityGroupName *string // 安全组名称。 长度为2~128个英文或中文字符。必须以大小字母或中文开头，不能以http://和https://开头。可以包含数字、半角冒号（:）、下划线（_）或者连字符（-）。 默认值：空
	Description       *string // 安全组描述信息。长度为2~256个英文或中文字符，不能以http://和https://开头。 默认值：空
}

type SecurityGroupGetArg struct {
	SecurityGroupID string // 安全组ID。

	// 安全组规则授权方向。取值范围：
	// - egress：安全组出方向
	// - ingress：安全组入方向
	// - all：不区分方向
	//
	// 默认值：all
	Direction *common.Direction
}

type SecurityGroupListArg struct {
	ProjectID        string
	SecurityGroupIDs []string // 安全组ID列表
}

type SecurityGroupCreateResp struct {
	SecurityGroup *SecurityGroup // 安全组的信息
}

type SecurityGroupDeleteResp struct{}

type SecurityGroupUpdateResp struct{}

type SecurityGroupGetResp struct {
	SecurityGroup SecurityGroup
}

type SecurityGroupListResp struct {
	SecurityGroups []SecurityGroup
}

type SecurityGroupRuleCreateArg struct {
	SecurityGroupID       string
	Direction             common.Direction
	Type                  common.SecurityRuleGrantType // 安全组授权类型，1 表示安全组，2 表示 ip 段
	RemoteSecurityGroupID *string                      // 远端安全组 ID
	CIDR                  *string                      // 远端 IP 段

	// 传输层协议。取值大小写敏感。取值范围：
	// - icmp
	// - gre
	// - tcp
	// - udp
	// - all：支持所有协议
	Protocol common.NetworkProtocol // enum tcp | udp | icmp | gre | all

	// 目的端安全组开放的传输层协议相关的端口范围
	FromPort int // 最小端口号
	ToPort   int // 最大端口号

	// 安全组规则的描述信息。长度为1~512个字符。
	Description *string
}

type SecurityGroupRuleDeleteArg struct {
	SecurityGroupRuleID string
	SecurityGroupID     string
}

type SecurityGroupRuleUpdateArg struct {
}

type SecurityGroupRuleGetArg struct {
}

type SecurityGroupRuleListArg struct {
	SecurityGroupID string
}

type SecurityGroupRuleCreateResp struct {
	SecurityGroupRule *SecurityGroupRule // 安全组的信息
}

type SecurityGroupRuleDeleteResp struct{}

type SecurityGroupRuleUpdateResp struct{}

type SecurityGroupRuleGetResp struct{}

type SecurityGroupRuleListResp struct {
	SecurityGroupRoles []SecurityGroupRule
}
