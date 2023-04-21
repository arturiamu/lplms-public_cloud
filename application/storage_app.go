package application

import (
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/arturiamu/lplms-public_cloud/domain/repository"
)

var _ StorageAppInterface = &storageApp{}

type StorageAppInterface interface {
	DiskRepositoryInterface
	SnapshotRepositoryInterface
}

type DiskRepositoryInterface interface {
	CreateDisk(arg *entity.DiskCreateArg) (*entity.Disk, error)
	DeleteDisk(arg *entity.DiskDeleteArg) (*entity.Disk, error)
	UpdateDisk(arg *entity.DiskUpdateArg) (*entity.Disk, error)
	GetDisk(arg *entity.DiskGetArg) (*entity.Disk, error)
	ListDisk(arg *entity.DiskListArg) ([]*entity.Disk, error)
}

type SnapshotRepositoryInterface interface {
	CreateSnapshot(arg *entity.SnapshotCreateArg) (*entity.Snapshot, error)
	DeleteSnapshot(arg *entity.SnapshotDeleteArg) (*entity.Snapshot, error)
	UpdateSnapshot(arg *entity.SnapshotUpdateArg) (*entity.Snapshot, error)
	GetSnapshot(arg *entity.SnapshotGetArg) (*entity.Snapshot, error)
	ListSnapshot(arg *entity.SnapshotListArg) ([]*entity.Snapshot, error)
}

type storageApp struct {
	sr repository.StorageRepositoryInterface
}

func (s storageApp) CreateDisk(arg *entity.DiskCreateArg) (*entity.Disk, error) {
	//TODO implement me
	panic("implement me")
}

func (s storageApp) DeleteDisk(arg *entity.DiskDeleteArg) (*entity.Disk, error) {
	//TODO implement me
	panic("implement me")
}

func (s storageApp) UpdateDisk(arg *entity.DiskUpdateArg) (*entity.Disk, error) {
	//TODO implement me
	panic("implement me")
}

func (s storageApp) GetDisk(arg *entity.DiskGetArg) (*entity.Disk, error) {
	//TODO implement me
	panic("implement me")
}

func (s storageApp) ListDisk(arg *entity.DiskListArg) ([]*entity.Disk, error) {
	//TODO implement me
	panic("implement me")
}

func (s storageApp) CreateSnapshot(arg *entity.SnapshotCreateArg) (*entity.Snapshot, error) {
	//TODO implement me
	panic("implement me")
}

func (s storageApp) DeleteSnapshot(arg *entity.SnapshotDeleteArg) (*entity.Snapshot, error) {
	//TODO implement me
	panic("implement me")
}

func (s storageApp) UpdateSnapshot(arg *entity.SnapshotUpdateArg) (*entity.Snapshot, error) {
	//TODO implement me
	panic("implement me")
}

func (s storageApp) GetSnapshot(arg *entity.SnapshotGetArg) (*entity.Snapshot, error) {
	//TODO implement me
	panic("implement me")
}

func (s storageApp) ListSnapshot(arg *entity.SnapshotListArg) ([]*entity.Snapshot, error) {
	//TODO implement me
	panic("implement me")
}
