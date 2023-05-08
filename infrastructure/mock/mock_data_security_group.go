package mock

import "github.com/arturiamu/lplms-public_cloud/domain/entity"

var mockSecurityGroupList = []entity.SecurityGroup{
	{
		SecurityGroupID:    "",
		SecurityGroupName:  "",
		Description:        "",
		SecurityGroupRules: "",
	},
}

var mockSecurityGroupRuleList = []entity.SecurityGroupRule{
	{
		SecurityGroupRuleID:   "",
		Direction:             "",
		Type:                  "", // 安全组授权类型，1 表示安全组，2 表示 ip 段
		RemoteSecurityGroupID: "", // 远端安全组 ID
		CIDR:                  "", // 远端 IP 段

		// 传输层协议。取值大小写敏感。取值范围：
		// - icmp
		// - gre
		// - tcp
		// - udp
		// - all：支持所有协议
		Protocol: "", // enum tcp | udp | icmp | gre | all

		// 目的端安全组开放的传输层协议相关的端口范围
		FromPort: "", // 最小端口号
		ToPort:   "", // 最大端口号

		// 安全组规则的描述信息。长度为1~512个字符。
		Description: "",

		CreatedAt: "",
	},
}
