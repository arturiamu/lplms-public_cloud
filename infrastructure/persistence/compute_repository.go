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
	var computeRepo = ComputeRepo{
		k8Virt: kubevirtClient,
		ovn:    ovnClient,
	}
	stk.C = &computeRepo
	return &computeRepo
}

func (c *ComputeRepo) CreateServer(arg *entity.ServerCreateArg) (*entity.Server, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) DeleteServer(arg *entity.ServerDeleteArg) (*entity.Server, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) UpdateServer(arg *entity.ServerUpdateArg) (*entity.Server, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) GetServer(arg *entity.ServerGetArg) (*entity.Server, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) ListServer(arg *entity.ServerListArg) ([]*entity.Server, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) CreateFlavor(arg *entity.FlavorCreateArg) (*entity.Flavor, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) DeleteFlavor(arg *entity.FlavorDeleteArg) (*entity.Flavor, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) UpdateFlavor(arg *entity.FlavorUpdateArg) (*entity.Flavor, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) GetFlavor(arg *entity.FlavorGetArg) (*entity.Flavor, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) ListFlavor(arg *entity.FlavorListArg) ([]*entity.Flavor, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) CreateImage(arg *entity.ImageCreateArg) (*entity.Image, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) DeleteImage(arg *entity.ImageDeleteArg) (*entity.Image, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) UpdateImage(arg *entity.ImageUpdateArg) (*entity.Image, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) GetImage(arg *entity.ImageGetArg) (*entity.Image, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) ListImage(arg *entity.ImageListArg) ([]*entity.Image, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) CreateKeypair(arg *entity.KeypairCreateArg) (*entity.Keypair, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) DeleteKeypair(arg *entity.KeypairDeleteArg) (*entity.Keypair, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) UpdateKeypair(arg *entity.KeypairUpdateArg) (*entity.Keypair, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) GetKeypair(arg *entity.KeypairGetArg) (*entity.Keypair, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) ListKeypair(arg *entity.KeypairListArg) ([]*entity.Keypair, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) CreateSecurityGroup(arg *entity.SecurityGroupCreateArg) (*entity.SecurityGroup, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) DeleteSecurityGroup(arg *entity.SecurityGroupDeleteArg) (*entity.SecurityGroup, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) UpdateSecurityGroup(arg *entity.SecurityGroupUpdateArg) (*entity.SecurityGroup, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) GetSecurityGroup(arg *entity.SecurityGroupGetArg) (*entity.SecurityGroup, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) ListSecurityGroup(arg *entity.SecurityGroupListArg) ([]*entity.SecurityGroup, error) {
	//TODO implement me
	panic("implement me")
}
