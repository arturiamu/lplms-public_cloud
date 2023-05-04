package application

import (
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/arturiamu/lplms-public_cloud/domain/repository"
)

var _ ComputeAppInterface = &computeApp{}

type ComputeAppInterface interface {
	ServerAppInterface
	FlavorAppInterface
	ImageAppInterface
	KeyPairAppInterface
	SecurityGroupAppInterface
}

type ServerAppInterface interface {
	CreateServer(args *entity.ServerCreateArg) (*entity.ServerCreateResp, error)
	DeleteServer(args *entity.ServerDeleteArg) (*entity.ServerDeleteResp, error)
	UpdateServer(args *entity.ServerUpdateArg) (*entity.ServerUpdateResp, error)
	GetServer(args *entity.ServerGetArg) (*entity.ServerGetResp, error)
	ListServer(args *entity.ServerListArg) (*entity.ServerListResp, error)
	GetServerDisks(args *entity.ServerDisksGetArg) (*entity.ServerDisksGetResp, error)
}

type FlavorAppInterface interface {
	CreateFlavor(args *entity.FlavorCreateArg) (*entity.FlavorCreateResp, error)
	DeleteFlavor(args *entity.FlavorDeleteArg) (*entity.FlavorDeleteResp, error)
	UpdateFlavor(args *entity.FlavorUpdateArg) (*entity.FlavorUpdateResp, error)
	GetFlavor(args *entity.FlavorGetArg) (*entity.FlavorGetResp, error)
	ListFlavor(args *entity.FlavorListArg) (*entity.FlavorListResp, error)
}

type ImageAppInterface interface {
	CreateImage(args *entity.ImageCreateArg) (*entity.ImageCreateResp, error)
	DeleteImage(args *entity.ImageDeleteArg) (*entity.ImageDeleteResp, error)
	UpdateImage(args *entity.ImageUpdateArg) (*entity.ImageUpdateResp, error)
	GetImage(args *entity.ImageGetArg) (*entity.ImageGetResp, error)
	ListImage(args *entity.ImageListArg) (*entity.ImageListResp, error)
}

type KeyPairAppInterface interface {
	CreateKeypair(args *entity.KeypairCreateArg) (*entity.KeypairCreateResp, error)
	DeleteKeypair(args *entity.KeypairDeleteArg) (*entity.KeypairDeleteResp, error)
	UpdateKeypair(args *entity.KeypairUpdateArg) (*entity.KeypairUpdateResp, error)
	GetKeypair(args *entity.KeypairGetArg) (*entity.KeypairGetResp, error)
	ListKeypair(args *entity.KeypairListArg) (*entity.KeypairListResp, error)
	DetachKeyPair(args *entity.KeypairDetachArg) (*entity.KeypairDetachResp, error)
	AttachKeyPair(args *entity.KeypairAttachArg) (*entity.KeypairAttachResp, error)
}

type SecurityGroupAppInterface interface {
	CreateSecurityGroup(args *entity.SecurityGroupCreateArg) (*entity.SecurityGroupCreateResp, error)
	DeleteSecurityGroup(args *entity.SecurityGroupDeleteArg) (*entity.SecurityGroupDeleteResp, error)
	UpdateSecurityGroup(args *entity.SecurityGroupUpdateArg) (*entity.SecurityGroupUpdateResp, error)
	GetSecurityGroup(args *entity.SecurityGroupGetArg) (*entity.SecurityGroupGetResp, error)
	ListSecurityGroup(args *entity.SecurityGroupListArg) (*entity.SecurityGroupListResp, error)
	CreateSecurityGroupRule(args *entity.SecurityGroupRuleCreateArg) (*entity.SecurityGroupRuleCreateResp, error)
	DeleteSecurityGroupRule(args *entity.SecurityGroupRuleDeleteArg) (*entity.SecurityGroupRuleDeleteResp, error)
	GetSecurityGroupRule(args *entity.SecurityGroupRuleGetArg) (*entity.SecurityGroupRuleGetResp, error)
	ListSecurityGroupRule(args *entity.SecurityGroupRuleListArg) (*entity.SecurityGroupRuleListResp, error)
}

type computeApp struct {
	cr repository.ComputeRepositoryInterface
}

func (c *computeApp) GetServerDisks(args *entity.ServerDisksGetArg) (*entity.ServerDisksGetResp, error) {
	return c.cr.GetServerDisks(args)
}

func (c *computeApp) ListSecurityGroupRule(args *entity.SecurityGroupRuleListArg) (*entity.SecurityGroupRuleListResp, error) {
	return c.cr.ListSecurityGroupRule(args)
}

func (c *computeApp) CreateSecurityGroupRule(args *entity.SecurityGroupRuleCreateArg) (*entity.SecurityGroupRuleCreateResp, error) {
	return c.cr.CreateSecurityGroupRule(args)
}

func (c *computeApp) DeleteSecurityGroupRule(args *entity.SecurityGroupRuleDeleteArg) (*entity.SecurityGroupRuleDeleteResp, error) {
	return c.cr.DeleteSecurityGroupRule(args)
}

func (c *computeApp) GetSecurityGroupRule(args *entity.SecurityGroupRuleGetArg) (*entity.SecurityGroupRuleGetResp, error) {
	return c.cr.GetSecurityGroupRule(args)
}

func (c *computeApp) CreateServer(args *entity.ServerCreateArg) (*entity.ServerCreateResp, error) {
	return c.cr.CreateServer(args)
}

func (c *computeApp) DeleteServer(args *entity.ServerDeleteArg) (*entity.ServerDeleteResp, error) {
	return c.cr.DeleteServer(args)
}

func (c *computeApp) UpdateServer(args *entity.ServerUpdateArg) (*entity.ServerUpdateResp, error) {
	return c.cr.UpdateServer(args)
}

func (c *computeApp) GetServer(args *entity.ServerGetArg) (*entity.ServerGetResp, error) {
	return c.cr.GetServer(args)
}

func (c *computeApp) ListServer(args *entity.ServerListArg) (*entity.ServerListResp, error) {
	return c.cr.ListServer(args)
}

func (c *computeApp) CreateFlavor(args *entity.FlavorCreateArg) (*entity.FlavorCreateResp, error) {
	return c.cr.CreateFlavor(args)
}

func (c *computeApp) DeleteFlavor(args *entity.FlavorDeleteArg) (*entity.FlavorDeleteResp, error) {
	return c.cr.DeleteFlavor(args)
}

func (c *computeApp) UpdateFlavor(args *entity.FlavorUpdateArg) (*entity.FlavorUpdateResp, error) {
	return c.cr.UpdateFlavor(args)
}

func (c *computeApp) GetFlavor(args *entity.FlavorGetArg) (*entity.FlavorGetResp, error) {
	return c.cr.GetFlavor(args)
}

func (c *computeApp) ListFlavor(args *entity.FlavorListArg) (*entity.FlavorListResp, error) {
	return c.cr.ListFlavor(args)
}

func (c *computeApp) CreateImage(args *entity.ImageCreateArg) (*entity.ImageCreateResp, error) {
	return c.cr.CreateImage(args)
}

func (c *computeApp) DeleteImage(args *entity.ImageDeleteArg) (*entity.ImageDeleteResp, error) {
	return c.cr.DeleteImage(args)
}

func (c *computeApp) UpdateImage(args *entity.ImageUpdateArg) (*entity.ImageUpdateResp, error) {
	return c.cr.UpdateImage(args)
}

func (c *computeApp) GetImage(args *entity.ImageGetArg) (*entity.ImageGetResp, error) {
	return c.cr.GetImage(args)
}

func (c *computeApp) ListImage(args *entity.ImageListArg) (*entity.ImageListResp, error) {
	return c.cr.ListImage(args)
}

func (c *computeApp) CreateKeypair(args *entity.KeypairCreateArg) (*entity.KeypairCreateResp, error) {
	return c.cr.CreateKeypair(args)
}

func (c *computeApp) DeleteKeypair(args *entity.KeypairDeleteArg) (*entity.KeypairDeleteResp, error) {
	return c.cr.DeleteKeypair(args)
}

func (c *computeApp) UpdateKeypair(args *entity.KeypairUpdateArg) (*entity.KeypairUpdateResp, error) {
	return c.cr.UpdateKeypair(args)
}

func (c *computeApp) GetKeypair(args *entity.KeypairGetArg) (*entity.KeypairGetResp, error) {
	return c.cr.GetKeypair(args)
}

func (c *computeApp) ListKeypair(args *entity.KeypairListArg) (*entity.KeypairListResp, error) {
	return c.cr.ListKeypair(args)
}

func (c *computeApp) DetachKeyPair(args *entity.KeypairDetachArg) (*entity.KeypairDetachResp, error) {
	return c.cr.DetachKeyPair(args)
}

func (c *computeApp) AttachKeyPair(args *entity.KeypairAttachArg) (*entity.KeypairAttachResp, error) {
	return c.cr.AttachKeyPair(args)
}

func (c *computeApp) CreateSecurityGroup(args *entity.SecurityGroupCreateArg) (*entity.SecurityGroupCreateResp, error) {
	return c.cr.CreateSecurityGroup(args)
}

func (c *computeApp) DeleteSecurityGroup(args *entity.SecurityGroupDeleteArg) (*entity.SecurityGroupDeleteResp, error) {
	return c.cr.DeleteSecurityGroup(args)
}

func (c *computeApp) UpdateSecurityGroup(args *entity.SecurityGroupUpdateArg) (*entity.SecurityGroupUpdateResp, error) {
	return c.cr.UpdateSecurityGroup(args)
}

func (c *computeApp) GetSecurityGroup(args *entity.SecurityGroupGetArg) (*entity.SecurityGroupGetResp, error) {
	return c.cr.GetSecurityGroup(args)
}

func (c *computeApp) ListSecurityGroup(args *entity.SecurityGroupListArg) (*entity.SecurityGroupListResp, error) {
	return c.cr.ListSecurityGroup(args)
}
