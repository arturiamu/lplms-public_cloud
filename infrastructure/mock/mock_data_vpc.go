package mock

import "github.com/arturiamu/lplms-public_cloud/domain/entity"

var mockVpcList = []entity.Vpc{
	{
		// VPC的ID。
		VPCID: "",

		// VPC的状态，取值：
		// - Pending：配置中。
		// - Available：可用。
		Status: "", // enum Pending | Available

		// VPC的名称。
		VPCName: "",

		// VPC中交换机的列表。
		VSwitchIDs: "",

		// VPC的IPv4网段。
		CIDR: "",

		// VPC的描述信息。
		Description: "",

		// VPC的创建时间。
		CreatedAt: "",
	},
}
