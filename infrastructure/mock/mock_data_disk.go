package mock

import (
	"github.com/arturiamu/lplms-public_cloud/common"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/arturiamu/lplms-public_cloud/utils/pointerutil"
	"github.com/arturiamu/lplms-public_cloud/utils/uuid"
)

var mockDiskMap map[string]entity.Disk

var mockDiskList []entity.Disk

func initDisk() {
	mockDiskMap = make(map[string]entity.Disk)
	mockDiskList = []entity.Disk{
		{
			DiskID:           uuid.GenerateUUID(),
			ImageID:          mockImageList[0].ImageID,
			Device:           "",
			DetachedAt:       0,
			DiskType:         pointerutil.Pointer(common.SystemDiskType),
			Bootable:         false,
			ServerID:         "600c674b-9753-47bc-b092-63157c9356e2", //mockServerList[0].ServerID,
			ServerName:       "测试机器1",                                //mockServerList[0].ServerName,
			AttachedTime:     0,
			SourceSnapshotID: "",
			Size:             20,
			Description:      "测试机器1-系统盘", //mockServerList[0].ServerName + "-系统盘",
			Portable:         false,
			DiskName:         "测试系统盘1",
			CreatedAt:        1680594796000,
			Status:           common.DiskInUse,
			Category:         common.CloudSSDiskCategory,
			DeleteWithServer: true,
			OSType:           mockImageList[0].OSType,
			OSName:           mockImageList[0].OSName,
			ImageName:        mockImageList[0].ImageName,
		},
		{
			DiskID:           uuid.GenerateUUID(),
			ImageID:          mockImageList[1].ImageID,
			Device:           "",
			DetachedAt:       0,
			DiskType:         pointerutil.Pointer(common.SystemDiskType),
			Bootable:         false,
			ServerID:         "a228ab47-ebb5-40a7-9f0c-7ac0a88b3b05", //mockServerList[1].ServerID,
			ServerName:       "测试机器2",                                //mockServerList[1].ServerName,
			AttachedTime:     0,
			SourceSnapshotID: "",
			Size:             20,
			Description:      "测试机器2-系统盘", //mockServerList[1].ServerName + "-系统盘",
			Portable:         false,
			DiskName:         "测试系统盘2",
			CreatedAt:        1680594796000,
			Status:           common.DiskInUse,
			Category:         common.CloudSSDiskCategory,
			DeleteWithServer: true,
			OSType:           mockImageList[1].OSType,
			OSName:           mockImageList[1].OSName,
			ImageName:        mockImageList[1].ImageName,
		},
		// 数据盘
		{
			DiskID:           uuid.GenerateUUID(),
			ImageID:          "",
			Device:           "",
			DetachedAt:       0,
			DiskType:         pointerutil.Pointer(common.DataDiskType),
			Bootable:         false,
			ServerID:         "",
			ServerName:       "",
			AttachedTime:     0,
			SourceSnapshotID: "",
			Size:             500,
			Description:      "这个一个测试数据盘",
			Portable:         false,
			DiskName:         "测试数据盘1",
			CreatedAt:        1680594796000,
			Status:           common.DiskAvailable,
			Category:         common.CloudSSDiskCategory,
			DeleteWithServer: true,
			OSType:           "",
			OSName:           "",
			ImageName:        "",
		},
		{
			DiskID:           uuid.GenerateUUID(),
			ImageID:          "",
			Device:           "",
			DetachedAt:       0,
			DiskType:         pointerutil.Pointer(common.DataDiskType),
			Bootable:         false,
			ServerID:         "",
			ServerName:       "",
			AttachedTime:     0,
			SourceSnapshotID: "",
			Size:             300,
			Description:      "这个一个测试数据盘",
			Portable:         false,
			DiskName:         "测试数据盘2",
			CreatedAt:        1680594796000,
			Status:           common.DiskAvailable,
			Category:         common.CloudSSDiskCategory,
			DeleteWithServer: true,
			OSType:           "",
			OSName:           "",
			ImageName:        "",
		},
	}
	for _, flavor := range mockFlavorList {
		mockFlavorMap[flavor.FlavorID] = flavor
	}
}

func getDiskList() []entity.Disk {
	return mockDiskList
}

func getDisk(id string) entity.Disk {
	return mockDiskMap[id]
}
