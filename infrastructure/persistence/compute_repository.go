package persistence

import (
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

func NewComputeRepo(path *string) repository.ComputeRepositoryInterface {
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
