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

func (rc *RouterCompute) CreateSecurityGroup(c *gin.Context) {

}

type DeleteSecurityGroupArgs struct {
}

func (args *DeleteSecurityGroupArgs) toEntityArgs(u *entity.User) *entity.SecurityGroupDeleteArg {
	return nil
}

func (rc *RouterCompute) DeleteSecurityGroup(c *gin.Context) {

}

type UpdateSecurityGroupArgs struct {
}

func (args *UpdateSecurityGroupArgs) toEntityArgs(u *entity.User) *entity.SecurityGroupUpdateArg {
	return nil
}

func (rc *RouterCompute) UpdateSecurityGroup(c *gin.Context) {

}

type GetSecurityGroupArgs struct {
}

func (args *GetSecurityGroupArgs) toEntityArgs(u *entity.User) *entity.SecurityGroupGetArg {
	return nil
}

func (rc *RouterCompute) GetSecurityGroup(c *gin.Context) {

}

type ListSecurityGroupArgs struct {
}

func (args *ListSecurityGroupArgs) toEntityArgs(u *entity.User) *entity.SecurityGroupListArg {
	return nil
}

func (rc *RouterCompute) ListSecurityGroup(c *gin.Context) {

}

type CreateSecurityGroupRuleArgs struct {
}

func (args *CreateSecurityGroupRuleArgs) toEntityArgs(u *entity.User) *entity.SecurityGroupCreateArg {
	return nil
}

func (rc *RouterCompute) CreateSecurityGroupRule(c *gin.Context) {

}

type DeleteSecurityGroupRuleArgs struct {
}

func (args *DeleteSecurityGroupRuleArgs) toEntityArgs(u *entity.User) *entity.SecurityGroupDeleteArg {
	return nil
}

func (rc *RouterCompute) DeleteSecurityGroupRule(c *gin.Context) {

}

type GetSecurityGroupRuleArgs struct {
}

func (args *GetSecurityGroupRuleArgs) toEntityArgs(u *entity.User) *entity.SecurityGroupRuleGetArg {
	return nil
}

func (rc *RouterCompute) GetSecurityGroupRule(c *gin.Context) {

}

type ListSecurityGroupRuleArgs struct {
}

func (args *ListSecurityGroupRuleArgs) toEntityArgs(u *entity.User) *entity.SecurityGroupListArg {
	return nil
}

func (rc *RouterCompute) ListSecurityGroupRule(c *gin.Context) {

}
