package network

import (
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/gin-gonic/gin"
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

}

type ListEipArgs struct {
}

func (args *ListEipArgs) toEntityArgs(u *entity.User) *entity.EipListArg {
	return nil
}

func (rn *RouterNetwork) ListEip(c *gin.Context) {

}
