package storage

import (
	"github.com/arturiamu/lplms-public_cloud/application"
	"github.com/gin-gonic/gin"
)

type Storage struct {
	si application.StorageAppInterface
}

func NewStorage(si application.StorageAppInterface) *Storage {
	return &Storage{si: si}
}

type CreateDiskArg struct {
	Size             int    `json:"size"`
	Category         string `json:"category"`
	SnapshotID       string `json:"snapshot_id"`
	DiskName         string `json:"disk_name"`
	Description      string `json:"description"`
	DeleteWithServer bool   `json:"delete_with_server"`
}

func (s *Storage) CreateDisk(c *gin.Context) {

}
