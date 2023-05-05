package compute

import (
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/gin-gonic/gin"
)

type CreateSecurityGroupArgs struct {
}

func (args *CreateSecurityGroupArgs) toEntityArgs(u *entity.User) *entity.SecurityGroupCreateArg {
	return nil
}

func (co *Compute) CreateSecurityGroup(c *gin.Context) {

}

type DeleteSecurityGroupArgs struct {
}

func (args *DeleteSecurityGroupArgs) toEntityArgs(u *entity.User) *entity.SecurityGroupDeleteArg {
	return nil
}

func (co *Compute) DeleteSecurityGroup(c *gin.Context) {

}

type UpdateSecurityGroupArgs struct {
}

func (args *UpdateSecurityGroupArgs) toEntityArgs(u *entity.User) *entity.SecurityGroupUpdateArg {
	return nil
}

func (co *Compute) UpdateSecurityGroup(c *gin.Context) {

}

type GetSecurityGroupArgs struct {
}

func (args *GetSecurityGroupArgs) toEntityArgs(u *entity.User) *entity.SecurityGroupGetArg {
	return nil
}

func (co *Compute) GetSecurityGroup(c *gin.Context) {

}

type ListSecurityGroupArgs struct {
}

func (args *ListSecurityGroupArgs) toEntityArgs(u *entity.User) *entity.SecurityGroupListArg {
	return nil
}

func (co *Compute) ListSecurityGroup(c *gin.Context) {

}

type CreateSecurityGroupRuleArgs struct {
}

func (args *CreateSecurityGroupRuleArgs) toEntityArgs(u *entity.User) *entity.SecurityGroupCreateArg {
	return nil
}

func (co *Compute) CreateSecurityGroupRule(c *gin.Context) {

}

type DeleteSecurityGroupRuleArgs struct {
}

func (args *DeleteSecurityGroupRuleArgs) toEntityArgs(u *entity.User) *entity.SecurityGroupDeleteArg {
	return nil
}

func (co *Compute) DeleteSecurityGroupRule(c *gin.Context) {

}

type GetSecurityGroupRuleArgs struct {
}

func (args *GetSecurityGroupRuleArgs) toEntityArgs(u *entity.User) *entity.SecurityGroupRuleGetArg {
	return nil
}

func (co *Compute) GetSecurityGroupRule(c *gin.Context) {

}

type ListSecurityGroupRuleArgs struct {
}

func (args *ListSecurityGroupRuleArgs) toEntityArgs(u *entity.User) *entity.SecurityGroupListArg {
	return nil
}

func (co *Compute) ListSecurityGroupRule(c *gin.Context) {

}
