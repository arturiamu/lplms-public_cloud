package network

import (
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/gin-gonic/gin"
)

type CreateVpcArgs struct {
}

func (args *CreateVpcArgs) toEntityArgs(u *entity.User) *entity.VpcCreateArg {
	return nil
}

func (rn *RouterNetwork) CreateVpc(c *gin.Context) {

}

type DeleteVpcArgs struct {
}

func (args *DeleteVpcArgs) toEntityArgs(u *entity.User) *entity.VpcDeleteArg {
	return nil
}

func (rn *RouterNetwork) DeleteVpc(c *gin.Context) {

}

type UpdateVpcArgs struct {
}

func (args *UpdateVpcArgs) toEntityArgs(u *entity.User) *entity.VpcUpdateArg {
	return nil
}

func (rn *RouterNetwork) UpdateVpc(c *gin.Context) {

}

type GetVpcArgs struct {
}

func (args *GetVpcArgs) toEntityArgs(u *entity.User) *entity.VpcGetArg {
	return nil
}

func (rn *RouterNetwork) GetVpc(c *gin.Context) {

}

type ListVpcArgs struct {
}

func (args *ListVpcArgs) toEntityArgs(u *entity.User) *entity.VpcListArg {
	return nil
}

func (rn *RouterNetwork) ListVpc(c *gin.Context) {

}
