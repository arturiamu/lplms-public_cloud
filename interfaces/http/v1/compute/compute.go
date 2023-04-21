package compute

import (
	"github.com/arturiamu/lplms-public_cloud/application"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/gin-gonic/gin"
)

type Compute struct {
	ci application.ComputeAppInterface
}

func NewCompute(ci application.ComputeAppInterface) *Compute {
	return &Compute{
		ci: ci,
	}
}

type CreateServerArg struct {
}

func (co *Compute) CreateServer(c *gin.Context) {
	co.ci.SaveServer(&entity.Server{})
}

func (co *Compute) DeleteServer(c *gin.Context) {

}
