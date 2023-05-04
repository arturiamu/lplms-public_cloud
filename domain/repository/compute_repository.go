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
	CreateServer(args *entity.ServerCreateArg) (*entity.ServerCreateResp, error)
	DeleteServer(args *entity.ServerDeleteArg) (*entity.ServerDeleteResp, error)
	UpdateServer(args *entity.ServerUpdateArg) (*entity.ServerUpdateResp, error)
	GetServer(args *entity.ServerGetArg) (*entity.ServerGetResp, error)
	ListServer(args *entity.ServerListArg) (*entity.ServerListResp, error)
}

type ImageRepositoryInterface interface {
	CreateFlavor(args *entity.FlavorCreateArg) (*entity.FlavorCreateResp, error)
	DeleteFlavor(args *entity.FlavorDeleteArg) (*entity.FlavorDeleteResp, error)
	UpdateFlavor(args *entity.FlavorUpdateArg) (*entity.FlavorUpdateResp, error)
	GetFlavor(args *entity.FlavorGetArg) (*entity.FlavorGetResp, error)
	ListFlavor(args *entity.FlavorListArg) (*entity.FlavorListResp, error)
}

type FlavorRepositoryInterface interface {
	CreateImage(args *entity.ImageCreateArg) (*entity.ImageCreateResp, error)
	DeleteImage(args *entity.ImageDeleteArg) (*entity.ImageDeleteResp, error)
	UpdateImage(args *entity.ImageUpdateArg) (*entity.ImageUpdateResp, error)
	GetImage(args *entity.ImageGetArg) (*entity.ImageGetResp, error)
	ListImage(args *entity.ImageListArg) (*entity.ImageListResp, error)
}

type KeypairRepositoryInterface interface {
	CreateKeypair(args *entity.KeypairCreateArg) (*entity.KeypairCreateResp, error)
	DeleteKeypair(args *entity.KeypairDeleteArg) (*entity.KeypairDeleteResp, error)
	UpdateKeypair(args *entity.KeypairUpdateArg) (*entity.KeypairUpdateResp, error)
	GetKeypair(args *entity.KeypairGetArg) (*entity.KeypairGetResp, error)
	ListKeypair(args *entity.KeypairListArg) (*entity.KeypairListResp, error)
	DetachKeyPair(args *entity.KeypairDetachArg) (*entity.KeypairDetachResp, error)
	AttachKeyPair(args *entity.KeypairAttachArg) (*entity.KeypairAttachResp, error)
}

type SecurityGroupRepositoryInterface interface {
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
