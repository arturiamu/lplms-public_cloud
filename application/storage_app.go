package application

import (
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/arturiamu/lplms-public_cloud/domain/repository"
	"github.com/arturiamu/lplms-public_cloud/infrastructure/mock"
)

var _ StorageAppInterface = &storageApp{}
var _ StorageAppInterface = &mock.BizMock{}

type StorageAppInterface interface {
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

type storageApp struct {
	sr repository.StorageRepositoryInterface
}

func (s *storageApp) AttachDisk(args *entity.DiskAttachArg) (*entity.DiskAttachResp, error) {
	return s.sr.AttachDisk(args)
}

func (s *storageApp) DetachDisk(args *entity.DiskDetachArg) (*entity.DiskDetachResp, error) {
	return s.sr.DetachDisk(args)
}

func (s *storageApp) ResizeDisk(args *entity.DiskResizeArg) (*entity.DiskResizeResp, error) {
	return s.sr.ResizeDisk(args)
}

func (s *storageApp) ResetDisk(args *entity.DiskResetArg) (*entity.DiskResetResp, error) {
	return s.sr.ResetDisk(args)
}

func (s *storageApp) CreateDisk(args *entity.DiskCreateArg) (*entity.DiskCreateResp, error) {
	return s.sr.CreateDisk(args)
}

func (s *storageApp) DeleteDisk(args *entity.DiskDeleteArg) (*entity.DiskDeleteResp, error) {
	return s.sr.DeleteDisk(args)
}

func (s *storageApp) UpdateDisk(args *entity.DiskUpdateArg) (*entity.DiskUpdateResp, error) {
	return s.sr.UpdateDisk(args)
}

func (s *storageApp) GetDisk(args *entity.DiskGetArg) (*entity.DiskGetResp, error) {
	return s.sr.GetDisk(args)
}

func (s *storageApp) ListDisk(args *entity.DiskListArg) (*entity.DiskListResp, error) {
	return s.sr.ListDisk(args)
}

func (s *storageApp) CreateSnapshot(args *entity.SnapshotCreateArg) (*entity.SnapshotCreateResp, error) {
	return s.sr.CreateSnapshot(args)
}

func (s *storageApp) DeleteSnapshot(args *entity.SnapshotDeleteArg) (*entity.SnapshotDeleteResp, error) {
	return s.sr.DeleteSnapshot(args)
}

func (s *storageApp) UpdateSnapshot(args *entity.SnapshotUpdateArg) (*entity.SnapshotUpdateResp, error) {
	return s.sr.UpdateSnapshot(args)
}

func (s *storageApp) GetSnapshot(args *entity.SnapshotGetArg) (*entity.SnapshotGetResp, error) {
	return s.sr.GetSnapshot(args)
}

func (s *storageApp) ListSnapshot(args *entity.SnapshotListArg) (*entity.SnapshotListResp, error) {
	return s.sr.ListSnapshot(args)
}
