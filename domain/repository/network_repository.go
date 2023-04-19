package repository

import "github.com/arturiamu/lplms-public_cloud/domain/entity"

type NetworkRepositoryInterface interface {
	EIPInterface
	VPCInterface
	VSwitchInterface
	SLBInterface
}

type EIPInterface interface {
	SaveEip(*entity.Eip) (*entity.Eip, error)
	GetEip(uint64) (*entity.Eip, error)
	ListEip() (*entity.Eip, error)
	GetEipBy(uint64) (*entity.Eip, error)
}

type VPCInterface interface {
	SaveVpc(*entity.Vpc) (*entity.Vpc, error)
	GetVpc(uint64) (*entity.Vpc, error)
	ListVpc() (*entity.Vpc, error)
	GetVpcBy(uint64) (*entity.Vpc, error)
}

type VSwitchInterface interface {
	SaveVSwitch(*entity.VSwitch) (*entity.VSwitch, error)
	GetVSwitch(uint64) (*entity.VSwitch, error)
	ListVSwitch() (*entity.VSwitch, error)
	GetVSwitchBy(uint64) (*entity.VSwitch, error)
}

type SLBInterface interface {
	SaveSlb(*entity.Slb) (*entity.Slb, error)
	GetSlb(uint64) (*entity.Slb, error)
	ListSlb() (*entity.Slb, error)
	GetSlbBy(uint64) (*entity.Slb, error)
}
