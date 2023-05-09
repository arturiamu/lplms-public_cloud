package mock

import (
	"github.com/arturiamu/lplms-public_cloud/common"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/arturiamu/lplms-public_cloud/utils/uuid"
)

var mockServerMap map[string]entity.Server

var mockServerList []entity.Server

func initServer() {
	mockServerMap = make(map[string]entity.Server)
	mockServerList = []entity.Server{
		{
			ServerID:                "600c674b-9753-47bc-b092-63157c9356e2",
			FlavorID:                mockFlavorList[0].FlavorID,
			ServerName:              "测试机器1",
			Description:             "这是一台测试机器",
			CPU:                     int64(mockFlavorList[0].CPU),
			Memory:                  int64(mockFlavorList[0].Memory),
			GPUSpec:                 mockFlavorList[0].GPUSpec,
			GPUAmount:               int64(mockFlavorList[0].GPUAmount),
			OSType:                  mockImageList[0].OSType,
			OSName:                  mockImageList[0].OSName,
			ImageID:                 mockImageList[0].ImageID,
			ImageName:               mockImageList[0].ImageName,
			VlanID:                  "",
			InternetMaxBandwidthIn:  0,
			InternetMaxBandwidthOut: 0,
			VPCAttributes: []*entity.VPCAttribute{
				{
					PrivateIPAddress: "192.223.147.1",
					Mac:              uuid.GenerateUUID(),
					VSwitchId:        mockVpcList[0].VSwitchIDs[0],
					VPCId:            mockVpcList[0].VPCID,
					VPCName:          mockVpcList[0].VPCName,
					VSwitchName:      getVSwitch(mockVpcList[0].VSwitchIDs[0]).VSwitchName,
					Type:             "",
					EipAddress: &entity.EipAddress{
						Name:         mockEipList[0].EIPName,
						AllocationID: mockEipList[0].EIPID,
						IPAddress:    mockEipList[0].EIPAddress,
						Bandwidth:    int64(mockEipList[0].Bandwidth),
					},
				},
			},
			DeviceAvailable: true,
			SecurityGroupInfos: []*entity.SecurityGroupInfo{
				{
					ID:   mockSecurityGroupList[0].SecurityGroupID,
					Name: mockSecurityGroupList[0].SecurityGroupName,
				},
			},
			HostName:   "",
			ServerType: "",
			Status:     common.RunningServerStatus,
			Disks: []*entity.ServerDiskInfo{
				{
					ID:                  mockDiskList[0].DiskID,
					Size:                int(mockDiskList[0].Size),
					IsBoot:              true,
					DeleteOnTermination: true,
					DiskCategory:        mockDiskList[0].Category.String(),
					DiskType:            mockDiskList[0].DiskType,
				},
			},
			Recyclable:  true,
			KeyPairName: mockKeypairList[0].KeyPairName,
			CreatedAt:   1682870400, // 2013-05-01
		},
		{
			ServerID:                "a228ab47-ebb5-40a7-9f0c-7ac0a88b3b05",
			FlavorID:                mockFlavorList[1].FlavorID,
			ServerName:              "测试机器2",
			Description:             "这是一台测试机器",
			CPU:                     int64(mockFlavorList[1].CPU),
			Memory:                  int64(mockFlavorList[1].Memory),
			GPUSpec:                 mockFlavorList[1].GPUSpec,
			GPUAmount:               int64(mockFlavorList[1].GPUAmount),
			OSType:                  mockImageList[1].OSType,
			OSName:                  mockImageList[1].OSName,
			ImageID:                 mockImageList[1].ImageID,
			ImageName:               mockImageList[1].ImageName,
			VlanID:                  "",
			InternetMaxBandwidthIn:  0,
			InternetMaxBandwidthOut: 0,
			VPCAttributes: []*entity.VPCAttribute{
				{
					PrivateIPAddress: "192.223.147.2",
					Mac:              uuid.GenerateUUID(),
					VSwitchId:        mockVpcList[0].VSwitchIDs[0],
					VPCId:            mockVpcList[0].VPCID,
					VPCName:          mockVpcList[0].VPCName,
					VSwitchName:      getVSwitch(mockVpcList[0].VSwitchIDs[0]).VSwitchName,
					Type:             "",
					EipAddress: &entity.EipAddress{
						Name:         mockEipList[1].EIPName,
						AllocationID: mockEipList[1].EIPID,
						IPAddress:    mockEipList[1].EIPAddress,
						Bandwidth:    int64(mockEipList[1].Bandwidth),
					},
				},
			},
			DeviceAvailable: true,
			SecurityGroupInfos: []*entity.SecurityGroupInfo{
				{
					ID:   mockSecurityGroupList[1].SecurityGroupID,
					Name: mockSecurityGroupList[1].SecurityGroupName,
				},
			},
			HostName:   "",
			ServerType: "",
			Status:     common.StoppedServerStatus,
			Disks: []*entity.ServerDiskInfo{
				{
					ID:                  mockDiskList[1].DiskID,
					Size:                int(mockDiskList[1].Size),
					IsBoot:              true,
					DeleteOnTermination: true,
					DiskCategory:        mockDiskList[1].Category.String(),
					DiskType:            mockDiskList[1].DiskType,
				},
			},
			Recyclable:  true,
			KeyPairName: mockKeypairList[1].KeyPairName,
			CreatedAt:   1682870400, // 2013-05-01
		},
	}
	for _, server := range mockServerList {
		mockServerMap[server.ServerID] = server
	}
}

func getServerList() []entity.Server {
	return mockServerList
}

func getServer(id string) entity.Server {
	return mockServerMap[id]
}

func getServerDisk(id string) []*entity.ServerDiskInfo {
	return mockServerMap[id].Disks
}
