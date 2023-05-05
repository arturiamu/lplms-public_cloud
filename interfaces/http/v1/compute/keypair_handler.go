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

func (rc *RouterCompute) CreateKeypair(c *gin.Context) {

}

type DeleteKeypairArgs struct {
}

func (args *DeleteKeypairArgs) toEntityArgs(u *entity.User) *entity.KeypairDeleteArg {
	return nil
}

func (rc *RouterCompute) DeleteKeypair(c *gin.Context) {

}

type UpdateKeypairArgs struct {
}

func (args *UpdateKeypairArgs) toEntityArgs(u *entity.User) *entity.KeypairUpdateArg {
	return nil
}

func (rc *RouterCompute) UpdateKeypair(c *gin.Context) {

}

type GetKeypairArgs struct {
}

func (args *GetKeypairArgs) toEntityArgs(u *entity.User) *entity.KeypairGetArg {
	return nil
}

func (rc *RouterCompute) GetKeypair(c *gin.Context) {

}

type ListKeypairArgs struct {
}

func (args *ListKeypairArgs) toEntityArgs(u *entity.User) *entity.KeypairListArg {
	return nil
}

func (rc *RouterCompute) ListKeypair(c *gin.Context) {

}

type DetachKeyPairArgs struct {
}

func (args *DetachKeyPairArgs) toEntityArgs(u *entity.User) *entity.KeypairDetachArg {
	return nil
}

func (rc *RouterCompute) DetachKeyPair(c *gin.Context) {

}

type AttachKeyPairArgs struct {
}

func (args *AttachKeyPairArgs) toEntityArgs(u *entity.User) *entity.KeypairAttachArg {
	return nil
}

func (rc *RouterCompute) AttachKeyPair(c *gin.Context) {

}
