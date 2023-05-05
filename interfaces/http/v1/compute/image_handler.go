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

func (rc *RouterCompute) CreateImage(c *gin.Context) {

}

type DeleteImageArgs struct {
}

func (args *DeleteImageArgs) toEntityArgs(u *entity.User) *entity.ImageDeleteArg {
	return nil
}

func (rc *RouterCompute) DeleteImage(c *gin.Context) {

}

type UpdateImageArgs struct {
}

func (args *UpdateImageArgs) toEntityArgs(u *entity.User) *entity.ImageUpdateArg {
	return nil
}

func (rc *RouterCompute) UpdateImage(c *gin.Context) {

}

type GetImageArgs struct {
}

func (args *GetImageArgs) toEntityArgs(u *entity.User) *entity.ImageGetArg {
	return nil
}

func (rc *RouterCompute) GetImage(c *gin.Context) {

}

type ListImageArgs struct {
}

func (args *ListImageArgs) toEntityArgs(u *entity.User) *entity.ImageListArg {
	return nil
}

func (rc *RouterCompute) ListImage(c *gin.Context) {

}
