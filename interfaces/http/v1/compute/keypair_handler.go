package compute

import (
	"github.com/arturiamu/lplms-public_cloud/common"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateKeypairArgs struct {
	KeyPairName   string  `json:"key_pair_name"`
	PublicKeyBody *string `json:"public_key_body"`
}

func (args *CreateKeypairArgs) toEntityArgs(u *entity.User) *entity.KeypairCreateArg {
	var arg = entity.KeypairCreateArg{
		ProjectID:     u.ProjectID,
		KeyPairName:   args.KeyPairName,
		PublicKeyBody: args.PublicKeyBody,
	}

	return &arg
}

func (rc *RouterCompute) CreateKeypair(c *gin.Context) {
	var (
		args   CreateKeypairArgs
		ctxUid = common.GetUid(c)
		ctxPid = common.GetProject(c)
		u      = &entity.User{
			UID:       ctxUid,
			ProjectID: ctxPid,
		}
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	createArgs := args.toEntityArgs(u)
	_, err := rc.ci.CreateKeypair(createArgs)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.Success())
}

type DeleteKeypairArgs struct {
	KeyPairName string `json:"key_pair_name"`
}

func (args *DeleteKeypairArgs) toEntityArgs(u *entity.User) *entity.KeypairDeleteArg {
	var arg = entity.KeypairDeleteArg{
		ProjectID:   u.ProjectID,
		KeyPairName: args.KeyPairName,
	}

	return &arg
}

func (rc *RouterCompute) DeleteKeypair(c *gin.Context) {
	var (
		args   DeleteKeypairArgs
		ctxUid = common.GetUid(c)
		ctxPid = common.GetProject(c)
		u      = &entity.User{
			UID:       ctxUid,
			ProjectID: ctxPid,
		}
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	createArgs := args.toEntityArgs(u)
	_, err := rc.ci.DeleteKeypair(createArgs)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.Success())
}

type UpdateKeypairArgs struct {
}

func (args *UpdateKeypairArgs) toEntityArgs(u *entity.User) *entity.KeypairUpdateArg {
	return nil
}

func (rc *RouterCompute) UpdateKeypair(c *gin.Context) {

}

type GetKeypairArgs struct {
	KeyPairName string `json:"key_pair_name"`
}

func (args *GetKeypairArgs) toEntityArgs(u *entity.User) *entity.KeypairGetArg {
	var arg = entity.KeypairGetArg{
		ProjectID:   u.ProjectID,
		KeyPairName: &args.KeyPairName,
	}

	return &arg
}

func (rc *RouterCompute) GetKeypair(c *gin.Context) {
	var (
		args   GetKeypairArgs
		ctxUid = common.GetUid(c)
		ctxPid = common.GetProject(c)
		u      = &entity.User{
			UID:       ctxUid,
			ProjectID: ctxPid,
		}
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	createArgs := args.toEntityArgs(u)
	_, err := rc.ci.GetKeypair(createArgs)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.Success())
}

type ListKeypairArgs struct {
	KeyPairName *string `json:"key_pair_name"`
}

func (args *ListKeypairArgs) toEntityArgs(u *entity.User) *entity.KeypairListArg {
	var arg = entity.KeypairListArg{
		ProjectID:   u.ProjectID,
		KeyPairName: args.KeyPairName,
	}

	return &arg
}

func (rc *RouterCompute) ListKeypair(c *gin.Context) {
	var (
		args   ListKeypairArgs
		ctxUid = common.GetUid(c)
		ctxPid = common.GetProject(c)
		u      = &entity.User{
			UID:       ctxUid,
			ProjectID: ctxPid,
		}
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	createArgs := args.toEntityArgs(u)
	_, err := rc.ci.ListKeypair(createArgs)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.Success())
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
