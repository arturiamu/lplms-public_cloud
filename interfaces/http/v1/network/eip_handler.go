package network

import (
	"github.com/arturiamu/lplms-public_cloud/common"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateEipArgs struct {
}

func (args *CreateEipArgs) toEntityArgs(u *entity.User) *entity.EipCreateArg {
	return nil
}

func (rn *RouterNetwork) CreateEip(c *gin.Context) {

}

type DeleteEipArgs struct {
}

func (args *DeleteEipArgs) toEntityArgs(u *entity.User) *entity.EipDeleteArg {
	return nil
}

func (rn *RouterNetwork) DeleteEip(c *gin.Context) {

}

type UpdateEipArgs struct {
}

func (args *UpdateEipArgs) toEntityArgs(u *entity.User) *entity.EipUpdateArg {
	return nil
}

func (rn *RouterNetwork) UpdateEip(c *gin.Context) {

}

type GetEipArgs struct {
}

func (args *GetEipArgs) toEntityArgs(u *entity.User) *entity.EipGetArg {
	return nil
}

func (rn *RouterNetwork) GetEip(c *gin.Context) {
	var (
		id = c.Param("id")
	)

	resp, err := rn.ni.GetEip(&entity.EipGetArg{EIPID: id})
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.SuccessWith("", resp.Eip))
}

type ListEipArgs struct {
}

func (args *ListEipArgs) toEntityArgs(u *entity.User) *entity.EipListArg {
	return nil
}

func (rn *RouterNetwork) ListEip(c *gin.Context) {
	resp, err := rn.ni.ListEip(&entity.EipListArg{})
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.SuccessWith("", resp.EIPs))
}
