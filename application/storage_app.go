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
	SaveDisk(*entity.Disk) (*entity.Disk, error)
	GetDisk(uint64) (*entity.Disk, error)
	ListDisk() (*entity.Disk, error)
	GetDiskBy(uint64) (*entity.Disk, error)
}

type SnapshotRepositoryInterface interface {
	SaveSnapshot(*entity.Disk) (*entity.Disk, error)
	GetSnapshot(uint64) (*entity.Disk, error)
	ListSnapshot() (*entity.Disk, error)
	GetSnapshotBy(uint64) (*entity.Disk, error)
}

type storageApp struct {
	sr repository.StorageRepositoryInterface
}

func (s storageApp) SaveDisk(disk *entity.Disk) (*entity.Disk, error) {
	return s.sr.SaveDisk(disk)
}

func (s storageApp) GetDisk(u uint64) (*entity.Disk, error) {
	//TODO implement me
	panic("implement me")
}

func (s storageApp) ListDisk() (*entity.Disk, error) {
	//TODO implement me
	panic("implement me")
}

func (s storageApp) GetDiskBy(u uint64) (*entity.Disk, error) {
	//TODO implement me
	panic("implement me")
}

func (s storageApp) SaveSnapshot(disk *entity.Disk) (*entity.Disk, error) {
	//TODO implement me
	panic("implement me")
}

func (s storageApp) GetSnapshot(u uint64) (*entity.Disk, error) {
	//TODO implement me
	panic("implement me")
}

func (s storageApp) ListSnapshot() (*entity.Disk, error) {
	//TODO implement me
	panic("implement me")
}

func (s storageApp) GetSnapshotBy(u uint64) (*entity.Disk, error) {
	//TODO implement me
	panic("implement me")
}
