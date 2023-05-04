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
	CreateServer(arg *entity.ServerCreateArg) (*entity.ServerCreateResp, error)
	DeleteServer(arg *entity.ServerDeleteArg) (*entity.ServerDeleteResp, error)
	UpdateServer(arg *entity.ServerUpdateArg) (*entity.ServerUpdateResp, error)
	GetServer(arg *entity.ServerGetArg) (*entity.ServerGetResp, error)
	ListServer(arg *entity.ServerListArg) (*entity.ServerListResp, error)
}

type FlavorAppInterface interface {
	CreateFlavor(arg *entity.FlavorCreateArg) (*entity.FlavorCreateResp, error)
	DeleteFlavor(arg *entity.FlavorDeleteArg) (*entity.FlavorDeleteResp, error)
	UpdateFlavor(arg *entity.FlavorUpdateArg) (*entity.FlavorUpdateResp, error)
	GetFlavor(arg *entity.FlavorGetArg) (*entity.FlavorGetResp, error)
	ListFlavor(arg *entity.FlavorListArg) (*entity.FlavorListResp, error)
}

type ImageAppInterface interface {
	CreateImage(arg *entity.ImageCreateArg) (*entity.ImageCreateResp, error)
	DeleteImage(arg *entity.ImageDeleteArg) (*entity.ImageDeleteResp, error)
	UpdateImage(arg *entity.ImageUpdateArg) (*entity.ImageUpdateResp, error)
	GetImage(arg *entity.ImageGetArg) (*entity.ImageGetResp, error)
	ListImage(arg *entity.ImageListArg) (*entity.ImageListResp, error)
}

type KeyPairAppInterface interface {
	CreateKeypair(arg *entity.KeypairCreateArg) (*entity.KeypairCreateResp, error)
	DeleteKeypair(arg *entity.KeypairDeleteArg) (*entity.KeypairDeleteResp, error)
	UpdateKeypair(arg *entity.KeypairUpdateArg) (*entity.KeypairUpdateResp, error)
	GetKeypair(arg *entity.KeypairGetArg) (*entity.KeypairGetResp, error)
	ListKeypair(arg *entity.KeypairListArg) (*entity.KeypairListResp, error)
	DetachKeyPair(args *entity.KeypairDetachArg) (*entity.KeypairDetachResp, error)
	AttachKeyPair(arg *entity.KeypairAttachArg) (*entity.KeypairAttachResp, error)
}

type SecurityGroupAppInterface interface {
	CreateSecurityGroup(arg *entity.SecurityGroupCreateArg) (*entity.SecurityGroupCreateResp, error)
	DeleteSecurityGroup(arg *entity.SecurityGroupDeleteArg) (*entity.SecurityGroupDeleteResp, error)
	UpdateSecurityGroup(arg *entity.SecurityGroupUpdateArg) (*entity.SecurityGroupUpdateResp, error)
	GetSecurityGroup(arg *entity.SecurityGroupGetArg) (*entity.SecurityGroupGetResp, error)
	ListSecurityGroup(arg *entity.SecurityGroupListArg) (*entity.SecurityGroupListResp, error)
}

type computeApp struct {
	cr repository.ComputeRepositoryInterface
}

func (c computeApp) CreateServer(arg *entity.ServerCreateArg) (*entity.ServerCreateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) DeleteServer(arg *entity.ServerDeleteArg) (*entity.ServerDeleteResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) UpdateServer(arg *entity.ServerUpdateArg) (*entity.ServerUpdateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) GetServer(arg *entity.ServerGetArg) (*entity.ServerGetResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) ListServer(arg *entity.ServerListArg) (*entity.ServerListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) CreateFlavor(arg *entity.FlavorCreateArg) (*entity.FlavorCreateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) DeleteFlavor(arg *entity.FlavorDeleteArg) (*entity.FlavorDeleteResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) UpdateFlavor(arg *entity.FlavorUpdateArg) (*entity.FlavorUpdateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) GetFlavor(arg *entity.FlavorGetArg) (*entity.FlavorGetResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) ListFlavor(arg *entity.FlavorListArg) (*entity.FlavorListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) CreateImage(arg *entity.ImageCreateArg) (*entity.ImageCreateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) DeleteImage(arg *entity.ImageDeleteArg) (*entity.ImageDeleteResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) UpdateImage(arg *entity.ImageUpdateArg) (*entity.ImageUpdateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) GetImage(arg *entity.ImageGetArg) (*entity.ImageGetResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) ListImage(arg *entity.ImageListArg) (*entity.ImageListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) CreateKeypair(arg *entity.KeypairCreateArg) (*entity.KeypairCreateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) DeleteKeypair(arg *entity.KeypairDeleteArg) (*entity.KeypairDeleteResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) UpdateKeypair(arg *entity.KeypairUpdateArg) (*entity.KeypairUpdateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) GetKeypair(arg *entity.KeypairGetArg) (*entity.KeypairGetResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) ListKeypair(arg *entity.KeypairListArg) (*entity.KeypairListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) DetachKeyPair(args *entity.KeypairDetachArg) (*entity.KeypairDetachResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) AttachKeyPair(arg *entity.KeypairAttachArg) (*entity.KeypairAttachResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) CreateSecurityGroup(arg *entity.SecurityGroupCreateArg) (*entity.SecurityGroupCreateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) DeleteSecurityGroup(arg *entity.SecurityGroupDeleteArg) (*entity.SecurityGroupDeleteResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) UpdateSecurityGroup(arg *entity.SecurityGroupUpdateArg) (*entity.SecurityGroupUpdateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) GetSecurityGroup(arg *entity.SecurityGroupGetArg) (*entity.SecurityGroupGetResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) ListSecurityGroup(arg *entity.SecurityGroupListArg) (*entity.SecurityGroupListResp, error) {
	//TODO implement me
	panic("implement me")
}
