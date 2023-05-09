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
	CreateEip(args *entity.EipCreateArg) (*entity.EipCreateResp, error)
	DeleteEip(args *entity.EipDeleteArg) (*entity.EipDeleteResp, error)
	UpdateEip(args *entity.EipUpdateArg) (*entity.EipUpdateResp, error)
	GetEip(args *entity.EipGetArg) (*entity.EipGetResp, error)
	ListEip(args *entity.EipListArg) (*entity.EipListResp, error)
}

type VPCInterface interface {
	CreateVpc(args *entity.VpcCreateArg) (*entity.VpcCreateResp, error)
	DeleteVpc(args *entity.VpcDeleteArg) (*entity.VpcDeleteResp, error)
	UpdateVpc(args *entity.VpcUpdateArg) (*entity.VpcUpdateResp, error)
	GetVpc(args *entity.VpcGetArg) (*entity.VpcGetResp, error)
	ListVpc(args *entity.VpcListArg) (*entity.VpcListResp, error)
}

type VSwitchInterface interface {
	CreateVSwitch(args *entity.VSwitchCreateArg) (*entity.VSwitchCreateResp, error)
	DeleteVSwitch(args *entity.VSwitchDeleteArg) (*entity.VSwitchDeleteResp, error)
	UpdateVSwitch(args *entity.VSwitchUpdateArg) (*entity.VSwitchUpdateResp, error)
	GetVSwitch(args *entity.VSwitchGetArg) (*entity.VSwitchGetResp, error)
	ListVSwitch(args *entity.VSwitchListArg) (*entity.VSwitchListResp, error)
}

type SLBInterface interface {
	CreateSlb(args *entity.SlbCreateArg) (*entity.SlbCreateResp, error)
	DeleteSlb(args *entity.SlbDeleteArg) (*entity.SlbDeleteResp, error)
	UpdateSlb(args *entity.SlbUpdateArg) (*entity.SlbUpdateResp, error)
	GetSlb(args *entity.SlbGetArg) (*entity.SlbGetResp, error)
	ListSlb(args *entity.SlbListArg) (*entity.SlbListResp, error)
}

type networkApp struct {
	nr repository.NetworkRepositoryInterface
}

func (n networkApp) CreateEip(args *entity.EipCreateArg) (*entity.EipCreateResp, error) {
	return n.nr.CreateEip(args)
}

func (n networkApp) DeleteEip(args *entity.EipDeleteArg) (*entity.EipDeleteResp, error) {
	return n.nr.DeleteEip(args)
}

func (n networkApp) UpdateEip(args *entity.EipUpdateArg) (*entity.EipUpdateResp, error) {
	return n.nr.UpdateEip(args)
}

func (n networkApp) GetEip(args *entity.EipGetArg) (*entity.EipGetResp, error) {
	return n.nr.GetEip(args)
}

func (n networkApp) ListEip(args *entity.EipListArg) (*entity.EipListResp, error) {
	return n.nr.ListEip(args)
}

func (n networkApp) CreateVpc(args *entity.VpcCreateArg) (*entity.VpcCreateResp, error) {
	return n.nr.CreateVpc(args)
}

func (n networkApp) DeleteVpc(args *entity.VpcDeleteArg) (*entity.VpcDeleteResp, error) {
	return n.nr.DeleteVpc(args)
}

func (n networkApp) UpdateVpc(args *entity.VpcUpdateArg) (*entity.VpcUpdateResp, error) {
	return n.nr.UpdateVpc(args)
}

func (n networkApp) GetVpc(args *entity.VpcGetArg) (*entity.VpcGetResp, error) {
	return n.nr.GetVpc(args)
}

func (n networkApp) ListVpc(args *entity.VpcListArg) (*entity.VpcListResp, error) {
	return n.nr.ListVpc(args)
}

func (n networkApp) CreateVSwitch(args *entity.VSwitchCreateArg) (*entity.VSwitchCreateResp, error) {
	return n.nr.CreateVSwitch(args)
}

func (n networkApp) DeleteVSwitch(args *entity.VSwitchDeleteArg) (*entity.VSwitchDeleteResp, error) {
	return n.nr.DeleteVSwitch(args)
}

func (n networkApp) UpdateVSwitch(args *entity.VSwitchUpdateArg) (*entity.VSwitchUpdateResp, error) {
	return n.nr.UpdateVSwitch(args)
}

func (n networkApp) GetVSwitch(args *entity.VSwitchGetArg) (*entity.VSwitchGetResp, error) {
	return n.nr.GetVSwitch(args)
}

func (n networkApp) ListVSwitch(args *entity.VSwitchListArg) (*entity.VSwitchListResp, error) {
	return n.nr.ListVSwitch(args)
}

func (n networkApp) CreateSlb(args *entity.SlbCreateArg) (*entity.SlbCreateResp, error) {
	return n.nr.CreateSlb(args)
}

func (n networkApp) DeleteSlb(args *entity.SlbDeleteArg) (*entity.SlbDeleteResp, error) {
	return n.nr.DeleteSlb(args)
}

func (n networkApp) UpdateSlb(args *entity.SlbUpdateArg) (*entity.SlbUpdateResp, error) {
	return n.nr.UpdateSlb(args)
}

func (n networkApp) GetSlb(args *entity.SlbGetArg) (*entity.SlbGetResp, error) {
	return n.nr.GetSlb(args)
}

func (n networkApp) ListSlb(args *entity.SlbListArg) (*entity.SlbListResp, error) {
	return n.nr.ListSlb(args)
}
