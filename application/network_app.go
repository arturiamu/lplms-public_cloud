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
	CreateEip(arg *entity.EipCreateArg) (*entity.Eip, error)
	DeleteEip(arg *entity.EipDeleteArg) (*entity.Eip, error)
	UpdateEip(arg *entity.EipUpdateArg) (*entity.Eip, error)
	GetEip(arg *entity.EipGetArg) (*entity.Eip, error)
	ListEip(arg *entity.EipListArg) ([]*entity.Eip, error)
}

type VPCInterface interface {
	CreateVpc(arg *entity.VpcCreateArg) (*entity.Vpc, error)
	DeleteVpc(arg *entity.VpcDeleteArg) (*entity.Vpc, error)
	UpdateVpc(arg *entity.VpcUpdateArg) (*entity.Vpc, error)
	GetVpc(arg *entity.VpcGetArg) (*entity.Vpc, error)
	ListVpc(arg *entity.VpcListArg) ([]*entity.Vpc, error)
}

type VSwitchInterface interface {
	CreateVSwitch(arg *entity.VSwitchCreateArg) (*entity.VSwitch, error)
	DeleteVSwitch(arg *entity.VSwitchDeleteArg) (*entity.VSwitch, error)
	UpdateVSwitch(arg *entity.VSwitchUpdateArg) (*entity.VSwitch, error)
	GetVSwitch(arg *entity.VSwitchGetArg) (*entity.VSwitch, error)
	ListVSwitch(arg *entity.VSwitchListArg) ([]*entity.VSwitch, error)
}

type SLBInterface interface {
	CreateSlb(arg *entity.SlbCreateArg) (*entity.Slb, error)
	DeleteSlb(arg *entity.SlbDeleteArg) (*entity.Slb, error)
	UpdateSlb(arg *entity.SlbUpdateArg) (*entity.Slb, error)
	GetSlb(arg *entity.SlbGetArg) (*entity.Slb, error)
	ListSlb(arg *entity.SlbListArg) ([]*entity.Slb, error)
}

type networkApp struct {
	nr repository.NetworkRepositoryInterface
}

func (n *networkApp) CreateEip(arg *entity.EipCreateArg) (*entity.Eip, error) {
	return n.nr.CreateEip(arg)
}

func (n *networkApp) DeleteEip(arg *entity.EipDeleteArg) (*entity.Eip, error) {
	//TODO implement me
	panic("implement me")
}

func (n *networkApp) UpdateEip(arg *entity.EipUpdateArg) (*entity.Eip, error) {
	//TODO implement me
	panic("implement me")
}

func (n *networkApp) GetEip(arg *entity.EipGetArg) (*entity.Eip, error) {
	//TODO implement me
	panic("implement me")
}

func (n *networkApp) ListEip(arg *entity.EipListArg) ([]*entity.Eip, error) {
	//TODO implement me
	panic("implement me")
}

func (n *networkApp) CreateVpc(arg *entity.VpcCreateArg) (*entity.Vpc, error) {
	//TODO implement me
	panic("implement me")
}

func (n *networkApp) DeleteVpc(arg *entity.VpcDeleteArg) (*entity.Vpc, error) {
	//TODO implement me
	panic("implement me")
}

func (n *networkApp) UpdateVpc(arg *entity.VpcUpdateArg) (*entity.Vpc, error) {
	//TODO implement me
	panic("implement me")
}

func (n *networkApp) GetVpc(arg *entity.VpcGetArg) (*entity.Vpc, error) {
	//TODO implement me
	panic("implement me")
}

func (n *networkApp) ListVpc(arg *entity.VpcListArg) ([]*entity.Vpc, error) {
	//TODO implement me
	panic("implement me")
}

func (n *networkApp) CreateVSwitch(arg *entity.VSwitchCreateArg) (*entity.VSwitch, error) {
	//TODO implement me
	panic("implement me")
}

func (n *networkApp) DeleteVSwitch(arg *entity.VSwitchDeleteArg) (*entity.VSwitch, error) {
	//TODO implement me
	panic("implement me")
}

func (n *networkApp) UpdateVSwitch(arg *entity.VSwitchUpdateArg) (*entity.VSwitch, error) {
	//TODO implement me
	panic("implement me")
}

func (n *networkApp) GetVSwitch(arg *entity.VSwitchGetArg) (*entity.VSwitch, error) {
	//TODO implement me
	panic("implement me")
}

func (n *networkApp) ListVSwitch(arg *entity.VSwitchListArg) ([]*entity.VSwitch, error) {
	//TODO implement me
	panic("implement me")
}

func (n *networkApp) CreateSlb(arg *entity.SlbCreateArg) (*entity.Slb, error) {
	//TODO implement me
	panic("implement me")
}

func (n *networkApp) DeleteSlb(arg *entity.SlbDeleteArg) (*entity.Slb, error) {
	//TODO implement me
	panic("implement me")
}

func (n *networkApp) UpdateSlb(arg *entity.SlbUpdateArg) (*entity.Slb, error) {
	//TODO implement me
	panic("implement me")
}

func (n *networkApp) GetSlb(arg *entity.SlbGetArg) (*entity.Slb, error) {
	//TODO implement me
	panic("implement me")
}

func (n *networkApp) ListSlb(arg *entity.SlbListArg) ([]*entity.Slb, error) {
	//TODO implement me
	panic("implement me")
}
