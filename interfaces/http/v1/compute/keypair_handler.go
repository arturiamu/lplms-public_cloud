package compute

import (
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/gin-gonic/gin"
)

type CreateKeypairArgs struct {
}

func (args *CreateKeypairArgs) toEntityArgs(u *entity.User) *entity.KeypairCreateArg {
	return nil
}

func (co *Compute) CreateKeypair(c *gin.Context) {

}

type DeleteKeypairArgs struct {
}

func (args *DeleteKeypairArgs) toEntityArgs(u *entity.User) *entity.KeypairDeleteArg {
	return nil
}

func (co *Compute) DeleteKeypair(c *gin.Context) {

}

type UpdateKeypairArgs struct {
}

func (args *UpdateKeypairArgs) toEntityArgs(u *entity.User) *entity.KeypairUpdateArg {
	return nil
}

func (co *Compute) UpdateKeypair(c *gin.Context) {

}

type GetKeypairArgs struct {
}

func (args *GetKeypairArgs) toEntityArgs(u *entity.User) *entity.KeypairGetArg {
	return nil
}

func (co *Compute) GetKeypair(c *gin.Context) {

}

type ListKeypairArgs struct {
}

func (args *ListKeypairArgs) toEntityArgs(u *entity.User) *entity.KeypairListArg {
	return nil
}

func (co *Compute) ListKeypair(c *gin.Context) {

}

type DetachKeyPairArgs struct {
}

func (args *DetachKeyPairArgs) toEntityArgs(u *entity.User) *entity.KeypairDetachArg {
	return nil
}

func (co *Compute) DetachKeyPair(c *gin.Context) {

}

type AttachKeyPairArgs struct {
}

func (args *AttachKeyPairArgs) toEntityArgs(u *entity.User) *entity.KeypairAttachArg {
	return nil
}

func (co *Compute) AttachKeyPair(c *gin.Context) {

}
