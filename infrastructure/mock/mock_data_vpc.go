package mock

import (
	"github.com/arturiamu/lplms-public_cloud/common"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
)

var mockVpcMap map[string]entity.Vpc
var mockVpcList = []entity.Vpc{
	{
		// VPC的ID。
		VPCID: "c037c4ad-f9d8-48cc-828a-39c52721d742",

		// VPC的状态，取值：
		// - Pending：配置中。
		// - Available：可用。
		Status: common.AvailableVPCStatus, // enum Pending | Available

		// VPC的名称。
		VPCName: "测试vpc1",

		// VPC中交换机的列表。
		VSwitchIDs: []string{
			"bc756a62-0bfe-4ef0-a3d6-3c2c5c8b7ed6", //mockVSwitchList[0].VSwitchID,
			"ac734ae7-8b94-495e-a238-0530a1ba4a93", //mockVSwitchList[1].VSwitchID,
		},

		// VPC的IPv4网段。
		CIDR: "192.223.200.0/16",

		// VPC的描述信息。
		Description: "这是一个测试vpc",

		// VPC的创建时间。
		CreatedAt: 1682870400, // 2023-05-01
	},
	{
		// VPC的ID。
		VPCID: "498a8e2f-9b1b-4928-9cf3-d22676f3656d",

		// VPC的状态，取值：
		// - Pending：配置中。
		// - Available：可用。
		Status: common.PendingVPCStatus, // enum Pending | Available

		// VPC的名称。
		VPCName: "测试vpc2",

		// VPC中交换机的列表。
		VSwitchIDs: []string{
			"2207a19b-24a8-40be-8128-d8a3581eb96a", //mockVSwitchList[2].VSwitchID,
		},

		// VPC的IPv4网段。
		CIDR: "192.224.148.0/16",

		// VPC的描述信息。
		Description: "这是一个测试vpc",

		// VPC的创建时间。
		CreatedAt: 1682870400, // 2023-05-01
	},
}

func initVpc() {
	mockVpcMap = make(map[string]entity.Vpc)
	for _, vpc := range mockVpcList {
		mockVpcMap[vpc.VPCID] = vpc
	}
}

func getVpc(id string) entity.Vpc {
	return mockVpcMap[id]
}

func getVpcList() []entity.Vpc {
	return mockVpcList
}
