package network

import (
	"errors"
	"github.com/arturiamu/lplms-public_cloud/common"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateVPCWithVSwitchArgs struct {
	VPC     *CreateVPCArgs     `json:"vpc"`
	VSwitch *CreateVSwitchArgs `json:"v_switch"`
}

type CreateVPCArgs struct {
	CIDR        *string `json:"cidr"`
	VPCName     *string `json:"vpc_name"`
	Description *string `json:"description"`
}

func (args *CreateVPCWithVSwitchArgs) toEntityArgs(u *entity.User) *entity.VpcCreateArg {
	var arg = entity.VpcCreateArg{
		VPCName:   args.VPC.VPCName,
		CIDR:      &args.VSwitch.VSwitchCidr,
		ProjectID: u.ProjectID,
	}
	return &arg
}

func (args *CreateVPCWithVSwitchArgs) vpcCreateBaseCheck() (err error) {
	if args.VPC == nil {
		return errors.New("empty vpc")
	}
	if args.VPC.VPCName == nil {
		return errors.New("empty vpc name")
	}
	if args.VSwitch == nil {
		return errors.New("empty vswitch")
	}
	if args.VSwitch.VSwitchName == nil {
		return errors.New("empty vswitch name")
	}
	if args.VSwitch.VSwitchCidr == "" {
		return errors.New("empty cidr")
	}
	return
}

func (rn *RouterNetwork) CreateVpc(c *gin.Context) {
	var (
		args CreateVPCWithVSwitchArgs
		u    = common.GetUser(c)
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	if err := args.vpcCreateBaseCheck(); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	// create vpc
	resp, err := rn.ni.CreateVpc(&entity.VpcCreateArg{
		VPCName:   args.VPC.VPCName,
		CIDR:      &args.VSwitch.VSwitchCidr,
		ProjectID: u.ProjectID,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	if args.VSwitch != nil {
		_, err = rn.ni.CreateVSwitch(&entity.VSwitchCreateArg{
			VPCID:       resp.VPCID,
			CIDR:        args.VSwitch.VSwitchCidr,
			VSwitchName: args.VSwitch.VSwitchName,
			ProjectID:   u.ProjectID,
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
			return
		}
	}
	c.JSON(http.StatusOK, common.SuccessWith("", nil))
}

type DeleteVpcArgs struct {
	VPCID string `json:"vpc_id"`
}

func (args *DeleteVpcArgs) toEntityArgs(u *entity.User) *entity.VpcDeleteArg {
	var arg = entity.VpcDeleteArg{
		ProjectID: u.ProjectID,
		VPCID:     args.VPCID,
	}
	return &arg
}

func (rn *RouterNetwork) DeleteVpc(c *gin.Context) {
	var (
		args DeleteVpcArgs
		u    = common.GetUser(c)
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	vws, err := rn.ni.ListVSwitch(&entity.VSwitchListArg{
		ProjectID: u.ProjectID,
		VPCID:     &args.VPCID,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	if vws.VSwitches != nil {
		c.JSON(http.StatusBadRequest, common.FailWith("please delete virtual switch first", nil))
		return
	}

	_, err = rn.ni.DeleteVpc(args.toEntityArgs(u))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.SuccessWith("", nil))
}

type UpdateVpcArgs struct {
	VPCID string `json:"vpc_id"`

	VPCName     *string `json:"vpc_name"`
	Description *string `json:"description"`
}

func (args *UpdateVpcArgs) toEntityArgs(u *entity.User) *entity.VpcUpdateArg {
	var arg = entity.VpcUpdateArg{
		VPCID:       args.VPCID,
		VPCName:     args.VPCName,
		Description: args.Description,
	}
	return &arg
}

func (rn *RouterNetwork) UpdateVpc(c *gin.Context) {
	var (
		args UpdateVpcArgs
		u    = common.GetUser(c)
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	//if err := validate.Struct(&arg); err != nil {
	//	c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
	//	return
	//}

	entityArgs := args.toEntityArgs(u)

	_, err := rn.ni.UpdateVpc(entityArgs)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.Success())
}

type GetVpcArgs struct {
}

func (args *GetVpcArgs) toEntityArgs(u *entity.User) *entity.VpcGetArg {
	return nil
}

func (rn *RouterNetwork) GetVpc(c *gin.Context) {

}

type ListVpcArgs struct {
	VPCIDs []string `json:"vpc_ids"`
}

func (args *ListVpcArgs) toEntityArgs(u *entity.User) *entity.VpcListArg {
	var arg = entity.VpcListArg{
		ProjectID: u.ProjectID,
		VPCIDs:    args.VPCIDs,
	}
	return &arg
}

func (rn *RouterNetwork) ListVpc(c *gin.Context) {
	var (
		args ListVpcArgs
		u    = common.GetUser(c)
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	resp, err := rn.ni.ListVpc(&entity.VpcListArg{
		VPCIDs:    args.VPCIDs,
		ProjectID: u.ProjectID,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	vpcs := make([]*entity.Vpc, 0, len(resp.VPCs))
	for _, v := range resp.VPCs {
		vpcs = append(vpcs, &entity.Vpc{
			VPCID:       v.VPCID,
			VPCName:     v.VPCName,
			Description: v.Description,
			Status:      v.Status,
			VSwitchIDs:  v.VSwitchIDs,
			CreatedAt:   v.CreatedAt,
			CIDR:        v.CIDR,
		})
	}

	c.JSON(http.StatusOK, common.SuccessWith("", vpcs))
}
