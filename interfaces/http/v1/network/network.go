package network

import (
	"github.com/arturiamu/lplms-public_cloud/application"
	"github.com/gin-gonic/gin"
)

type Network struct {
	ni application.NetworkAppInterface
}

func NewNetwork(ni application.NetworkAppInterface) *Network {
	return &Network{ni: ni}
}

func (s *Network) CreateVpc(c *gin.Context) {

}
