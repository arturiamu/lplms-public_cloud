package repository

import "github.com/arturiamu/lplms-public_cloud/domain/entity"

type StorageRepositoryInterface interface {
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
