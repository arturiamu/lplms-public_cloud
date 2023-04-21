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

func (n *NetworkRepo) CreateEip(arg *entity.EipCreateArg) (*entity.Eip, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) DeleteEip(arg *entity.EipDeleteArg) (*entity.Eip, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) UpdateEip(arg *entity.EipUpdateArg) (*entity.Eip, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) GetEip(arg *entity.EipGetArg) (*entity.Eip, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) ListEip(arg *entity.EipListArg) ([]*entity.Eip, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) CreateVpc(arg *entity.VpcCreateArg) (*entity.Vpc, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) DeleteVpc(arg *entity.VpcDeleteArg) (*entity.Vpc, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) UpdateVpc(arg *entity.VpcUpdateArg) (*entity.Vpc, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) GetVpc(arg *entity.VpcGetArg) (*entity.Vpc, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) ListVpc(arg *entity.VpcListArg) ([]*entity.Vpc, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) CreateVSwitch(arg *entity.VSwitchCreateArg) (*entity.VSwitch, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) DeleteVSwitch(arg *entity.VSwitchDeleteArg) (*entity.VSwitch, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) UpdateVSwitch(arg *entity.VSwitchUpdateArg) (*entity.VSwitch, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) GetVSwitch(arg *entity.VSwitchGetArg) (*entity.VSwitch, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) ListVSwitch(arg *entity.VSwitchListArg) ([]*entity.VSwitch, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) CreateSlb(arg *entity.SlbCreateArg) (*entity.Slb, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) DeleteSlb(arg *entity.SlbDeleteArg) (*entity.Slb, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) UpdateSlb(arg *entity.SlbUpdateArg) (*entity.Slb, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) GetSlb(arg *entity.SlbGetArg) (*entity.Slb, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepo) ListSlb(arg *entity.SlbListArg) ([]*entity.Slb, error) {
	//TODO implement me
	panic("implement me")
}
