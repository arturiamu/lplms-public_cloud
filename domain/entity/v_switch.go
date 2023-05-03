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
}

type VSwitchDeleteArg struct {
}

type VSwitchUpdateArg struct {
}

type VSwitchGetArg struct {
}

type VSwitchListArg struct {
	VPCID      *string  // 要查询的交换机所属VPC的ID。
	VSwitchIDs []string // 要查询的交换机的ID。
	ProjectID  string
}

type VSwitchCreateResp struct {
}

type VSwitchDeleteResp struct{}

type VSwitchUpdateResp struct{}

type VSwitchGetResp struct{}

type VSwitchListResp struct {
	VSwitches []VSwitch
}
