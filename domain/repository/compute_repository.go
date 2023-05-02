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
	CreateServer(arg *entity.ServerCreateArg) (*entity.ServerCreateResp, error)
	DeleteServer(arg *entity.ServerDeleteArg) (*entity.ServerDeleteResp, error)
	UpdateServer(arg *entity.ServerUpdateArg) (*entity.ServerUpdateResp, error)
	GetServer(arg *entity.ServerGetArg) (*entity.ServerGetResp, error)
	ListServer(arg *entity.ServerListArg) (*entity.ServerListResp, error)
}

type ImageRepositoryInterface interface {
	CreateFlavor(arg *entity.FlavorCreateArg) (*entity.FlavorCreateResp, error)
	DeleteFlavor(arg *entity.FlavorDeleteArg) (*entity.FlavorDeleteResp, error)
	UpdateFlavor(arg *entity.FlavorUpdateArg) (*entity.FlavorUpdateResp, error)
	GetFlavor(arg *entity.FlavorGetArg) (*entity.FlavorGetResp, error)
	ListFlavor(arg *entity.FlavorListArg) (*entity.FlavorListResp, error)
}

type FlavorRepositoryInterface interface {
	CreateImage(arg *entity.ImageCreateArg) (*entity.ImageCreateResp, error)
	DeleteImage(arg *entity.ImageDeleteArg) (*entity.ImageDeleteResp, error)
	UpdateImage(arg *entity.ImageUpdateArg) (*entity.ImageUpdateResp, error)
	GetImage(arg *entity.ImageGetArg) (*entity.ImageGetResp, error)
	ListImage(arg *entity.ImageListArg) (*entity.ImageListResp, error)
}

type KeypairRepositoryInterface interface {
	CreateKeypair(arg *entity.KeypairCreateArg) (*entity.KeypairCreateResp, error)
	DeleteKeypair(arg *entity.KeypairDeleteArg) (*entity.KeypairDeleteResp, error)
	UpdateKeypair(arg *entity.KeypairUpdateArg) (*entity.KeypairUpdateResp, error)
	GetKeypair(arg *entity.KeypairGetArg) (*entity.KeypairGetResp, error)
	ListKeypair(arg *entity.KeypairListArg) (*entity.KeypairListResp, error)
}

type SecurityGroupRepositoryInterface interface {
	CreateSecurityGroup(arg *entity.SecurityGroupCreateArg) (*entity.SecurityGroupCreateResp, error)
	DeleteSecurityGroup(arg *entity.SecurityGroupDeleteArg) (*entity.SecurityGroupDeleteResp, error)
	UpdateSecurityGroup(arg *entity.SecurityGroupUpdateArg) (*entity.SecurityGroupUpdateResp, error)
	GetSecurityGroup(arg *entity.SecurityGroupGetArg) (*entity.SecurityGroupGetResp, error)
	ListSecurityGroup(arg *entity.SecurityGroupListArg) (*entity.SecurityGroupListResp, error)
}
