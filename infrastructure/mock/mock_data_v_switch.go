package mock

import (
	"github.com/arturiamu/lplms-public_cloud/common"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
)

var mockVSwitchMap map[string]entity.VSwitch
var mockVSwitchList []entity.VSwitch

func initVSwitch() {
	mockVSwitchMap = make(map[string]entity.VSwitch)
	mockVSwitchList = []entity.VSwitch{
		{
			// 交换机的ID。
			VSwitchID: "bc756a62-0bfe-4ef0-a3d6-3c2c5c8b7ed6",

			// 交换机所属VPC的ID。
			VPCID: "c037c4ad-f9d8-48cc-828a-39c52721d742", //mockVpcList[0].VPCID,

			// 交换机的状态，取值：
			// - Pending：配置中。
			// - Available：可用。
			Status: common.AvailableVSwitchStatus, // enum Pending | Available

			// 交换机的IPv4网段。
			CIDR: "192.223.201.0/24",

			// 交换机中可用的IP地址数量。
			AvailableIPCount: 250,

			// ResourceCount 资源数量
			ResourceCount: 1,

			// 交换机的描述信息。
			Description: "这是一台测试交换机",

			// 交换机的名称。
			VSwitchName: "测试交换机1",

			// 交换机的创建时间。
			CreatedAt: 1682870400, // 2023-05-01
		},
		{
			// 交换机的ID。
			VSwitchID: "ac734ae7-8b94-495e-a238-0530a1ba4a93",

			// 交换机所属VPC的ID。
			VPCID: "498a8e2f-9b1b-4928-9cf3-d22676f3656d", // mockVpcList[1].VPCID,

			// 交换机的状态，取值：
			// - Pending：配置中。
			// - Available：可用。
			Status: common.AvailableVSwitchStatus, // enum Pending | Available

			// 交换机的IPv4网段。
			CIDR: "192.223.201.0/24",

			// 交换机中可用的IP地址数量。
			AvailableIPCount: 250,

			// ResourceCount 资源数量
			ResourceCount: 1,

			// 交换机的描述信息。
			Description: "这是一台测试交换机",

			// 交换机的名称。
			VSwitchName: "测试交换机2",

			// 交换机的创建时间。
			CreatedAt: 1682870400, // 2023-05-01
		},
		{
			// 交换机的ID。
			VSwitchID: "2207a19b-24a8-40be-8128-d8a3581eb96a",

			// 交换机所属VPC的ID。
			VPCID: "498a8e2f-9b1b-4928-9cf3-d22676f3656d", // mockVpcList[1].VPCID,

			// 交换机的状态，取值：
			// - Pending：配置中。
			// - Available：可用。
			Status: common.PendingVSwitchStatus, // enum Pending | Available

			// 交换机的IPv4网段。
			CIDR: "192.223.202.0/24",

			// 交换机中可用的IP地址数量。
			AvailableIPCount: 250,

			// ResourceCount 资源数量
			ResourceCount: 1,

			// 交换机的描述信息。
			Description: "这是一台测试交换机",

			// 交换机的名称。
			VSwitchName: "测试交换机3",

			// 交换机的创建时间。
			CreatedAt: 1682870400, // 2023-05-01
		},
	}
	for _, vSwitch := range mockVSwitchList {
		mockVSwitchMap[vSwitch.VSwitchID] = vSwitch
	}
}

func getVSwitchList() []entity.VSwitch {
	return mockVSwitchList
}

func getVSwitch(id string) entity.VSwitch {
	return mockVSwitchMap[id]
}
