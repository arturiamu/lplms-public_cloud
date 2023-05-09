package storage

import (
	"errors"
	"github.com/arturiamu/lplms-public_cloud/common"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateDiskArgs struct {
	// 容量大小，以GiB为单位。指定该参数后，其取值必须≥指定快照ID的容量大小。取值范围：
	// - cloud：5~2000
	// - cloud_efficiency：20~32768
	// - cloud_ssd：20~32768
	// - cloud_essd：20~32768
	Size *int64 `json:"size,omitempty" oplog:"size,omitempty" quota:"disk_num:1, disk_$category$_num:1, disk_size, disk_$category$_size"`

	// 云盘名称。长度为2~128个英文或中文字符。必须以大小字母或中文开头，不能以 http:// 和 https:// 开头。可以包含数字、半角冒号（:）、下划线（_）或者连字符（-）。
	//
	// 默认值：空
	DiskName *string `json:"disk_name,omitempty" oplog:"disk_name,omitempty"`

	// 云盘描述。长度为2~256个英文或中文字符，不能以 http:// 和 https:// 开头。
	//
	// 默认值：空
	Description *string `json:"description,omitempty" oplog:"description,omitempty"`

	// 云盘种类。取值范围：
	// - cloud：普通云盘
	// - cloud_ssd：SSD云盘
	// 默认值：cloud
	Category *common.DiskCategory `json:"category,omitempty" oplog:"category,omitempty"`

	// 创建云盘使用的快照。指定该参数后，Size会被忽略，实际创建的云盘大小为指定快照的大小。
	SnapshotID *string `json:"snapshot_id,omitempty" oplog:"snapshot_id,omitempty"`

	// 创建云盘使用的镜像
	ImageID *string `json:"image_id,omitempty" oplog:"image_id,omitempty"`
}

func (args *CreateDiskArgs) IsValid() (err error) {
	if args.SnapshotID != nil && args.ImageID != nil {
		err = errors.New("请指定唯一的创建来源")
		return
	}

	if args.Size == nil || *args.Size < 0 {
		err = errors.New("错误的磁盘容量")
	}

	if args.Category != nil && !args.Category.IsValid() {
		err = errors.New("错误的磁盘类型")
	}

	return
}

func (args *CreateDiskArgs) toEntityArgs(u *entity.User) *entity.DiskCreateArg {
	var arg = entity.DiskCreateArg{
		ProjectID:   u.ProjectID,
		Category:    args.Category.String(),
		Size:        args.Size,
		DiskName:    args.DiskName,
		Description: args.Description,
		SnapshotID:  args.SnapshotID,
		ImageID:     args.ImageID,
	}
	return &arg
}

func (rs *RouterStorage) CreateDisk(c *gin.Context) {
	var (
		args   CreateDiskArgs
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

	if err := args.IsValid(); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	// 目前暂时不支持 非 SSD 磁盘创建, HDD 盘目前没有冗余
	if args.Category != nil && *args.Category == common.CloudDiskCategory {
		c.JSON(http.StatusBadRequest, common.FailWith("disk category not support", nil))
		return
	}

	// create v_switch
	_, err := rs.si.CreateDisk(args.toEntityArgs(u))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, common.SuccessWith("", nil))
}

type DeleteDiskArgs struct {
	DiskID string `json:"disk_id"`
}

func (args *DeleteDiskArgs) toEntityArgs(u *entity.User) *entity.DiskDeleteArg {
	var arg = entity.DiskDeleteArg{
		DiskID:    args.DiskID,
		ProjectID: u.ProjectID,
	}
	return &arg
}

func (rs *RouterStorage) DeleteDisk(c *gin.Context) {
	var (
		args   DeleteDiskArgs
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

	// create v_switch
	_, err := rs.si.DeleteDisk(args.toEntityArgs(u))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, common.SuccessWith("", nil))
}

type UpdateDiskArgs struct {
	DiskID string `pos:"path:id"`
	// 磁盘名称。长度为2~128个英文或中文字符。必须以大小字母或中文开头，不能以 `http://` 和 `https://` 开头。可以包含数字、半角冒号（:）、下划线（_）或者连字符（-）。
	DiskName *string `json:"disk_name,omitempty" oplog:"disk_name,omitempty"`
	// 磁盘描述。 长度为2~256个英文或中文字符，不能以 `http://` 和 `https://` 开头。
	Description *string `json:"description,omitempty" oplog:"description,omitempty"`
	// 磁盘是否随实例释放。默认值：无，无表示不改变当前的值。
	//
	// 在下列两种情况下，将参数DeleteWithServer设置成false时会报错。
	// - 磁盘的种类（category）为本地盘（ephemeral）时。
	// - 磁盘的种类（category）为普通云盘（cloud），且不可以卸载（Portable=false）时。
	DeleteWithServer *bool `json:"delete_with_server,omitempty" oplog:"delete_with_server,omitempty"`
}

func (args *UpdateDiskArgs) toEntityArgs(u *entity.User) *entity.DiskUpdateArg {
	var arg = entity.DiskUpdateArg{
		ProjectID:        u.ProjectID,
		DiskID:           args.DiskID,
		DiskName:         args.DiskName,
		DeleteWithServer: args.DeleteWithServer,
		Description:      args.Description,
	}
	return &arg
}

func (rs *RouterStorage) UpdateDisk(c *gin.Context) {
	var (
		args   UpdateDiskArgs
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

	// create v_switch
	_, err := rs.si.UpdateDisk(args.toEntityArgs(u))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, common.SuccessWith("", nil))
}

type GetDiskArgs struct {
	DiskID string `json:"disk_id"`
}

func (args *GetDiskArgs) toEntityArgs(u *entity.User) *entity.DiskGetArg {
	var arg = entity.DiskGetArg{
		ProjectID: u.ProjectID,
		DiskID:    args.DiskID,
	}
	return &arg
}

func (rs *RouterStorage) GetDisk(c *gin.Context) {
	var (
		id     = c.Param("id")
		args   GetDiskArgs
		ctxUid = common.GetUid(c)
		ctxPid = common.GetProject(c)
		u      = &entity.User{
			UID:       ctxUid,
			ProjectID: ctxPid,
		}
	)
	entityArgs := args.toEntityArgs(u)
	entityArgs.DiskID = id
	resp, err := rs.si.GetDisk(entityArgs)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, common.SuccessWith("", resp.Disk))
}

type ListDiskArgs struct {
	DiskID   *string          `pos:"query:disk_id"`
	DiskName *string          `pos:"query:disk_name"`
	ServerID *string          `pos:"query:server_id"`
	DiskType *common.DiskType `pos:"query:disk_type"`
	// DiskCategory *common.DiskCategory   `pos:"query:disk_category"`
	Status *common.DiskStatusType `pos:"query:status"`
}

func (args *ListDiskArgs) toEntityArgs(u *entity.User) *entity.DiskListArg {
	var arg = entity.DiskListArg{
		ProjectID: u.ProjectID,
		ServerID:  args.ServerID,
		DiskName:  args.DiskName,
		DiskType:  args.DiskType,
		Status:    args.Status,
	}
	return &arg
}

func (rs *RouterStorage) ListDisk(c *gin.Context) {
	var (
		args   ListDiskArgs
		ctxUid = common.GetUid(c)
		ctxPid = common.GetProject(c)
		u      = &entity.User{
			UID:       ctxUid,
			ProjectID: ctxPid,
		}
	)
	listArgs := args.toEntityArgs(u)
	if args.DiskID != nil {
		listArgs.DiskIDs = []string{*args.DiskID}
	}

	listRes, err := rs.si.ListDisk(args.toEntityArgs(u))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	data := make([]entity.DiskInfo, len(listRes.Disks))
	for _, disk := range listRes.Disks {
		data = append(data, *entity.BuildDiskInfo(disk))
	}
	c.JSON(http.StatusOK, common.SuccessWith("", listRes.Disks))
}

type DetachDiskArgs struct {
	DiskID   *string `pos:"path:id"`
	ServerID *string `json:"server_id,omitempty" oplog:"server_id,omitempty"`
}

func (args *DetachDiskArgs) toEntityArgs(u *entity.User) *entity.DiskDetachArg {
	var arg = entity.DiskDetachArg{
		ProjectID: u.ProjectID,
		DiskID:    *args.DiskID,
		ServerID:  *args.ServerID,
	}
	return &arg
}

func (rs *RouterStorage) DetachDisk(c *gin.Context) {
	var (
		args   DetachDiskArgs
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

	// create v_switch
	_, err := rs.si.DetachDisk(args.toEntityArgs(u))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, common.SuccessWith("", nil))
}

type AttachDiskArgs struct {
	DiskID *string `pos:"path:id"`

	ServerID         *string `json:"server_id,omitempty"          oplog:"server_id,omitempty"`
	DeleteWithServer *bool   `json:"delete_with_server,omitempty" oplog:"delete_with_server,omitempty"`
	// Bootable         *bool   `json:"bootable,omitempty"`
	// 挂载系统盘时，设置实例的用户名密码，仅对administrator和root用户名生效，其他用户名不生效。
	// 长度为8至30个字符，必须同时包含大小写英文字母、数字和特殊符号中的三类字符。特殊符号可以是：
	// ```
	// ()`~!@#$%^&*-_+=|{}[]:;'<>,.?/
	// ```
	// 其中，Windows实例不能以斜线号（/）为密码首字符。
	//
	// 默认值：保持不变。
	//
	// @GSD:NOTE 说明 如果传入Password参数，建议您使用HTTPS协议发送请求，避免密码泄露。
	// Password *string `json:"password,omitempty"`
	// 密钥对名称。
	// 挂载系统盘时，为Linux系统ECS实例绑定的SSH密钥对的名称。
	//
	// @GSD:NOTE 说明 Windows Server系统：不支持SSH密钥对。即使填写了该参数，只执行Password的配置。Linux系统：密码登录方式会被初始化成禁止。
	// KeyPairName *string `json:"key_pair_name,omitempty"`
}

func (args *AttachDiskArgs) toEntityArgs(u *entity.User) *entity.DiskAttachArg {
	var arg = entity.DiskAttachArg{
		ProjectID:        u.ProjectID,
		DiskID:           *args.DiskID,
		DeleteWithServer: args.DeleteWithServer,
	}
	return &arg
}

func (rs *RouterStorage) AttachDisk(c *gin.Context) {
	var (
		args   AttachDiskArgs
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

	if args.ServerID == nil {
		c.JSON(http.StatusBadRequest, common.FailWith("empty server id", nil))
		return
	}

	// create v_switch
	_, err := rs.si.AttachDisk(args.toEntityArgs(u))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, common.SuccessWith("", nil))
}

type ResizeDiskType string

var (
	ResizeDiskTypeOffline ResizeDiskType = "offline"
	ResizeDiskTypeOnline  ResizeDiskType = "online"
)

type ResizeDiskArgs struct {
	DiskID string `pos:"path:id"`
	// 希望扩容到的云盘容量大小。单位为GiB。取值范围：
	// - 高效云盘（cloud_efficiency）：20~32768
	// - SSD云盘（cloud_ssd）：20~32768
	// - ESSD云盘（cloud_essd）：20~32768
	// - 普通云盘（cloud）：5~2000
	// 指定的新云盘容量必须比原云盘容量大。
	NewSize *int64 `json:"new_size,omitempty" oplog:"new_size,omitempty" quota:"disk_resize($zone_id, $disk_id, $new_size)"`
	// 扩容云盘的方式。取值范围：
	// - **offline**（默认）：离线扩容。扩容后，您必须在控制台重启实例或者调用API RebootInstance使操作生效。
	// - **online**：在线扩容，无需重启实例即可完成扩容。云盘类型支持高效云盘、SSD云盘和ESSD云盘。
	Type *ResizeDiskType `json:"resize_disk_type,omitempty" oplog:"type,omitempty"`
}

func (args *ResizeDiskArgs) toEntityArgs(u *entity.User) *entity.DiskResizeArg {
	var arg = entity.DiskResizeArg{
		ProjectID: u.ProjectID,
		DiskID:    args.DiskID,
		NewSize:   *args.NewSize,
	}
	return &arg
}

func (rs *RouterStorage) ResizeDisk(c *gin.Context) {
	var (
		args   ResizeDiskArgs
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

	// create v_switch
	_, err := rs.si.ResizeDisk(args.toEntityArgs(u))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, common.SuccessWith("", nil))
}

type ResetDiskArgs struct {
	DiskID string `pos:"path:id"`
	// 需要恢复到某一磁盘阶段的历史快照ID。
	SnapshotID string `json:"snapshot_id" oplog:"snapshot_id,omitempty"`
}

func (args *ResetDiskArgs) toEntityArgs(u *entity.User) *entity.DiskResetArg {
	var arg = entity.DiskResetArg{
		ProjectID:  u.ProjectID,
		SnapshotID: args.SnapshotID,
		DiskID:     args.DiskID,
	}
	return &arg
}

func (rs *RouterStorage) ResetDisk(c *gin.Context) {
	var (
		args   ResetDiskArgs
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

	// create v_switch
	_, err := rs.si.ResetDisk(args.toEntityArgs(u))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, common.SuccessWith("", nil))
}
