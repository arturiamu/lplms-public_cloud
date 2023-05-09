package mock

import (
	"github.com/arturiamu/lplms-public_cloud/common"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/arturiamu/lplms-public_cloud/utils/uuid"
)

var mockSnapshotMap map[string]entity.SnapshotInfo

var mockSnapshotList []entity.SnapshotInfo

func initSnapshot() {
	mockSnapshotMap = make(map[string]entity.SnapshotInfo)
	mockSnapshotList = []entity.SnapshotInfo{
		{
			SnapshotID:       uuid.GenerateUUID(),
			SnapshotName:     "测试快照1",
			Usage:            "none",
			Status:           common.AccomplishedSnapshotStatus,
			Progress:         "100",
			RemainTime:       0,
			Description:      "这是一个测试快照",
			SourceDiskID:     "",
			SourceDiskSize:   mockDiskList[0].Size,
			SourceDiskName:   mockDiskList[0].DiskName,
			SourceDiskStatus: mockDiskList[0].Status,
			CreatedAt:        1682870400,
		},
		{
			SnapshotID:       uuid.GenerateUUID(),
			SnapshotName:     "测试快照2",
			Usage:            "none",
			Status:           common.AccomplishedSnapshotStatus,
			Progress:         "100",
			RemainTime:       0,
			Description:      "这是一个测试快照",
			SourceDiskID:     "",
			SourceDiskSize:   mockDiskList[1].Size,
			SourceDiskName:   mockDiskList[1].DiskName,
			SourceDiskStatus: mockDiskList[1].Status,
			CreatedAt:        1682870400,
		},
		{
			SnapshotID:       uuid.GenerateUUID(),
			SnapshotName:     "测试快照3",
			Usage:            "none",
			Status:           common.ProcessingSnapshotStatus,
			Progress:         "55",
			RemainTime:       20,
			Description:      "这是一个测试快照",
			SourceDiskID:     "",
			SourceDiskSize:   mockDiskList[2].Size,
			SourceDiskName:   mockDiskList[2].DiskName,
			SourceDiskStatus: mockDiskList[2].Status,
			CreatedAt:        1682870400,
		},
		{
			SnapshotID:       uuid.GenerateUUID(),
			SnapshotName:     "测试快照4",
			Usage:            "none",
			Status:           common.FailedSnapshotStatus,
			Progress:         "100",
			RemainTime:       0,
			Description:      "这是一个测试快照",
			SourceDiskID:     "",
			SourceDiskSize:   mockDiskList[3].Size,
			SourceDiskName:   mockDiskList[3].DiskName,
			SourceDiskStatus: mockDiskList[3].Status,
			CreatedAt:        1682870400,
		},
	}
	for _, info := range mockSnapshotList {
		mockSnapshotMap[info.SnapshotID] = info
	}
}

func getSnapshot(id string) entity.SnapshotInfo {
	return mockSnapshotMap[id]
}

func getSnapshotList() []entity.SnapshotInfo {
	return mockSnapshotList
}
