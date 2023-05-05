package repository

import "github.com/arturiamu/lplms-public_cloud/domain/entity"

type StorageRepositoryInterface interface {
	DiskRepositoryInterface
	SnapshotRepositoryInterface
}

type DiskRepositoryInterface interface {
	CreateDisk(args *entity.DiskCreateArg) (*entity.DiskCreateResp, error)
	DeleteDisk(args *entity.DiskDeleteArg) (*entity.DiskDeleteResp, error)
	UpdateDisk(args *entity.DiskUpdateArg) (*entity.DiskUpdateResp, error)
	GetDisk(args *entity.DiskGetArg) (*entity.DiskGetResp, error)
	ListDisk(args *entity.DiskListArg) (*entity.DiskListResp, error)

	AttachDisk(args *entity.DiskAttachArg) (*entity.DiskAttachResp, error)
	DetachDisk(args *entity.DiskDetachArg) (*entity.DiskDetachResp, error)
	ResizeDisk(args *entity.DiskResizeArg) (*entity.DiskResizeResp, error)
	ResetDisk(args *entity.DiskResetArg) (*entity.DiskResetResp, error)
}

type SnapshotRepositoryInterface interface {
	CreateSnapshot(args *entity.SnapshotCreateArg) (*entity.SnapshotCreateResp, error)
	DeleteSnapshot(args *entity.SnapshotDeleteArg) (*entity.SnapshotDeleteResp, error)
	UpdateSnapshot(args *entity.SnapshotUpdateArg) (*entity.SnapshotUpdateResp, error)
	GetSnapshot(args *entity.SnapshotGetArg) (*entity.SnapshotGetResp, error)
	ListSnapshot(args *entity.SnapshotListArg) (*entity.SnapshotListResp, error)
}
