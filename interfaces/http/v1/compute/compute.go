package compute

import (
	"github.com/arturiamu/lplms-public_cloud/application"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type RouterCompute struct {
	ci application.ComputeAppInterface
}

func NewCompute(ci application.ComputeAppInterface) *RouterCompute {
	return &RouterCompute{
		ci: ci,
	}
}
