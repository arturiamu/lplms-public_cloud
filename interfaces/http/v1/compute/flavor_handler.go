package compute

import (
	"github.com/arturiamu/lplms-public_cloud/common"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateFlavorArgs struct {
}

func (args *CreateFlavorArgs) toEntityArgs(u *entity.User) *entity.FlavorCreateArg {
	return nil
}

func (rc *RouterCompute) CreateFlavor(c *gin.Context) {

}

type DeleteFlavorArgs struct {
}

func (args *DeleteFlavorArgs) toEntityArgs(u *entity.User) *entity.FlavorDeleteArg {
	return nil
}

func (rc *RouterCompute) DeleteFlavor(c *gin.Context) {

}

type UpdateFlavorArgs struct {
}

func (args *UpdateFlavorArgs) toEntityArgs(u *entity.User) *entity.FlavorUpdateArg {
	return nil
}

func (rc *RouterCompute) UpdateFlavor(c *gin.Context) {

}

type GetFlavorArgs struct {
}

func (args *GetFlavorArgs) toEntityArgs(u *entity.User) *entity.FlavorGetArg {
	return nil
}

func (rc *RouterCompute) GetFlavor(c *gin.Context) {
	var (
		id = c.Param("id")
	)

	resp, err := rc.ci.GetFlavor(&entity.FlavorGetArg{FlavorID: id})
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.SuccessWith("", resp.Flavor))
}

func (rc *RouterCompute) ListFlavor(c *gin.Context) {
	resp, err := rc.ci.ListFlavor(&entity.FlavorListArg{})
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.SuccessWith("", resp.Flavors))
}
