package persistence

import (
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

func NewNetworkRepo(path *string) repository.NetworkRepositoryInterface {
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
