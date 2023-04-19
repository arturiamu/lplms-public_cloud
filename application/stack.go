package application

import (
	"github.com/arturiamu/lplms-public_cloud/config"
	"github.com/arturiamu/lplms-public_cloud/infrastructure/persistence"
)

type StackInterface interface {
	User() UserAppInterface
	Compute() ComputeAppInterface
	Storage() StorageAppInterface
	Network() NetworkAppInterface
}

func NewStack(config *config.AppConfig, path *string) StackInterface {
	stk, err := persistence.NewRepositories(nil, nil)
	if err != nil {
		return nil
	}
	return stk
}
