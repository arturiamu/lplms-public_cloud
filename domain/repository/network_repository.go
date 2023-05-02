package repository

import "github.com/arturiamu/lplms-public_cloud/domain/entity"

type NetworkRepositoryInterface interface {
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
