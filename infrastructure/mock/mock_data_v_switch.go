package mock

import "github.com/arturiamu/lplms-public_cloud/domain/entity"

var mockVSwitchList = []entity.VSwitch{
	{
		// 交换机的ID。
		VSwitchID: "",

		// 交换机所属VPC的ID。
		VPCID: "",

		// 交换机的状态，取值：
		// - Pending：配置中。
		// - Available：可用。
		Status: "", // enum Pending | Available

		// 交换机的IPv4网段。
		CIDR: "",

		// 交换机中可用的IP地址数量。
		AvailableIPCount: "",

		// ResourceCount 资源数量
		ResourceCount: "",

		// 交换机的描述信息。
		Description: "",

		// 交换机的名称。
		VSwitchName: "",

		// 交换机的创建时间。
		CreatedAt: "",
	},
}
