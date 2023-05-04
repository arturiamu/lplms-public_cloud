package persistence

import (
	"github.com/arturiamu/lplms-public_cloud/common"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
)

type ImageInfo struct {
	*entity.Image
	Addr string `json:"-"`
}

var imageMap = map[string]ImageInfo{
	"45fdcaa4-ec8f-499e-96e1-7b5b44a82e2f": {
		Image: &entity.Image{
			ImageID:            "45fdcaa4-ec8f-499e-96e1-7b5b44a82e2f",
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
		Addr: "docker://docker.qiniu.io:32500/vmidisks/centos:7.6",
	},
	"55fdcaa4-ec8f-499e-96e1-7b5b44a82e2f": {
		Image: &entity.Image{
			ImageID:            "55fdcaa4-ec8f-499e-96e1-7b5b44a82e2f",
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
		Addr: "docker://docker.qiniu.io:32500/vmidisks/centos:7.9",
	},
	"55fdcaa4-ec8f-499e-96e1-7b5b44a82211": {
		Image: &entity.Image{
			ImageID:            "55fdcaa4-ec8f-499e-96e1-7b5b44a82211",
			ImageName:          "UbuntuServer 18.04",
			OSType:             common.LinuxOSType,
			Architecture:       common.ArchitectureType86,
			ImageOwnerAlias:    "system",
			Platform:           "Ubuntu",
			CreatedAt:          1660038097000,
			Size:               1209401344,
			Status:             common.ImageStatusAvailable,
			IsSupportCloudInit: true,
			ImageVersion:       "UbuntuServer 18.04",
		},
		Addr: "docker://docker.qiniu.io:32500/vmidisks/ubuntu:18.04",
	},
}

// GetImageInfo ..
func GetImageInfo(imageID string) ImageInfo {
	return imageMap[imageID]
}

// CreateImage
// 调用 CreateImage 创建一份自定义镜像。您可以使用创建的自定义镜像创建 Server 实例，或者更换实例的系统盘（ReplaceSystemDisk）。
func (c *ComputeRepo) CreateImage(args *entity.ImageCreateArg) (*entity.ImageCreateResp, error) {
	//TODO implement me
	panic("implement me")
}

// DeleteImage
// 调用 DeleteImage 删除一份自定义镜像。
func (c *ComputeRepo) DeleteImage(args *entity.ImageDeleteArg) (*entity.ImageDeleteResp, error) {
	//TODO implement me
	panic("implement me")
}

// UpdateImage
// 调用 ModifyImage 修改一份自定义镜像的名称、描述信息、状态或镜像族系。
func (c *ComputeRepo) UpdateImage(args *entity.ImageUpdateArg) (*entity.ImageUpdateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) GetImage(args *entity.ImageGetArg) (*entity.ImageGetResp, error) {
	//TODO implement me
	panic("implement me")
}

// ListImage
// 调用 DescribeImages 查询您可以使用的镜像资源。
func (c *ComputeRepo) ListImage(args *entity.ImageListArg) (*entity.ImageListResp, error) {
	//TODO implement me
	panic("implement me")
}

///////////////////////////// help functions /////////////
