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

func (s *Storage) CreateDisk(c *gin.Context) {

}
