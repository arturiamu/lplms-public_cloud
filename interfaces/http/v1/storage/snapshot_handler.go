package storage

import (
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/gin-gonic/gin"
)

type CreateSnapshotArgs struct {
}

func (args *CreateSnapshotArgs) toEntityArgs(u *entity.User) *entity.SnapshotCreateArg {
	return nil
}

func (s *Storage) CreateSnapshot(c *gin.Context) {

}

type DeleteSnapshotArgs struct {
}

func (args *DeleteSnapshotArgs) toEntityArgs(u *entity.User) *entity.SnapshotDeleteArg {
	return nil
}

func (s *Storage) DeleteSnapshot(c *gin.Context) {

}

type UpdateSnapshotArgs struct {
}

func (args *UpdateSnapshotArgs) toEntityArgs(u *entity.User) *entity.SnapshotUpdateArg {
	return nil
}

func (s *Storage) UpdateSnapshot(c *gin.Context) {

}

type GetSnapshotArgs struct {
}

func (args *GetSnapshotArgs) toEntityArgs(u *entity.User) *entity.SnapshotGetArg {
	return nil
}

func (s *Storage) GetSnapshot(c *gin.Context) {

}

type ListSnapshotArgs struct {
}

func (args *ListSnapshotArgs) toEntityArgs(u *entity.User) *entity.SnapshotListArg {
	return nil
}

func (s *Storage) ListSnapshot(c *gin.Context) {

}
