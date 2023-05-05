package network

import (
	"github.com/arturiamu/lplms-public_cloud/application"
)

type RouterNetwork struct {
	ni application.NetworkAppInterface
}

func NewNetwork(ni application.NetworkAppInterface) *RouterNetwork {
	return &RouterNetwork{ni: ni}
}
