package storage

import (
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/gin-gonic/gin"
)

type CreateDiskArgs struct {
}

func (args *CreateDiskArgs) toEntityArgs(u *entity.User) *entity.DiskCreateArg {
	return nil
}

func (rs *RouterStorage) CreateDisk(c *gin.Context) {

}

type DeleteDiskArgs struct {
}

func (args *DeleteDiskArgs) toEntityArgs(u *entity.User) *entity.DiskDeleteArg {
	return nil
}

func (rs *RouterStorage) DeleteDisk(c *gin.Context) {

}

type UpdateDiskArgs struct {
}

func (args *UpdateDiskArgs) toEntityArgs(u *entity.User) *entity.DiskUpdateArg {
	return nil
}

func (rs *RouterStorage) UpdateDisk(c *gin.Context) {

}

type GetDiskArgs struct {
}

func (args *GetDiskArgs) toEntityArgs(u *entity.User) *entity.DiskGetArg {
	return nil
}

func (rs *RouterStorage) GetDisk(c *gin.Context) {

}

type ListDiskArgs struct {
}

func (args *ListDiskArgs) toEntityArgs(u *entity.User) *entity.DiskListArg {
	return nil
}

func (rs *RouterStorage) ListDisk(c *gin.Context) {

}

type DetachDiskArgs struct {
}

func (args *DetachDiskArgs) toEntityArgs(u *entity.User) *entity.DiskDetachArg {
	return nil
}

func (rs *RouterStorage) DetachDisk(c *gin.Context) {

}

type AttachDiskArgs struct {
}

func (args *AttachDiskArgs) toEntityArgs(u *entity.User) *entity.DiskAttachArg {
	return nil
}

func (rs *RouterStorage) AttachDisk(c *gin.Context) {

}

type ResizeDiskArgs struct {
}

func (args *ResizeDiskArgs) toEntityArgs(u *entity.User) *entity.DiskResizeArg {
	return nil
}

func (rs *RouterStorage) ResizeDisk(c *gin.Context) {

}

type ResetDiskArgs struct {
}

func (args *ResetDiskArgs) toEntityArgs(u *entity.User) *entity.DiskResetArg {
	return nil
}

func (rs *RouterStorage) ResetDisk(c *gin.Context) {

}
