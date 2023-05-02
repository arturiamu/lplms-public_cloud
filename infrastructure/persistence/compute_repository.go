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

func (c *ComputeRepo) CreateKeypair(arg *entity.KeypairCreateArg) (*entity.KeypairCreateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) DeleteKeypair(arg *entity.KeypairDeleteArg) (*entity.KeypairDeleteResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) UpdateKeypair(arg *entity.KeypairUpdateArg) (*entity.KeypairUpdateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) GetKeypair(arg *entity.KeypairGetArg) (*entity.KeypairGetResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) ListKeypair(arg *entity.KeypairListArg) (*entity.KeypairListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) CreateSecurityGroup(arg *entity.SecurityGroupCreateArg) (*entity.SecurityGroupCreateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) DeleteSecurityGroup(arg *entity.SecurityGroupDeleteArg) (*entity.SecurityGroupDeleteResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) UpdateSecurityGroup(arg *entity.SecurityGroupUpdateArg) (*entity.SecurityGroupUpdateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) GetSecurityGroup(arg *entity.SecurityGroupGetArg) (*entity.SecurityGroupGetResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) ListSecurityGroup(arg *entity.SecurityGroupListArg) (*entity.SecurityGroupListResp, error) {
	//TODO implement me
	panic("implement me")
}
