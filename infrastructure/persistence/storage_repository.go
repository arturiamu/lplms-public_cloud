package persistence

import (
	"github.com/arturiamu/lplms-public_cloud/domain/repository"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/tools/clientcmd"
	"kubevirt.io/client-go/kubecli"
)

var _ repository.StorageRepositoryInterface = &StorageRepo{}

type StorageRepo struct {
	k8Virt kubecli.KubevirtClient
}

func NewStorageRepo(path *string) repository.StorageRepositoryInterface {
	config, err := clientcmd.BuildConfigFromFlags("", *path)
	if err != nil {
		panic(err)
	}
	config.NegotiatedSerializer = serializer.WithoutConversionCodecFactory{CodecFactory: serializer.NewCodecFactory(runtime.NewScheme())}
	kubevirtClient, err := kubecli.GetKubevirtClientFromRESTConfig(config)
	var storageRepo = StorageRepo{
		k8Virt: kubevirtClient,
	}
	stk.S = &storageRepo
	return &storageRepo
}
