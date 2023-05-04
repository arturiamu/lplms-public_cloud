package entity

import "github.com/arturiamu/lplms-public_cloud/common"

type Eip struct {
	// EIP所在的地域。
	//ZoneID enums.Zone

	// EIP的ID。
	EIPID string

	// 当前绑定的实例的ID。
	ServerID *string

	// 当前绑定的实例的名称。
	ServerName *string

	// 当前绑定的实例的类型
	InstanceType common.AttachInstanceType

	// 弹性公网IP的名称。
	EIPName string

	// EIP的状态。
	// - Associating：绑定中。
	// - Unassociating：解绑中。
	// - InUse：已分配。
	// - Available：可用。
	Status common.EIPStatus //enum Associating | Unassociating | InUse | Available

	// 运营商类型
	MNO common.MNOType

	// EIP的创建时间。
	CreatedAt int64

	// EIP的更新时间。
	UpdatedAt int64

	// 弹性公网IP的IP地址。
	EIPAddress string

	// EIP的带宽峰值，单位为Mbps。
	Bandwidth uint
}

type EipCreateArg struct {
	EIPName   string // 分配名称
	Bandwidth uint   // EIP的带宽峰值，取值范围：1~200，单位为Mbps。 默认值为5。

	ProjectID *string

	Description string         // 暂时只用于存运营商类型
	MNO         common.MNOType // 运营商类型
}

type EipDeleteArg struct {
}

type EipUpdateArg struct {
}

type EipGetArg struct {
}

type EipListArg struct {
	ProjectID string
	//ZoneID    enums.Zone // EIP所在的地域。您可以通过调用DescribeRegions接口获取地域ID。

	// 要查询的EIP实例的ID。
	//
	// 最多支持输入50个EIP实例ID，实例ID之间用逗号（,）分隔。
	//
	// @GSD:NOTE 说明 如果同时传入EipAddress和AllocationId参数，AllocationId可输入50个EIP实例ID，EipAddress也可同时输入50个EIP的IP地址。
	EIPID *string

	// EIP的状态，取值：
	// - Associating：绑定中。
	// - Unassociating：解绑中。
	// - InUse：已分配。
	// - Available：可用。
	Status *common.EIPStatus //enum Associating | Unassociating | InUse | Available

	// 要查询的EIP的IP地址。
	//
	// 最多支持输入50个EIP的IP地址，IP地址之间用逗号（,）分隔。
	//
	// @GSD:NOTE 说明 如果同时传入EipAddress和AllocationId参数，EipAddress可输入50个EIP的IP地址，AllocationId也可同时输入50个EIP实例ID。
	EIPAddress *string

	// 云产品的实例ID。
	ServerID *string
}

type EipCreateResp struct {
	EIPID      string
	EIPAddress string
}

type EipDeleteResp struct{}

type EipUpdateResp struct{}

type EipGetResp struct{}

type EipListResp struct {
	EIPs []Eip
}
