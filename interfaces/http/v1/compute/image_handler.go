package compute

import (
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/gin-gonic/gin"
)

type CreateImageArgs struct {
}

func (args *CreateImageArgs) toEntityArgs(u *entity.User) *entity.ImageCreateArg {
	return nil
}

func (co *Compute) CreateImage(c *gin.Context) {

}

type DeleteImageArgs struct {
}

func (args *DeleteImageArgs) toEntityArgs(u *entity.User) *entity.ImageDeleteArg {
	return nil
}

func (co *Compute) DeleteImage(c *gin.Context) {

}

type UpdateImageArgs struct {
}

func (args *UpdateImageArgs) toEntityArgs(u *entity.User) *entity.ImageUpdateArg {
	return nil
}

func (co *Compute) UpdateImage(c *gin.Context) {

}

type GetImageArgs struct {
}

func (args *GetImageArgs) toEntityArgs(u *entity.User) *entity.ImageGetArg {
	return nil
}

func (co *Compute) GetImage(c *gin.Context) {

}

type ListImageArgs struct {
}

func (args *ListImageArgs) toEntityArgs(u *entity.User) *entity.ImageListArg {
	return nil
}

func (co *Compute) ListImage(c *gin.Context) {

}
