package mock

import (
	"github.com/arturiamu/lplms-public_cloud/common"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/arturiamu/lplms-public_cloud/utils/pointerutil"
)

var mockServerMap map[string]entity.Server

var mockServerLIst = []entity.Server{
	{
		ServerID:                "a5fb8fb5-109d-4177-b776-bf49cb77e1f2",
		ServerName:              "测试机器1",
		Description:             "这是一台测试机器",
		CPU:                     1,
		Memory:                  2 * 1024,
		GPUSpec:                 "",
		GPUAmount:               0,
		OSType:                  "linux",
		OSName:                  "CentOS",
		ImageID:                 "1aba5759-c11c-4990-9278-ec70879819f0",
		ImageName:               "CentOS 7.6",
		VlanID:                  "",
		InternetMaxBandwidthIn:  0,
		InternetMaxBandwidthOut: 0,
		VPCAttributes: []*entity.VPCAttribute{
			{
				PrivateIPAddress: "192.168.100.1",
				Mac:              "49681a1b-bba9-4d2b-bf0c-71d6183a10ec",
				VSwitchId:        "2fb6ee51-36c5-42a1-9a6a-db14075079cb",
				VPCId:            "f0ea2036-a8e8-49f0-9786-f504b5151d53",
				VPCName:          "测试子网1",
				VSwitchName:      "测试交换机1",
				Type:             "",
				EipAddress: &entity.EipAddress{
					Name:         "模拟eip1",
					AllocationID: "7c3e45cb-ae52-4a0e-aaa4-24d69c9de6d6",
					IPAddress:    "10.23.116.45",
					Bandwidth:    1,
				},
			},
		},
		DeviceAvailable: true,
		SecurityGroupInfos: []*entity.SecurityGroupInfo{
			{
				ID:   "4e440ed4-2c07-44bc-9765-4359e649fe89",
				Name: "测试安全组1",
			},
		},
		HostName:   "",
		FlavorID:   "",
		ServerType: "",
		Status:     "",
		Disks: []*entity.ServerDiskInfo{
			{
				ID:                  "477e174e-b323-41e9-b5fc-5b77553d921f",
				Size:                20,
				IsBoot:              true,
				DeleteOnTermination: true,
				DiskCategory:        "",
				DiskType:            pointerutil.Pointer(common.SystemDiskType),
			},
		},
		Recyclable:  true,
		KeyPairName: "",
		CreatedAt:   1682870400, // 2013-05-01
	},
	{
		ServerID:                "0ca28efc-03f5-4eb8-b107-46b0307e3328",
		ServerName:              "测试机器2",
		Description:             "这是一台测试机器",
		CPU:                     100,
		Memory:                  200 * 1024,
		GPUSpec:                 "NVIDIA-3090-24G",
		GPUAmount:               8,
		OSType:                  "linux",
		OSName:                  "CentOS",
		ImageID:                 "1aba5759-c11c-4990-9278-ec70879819f0",
		ImageName:               "CentOS 7.6",
		VlanID:                  "",
		InternetMaxBandwidthIn:  0,
		InternetMaxBandwidthOut: 0,
		VPCAttributes: []*entity.VPCAttribute{
			{
				PrivateIPAddress: "192.168.100.4",
				Mac:              "49681a1b-bba9-4d2b-bf0c-71d6183a10ec",
				VSwitchId:        "2fb6ee51-36c5-42a1-9a6a-db14075079cb",
				VPCId:            "f0ea2036-a8e8-49f0-9786-f504b5151d53",
				VPCName:          "测试子网1",
				VSwitchName:      "测试交换机1",
				Type:             "",
				EipAddress: &entity.EipAddress{
					Name:         "模拟eip2",
					AllocationID: "109672ca-de29-48eb-9746-40a453d9fe5a",
					IPAddress:    "10.23.21.12",
					Bandwidth:    10,
				},
			},
		},
		DeviceAvailable: true,
		SecurityGroupInfos: []*entity.SecurityGroupInfo{
			{
				ID:   "4e440ed4-2c07-44bc-9765-4359e649fe89",
				Name: "测试安全组1",
			},
		},
		HostName:   "",
		FlavorID:   "",
		ServerType: "",
		Status:     "",
		Disks: []*entity.ServerDiskInfo{
			{
				ID:                  "5e88609b-0f5b-43db-9cdb-76c947473447",
				Size:                250,
				IsBoot:              true,
				DeleteOnTermination: true,
				DiskCategory:        "",
				DiskType:            pointerutil.Pointer(common.SystemDiskType),
			},
			{
				ID:                  "0ee5be68-93ea-4e49-bf0e-e81bb025b48a",
				Size:                500,
				IsBoot:              false,
				DeleteOnTermination: true,
				DiskCategory:        "",
				DiskType:            pointerutil.Pointer(common.DataDiskType),
			},
		},
		Recyclable:  true,
		KeyPairName: "",
		CreatedAt:   1683388800, // 2013-05-07
	},
}

func init() {
	mockServerMap = make(map[string]entity.Server)
	for _, server := range mockServerLIst {
		mockServerMap[server.ServerID] = server
	}
}

func getServerList() []entity.Server {
	return mockServerLIst
}

func getServer(id string) entity.Server {
	return mockServerMap[id]
}

func getServerDisk(id string) []*entity.ServerDiskInfo {
	return mockServerMap[id].Disks
}
