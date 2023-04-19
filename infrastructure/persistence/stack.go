package persistence

import (
	"fmt"
	"github.com/arturiamu/lplms-public_cloud/application"
	"github.com/arturiamu/lplms-public_cloud/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const MSQDSN = "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local"

var _ application.StackInterface = &Repositories{}

type Repositories struct {
	computeRepo ComputeRepo
	userRepo    UserRepo
	storageRepo StorageRepo
	networkRepo NetworkRepo
}

func NewRepositories(sql *config.MySQL, path *string) (repos *Repositories, err error) {
	dsn := fmt.Sprintf(MSQDSN, sql.Username, sql.Password, sql.Host, sql.Port, sql.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	return &Repositories{
		userRepo:    *NewUserRepository(db),
		computeRepo: *NewComputeRepo(path),
		storageRepo: *NewStorageRepo(path),
		networkRepo: *NewNetworkRepo(path),
	}, nil
}

func (r Repositories) User() application.UserAppInterface {
	return &r.userRepo
}

func (r Repositories) Compute() application.ComputeAppInterface {
	return &r.computeRepo
}

func (r Repositories) Storage() application.StorageAppInterface {
	return &r.storageRepo
}

func (r Repositories) Network() application.NetworkAppInterface {
	return &r.networkRepo
}
