package entity

import "github.com/arturiamu/lplms-public_cloud/common"

type Vpc struct {

	// VPC的ID。
	VPCID string

	// VPC的状态，取值：
	// - Pending：配置中。
	// - Available：可用。
	Status common.VPCStatus // enum Pending | Available

	// VPC的名称。
	VPCName string

	// VPC中交换机的列表。
	VSwitchIDs []string

	// VPC的IPv4网段。
	CIDR string

	// VPC的描述信息。
	Description string

	// VPC的创建时间。
	CreatedAt int64
}

type EipAddress struct {
	Name         string `json:"name"`
	AllocationID string `json:"allocation_id"`
	IPAddress    string `json:"ip_address"`
	Bandwidth    int64  `json:"bandwidth"`
}

type VpcCreateArg struct {
	ProjectID string

	// VPC的网段。您可以使用以下网段或其子集作为VPC的网段：
	// - 172.16.0.0/12（默认值）。
	// - 10.0.0.0/8。
	// - 192.168.0.0/16。
	CIDR *string

	// VPC的名称。
	//
	// 长度为2~128个字符，必须以字母或中文开头，可包含数字、点号（.）、下划线（_）和短横线（-），但不能以 `http://` 或 `https://` 开头。
	VPCName *string

	// VPC的描述信息。
	//
	// 长度为2~256个字符，必须以字母或中文开头，但不能以 `http://` 或 `https://` 开头。
	Description *string
}

type VpcDeleteArg struct {
	ProjectID string
	VPCID     string // 要删除的VPC ID。
}

type VpcUpdateArg struct {
	// 要修改信息的VPC的ID。
	VPCID string

	// VPC的名称。名称长度为2~128个字符，必须以字母或中文开头，可包含数字、点号（.）、下划线（_）和短横线（-），但不能以http://或https://开头。
	VPCName *string

	// VPC的描述信息。描述长度为2~256个字符，必须以字母或中文开头，但不能以http://或https://开头。
	Description *string
}

type VpcGetArg struct {
	ProjectID string
}

type VpcListArg struct {
	VPCIDs    []string
	ProjectID string
}

type VpcCreateResp struct {
	VPCID string // 创建的VPC的ID。
}

type VpcDeleteResp struct{}

type VpcUpdateResp struct{}

type VpcGetResp struct{}

type VpcListResp struct {
	VPCs []Vpc
}
