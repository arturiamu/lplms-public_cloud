package persistence

import (
	"github.com/arturiamu/lplms-public_cloud/common"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
)

type imageInfo struct {
	*entity.Image
	Addr string `json:"-"`
}

var imageMap = map[string]imageInfo{
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
func GetImageInfo(imageID string) imageInfo {
	return imageMap[imageID]
}

func (c *ComputeRepo) CreateImage(arg *entity.ImageCreateArg) (*entity.ImageCreateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) DeleteImage(arg *entity.ImageDeleteArg) (*entity.ImageDeleteResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) UpdateImage(arg *entity.ImageUpdateArg) (*entity.ImageUpdateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) GetImage(arg *entity.ImageGetArg) (*entity.ImageGetResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) ListImage(arg *entity.ImageListArg) (*entity.ImageListResp, error) {
	//TODO implement me
	panic("implement me")
}
