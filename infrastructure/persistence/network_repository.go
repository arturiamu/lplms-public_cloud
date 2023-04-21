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

func (n NetworkRepo) SaveEip(eip *entity.Eip) (*entity.Eip, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetworkRepo) GetEip(u uint64) (*entity.Eip, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetworkRepo) ListEip() (*entity.Eip, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetworkRepo) GetEipBy(u uint64) (*entity.Eip, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetworkRepo) SaveVpc(vpc *entity.Vpc) (*entity.Vpc, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetworkRepo) GetVpc(u uint64) (*entity.Vpc, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetworkRepo) ListVpc() (*entity.Vpc, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetworkRepo) GetVpcBy(u uint64) (*entity.Vpc, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetworkRepo) SaveVSwitch(vSwitch *entity.VSwitch) (*entity.VSwitch, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetworkRepo) GetVSwitch(u uint64) (*entity.VSwitch, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetworkRepo) ListVSwitch() (*entity.VSwitch, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetworkRepo) GetVSwitchBy(u uint64) (*entity.VSwitch, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetworkRepo) SaveSlb(slb *entity.Slb) (*entity.Slb, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetworkRepo) GetSlb(u uint64) (*entity.Slb, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetworkRepo) ListSlb() (*entity.Slb, error) {
	//TODO implement me
	panic("implement me")
}

func (n NetworkRepo) GetSlbBy(u uint64) (*entity.Slb, error) {
	//TODO implement me
	panic("implement me")
}
