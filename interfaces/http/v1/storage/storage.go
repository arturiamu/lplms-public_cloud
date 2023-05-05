package storage

import (
	"github.com/arturiamu/lplms-public_cloud/application"
)

type RouterStorage struct {
	si application.StorageAppInterface
}

func NewStorage(si application.StorageAppInterface) *RouterStorage {
	return &RouterStorage{si: si}
}

type CreateDiskArg struct {
	Size             int    `json:"size"`
	Category         string `json:"category"`
	SnapshotID       string `json:"snapshot_id"`
	DiskName         string `json:"disk_name"`
	Description      string `json:"description"`
	DeleteWithServer bool   `json:"delete_with_server"`
}
