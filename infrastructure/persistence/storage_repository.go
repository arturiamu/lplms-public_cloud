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

func (s *StorageRepo) CreateDisk(arg *entity.DiskCreateArg) (*entity.DiskCreateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *StorageRepo) DeleteDisk(arg *entity.DiskDeleteArg) (*entity.DiskDeleteResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *StorageRepo) UpdateDisk(arg *entity.DiskUpdateArg) (*entity.DiskUpdateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *StorageRepo) GetDisk(arg *entity.DiskGetArg) (*entity.DiskGetResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *StorageRepo) ListDisk(arg *entity.DiskListArg) (*entity.DiskListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *StorageRepo) CreateSnapshot(arg *entity.SnapshotCreateArg) (*entity.SnapshotCreateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *StorageRepo) DeleteSnapshot(arg *entity.SnapshotDeleteArg) (*entity.SnapshotDeleteResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *StorageRepo) UpdateSnapshot(arg *entity.SnapshotUpdateArg) (*entity.SnapshotUpdateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *StorageRepo) GetSnapshot(arg *entity.SnapshotGetArg) (*entity.SnapshotGetResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *StorageRepo) ListSnapshot(arg *entity.SnapshotListArg) (*entity.SnapshotListResp, error) {
	//TODO implement me
	panic("implement me")
}
