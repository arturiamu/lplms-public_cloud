package mock

import (
	"github.com/arturiamu/lplms-public_cloud/common"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/arturiamu/lplms-public_cloud/utils/uuid"
)

var mockImageMap map[string]entity.Image

var mockImageList = []entity.Image{
	{
		ImageID:            uuid.GenerateUUID(),
		SnapshotID:         "",                     // 快照 ID
		Description:        "CentOS 7.6 官方镜像",      // 镜像描述
		OSName:             "CentOS 7.6",           // 操作系统的中文显示名称
		Progress:           "",                     // 对于导入中的镜像，返回导入任务的进度
		IsCopied:           false,                  // 是否是拷贝的镜像
		MinImageSize:       5 * 1024 * 1024 * 1024, // 镜像要求的最小系统盘容量，单位：Bytes
		ImageName:          "CentOS 7.6",
		OSType:             common.LinuxOSType,
		Architecture:       common.ArchitectureType86,
		ImageOwnerAlias:    "system",
		Platform:           "CentOS",
		CreatedAt:          1660038097000,
		Size:               4 * 1024 * 1024,
		Status:             common.ImageStatusAvailable,
		IsSupportCloudInit: true,
		ImageVersion:       "CentOS 7.6",
	},
	{
		ImageID:            uuid.GenerateUUID(),
		SnapshotID:         "",                     // 快照 ID
		Description:        "CentOS 7.9 官方镜像",      // 镜像描述
		OSName:             "CentOS 7.9",           // 操作系统的中文显示名称
		Progress:           "",                     // 对于导入中的镜像，返回导入任务的进度
		IsCopied:           false,                  // 是否是拷贝的镜像
		MinImageSize:       5 * 1024 * 1024 * 1024, // 镜像要求的最小系统盘容量，单位：Bytes
		ImageName:          "CentOS 7.9",
		OSType:             common.LinuxOSType,
		Architecture:       common.ArchitectureType86,
		ImageOwnerAlias:    "system",
		Platform:           "CentOS",
		CreatedAt:          1660038097000,
		Size:               4 * 1024 * 1024,
		Status:             common.ImageStatusAvailable,
		IsSupportCloudInit: true,
		ImageVersion:       "CentOS 7.9",
	},
	{
		ImageID:            uuid.GenerateUUID(),
		SnapshotID:         "",                        // 快照 ID
		Description:        "UbuntuServer 18.04 官方镜像", // 镜像描述
		OSName:             "乌班图服务器版",                 // 操作系统的中文显示名称
		Progress:           "",                        // 对于导入中的镜像，返回导入任务的进度
		IsCopied:           false,                     // 是否是拷贝的镜像
		MinImageSize:       10 * 1024 * 1024 * 1024,   // 镜像要求的最小系统盘容量，单位：Bytes
		ImageName:          "UbuntuServer 18.04",
		OSType:             common.LinuxOSType,
		Architecture:       common.ArchitectureType86,
		ImageOwnerAlias:    "system",
		Platform:           "Ubuntu",
		CreatedAt:          1660038097000,
		Size:               8 * 1024 * 1024 * 1024,
		Status:             common.ImageStatusAvailable,
		IsSupportCloudInit: true,
		ImageVersion:       "UbuntuServer 18.04",
	},
	{
		ImageID:            uuid.GenerateUUID(),
		SnapshotID:         "",                      // 快照 ID
		Description:        "Ubuntu 20.04 官方镜像",     // 镜像描述
		OSName:             "乌班图",                   // 操作系统的中文显示名称
		Progress:           "",                      // 对于导入中的镜像，返回导入任务的进度
		IsCopied:           false,                   // 是否是拷贝的镜像
		MinImageSize:       10 * 1024 * 1024 * 1024, // 镜像要求的最小系统盘容量，单位：Bytes
		ImageName:          "Ubuntu 20.04",
		OSType:             common.LinuxOSType,
		Architecture:       common.ArchitectureType86,
		ImageOwnerAlias:    "system",
		Platform:           "Ubuntu",
		CreatedAt:          1660038097000,
		Size:               9 * 1024 * 1024 * 1024,
		Status:             common.ImageStatusAvailable,
		IsSupportCloudInit: true,
		ImageVersion:       "Ubuntu 20.04",
	},
}

func initImage() {
	mockImageMap = make(map[string]entity.Image)
	for _, image := range mockImageList {
		mockImageMap[image.ImageID] = image
	}
}

func getImageList() []entity.Image {
	return mockImageList
}

func getImage(id string) entity.Image {
	return mockImageMap[id]
}
