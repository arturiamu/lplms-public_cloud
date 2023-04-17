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
	SaveServer(*entity.Server) (*entity.Server, error)
	GetServer(uint64) (*entity.Server, error)
	ListServer() (*entity.Server, error)
	GetServerBy(uint64) (*entity.Server, error)
}

type ImageRepositoryInterface interface {
	SaveImage(*entity.Image) (*entity.Image, error)
	GetImage(uint64) (*entity.Image, error)
	ListImage() (*entity.Image, error)
	GetImageBy(uint64) (*entity.Image, error)
}

type FlavorRepositoryInterface interface {
	SaveFlavor(*entity.Flavor) (*entity.Flavor, error)
	GetFlavor(uint64) (*entity.Flavor, error)
	ListFlavor() (*entity.Flavor, error)
	GetFlavorBy(uint64) (*entity.Flavor, error)
}

type KeypairRepositoryInterface interface {
	SaveKeyPair(*entity.KeyPair) (*entity.KeyPair, error)
	GetKeyPair(uint64) (*entity.KeyPair, error)
	ListKeyPair() (*entity.KeyPair, error)
	GetKeyPairBy(uint64) (*entity.KeyPair, error)
}

type SecurityGroupRepositoryInterface interface {
	SaveSecurityGroup(*entity.SecurityGroup) (*entity.SecurityGroup, error)
	GetSecurityGroup(uint64) (*entity.SecurityGroup, error)
	ListSecurityGroup() (*entity.SecurityGroup, error)
	GetSecurityGroupBy(uint64) (*entity.SecurityGroup, error)
}
