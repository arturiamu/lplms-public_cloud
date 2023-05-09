package mock

import "github.com/arturiamu/lplms-public_cloud/domain/entity"

func (b BizMock) CreateEip(args *entity.EipCreateArg) (*entity.EipCreateResp, error) {
	return nil, nil
}

func (b BizMock) DeleteEip(args *entity.EipDeleteArg) (*entity.EipDeleteResp, error) {
	return nil, nil
}

func (b BizMock) UpdateEip(args *entity.EipUpdateArg) (*entity.EipUpdateResp, error) {
	return nil, nil
}

func (b BizMock) GetEip(args *entity.EipGetArg) (*entity.EipGetResp, error) {
	return &entity.EipGetResp{
		Eip: getEip(args.EIPID),
	}, nil
}

func (b BizMock) ListEip(args *entity.EipListArg) (*entity.EipListResp, error) {
	return &entity.EipListResp{
		EIPs: getEipList(),
	}, nil
}

func (b BizMock) CreateVpc(args *entity.VpcCreateArg) (*entity.VpcCreateResp, error) {
	return nil, nil
}

func (b BizMock) DeleteVpc(args *entity.VpcDeleteArg) (*entity.VpcDeleteResp, error) {
	return nil, nil
}

func (b BizMock) UpdateVpc(args *entity.VpcUpdateArg) (*entity.VpcUpdateResp, error) {
	return nil, nil
}

func (b BizMock) GetVpc(args *entity.VpcGetArg) (*entity.VpcGetResp, error) {
	return &entity.VpcGetResp{
		VPC: getVpc(args.VPCID),
	}, nil
}

func (b BizMock) ListVpc(args *entity.VpcListArg) (*entity.VpcListResp, error) {
	return &entity.VpcListResp{
		VPCs: getVpcList(),
	}, nil
}

func (b BizMock) CreateVSwitch(args *entity.VSwitchCreateArg) (*entity.VSwitchCreateResp, error) {
	return nil, nil
}

func (b BizMock) DeleteVSwitch(args *entity.VSwitchDeleteArg) (*entity.VSwitchDeleteResp, error) {
	return nil, nil
}

func (b BizMock) UpdateVSwitch(args *entity.VSwitchUpdateArg) (*entity.VSwitchUpdateResp, error) {
	return nil, nil
}

func (b BizMock) GetVSwitch(args *entity.VSwitchGetArg) (*entity.VSwitchGetResp, error) {
	return &entity.VSwitchGetResp{
		VSwitch: getVSwitch(args.VSwitchID),
	}, nil
}

func (b BizMock) ListVSwitch(args *entity.VSwitchListArg) (*entity.VSwitchListResp, error) {
	return &entity.VSwitchListResp{
		VSwitches: getVSwitchList(),
	}, nil
}

func (b BizMock) CreateSlb(args *entity.SlbCreateArg) (*entity.SlbCreateResp, error) {
	return nil, nil
}

func (b BizMock) DeleteSlb(args *entity.SlbDeleteArg) (*entity.SlbDeleteResp, error) {
	return nil, nil
}

func (b BizMock) UpdateSlb(args *entity.SlbUpdateArg) (*entity.SlbUpdateResp, error) {
	return nil, nil
}

func (b BizMock) GetSlb(args *entity.SlbGetArg) (*entity.SlbGetResp, error) {
	return nil, nil
}

func (b BizMock) ListSlb(args *entity.SlbListArg) (*entity.SlbListResp, error) {
	return nil, nil
}
