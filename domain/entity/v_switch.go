package entity

import "github.com/arturiamu/lplms-public_cloud/common"

type VSwitch struct {
	// 交换机所属的可用区。
	//ZoneID enums.Zone

	// 交换机的ID。
	VSwitchID string

	// 交换机所属VPC的ID。
	VPCID string

	// 交换机的状态，取值：
	// - Pending：配置中。
	// - Available：可用。
	Status common.VSwitchStatus // enum Pending | Available

	// 交换机的IPv4网段。
	CIDR string

	// 交换机中可用的IP地址数量。
	AvailableIPCount int

	// ResourceCount 资源数量
	ResourceCount int

	// 交换机的描述信息。
	Description string

	// 交换机的名称。
	VSwitchName string

	// 交换机的创建时间。
	CreatedAt int64
}

type VSwitchCreateArg struct {
	// 交换机的网段。交换机网段要求如下：
	// - 交换机的网段的掩码长度范围为16~29位。
	// - 交换机的网段必须从属于所在VPC的网段。
	// - 交换机的网段不能与所在VPC中路由条目的目标网段相同，但可以是目标网段的子集。
	CIDR string

	ProjectID string

	// 要创建的交换机所属的VPC ID。
	VPCID string

	// 交换机的名称。 名称长度为2~128个字符，必须以字母或中文开头，但不能以 `http://` 或 `https://` 开头。
	VSwitchName *string

	// 交换机的描述信息。 描述长度为2~256个字符，必须以字母或中文开头，但不能以 `http://` 或 `https://` 开头。
	Description *string
}

type VSwitchDeleteArg struct {
	VPCID     string // 要创建的交换机所属的VPC ID。
	VSwitchID string // 要删除的交换机的ID。
}

type VSwitchUpdateArg struct {
	VPCID       string  // 交换机所属VPC的ID。
	VSwitchID   string  // 交换机的ID。
	VSwitchName *string // 交换机要修改的名称。名称长度为2~128个字符，必须以字母或中文开头，可包含数字、点号（.）、下划线（_）和短横线（-）。但不能以 `http://` 或 `https://` 开头。
	Description *string // 交换机要修改的描述信息。描述长度为2~256个字符，必须以字母或中文开头，但不能以 `http://` 或 `https://` 开头。
}

type VSwitchGetArg struct {
	VSwitchID string
	VPCID     *string
}

type VSwitchListArg struct {
	VPCID      *string  // 要查询的交换机所属VPC的ID。
	VSwitchIDs []string // 要查询的交换机的ID。
	ProjectID  string
}

type VSwitchCreateResp struct {
	VSwitchID string // 创建的交换机的ID。
}

type VSwitchDeleteResp struct{}

type VSwitchUpdateResp struct{}

type VSwitchGetResp struct {
	Resources []VSwitchResourceInfo
}

type VSwitchResourceInfo struct {
	ID   string
	Name string
	Type common.ResourceType
}

type VSwitchListResp struct {
	VSwitches []VSwitch
}
