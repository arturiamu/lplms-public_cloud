package storage

import (
	"github.com/arturiamu/lplms-public_cloud/common"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateSnapshotArgs struct {
	// 云盘ID
	DiskID *string `json:"disk_id,omitempty" oplog:"disk_id,omitempty"`
	// 快照的显示名称。长度为2~128个英文或中文字符。必须以大小字母或中文开头，不能以 `http://` 和 `https://` 开头。可以包含数字、半角冒号（:）、下划线（_）或者连字符（-）。
	//
	// 为防止和自动快照的名称冲突，不能以auto开头。
	SnapshotName string `json:"snapshot_name,omitempty" oplog:"snapshot_name,omitempty"`
	// 快照的描述。长度为2~256个英文或中文字符，不能以 `http://` 和 `https://` 开头。 默认值：空
	Description *string `json:"description,omitempty" oplog:"description,omitempty"`
}

func (args *CreateSnapshotArgs) toEntityArgs(u *entity.User) *entity.SnapshotCreateArg {
	var arg = entity.SnapshotCreateArg{
		ProjectID:    u.ProjectID,
		DiskID:       *args.DiskID,
		SnapshotName: args.SnapshotName,
		Description:  args.Description,
	}
	return &arg
}

func (rs *RouterStorage) CreateSnapshot(c *gin.Context) {
	var (
		args CreateSnapshotArgs
		u    = common.GetUser(c)
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	if args.DiskID == nil {
		c.JSON(http.StatusBadRequest, common.FailWith("lack of diskID", nil))
		return
	}

	// create v_switch
	_, err := rs.si.CreateSnapshot(args.toEntityArgs(u))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, common.SuccessWith("", nil))
}

type DeleteSnapshotArgs struct {
	SnapshotID string `pos:"path:id" quota:"snapshot_num:-1"`
	// 是否强制删除有磁盘关联的快照。说明 删除后该磁盘无法重新初始化。
	Force *bool `json:"force,omitempty"` // TODO 暂不支持
}

func (args *DeleteSnapshotArgs) toEntityArgs(u *entity.User) *entity.SnapshotDeleteArg {
	var arg = entity.SnapshotDeleteArg{
		ProjectID:  u.ProjectID,
		SnapshotID: args.SnapshotID,
	}
	return &arg
}

func (rs *RouterStorage) DeleteSnapshot(c *gin.Context) {
	var (
		args DeleteSnapshotArgs
		u    = common.GetUser(c)
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	// create v_switch
	_, err := rs.si.DeleteSnapshot(args.toEntityArgs(u))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, common.SuccessWith("", nil))
}

type UpdateSnapshotArgs struct {
	SnapshotID string `pos:"path:id"`
	// 快照的显示名称。长度为2~128个英文或中文字符。必须以大小字母或中文开头，不能以 `http://` 和 `https://` 开头。可以包含数字、半角冒号（:）、下划线（_）或者连字符（-）。
	//
	// 为防止和自动快照的名称冲突，不能以auto开头。
	SnapshotName *string `json:"snapshot_name,omitempty" oplog:"snapshot_name,omitempty"`
	// 快照的描述。长度为2~256个英文或中文字符，不能以 `http://` 和 `https://` 开头。 默认值：空
	Description *string `json:"description,omitempty" oplog:"description,omitempty"`
}

func (args *UpdateSnapshotArgs) toEntityArgs(u *entity.User) *entity.SnapshotUpdateArg {
	var arg = entity.SnapshotUpdateArg{
		ProjectID:    u.ProjectID,
		SnapshotID:   args.SnapshotID,
		SnapshotName: args.SnapshotName,
		Description:  args.Description,
	}
	return &arg
}

func (rs *RouterStorage) UpdateSnapshot(c *gin.Context) {
	var (
		args UpdateSnapshotArgs
		u    = common.GetUser(c)
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	// create v_switch
	_, err := rs.si.UpdateSnapshot(args.toEntityArgs(u))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, common.SuccessWith("", nil))
}

type GetSnapshotArgs struct {
	SnapshotID string `pos:"path:id"`
}

func (args *GetSnapshotArgs) toEntityArgs(u *entity.User) *entity.SnapshotGetArg {
	var arg = entity.SnapshotGetArg{
		ProjectID:  u.ProjectID,
		SnapshotID: args.SnapshotID,
	}
	return &arg
}

func (rs *RouterStorage) GetSnapshot(c *gin.Context) {
	var (
		args GetSnapshotArgs
		u    = common.GetUser(c)
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	// create v_switch
	_, err := rs.si.GetSnapshot(args.toEntityArgs(u))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, common.SuccessWith("", nil))
}

type ListSnapshotArgs struct {
	SnapshotID     *string          `pos:"query:snapshot_id"`
	SnapshotName   *string          `pos:"query:snapshot_name"`
	DiskID         *string          `pos:"query:disk_id"`
	ServerID       *string          `pos:"query:server_id"`
	SourceDiskType *common.DiskType `pos:"query:source_disk_type"`
}

func (args *ListSnapshotArgs) toEntityArgs(u *entity.User) *entity.SnapshotListArg {
	var arg = entity.SnapshotListArg{
		ProjectID:      u.ProjectID,
		SnapshotName:   args.SnapshotName,
		DiskID:         args.DiskID,
		ServerID:       args.ServerID,
		SourceDiskType: args.SourceDiskType.String(),
	}
	return &arg
}

func (rs *RouterStorage) ListSnapshot(c *gin.Context) {
	var (
		args ListSnapshotArgs
		u    = common.GetUser(c)
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	listArgs := args.toEntityArgs(u)
	if args.SnapshotID != nil {
		listArgs.SnapshotIDs = []string{*args.SnapshotID}
	}

	// create v_switch
	listRes, err := rs.si.ListSnapshot(args.toEntityArgs(u))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.SuccessWith("", listRes.Snapshots))
}
