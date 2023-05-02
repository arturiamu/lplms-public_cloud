package repository

import "github.com/arturiamu/lplms-public_cloud/domain/entity"

type StorageRepositoryInterface interface {
	DiskRepositoryInterface
	SnapshotRepositoryInterface
}

type DiskRepositoryInterface interface {
	CreateDisk(arg *entity.DiskCreateArg) (*entity.DiskCreateResp, error)
	DeleteDisk(arg *entity.DiskDeleteArg) (*entity.DiskDeleteResp, error)
	UpdateDisk(arg *entity.DiskUpdateArg) (*entity.DiskUpdateResp, error)
	GetDisk(arg *entity.DiskGetArg) (*entity.DiskGetResp, error)
	ListDisk(arg *entity.DiskListArg) (*entity.DiskListResp, error)
}

type SnapshotRepositoryInterface interface {
	CreateSnapshot(arg *entity.SnapshotCreateArg) (*entity.SnapshotCreateResp, error)
	DeleteSnapshot(arg *entity.SnapshotDeleteArg) (*entity.SnapshotDeleteResp, error)
	UpdateSnapshot(arg *entity.SnapshotUpdateArg) (*entity.SnapshotUpdateResp, error)
	GetSnapshot(arg *entity.SnapshotGetArg) (*entity.SnapshotGetResp, error)
	ListSnapshot(arg *entity.SnapshotListArg) (*entity.SnapshotListResp, error)
}
