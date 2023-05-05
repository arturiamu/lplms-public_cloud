package network

import (
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/gin-gonic/gin"
)

type CreateVSwitchArgs struct {
}

func (args *CreateVSwitchArgs) toEntityArgs(u *entity.User) *entity.VSwitchCreateArg {
	return nil
}

func (ne *Network) CreateVSwitch(c *gin.Context) {

}

type DeleteVSwitchArgs struct {
}

func (args *DeleteVSwitchArgs) toEntityArgs(u *entity.User) *entity.VSwitchDeleteArg {
	return nil
}

func (ne *Network) DeleteVSwitch(c *gin.Context) {

}

type UpdateVSwitchArgs struct {
}

func (args *UpdateVSwitchArgs) toEntityArgs(u *entity.User) *entity.VSwitchUpdateArg {
	return nil
}

func (ne *Network) UpdateVSwitch(c *gin.Context) {

}

type GetVSwitchArgs struct {
}

func (args *GetVSwitchArgs) toEntityArgs(u *entity.User) *entity.VSwitchGetArg {
	return nil
}

func (ne *Network) GetVSwitch(c *gin.Context) {

}

type ListVSwitchArgs struct {
}

func (args *ListVSwitchArgs) toEntityArgs(u *entity.User) *entity.VSwitchListArg {
	return nil
}

func (ne *Network) ListVSwitch(c *gin.Context) {

}
