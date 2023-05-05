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

func (s *Storage) CreateDisk(c *gin.Context) {

}

type DeleteDiskArgs struct {
}

func (args *DeleteDiskArgs) toEntityArgs(u *entity.User) *entity.DiskDeleteArg {
	return nil
}

func (s *Storage) DeleteDisk(c *gin.Context) {

}

type UpdateDiskArgs struct {
}

func (args *UpdateDiskArgs) toEntityArgs(u *entity.User) *entity.DiskUpdateArg {
	return nil
}

func (s *Storage) UpdateDisk(c *gin.Context) {

}

type GetDiskArgs struct {
}

func (args *GetDiskArgs) toEntityArgs(u *entity.User) *entity.DiskGetArg {
	return nil
}

func (s *Storage) GetDisk(c *gin.Context) {

}

type ListDiskArgs struct {
}

func (args *ListDiskArgs) toEntityArgs(u *entity.User) *entity.DiskListArg {
	return nil
}

func (s *Storage) ListDisk(c *gin.Context) {

}

type DetachDiskArgs struct {
}

func (args *DetachDiskArgs) toEntityArgs(u *entity.User) *entity.DiskDetachArg {
	return nil
}

func (s *Storage) DetachDisk(c *gin.Context) {

}

type AttachDiskArgs struct {
}

func (args *AttachDiskArgs) toEntityArgs(u *entity.User) *entity.DiskAttachArg {
	return nil
}

func (s *Storage) AttachDisk(c *gin.Context) {

}

type ResizeDiskArgs struct {
}

func (args *ResizeDiskArgs) toEntityArgs(u *entity.User) *entity.DiskResizeArg {
	return nil
}

func (s *Storage) ResizeDisk(c *gin.Context) {

}

type ResetDiskArgs struct {
}

func (args *ResetDiskArgs) toEntityArgs(u *entity.User) *entity.DiskResetArg {
	return nil
}

func (s *Storage) ResetDisk(c *gin.Context) {

}
