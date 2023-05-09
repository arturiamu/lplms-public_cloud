package mock

import "github.com/arturiamu/lplms-public_cloud/domain/entity"

func (b BizMock) CreateDisk(args *entity.DiskCreateArg) (*entity.DiskCreateResp, error) {
	return nil, nil
}

func (b BizMock) DeleteDisk(args *entity.DiskDeleteArg) (*entity.DiskDeleteResp, error) {
	return nil, nil
}

func (b BizMock) UpdateDisk(args *entity.DiskUpdateArg) (*entity.DiskUpdateResp, error) {
	return nil, nil
}

func (b BizMock) GetDisk(args *entity.DiskGetArg) (*entity.DiskGetResp, error) {
	return &entity.DiskGetResp{
		Disk: getDisk(args.DiskID),
	}, nil
}

func (b BizMock) ListDisk(args *entity.DiskListArg) (*entity.DiskListResp, error) {
	return &entity.DiskListResp{
		Disks: getDiskList(),
	}, nil
}

func (b BizMock) AttachDisk(args *entity.DiskAttachArg) (*entity.DiskAttachResp, error) {
	return nil, nil
}

func (b BizMock) DetachDisk(args *entity.DiskDetachArg) (*entity.DiskDetachResp, error) {
	return nil, nil
}

func (b BizMock) ResizeDisk(args *entity.DiskResizeArg) (*entity.DiskResizeResp, error) {
	return nil, nil
}

func (b BizMock) ResetDisk(args *entity.DiskResetArg) (*entity.DiskResetResp, error) {
	return nil, nil
}

func (b BizMock) CreateSnapshot(args *entity.SnapshotCreateArg) (*entity.SnapshotCreateResp, error) {
	return nil, nil
}

func (b BizMock) DeleteSnapshot(args *entity.SnapshotDeleteArg) (*entity.SnapshotDeleteResp, error) {
	return nil, nil
}

func (b BizMock) UpdateSnapshot(args *entity.SnapshotUpdateArg) (*entity.SnapshotUpdateResp, error) {
	return nil, nil
}

func (b BizMock) GetSnapshot(args *entity.SnapshotGetArg) (*entity.SnapshotGetResp, error) {
	return &entity.SnapshotGetResp{
		Snapshot: getSnapshot(args.SnapshotID),
	}, nil
}

func (b BizMock) ListSnapshot(args *entity.SnapshotListArg) (*entity.SnapshotListResp, error) {
	return &entity.SnapshotListResp{
		Snapshots: getSnapshotList(),
	}, nil
}
