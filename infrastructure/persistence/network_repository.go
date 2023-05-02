package persistence

import (
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/arturiamu/lplms-public_cloud/domain/repository"
	"github.com/kubeovn/kube-ovn/pkg/client/clientset/versioned"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/tools/clientcmd"
)

var _ repository.NetworkRepositoryInterface = &NetworkRepo{}

type NetworkRepo struct {
	ovn *versioned.Clientset
}

func NewNetworkRepo(path *string) *NetworkRepo {
	config, err := clientcmd.BuildConfigFromFlags("", *path)
	if err != nil {
		panic(err)
	}
	config.NegotiatedSerializer = serializer.WithoutConversionCodecFactory{CodecFactory: serializer.NewCodecFactory(runtime.NewScheme())}
	ovnClient, err := versioned.NewForConfig(config)
	var networkRepo = NetworkRepo{
		ovn: ovnClient,
	}
	stk.N = &networkRepo
	return &networkRepo
}

func (n *NetworkRepo) CreateEip(arg *entity.EipCreateArg) (*entity.EipCreateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) DeleteEip(arg *entity.EipDeleteArg) (*entity.EipDeleteResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) UpdateEip(arg *entity.EipUpdateArg) (*entity.EipUpdateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) GetEip(arg *entity.EipGetArg) (*entity.EipGetResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) ListEip(arg *entity.EipListArg) (*entity.EipListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) CreateVpc(arg *entity.VpcCreateArg) (*entity.VpcCreateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) DeleteVpc(arg *entity.VpcDeleteArg) (*entity.VpcDeleteResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) UpdateVpc(arg *entity.VpcUpdateArg) (*entity.VpcUpdateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) GetVpc(arg *entity.VpcGetArg) (*entity.VpcGetResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) ListVpc(arg *entity.VpcListArg) (*entity.VpcListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) CreateVSwitch(arg *entity.VSwitchCreateArg) (*entity.VSwitchCreateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) DeleteVSwitch(arg *entity.VSwitchDeleteArg) (*entity.VSwitchDeleteResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) UpdateVSwitch(arg *entity.VSwitchUpdateArg) (*entity.VSwitchUpdateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) GetVSwitch(arg *entity.VSwitchGetArg) (*entity.VSwitchGetResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) ListVSwitch(arg *entity.VSwitchListArg) (*entity.VSwitchListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) CreateSlb(arg *entity.SlbCreateArg) (*entity.SlbCreateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) DeleteSlb(arg *entity.SlbDeleteArg) (*entity.SlbDeleteResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) UpdateSlb(arg *entity.SlbUpdateArg) (*entity.SlbUpdateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) GetSlb(arg *entity.SlbGetArg) (*entity.SlbGetResp, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) ListSlb(arg *entity.SlbListArg) (*entity.SlbListResp, error) {
	//TODO implement me
	panic("implement me")
}
