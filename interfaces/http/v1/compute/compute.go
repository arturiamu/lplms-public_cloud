package compute

import (
	"github.com/arturiamu/lplms-public_cloud/application"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Compute struct {
	ci application.ComputeAppInterface
}

func NewCompute(ci application.ComputeAppInterface) *Compute {
	return &Compute{
		ci: ci,
	}
}
