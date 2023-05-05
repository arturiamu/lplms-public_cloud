package network

import (
	"github.com/arturiamu/lplms-public_cloud/application"
)

type Network struct {
	ni application.NetworkAppInterface
}

func NewNetwork(ni application.NetworkAppInterface) *Network {
	return &Network{ni: ni}
}
