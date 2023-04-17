package persistence

import (
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/arturiamu/lplms-public_cloud/domain/repository"
	"github.com/kubeovn/kube-ovn/pkg/client/clientset/versioned"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/tools/clientcmd"
	"kubevirt.io/client-go/kubecli"
)

var _ repository.ComputeRepositoryInterface = &ComputeRepo{}

type ComputeRepo struct {
	k8Virt kubecli.KubevirtClient
	ovn    *versioned.Clientset
}

func NewComputeRepo(path *string) *ComputeRepo {
	config, err := clientcmd.BuildConfigFromFlags("", *path)
	if err != nil {
		panic(err)
	}
	config.NegotiatedSerializer = serializer.WithoutConversionCodecFactory{CodecFactory: serializer.NewCodecFactory(runtime.NewScheme())}
	kubevirtClient, err := kubecli.GetKubevirtClientFromRESTConfig(config)
	ovnClient, err := versioned.NewForConfig(config)
	return &ComputeRepo{
		k8Virt: kubevirtClient,
		ovn:    ovnClient,
	}
}

func (c *ComputeRepo) SaveServer(server *entity.Server) (*entity.Server, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) GetServer(u uint64) (*entity.Server, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) ListServer() (*entity.Server, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) GetServerBy(u uint64) (*entity.Server, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) SaveImage(image *entity.Image) (*entity.Image, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) GetImage(u uint64) (*entity.Image, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) ListImage() (*entity.Image, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) GetImageBy(u uint64) (*entity.Image, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) SaveFlavor(flavor *entity.Flavor) (*entity.Flavor, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) GetFlavor(u uint64) (*entity.Flavor, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) ListFlavor() (*entity.Flavor, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) GetFlavorBy(u uint64) (*entity.Flavor, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) SaveKeyPair(pair *entity.KeyPair) (*entity.KeyPair, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) GetKeyPair(u uint64) (*entity.KeyPair, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) ListKeyPair() (*entity.KeyPair, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) GetKeyPairBy(u uint64) (*entity.KeyPair, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) SaveSecurityGroup(group *entity.SecurityGroup) (*entity.SecurityGroup, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) GetSecurityGroup(u uint64) (*entity.SecurityGroup, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) ListSecurityGroup() (*entity.SecurityGroup, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) GetSecurityGroupBy(u uint64) (*entity.SecurityGroup, error) {
	//TODO implement me
	panic("implement me")
}
