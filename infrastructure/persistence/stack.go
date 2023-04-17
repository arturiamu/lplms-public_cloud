package persistence

import (
	"fmt"
	"github.com/arturiamu/lplms-public_cloud/config"
	"github.com/arturiamu/lplms-public_cloud/domain/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const MSQDSN = "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local"

type Repositories struct {
	User    repository.UserRepositoryInterface
	Compute repository.ComputeRepositoryInterface
}

func NewRepositories(sql *config.MySQL, path *string) (repos *Repositories, err error) {
	dsn := fmt.Sprintf(MSQDSN, sql.Username, sql.Password, sql.Host, sql.Port, sql.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	return &Repositories{
		User:    NewUserRepository(db),
		Compute: NewComputeRepo(path),
	}, nil
}
