package repository

import "github.com/arturiamu/lplms-public_cloud/domain/entity"

type NetworkRepositoryInterface interface {
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
