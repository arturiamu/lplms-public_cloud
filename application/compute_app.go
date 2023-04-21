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
	CreateServer(arg *entity.ServerCreateArg) (*entity.Server, error)
	DeleteServer(arg *entity.ServerDeleteArg) (*entity.Server, error)
	UpdateServer(arg *entity.ServerUpdateArg) (*entity.Server, error)
	GetServer(arg *entity.ServerGetArg) (*entity.Server, error)
	ListServer(arg *entity.ServerListArg) ([]*entity.Server, error)
}

type FlavorAppInterface interface {
	CreateFlavor(arg *entity.FlavorCreateArg) (*entity.Flavor, error)
	DeleteFlavor(arg *entity.FlavorDeleteArg) (*entity.Flavor, error)
	UpdateFlavor(arg *entity.FlavorUpdateArg) (*entity.Flavor, error)
	GetFlavor(arg *entity.FlavorGetArg) (*entity.Flavor, error)
	ListFlavor(arg *entity.FlavorListArg) ([]*entity.Flavor, error)
}

type ImageAppInterface interface {
	CreateImage(arg *entity.ImageCreateArg) (*entity.Image, error)
	DeleteImage(arg *entity.ImageDeleteArg) (*entity.Image, error)
	UpdateImage(arg *entity.ImageUpdateArg) (*entity.Image, error)
	GetImage(arg *entity.ImageGetArg) (*entity.Image, error)
	ListImage(arg *entity.ImageListArg) ([]*entity.Image, error)
}

type KeyPairAppInterface interface {
	CreateKeypair(arg *entity.KeypairCreateArg) (*entity.Keypair, error)
	DeleteKeypair(arg *entity.KeypairDeleteArg) (*entity.Keypair, error)
	UpdateKeypair(arg *entity.KeypairUpdateArg) (*entity.Keypair, error)
	GetKeypair(arg *entity.KeypairGetArg) (*entity.Keypair, error)
	ListKeypair(arg *entity.KeypairListArg) ([]*entity.Keypair, error)
}

type SecurityGroupAppInterface interface {
	CreateSecurityGroup(arg *entity.SecurityGroupCreateArg) (*entity.SecurityGroup, error)
	DeleteSecurityGroup(arg *entity.SecurityGroupDeleteArg) (*entity.SecurityGroup, error)
	UpdateSecurityGroup(arg *entity.SecurityGroupUpdateArg) (*entity.SecurityGroup, error)
	GetSecurityGroup(arg *entity.SecurityGroupGetArg) (*entity.SecurityGroup, error)
	ListSecurityGroup(arg *entity.SecurityGroupListArg) ([]*entity.SecurityGroup, error)
}

type computeApp struct {
	cr repository.ComputeRepositoryInterface
}

func (c *computeApp) CreateServer(arg *entity.ServerCreateArg) (*entity.Server, error) {
	return c.cr.CreateServer(arg)
}

func (c *computeApp) DeleteServer(arg *entity.ServerDeleteArg) (*entity.Server, error) {
	//TODO implement me
	panic("implement me")
}

func (c *computeApp) UpdateServer(arg *entity.ServerUpdateArg) (*entity.Server, error) {
	//TODO implement me
	panic("implement me")
}

func (c *computeApp) GetServer(arg *entity.ServerGetArg) (*entity.Server, error) {
	//TODO implement me
	panic("implement me")
}

func (c *computeApp) ListServer(arg *entity.ServerListArg) ([]*entity.Server, error) {
	//TODO implement me
	panic("implement me")
}

func (c *computeApp) CreateFlavor(arg *entity.FlavorCreateArg) (*entity.Flavor, error) {
	//TODO implement me
	panic("implement me")
}

func (c *computeApp) DeleteFlavor(arg *entity.FlavorDeleteArg) (*entity.Flavor, error) {
	//TODO implement me
	panic("implement me")
}

func (c *computeApp) UpdateFlavor(arg *entity.FlavorUpdateArg) (*entity.Flavor, error) {
	//TODO implement me
	panic("implement me")
}

func (c *computeApp) GetFlavor(arg *entity.FlavorGetArg) (*entity.Flavor, error) {
	//TODO implement me
	panic("implement me")
}

func (c *computeApp) ListFlavor(arg *entity.FlavorListArg) ([]*entity.Flavor, error) {
	//TODO implement me
	panic("implement me")
}

func (c *computeApp) CreateImage(arg *entity.ImageCreateArg) (*entity.Image, error) {
	//TODO implement me
	panic("implement me")
}

func (c *computeApp) DeleteImage(arg *entity.ImageDeleteArg) (*entity.Image, error) {
	//TODO implement me
	panic("implement me")
}

func (c *computeApp) UpdateImage(arg *entity.ImageUpdateArg) (*entity.Image, error) {
	//TODO implement me
	panic("implement me")
}

func (c *computeApp) GetImage(arg *entity.ImageGetArg) (*entity.Image, error) {
	//TODO implement me
	panic("implement me")
}

func (c *computeApp) ListImage(arg *entity.ImageListArg) ([]*entity.Image, error) {
	//TODO implement me
	panic("implement me")
}

func (c *computeApp) CreateKeypair(arg *entity.KeypairCreateArg) (*entity.Keypair, error) {
	//TODO implement me
	panic("implement me")
}

func (c *computeApp) DeleteKeypair(arg *entity.KeypairDeleteArg) (*entity.Keypair, error) {
	//TODO implement me
	panic("implement me")
}

func (c *computeApp) UpdateKeypair(arg *entity.KeypairUpdateArg) (*entity.Keypair, error) {
	//TODO implement me
	panic("implement me")
}

func (c *computeApp) GetKeypair(arg *entity.KeypairGetArg) (*entity.Keypair, error) {
	//TODO implement me
	panic("implement me")
}

func (c *computeApp) ListKeypair(arg *entity.KeypairListArg) ([]*entity.Keypair, error) {
	//TODO implement me
	panic("implement me")
}

func (c *computeApp) CreateSecurityGroup(arg *entity.SecurityGroupCreateArg) (*entity.SecurityGroup, error) {
	//TODO implement me
	panic("implement me")
}

func (c *computeApp) DeleteSecurityGroup(arg *entity.SecurityGroupDeleteArg) (*entity.SecurityGroup, error) {
	//TODO implement me
	panic("implement me")
}

func (c *computeApp) UpdateSecurityGroup(arg *entity.SecurityGroupUpdateArg) (*entity.SecurityGroup, error) {
	//TODO implement me
	panic("implement me")
}

func (c *computeApp) GetSecurityGroup(arg *entity.SecurityGroupGetArg) (*entity.SecurityGroup, error) {
	//TODO implement me
	panic("implement me")
}

func (c *computeApp) ListSecurityGroup(arg *entity.SecurityGroupListArg) ([]*entity.SecurityGroup, error) {
	//TODO implement me
	panic("implement me")
}
