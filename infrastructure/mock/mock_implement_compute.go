package mock

import "github.com/arturiamu/lplms-public_cloud/domain/entity"

func (b BizMock) CreateServer(args *entity.ServerCreateArg) (*entity.ServerCreateResp, error) {
	return nil, nil
}

func (b BizMock) DeleteServer(args *entity.ServerDeleteArg) (*entity.ServerDeleteResp, error) {
	return nil, nil
}

func (b BizMock) UpdateServer(args *entity.ServerUpdateArg) (*entity.ServerUpdateResp, error) {
	return nil, nil
}

func (b BizMock) GetServer(args *entity.ServerGetArg) (*entity.ServerGetResp, error) {
	return &entity.ServerGetResp{
		Server: getServer(args.ServerID),
	}, nil
}

func (b BizMock) ListServer(args *entity.ServerListArg) (*entity.ServerListResp, error) {
	return &entity.ServerListResp{
		Servers: getServerList(),
	}, nil
}

func (b BizMock) GetServerDisks(args *entity.ServerDisksGetArg) (*entity.ServerDisksGetResp, error) {
	return &entity.ServerDisksGetResp{
		Disks: getServerDisk(args.ServerID),
	}, nil
}

func (b BizMock) StartServer(args *entity.ServerStartArgs) (*entity.ServerStartResp, error) {
	return nil, nil
}

func (b BizMock) StopServer(args *entity.ServerStopArgs) (*entity.ServerStopResp, error) {
	return nil, nil
}

func (b BizMock) RestartServer(args *entity.ServerRestartArgs) (*entity.ServerRestartResp, error) {
	return nil, nil
}

func (b BizMock) CreateFlavor(args *entity.FlavorCreateArg) (*entity.FlavorCreateResp, error) {
	return nil, nil
}

func (b BizMock) DeleteFlavor(args *entity.FlavorDeleteArg) (*entity.FlavorDeleteResp, error) {
	return nil, nil
}

func (b BizMock) UpdateFlavor(args *entity.FlavorUpdateArg) (*entity.FlavorUpdateResp, error) {
	return nil, nil
}

func (b BizMock) GetFlavor(args *entity.FlavorGetArg) (*entity.FlavorGetResp, error) {
	return &entity.FlavorGetResp{
		Flavor: getFlavor(args.FlavorID),
	}, nil
}

func (b BizMock) ListFlavor(args *entity.FlavorListArg) (*entity.FlavorListResp, error) {
	return &entity.FlavorListResp{
		Flavors: getFlavorList(),
	}, nil
}

func (b BizMock) CreateImage(args *entity.ImageCreateArg) (*entity.ImageCreateResp, error) {
	return nil, nil
}

func (b BizMock) DeleteImage(args *entity.ImageDeleteArg) (*entity.ImageDeleteResp, error) {
	return nil, nil
}

func (b BizMock) UpdateImage(args *entity.ImageUpdateArg) (*entity.ImageUpdateResp, error) {
	return nil, nil
}

func (b BizMock) GetImage(args *entity.ImageGetArg) (*entity.ImageGetResp, error) {
	return &entity.ImageGetResp{
		Image: getImage(args.ImageID),
	}, nil
}

func (b BizMock) ListImage(args *entity.ImageListArg) (*entity.ImageListResp, error) {
	return &entity.ImageListResp{
		Images: getImageList(),
	}, nil
}

func (b BizMock) CreateKeypair(args *entity.KeypairCreateArg) (*entity.KeypairCreateResp, error) {
	return nil, nil
}

func (b BizMock) DeleteKeypair(args *entity.KeypairDeleteArg) (*entity.KeypairDeleteResp, error) {
	return nil, nil
}

func (b BizMock) UpdateKeypair(args *entity.KeypairUpdateArg) (*entity.KeypairUpdateResp, error) {
	return nil, nil
}

func (b BizMock) GetKeypair(args *entity.KeypairGetArg) (*entity.KeypairGetResp, error) {
	return &entity.KeypairGetResp{
		Keypair: getKeypair(args.KeyPairID),
	}, nil
}

func (b BizMock) ListKeypair(args *entity.KeypairListArg) (*entity.KeypairListResp, error) {
	return &entity.KeypairListResp{
		KeyPairs: getKeypairList(),
	}, nil
}

func (b BizMock) DetachKeyPair(args *entity.KeypairDetachArg) (*entity.KeypairDetachResp, error) {
	return nil, nil
}

func (b BizMock) AttachKeyPair(args *entity.KeypairAttachArg) (*entity.KeypairAttachResp, error) {
	return nil, nil
}

func (b BizMock) CreateSecurityGroup(args *entity.SecurityGroupCreateArg) (*entity.SecurityGroupCreateResp, error) {
	return nil, nil
}

func (b BizMock) DeleteSecurityGroup(args *entity.SecurityGroupDeleteArg) (*entity.SecurityGroupDeleteResp, error) {
	return nil, nil
}

func (b BizMock) UpdateSecurityGroup(args *entity.SecurityGroupUpdateArg) (*entity.SecurityGroupUpdateResp, error) {
	return nil, nil
}

func (b BizMock) GetSecurityGroup(args *entity.SecurityGroupGetArg) (*entity.SecurityGroupGetResp, error) {
	return &entity.SecurityGroupGetResp{
		SecurityGroup: getSecurityGroup(args.SecurityGroupID),
	}, nil
}

func (b BizMock) ListSecurityGroup(args *entity.SecurityGroupListArg) (*entity.SecurityGroupListResp, error) {
	return &entity.SecurityGroupListResp{
		SecurityGroups: getSecurityGroupList(),
	}, nil
}

func (b BizMock) CreateSecurityGroupRule(args *entity.SecurityGroupRuleCreateArg) (*entity.SecurityGroupRuleCreateResp, error) {
	return nil, nil
}

func (b BizMock) DeleteSecurityGroupRule(args *entity.SecurityGroupRuleDeleteArg) (*entity.SecurityGroupRuleDeleteResp, error) {
	return nil, nil
}

func (b BizMock) GetSecurityGroupRule(args *entity.SecurityGroupRuleGetArg) (*entity.SecurityGroupRuleGetResp, error) {
	return nil, nil
}

func (b BizMock) ListSecurityGroupRule(args *entity.SecurityGroupRuleListArg) (*entity.SecurityGroupRuleListResp, error) {
	return &entity.SecurityGroupRuleListResp{
		SecurityGroupRoles: getSecurityGroupRules(args.SecurityGroupID),
	}, nil
}
