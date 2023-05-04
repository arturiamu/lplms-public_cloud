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
	CreateEip(arg *entity.EipCreateArg) (*entity.EipCreateResp, error)
	DeleteEip(arg *entity.EipDeleteArg) (*entity.EipDeleteResp, error)
	UpdateEip(arg *entity.EipUpdateArg) (*entity.EipUpdateResp, error)
	GetEip(arg *entity.EipGetArg) (*entity.EipGetResp, error)
	ListEip(arg *entity.EipListArg) (*entity.EipListResp, error)
}

type VPCInterface interface {
	CreateVpc(arg *entity.VpcCreateArg) (*entity.VpcCreateResp, error)
	DeleteVpc(arg *entity.VpcDeleteArg) (*entity.VpcDeleteResp, error)
	UpdateVpc(arg *entity.VpcUpdateArg) (*entity.VpcUpdateResp, error)
	GetVpc(arg *entity.VpcGetArg) (*entity.VpcGetResp, error)
	ListVpc(arg *entity.VpcListArg) (*entity.VpcListResp, error)
}

type VSwitchInterface interface {
	CreateVSwitch(arg *entity.VSwitchCreateArg) (*entity.VSwitchCreateResp, error)
	DeleteVSwitch(arg *entity.VSwitchDeleteArg) (*entity.VSwitchDeleteResp, error)
	UpdateVSwitch(arg *entity.VSwitchUpdateArg) (*entity.VSwitchUpdateResp, error)
	GetVSwitch(arg *entity.VSwitchGetArg) (*entity.VSwitchGetResp, error)
	ListVSwitch(arg *entity.VSwitchListArg) (*entity.VSwitchListResp, error)
}

type SLBInterface interface {
	CreateSlb(arg *entity.SlbCreateArg) (*entity.SlbCreateResp, error)
	DeleteSlb(arg *entity.SlbDeleteArg) (*entity.SlbDeleteResp, error)
	UpdateSlb(arg *entity.SlbUpdateArg) (*entity.SlbUpdateResp, error)
	GetSlb(arg *entity.SlbGetArg) (*entity.SlbGetResp, error)
	ListSlb(arg *entity.SlbListArg) (*entity.SlbListResp, error)
}

type networkApp struct {
	nr repository.NetworkRepositoryInterface
}

func (n networkApp) CreateEip(arg *entity.EipCreateArg) (*entity.EipCreateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) DeleteEip(arg *entity.EipDeleteArg) (*entity.EipDeleteResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) UpdateEip(arg *entity.EipUpdateArg) (*entity.EipUpdateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) GetEip(arg *entity.EipGetArg) (*entity.EipGetResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) ListEip(arg *entity.EipListArg) (*entity.EipListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) CreateVpc(arg *entity.VpcCreateArg) (*entity.VpcCreateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) DeleteVpc(arg *entity.VpcDeleteArg) (*entity.VpcDeleteResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) UpdateVpc(arg *entity.VpcUpdateArg) (*entity.VpcUpdateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) GetVpc(arg *entity.VpcGetArg) (*entity.VpcGetResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) ListVpc(arg *entity.VpcListArg) (*entity.VpcListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) CreateVSwitch(arg *entity.VSwitchCreateArg) (*entity.VSwitchCreateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) DeleteVSwitch(arg *entity.VSwitchDeleteArg) (*entity.VSwitchDeleteResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) UpdateVSwitch(arg *entity.VSwitchUpdateArg) (*entity.VSwitchUpdateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) GetVSwitch(arg *entity.VSwitchGetArg) (*entity.VSwitchGetResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) ListVSwitch(arg *entity.VSwitchListArg) (*entity.VSwitchListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) CreateSlb(arg *entity.SlbCreateArg) (*entity.SlbCreateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) DeleteSlb(arg *entity.SlbDeleteArg) (*entity.SlbDeleteResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) UpdateSlb(arg *entity.SlbUpdateArg) (*entity.SlbUpdateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) GetSlb(arg *entity.SlbGetArg) (*entity.SlbGetResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n networkApp) ListSlb(arg *entity.SlbListArg) (*entity.SlbListResp, error) {
	//TODO implement me
	panic("implement me")
}
