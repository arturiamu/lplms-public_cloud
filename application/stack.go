package application

import (
	"github.com/arturiamu/lplms-public_cloud/config"
	"github.com/arturiamu/lplms-public_cloud/infrastructure/persistence"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type StackInterface interface {
	User() UserAppInterface
	Compute() ComputeAppInterface
	Storage() StorageAppInterface
	Network() NetworkAppInterface
}

var _ StackInterface = &stk{}

type stk struct {
	userRepo    *persistence.UserRepo
	computeRepo *persistence.ComputeRepo
	storageRepo *persistence.StorageRepo
	networkRepo *persistence.NetworkRepo
}

func NewStack(path *string) (stack StackInterface, err error) {
	var (
		db *gorm.DB
	)

	if config.AppConfig.App.Env == "dev" {
		db, err = gorm.Open(sqlite.Open("config/lplms.dev.db"), &gorm.Config{})
		if err != nil {
			return nil, err
		}
	} else {
		db, err = gorm.Open(mysql.Open(config.GetDSN()))
		if err != nil {
			return nil, err
		}
	}

	return stk{
		userRepo:    persistence.NewUserRepository(db),
		computeRepo: persistence.NewComputeRepo(path),
		storageRepo: persistence.NewStorageRepo(path),
		networkRepo: persistence.NewNetworkRepo(path),
	}, nil
}

func (s stk) User() UserAppInterface {
	return s.userRepo
}

func (s stk) Compute() ComputeAppInterface {
	return s.computeRepo
}

func (s stk) Storage() StorageAppInterface {
	return s.storageRepo
}

func (s stk) Network() NetworkAppInterface {
	return s.networkRepo
}
