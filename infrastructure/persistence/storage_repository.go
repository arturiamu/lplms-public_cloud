package persistence

import (
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
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

func NewStorageRepo(path *string) *StorageRepo {
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

func (s StorageRepo) SaveDisk(disk *entity.Disk) (*entity.Disk, error) {
	//TODO implement me
	panic("implement me")
}

func (s StorageRepo) GetDisk(u uint64) (*entity.Disk, error) {
	//TODO implement me
	panic("implement me")
}

func (s StorageRepo) ListDisk() (*entity.Disk, error) {
	//TODO implement me
	panic("implement me")
}

func (s StorageRepo) GetDiskBy(u uint64) (*entity.Disk, error) {
	//TODO implement me
	panic("implement me")
}

func (s StorageRepo) SaveSnapshot(disk *entity.Disk) (*entity.Disk, error) {
	//TODO implement me
	panic("implement me")
}

func (s StorageRepo) GetSnapshot(u uint64) (*entity.Disk, error) {
	//TODO implement me
	panic("implement me")
}

func (s StorageRepo) ListSnapshot() (*entity.Disk, error) {
	//TODO implement me
	panic("implement me")
}

func (s StorageRepo) GetSnapshotBy(u uint64) (*entity.Disk, error) {
	//TODO implement me
	panic("implement me")
}
