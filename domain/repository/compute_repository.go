package repository

import "github.com/arturiamu/lplms-public_cloud/domain/entity"

type ComputeRepositoryInterface interface {
	ServerRepositoryInterface
	ImageRepositoryInterface
	FlavorRepositoryInterface
	KeypairRepositoryInterface
	SecurityGroupRepositoryInterface
}

type ServerRepositoryInterface interface {
	CreateServer(arg *entity.ServerCreateArg) (*entity.Server, error)
	DeleteServer(arg *entity.ServerDeleteArg) (*entity.Server, error)
	UpdateServer(arg *entity.ServerUpdateArg) (*entity.Server, error)
	GetServer(arg *entity.ServerGetArg) (*entity.Server, error)
	ListServer(arg *entity.ServerListArg) ([]*entity.Server, error)
}

type ImageRepositoryInterface interface {
	CreateFlavor(arg *entity.FlavorCreateArg) (*entity.Flavor, error)
	DeleteFlavor(arg *entity.FlavorDeleteArg) (*entity.Flavor, error)
	UpdateFlavor(arg *entity.FlavorUpdateArg) (*entity.Flavor, error)
	GetFlavor(arg *entity.FlavorGetArg) (*entity.Flavor, error)
	ListFlavor(arg *entity.FlavorListArg) ([]*entity.Flavor, error)
}

type FlavorRepositoryInterface interface {
	CreateImage(arg *entity.ImageCreateArg) (*entity.Image, error)
	DeleteImage(arg *entity.ImageDeleteArg) (*entity.Image, error)
	UpdateImage(arg *entity.ImageUpdateArg) (*entity.Image, error)
	GetImage(arg *entity.ImageGetArg) (*entity.Image, error)
	ListImage(arg *entity.ImageListArg) ([]*entity.Image, error)
}

type KeypairRepositoryInterface interface {
	CreateKeypair(arg *entity.KeypairCreateArg) (*entity.Keypair, error)
	DeleteKeypair(arg *entity.KeypairDeleteArg) (*entity.Keypair, error)
	UpdateKeypair(arg *entity.KeypairUpdateArg) (*entity.Keypair, error)
	GetKeypair(arg *entity.KeypairGetArg) (*entity.Keypair, error)
	ListKeypair(arg *entity.KeypairListArg) ([]*entity.Keypair, error)
}

type SecurityGroupRepositoryInterface interface {
	CreateSecurityGroup(arg *entity.SecurityGroupCreateArg) (*entity.SecurityGroup, error)
	DeleteSecurityGroup(arg *entity.SecurityGroupDeleteArg) (*entity.SecurityGroup, error)
	UpdateSecurityGroup(arg *entity.SecurityGroupUpdateArg) (*entity.SecurityGroup, error)
	GetSecurityGroup(arg *entity.SecurityGroupGetArg) (*entity.SecurityGroup, error)
	ListSecurityGroup(arg *entity.SecurityGroupListArg) ([]*entity.SecurityGroup, error)
}
