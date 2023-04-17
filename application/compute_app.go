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
	SaveServer(*entity.Server) (*entity.Server, error)
	GetServer(uint64) (*entity.Server, error)
	ListServer() (*entity.Server, error)
	GetServerBy(uint64) (*entity.Server, error)
}

type FlavorAppInterface interface {
	SaveFlavor(*entity.Flavor) (*entity.Flavor, error)
	GetFlavor(uint64) (*entity.Flavor, error)
	ListFlavor() (*entity.Flavor, error)
	GetFlavorBy(uint64) (*entity.Flavor, error)
}

type ImageAppInterface interface {
	SaveImage(*entity.Image) (*entity.Image, error)
	GetImage(uint64) (*entity.Image, error)
	ListImage() (*entity.Image, error)
	GetImageBy(uint64) (*entity.Image, error)
}

type KeyPairAppInterface interface {
	SaveKeyPair(*entity.KeyPair) (*entity.KeyPair, error)
	GetKeyPair(uint64) (*entity.KeyPair, error)
	ListKeyPair() (*entity.KeyPair, error)
	GetKeyPairBy(uint64) (*entity.KeyPair, error)
}

type SecurityGroupAppInterface interface {
	SaveSecurityGroup(*entity.SecurityGroup) (*entity.SecurityGroup, error)
	GetSecurityGroup(uint64) (*entity.SecurityGroup, error)
	ListSecurityGroup() (*entity.SecurityGroup, error)
	GetSecurityGroupBy(uint64) (*entity.SecurityGroup, error)
}

type computeApp struct {
	cr repository.ComputeRepositoryInterface
}

func (c computeApp) SaveServer(server *entity.Server) (*entity.Server, error) {
	return c.cr.SaveServer(server)
}

func (c computeApp) GetServer(u uint64) (*entity.Server, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) ListServer() (*entity.Server, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) GetServerBy(u uint64) (*entity.Server, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) SaveFlavor(flavor *entity.Flavor) (*entity.Flavor, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) GetFlavor(u uint64) (*entity.Flavor, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) ListFlavor() (*entity.Flavor, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) GetFlavorBy(u uint64) (*entity.Flavor, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) SaveImage(image *entity.Image) (*entity.Image, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) GetImage(u uint64) (*entity.Image, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) ListImage() (*entity.Image, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) GetImageBy(u uint64) (*entity.Image, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) SaveKeyPair(pair *entity.KeyPair) (*entity.KeyPair, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) GetKeyPair(u uint64) (*entity.KeyPair, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) ListKeyPair() (*entity.KeyPair, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) GetKeyPairBy(u uint64) (*entity.KeyPair, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) SaveSecurityGroup(group *entity.SecurityGroup) (*entity.SecurityGroup, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) GetSecurityGroup(u uint64) (*entity.SecurityGroup, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) ListSecurityGroup() (*entity.SecurityGroup, error) {
	//TODO implement me
	panic("implement me")
}

func (c computeApp) GetSecurityGroupBy(u uint64) (*entity.SecurityGroup, error) {
	//TODO implement me
	panic("implement me")
}
