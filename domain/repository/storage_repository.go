package repository

import "github.com/arturiamu/lplms-public_cloud/domain/entity"

type StorageRepositoryInterface interface {
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
