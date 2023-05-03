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
}

type SecurityGroupDeleteArg struct {
}

type SecurityGroupUpdateArg struct {
}

type SecurityGroupGetArg struct {
}

type SecurityGroupListArg struct {
	ProjectID        string
	SecurityGroupIDs []string // 安全组ID列表
}

type SecurityGroupCreateResp struct{}

type SecurityGroupDeleteResp struct{}

type SecurityGroupUpdateResp struct{}

type SecurityGroupGetResp struct{}

type SecurityGroupListResp struct {
	SecurityGroups []SecurityGroup
}
