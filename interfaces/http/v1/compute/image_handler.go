package compute

import (
	"github.com/arturiamu/lplms-public_cloud/common"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateImageArgs struct {
	SnapshotID   string           `json:"snapshot_id"`
	ServerID     string           `json:"server_id"`
	ImageName    *string          `json:"image_name"`
	ImageVersion *string          `json:"image_version"`
	Description  *string          `json:"description"`
	DiskFormat   common.ImageDisk `json:"disk_format"`
}

func (args *CreateImageArgs) toEntityArgs(u *entity.User) *entity.ImageCreateArg {
	var arg = entity.ImageCreateArg{
		ProjectID:   u.ProjectID,
		SnapshotID:  args.SnapshotID,
		ServerID:    args.ServerID,
		ImageName:   args.ImageName,
		Description: args.Description,
	}

	return &arg
}

func (rc *RouterCompute) CreateImage(c *gin.Context) {
	var (
		args   CreateImageArgs
		ctxUid = common.GetUid(c)
		ctxPid = common.GetProject(c)
		u      = &entity.User{
			UID:       ctxUid,
			ProjectID: ctxPid,
		}
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	if args.ServerID == "" {
		c.JSON(http.StatusBadRequest, common.FailWith("empty server id", nil))
		return
	}
	if args.ImageName == nil {
		c.JSON(http.StatusBadRequest, common.FailWith("empty image name", nil))
		return
	}

	createArgs := args.toEntityArgs(u)
	_, err := rc.ci.CreateImage(createArgs)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.Success())
}

type DeleteImageArgs struct {
	ImgaeID               string `json:"imgae_id"`
	Force                 *bool  `json:"force"`
	DeleteRelatedSnapshot *bool  `json:"delete_related_snapshot"`
}

func (args *DeleteImageArgs) toEntityArgs(u *entity.User) *entity.ImageDeleteArg {
	var arg = entity.ImageDeleteArg{
		ProjectID:             u.ProjectID,
		ImageID:               args.ImgaeID,
		Force:                 args.Force,
		DeleteRelatedSnapshot: args.DeleteRelatedSnapshot,
	}

	return &arg
}

func (rc *RouterCompute) DeleteImage(c *gin.Context) {
	var (
		args   DeleteImageArgs
		ctxUid = common.GetUid(c)
		ctxPid = common.GetProject(c)
		u      = &entity.User{
			UID:       ctxUid,
			ProjectID: ctxPid,
		}
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	createArgs := args.toEntityArgs(u)
	_, err := rc.ci.DeleteImage(createArgs)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.Success())
}

type UpdateImageArgs struct {
	ImgaeID     string  `json:"imgae_id"`
	ImageName   *string `json:"image_name"`
	Version     *string `json:"version"`
	Description *string `json:"description"`
	ClientToken string  `json:"client_token"`
}

func (args *UpdateImageArgs) toEntityArgs(u *entity.User) *entity.ImageUpdateArg {
	var arg = entity.ImageUpdateArg{
		ImageID:     args.ImgaeID,
		ProjectID:   u.ProjectID,
		ImageName:   args.ImageName,
		Description: args.Description,
		Version:     args.Version,
	}

	return &arg
}

func (rc *RouterCompute) UpdateImage(c *gin.Context) {
	var (
		args   UpdateImageArgs
		ctxUid = common.GetUid(c)
		ctxPid = common.GetProject(c)
		u      = &entity.User{
			UID:       ctxUid,
			ProjectID: ctxPid,
		}
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	createArgs := args.toEntityArgs(u)
	_, err := rc.ci.UpdateImage(createArgs)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.Success())
}

type GetImageArgs struct {
}

func (args *GetImageArgs) toEntityArgs(u *entity.User) *entity.ImageGetArg {
	return nil
}

func (rc *RouterCompute) GetImage(c *gin.Context) {

}

type ListImageArgs struct {
	ImageID         string                 `json:"image_id"`
	ImageName       string                 `json:"image_name"`
	Platform        common.PlatformType    `json:"platform"`
	ImageOwnerAlias common.ImageOwnerAlias `json:"image_owner_alias"`
	ImageVersion    string                 `json:"image_version"`
	Status          common.ImageStatus     `json:"status"`
}

func (args *ListImageArgs) toEntityArgs(u *entity.User) *entity.ImageListArg {
	var arg = entity.ImageListArg{
		ProjectID:       u.ProjectID,
		ImageID:         args.ImageID,
		Status:          []common.ImageStatus{args.Status},
		ImageName:       args.ImageName,
		ImageOwnerAlias: args.ImageOwnerAlias,
	}

	return &arg
}

func (rc *RouterCompute) ListImage(c *gin.Context) {
	var (
		args   ListImageArgs
		ctxUid = common.GetUid(c)
		ctxPid = common.GetProject(c)
		u      = &entity.User{
			UID:       ctxUid,
			ProjectID: ctxPid,
		}
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	createArgs := args.toEntityArgs(u)
	_, err := rc.ci.ListImage(createArgs)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.Success())
}
