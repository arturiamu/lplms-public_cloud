package application

import (
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/arturiamu/lplms-public_cloud/domain/repository"
)

var _ NetworkAppInterface = &networkApp{}

type NetworkAppInterface interface {
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

type networkApp struct {
	nr repository.NetworkRepositoryInterface
}

func (n networkApp) SaveEip(eip *entity.Eip) (*entity.Eip, error) {
	return n.nr.SaveEip(eip)
}

func (n networkApp) GetEip(u uint64) (*entity.Eip, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) ListEip() (*entity.Eip, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) GetEipBy(u uint64) (*entity.Eip, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) SaveVpc(vpc *entity.Vpc) (*entity.Vpc, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) GetVpc(u uint64) (*entity.Vpc, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) ListVpc() (*entity.Vpc, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) GetVpcBy(u uint64) (*entity.Vpc, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) SaveVSwitch(vSwitch *entity.VSwitch) (*entity.VSwitch, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) GetVSwitch(u uint64) (*entity.VSwitch, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) ListVSwitch() (*entity.VSwitch, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) GetVSwitchBy(u uint64) (*entity.VSwitch, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) SaveSlb(slb *entity.Slb) (*entity.Slb, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) GetSlb(u uint64) (*entity.Slb, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) ListSlb() (*entity.Slb, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) GetSlbBy(u uint64) (*entity.Slb, error) {
	//TODO implement me
	panic("implement me")
}
