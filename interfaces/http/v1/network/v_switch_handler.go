package network

import (
	"github.com/arturiamu/lplms-public_cloud/common"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateVSwitchArgs struct {
	VPCID       string  `json:"vpc_id"`
	VSwitchName *string `json:"v_switch_name"`
	VSwitchCidr string  `json:"v_switch_cidr"`
	Description *string `json:"description"`
}

func (args *CreateVSwitchArgs) toEntityArgs(u *entity.User) *entity.VSwitchCreateArg {
	var arg = entity.VSwitchCreateArg{
		CIDR:        args.VSwitchCidr,
		VSwitchName: args.VSwitchName,
		VPCID:       args.VPCID,
		Description: args.Description,
		ProjectID:   u.ProjectID,
	}
	return &arg
}

func (rn *RouterNetwork) CreateVSwitch(c *gin.Context) {
	var (
		args CreateVSwitchArgs
		u    = common.GetUser(c)
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	if args.VSwitchName == nil {
		c.JSON(http.StatusBadRequest, common.FailWith("empty vswitch name", nil))
		return
	}
	if args.VSwitchCidr == "" {
		c.JSON(http.StatusBadRequest, common.FailWith("empty vswitch cidr", nil))
		return
	}

	// create v_switch
	_, err := rn.ni.CreateVSwitch(args.toEntityArgs(u))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, common.SuccessWith("", nil))
}

type DeleteVSwitchArgs struct {
	VPCID     *string `pos:"path:id"`
	VSwitchID string  `pos:"path:vswitch_id"`
}

func (args *DeleteVSwitchArgs) toEntityArgs(u *entity.User) *entity.VSwitchDeleteArg {
	var arg = entity.VSwitchDeleteArg{
		VPCID:     *args.VPCID,
		VSwitchID: args.VSwitchID,
	}
	return &arg
}

func (rn *RouterNetwork) DeleteVSwitch(c *gin.Context) {
	var (
		args DeleteVSwitchArgs
		u    = common.GetUser(c)
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	// create v_switch
	_, err := rn.ni.DeleteVSwitch(args.toEntityArgs(u))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, common.SuccessWith("", nil))
}

type UpdateVSwitchArgs struct {
	VSwitchID   *string `json:"v_switch_id"`
	VPCID       *string `json:"vpc_id"`
	VSwitchName *string `json:"v_switch_name"`
	Description *string `json:"description"`
}

func (args *UpdateVSwitchArgs) toEntityArgs(u *entity.User) *entity.VSwitchUpdateArg {
	var arg = entity.VSwitchUpdateArg{
		VSwitchID:   *args.VSwitchID,
		VSwitchName: args.VSwitchName,
		Description: args.Description,
	}
	return &arg
}

func (rn *RouterNetwork) UpdateVSwitch(c *gin.Context) {
	var (
		args UpdateVSwitchArgs
		u    = common.GetUser(c)
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	// create v_switch
	_, err := rn.ni.UpdateVSwitch(args.toEntityArgs(u))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, common.SuccessWith("", nil))
}

type GetVSwitchArgs struct {
}

func (args *GetVSwitchArgs) toEntityArgs(u *entity.User) *entity.VSwitchGetArg {
	return nil
}

func (rn *RouterNetwork) GetVSwitch(c *gin.Context) {

}

type ListVSwitchArgs struct {
	VPCID      *string  `json:"vpc_id"`
	VSwitchIDs []string `json:"v_switch_ids"`
}

func (args *ListVSwitchArgs) toEntityArgs(u *entity.User) *entity.VSwitchListArg {
	var arg = entity.VSwitchListArg{
		VPCID:      args.VPCID,
		VSwitchIDs: args.VSwitchIDs,
		ProjectID:  u.ProjectID,
	}
	return &arg
}

func (rn *RouterNetwork) ListVSwitch(c *gin.Context) {
	var (
		args ListVSwitchArgs
		u    = common.GetUser(c)
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	// create v_switch
	resp, err := rn.ni.ListVSwitch(args.toEntityArgs(u))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	vswitchs := make([]*entity.VSwitch, 0, len(resp.VSwitches))
	for _, v := range resp.VSwitches {
		vswitchs = append(vswitchs, &entity.VSwitch{
			VSwitchID:        v.VSwitchID,
			VSwitchName:      v.VSwitchName,
			VPCID:            v.VPCID,
			CIDR:             v.CIDR,
			Status:           v.Status,
			Description:      v.Description,
			CreatedAt:        v.CreatedAt,
			AvailableIPCount: v.AvailableIPCount,
			ResourceCount:    v.ResourceCount,
		})
	}

	c.JSON(http.StatusOK, common.SuccessWith("", vswitchs))
}
