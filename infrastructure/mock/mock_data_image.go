package mock

import (
	"github.com/arturiamu/lplms-public_cloud/common"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
)

var mockImageMap map[string]entity.Image

var mockImageList = []entity.Image{
	{
		ImageID:            "1aba5759-c11c-4990-9278-ec70879819f0",
		SnapshotID:         "",    // 快照 ID
		Description:        "",    // 镜像描述
		OSName:             "",    // 操作系统的中文显示名称
		Progress:           "",    // 对于导入中的镜像，返回导入任务的进度
		IsCopied:           false, // 是否是拷贝的镜像
		MinImageSize:       0,     // 镜像要求的最小系统盘容量，单位：Bytes
		ImageName:          "CentOS 7.6",
		OSType:             common.LinuxOSType,
		Architecture:       common.ArchitectureType86,
		ImageOwnerAlias:    "system",
		Platform:           "CentOS",
		CreatedAt:          1660038097000,
		Size:               1209401344,
		Status:             common.ImageStatusAvailable,
		IsSupportCloudInit: true,
		ImageVersion:       "CentOS 7.6",
	},
	{
		ImageID:            "43dda316-dd89-4d7a-b848-88ba2fefe402",
		SnapshotID:         "",    // 快照 ID
		Description:        "",    // 镜像描述
		OSName:             "",    // 操作系统的中文显示名称
		Progress:           "",    // 对于导入中的镜像，返回导入任务的进度
		IsCopied:           false, // 是否是拷贝的镜像
		MinImageSize:       0,     // 镜像要求的最小系统盘容量，单位：Bytes
		ImageName:          "CentOS 7.9",
		OSType:             common.LinuxOSType,
		Architecture:       common.ArchitectureType86,
		ImageOwnerAlias:    "system",
		Platform:           "CentOS",
		CreatedAt:          1660038097000,
		Size:               1209401344,
		Status:             common.ImageStatusAvailable,
		IsSupportCloudInit: true,
		ImageVersion:       "CentOS 7.9",
	},
	{
		ImageID:            "cebec4f6-e3e0-4e7b-8bbe-95dbff193af5",
		SnapshotID:         "",    // 快照 ID
		Description:        "",    // 镜像描述
		OSName:             "",    // 操作系统的中文显示名称
		Progress:           "",    // 对于导入中的镜像，返回导入任务的进度
		IsCopied:           false, // 是否是拷贝的镜像
		MinImageSize:       0,     // 镜像要求的最小系统盘容量，单位：Bytes
		ImageName:          "UbuntuServer 18.04",
		OSType:             common.LinuxOSType,
		Architecture:       common.ArchitectureType86,
		ImageOwnerAlias:    "system",
		Platform:           "Ubuntu",
		CreatedAt:          1660038097000,
		Size:               1209401344,
		Status:             common.ImageStatusAvailable,
		IsSupportCloudInit: true,
		ImageVersion:       "CentOS 7.6",
	},
}

func init() {
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
