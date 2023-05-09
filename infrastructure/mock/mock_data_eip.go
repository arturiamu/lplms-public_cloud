package mock

import (
	"github.com/arturiamu/lplms-public_cloud/common"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/arturiamu/lplms-public_cloud/utils/pointerutil"
	"github.com/arturiamu/lplms-public_cloud/utils/uuid"
)

var mockEipMap map[string]entity.Eip

var mockEipList []entity.Eip

func initEip() {
	mockEipMap = make(map[string]entity.Eip)
	mockEipList = []entity.Eip{
		{
			EIPID:      uuid.GenerateUUID(),                                         // EIP的ID。
			ServerID:   pointerutil.Pointer("600c674b-9753-47bc-b092-63157c9356e2"), //&mockServerList[0].ServerID,   // 当前绑定的实例的ID。
			ServerName: pointerutil.Pointer("测试机器1"),                                //&mockServerList[0].ServerName, // 当前绑定的实例的名称。

			InstanceType: "", // 当前绑定的实例的类型

			EIPName: "测试eip1", // 弹性公网IP的名称。
			// EIP的状态。
			// - Associating：绑定中。
			// - Unassociating：解绑中。
			// - InUse：已分配。
			// - Available：可用。
			Status:     common.InUseEIPStatus, //enum Associating | Unassociating | InUse | Available
			MNO:        common.MNOTypeMobile,  // 运营商类型
			CreatedAt:  0,                     // EIP的创建时间。
			UpdatedAt:  0,                     // EIP的更新时间。
			EIPAddress: "10.26.124.90",        // 弹性公网IP的IP地址。
			Bandwidth:  10,                    // EIP的带宽峰值，单位为Mbps。
		},
		{
			EIPID:      uuid.GenerateUUID(),                                         // EIP的ID。
			ServerID:   pointerutil.Pointer("a228ab47-ebb5-40a7-9f0c-7ac0a88b3b05"), //&mockServerList[0].ServerID,   // 当前绑定的实例的ID。
			ServerName: pointerutil.Pointer("测试机器2"),                                //&mockServerList[0].ServerName, // 当前绑定的实例的名称。

			InstanceType: "", // 当前绑定的实例的类型

			EIPName: "测试eip2", // 弹性公网IP的名称。
			// EIP的状态。
			// - Associating：绑定中。
			// - Unassociating：解绑中。
			// - InUse：已分配。
			// - Available：可用。
			Status:     common.InUseEIPStatus, //enum Associating | Unassociating | InUse | Available
			MNO:        common.MNOTypeMobile,  // 运营商类型
			CreatedAt:  0,                     // EIP的创建时间。
			UpdatedAt:  0,                     // EIP的更新时间。
			EIPAddress: "124.23.76.88",        // 弹性公网IP的IP地址。
			Bandwidth:  5,                     // EIP的带宽峰值，单位为Mbps。
		},
		{
			EIPID:        uuid.GenerateUUID(), // EIP的ID。
			ServerID:     nil,                 // 当前绑定的实例的ID。
			ServerName:   nil,                 // 当前绑定的实例的名称。
			InstanceType: "",                  // 当前绑定的实例的类型

			EIPName: "测试eip3", // 弹性公网IP的名称。
			// EIP的状态。
			// - Associating：绑定中。
			// - Unassociating：解绑中。
			// - InUse：已分配。
			// - Available：可用。
			Status:     common.AvailableEIPStatus, //enum Associating | Unassociating | InUse | Available
			MNO:        common.MNOTypeTelecom,     // 运营商类型
			CreatedAt:  0,                         // EIP的创建时间。
			UpdatedAt:  0,                         // EIP的更新时间。
			EIPAddress: "45.234.68.100",           // 弹性公网IP的IP地址。
			Bandwidth:  20,                        // EIP的带宽峰值，单位为Mbps。
		},
	}
	for _, eip := range mockEipList {
		mockEipMap[eip.EIPID] = eip
	}
}

func getEip(id string) entity.Eip {
	return mockEipMap[id]
}

func getEipList() []entity.Eip {
	return mockEipList
}
