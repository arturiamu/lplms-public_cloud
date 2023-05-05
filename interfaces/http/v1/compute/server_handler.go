package compute

import (
	"errors"
	"github.com/arturiamu/lplms-public_cloud/common"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateServerArg struct {
	ImageID         string `json:"image_id"`          // 镜像 ID，启动实例时选择的镜像资源。您可以通过 DescribeImages 查询您可以使用的镜像资源。
	FlavorID        string `json:"flavor_id"`         // 实例的资源规格ID。
	FlavorName      string `json:"flavor_name"`       // 实例的资源规格名称。
	SecurityGroupID string `json:"security_group_id"` // 指定新创建实例所属于的安全组 ID。同一个安全组内的实例之间可以互相访问，一个安全组能容纳的实例数量视安全组类型而定，具体请参见使用限制的安全组章节。

	// 虚拟交换机ID。如果您创建的是VPC类型ECS实例，必须指定虚拟交换机ID，且安全组和虚拟交换机在同一个专有网络VPC中。
	VSwitchID *string `json:"v_switch_id"`

	// 实例的名称。长度为2~128个英文或中文字符。必须以大小字母或中文开头，不能以 http:// 和 https:// 开头。
	// 可以包含数字、半角冒号（:）、下划线（_）、英文句号（.）或者连字符（-）。
	// 如果没有指定该参数，默认值为实例的ServerId。
	ServerName *string `json:"server_name"`

	// 云服务器的主机名。
	// - 英文句号（.）和短横线（-）不能作为首尾字符，更不能连续使用。
	// - Windows实例：字符长度为2~15，不支持英文句号（.），不能全是数字。允许大小写英文字母、数字和短横线（-）。
	// - 其他类型实例（Linux等）：字符长度为2~64，支持多个英文句号（.），英文句号之间为一段，每段允许大小写英文字母、数字和短横线（-）。
	HostName *string `json:"host_name"`

	// 实例的描述。长度为 2~256 个英文或中文字符，不能以 http:// 和 https:// 开头。
	Description *string `json:"description"`

	// 实例的密码。长度为 8 至 30 个字符，必须同时包含大小写英文字母、数字和特殊符号中的三类字符。特殊符号可以是：
	//
	// ```
	// ()`~!@#$%^&*-_+=|{}[]:;'<>,.?/
	// ```
	// 其中，Windows实例不能以斜线号（/）为密码首字符。
	//
	// @GSE:NOTE: 如果传入 Password 参数，建议您使用 HTTPS 协议发送请求，避免密码泄露。
	Password *string `json:"password"`

	// 实例自定义数据，需要以 Base64 方式编码，原始数据最多为16KB。
	UserData *string `json:"user_data"`

	// 密钥对名称。
	// - Windows实例，忽略该参数。默认为空。即使填写了该参数，仍旧只执行Password的内容。
	// - Linux实例的密码登录方式会被初始化成禁止。为提高实例安全性，强烈建议您使用密钥对的连接方式。
	KeyPairName *string `json:"key_pair_name"`

	VPCID *string `json:"vpc_id"` // VPC ID

	SystemDisk *ServerDiskArgs   `json:"system_disk,omitempty"` // 系统盘
	DataDisks  []*ServerDiskArgs `json:"data_disks,omitempty"`  // 数据盘
	EIP        *ipInfo           `json:"eip,omitempty"`
}

type ServerDiskArgs struct {
	Size             int                  `json:"size"`
	Category         *common.DiskCategory `json:"category"`
	SnapshotID       *string              `json:"snapshot_id"`
	DiskName         *string              `json:"disk_name"`
	Description      *string              `json:"description"`
	DeleteWithServer *bool                `json:"delete_with_server"`
}

func (a *ServerDiskArgs) toDiskArgs() *entity.DiskArgs {
	return &entity.DiskArgs{
		Size:             a.Size,
		Category:         a.Category.String(),
		SnapshotID:       a.SnapshotID,
		DiskName:         a.DiskName,
		Description:      a.Description,
		DeleteWithServer: a.DeleteWithServer,
	}
}

func (a *ServerDiskArgs) IsValid() bool {
	return a.Size > 0 && a.Category.IsValid()
}

type ipInfo struct {
	EIPName   string `json:"eip_name"`
	Bandwidth uint   `json:"bandwidth"`
	MNO       string `json:"mno"`
}

func (a *CreateServerArg) toEntityArgs(u *entity.User) (*entity.ServerCreateArg, error) {
	arg := &entity.ServerCreateArg{
		ProjectID:       u.ProjectID,
		ImageID:         a.ImageID,
		FlavorID:        a.FlavorID,
		SecurityGroupID: a.SecurityGroupID,
		VSwitchID:       a.VSwitchID,
		ServerName:      a.ServerName,
		HostName:        a.HostName,
		Description:     a.Description,
		Password:        a.Password,
		UserData:        a.UserData,
		KeyPairName:     a.KeyPairName,
		VPCID:           a.VPCID,
	}
	// 兼容前端空 keypair 的情况
	if a.KeyPairName != nil && *a.KeyPairName == "" {
		arg.KeyPairName = nil
	}
	if a.SystemDisk != nil {
		// 目前暂时不支持 非 SSD 磁盘创建, HDD 盘目前没有冗余
		if *a.SystemDisk.Category == common.CloudDiskCategory {
			return nil, errors.New("err disk category")
		}
		arg.SystemDisk = a.SystemDisk.toDiskArgs()
	}

	arg.DataDisks = make([]*entity.DiskArgs, len(a.DataDisks))
	for i, disk := range a.DataDisks {
		// 目前暂时不支持 非 SSD 磁盘创建, HDD 盘目前没有冗余
		if *disk.Category == common.CloudDiskCategory {
			return nil, errors.New("err disk category")
		}

		arg.DataDisks[i] = disk.toDiskArgs()
	}
	return arg, nil
}

func (a *CreateServerArg) serverCreateBaseCheck() error {
	if a.ImageID == "" {
		return errors.New("empty image id")
	}
	if a.FlavorName == "" && a.FlavorID == "" {
		return errors.New("empty flavor")
	}
	if a.SecurityGroupID == "" {
		return errors.New("empty security group id")
	}
	if a.VPCID == nil {
		return errors.New("empty vpc id")
	}
	if a.SystemDisk == nil {
		return errors.New("empty system disk")
	}
	if a.SystemDisk.Size <= 0 {
		return errors.New("err system disk size")
	}
	if !a.SystemDisk.Category.IsValid() {
		return errors.New("err system disk category")
	}
	if a.Password == nil {
		if a.KeyPairName == nil {
			return errors.New("password and Key Pair cannot be empty at the same time")
		}
		tp := ""
		a.Password = &tp
	}
	if a.DataDisks != nil {
		for _, disk := range a.DataDisks {
			if !disk.IsValid() {
				return errors.New("err data disk")
			}
		}
	}
	return nil
}

func (rc *RouterCompute) CreateServer(c *gin.Context) {
	var (
		arg CreateServerArg
		u   = common.GetUser(c)
	)
	if err := c.BindJSON(&arg); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	//if err := validate.Struct(&arg); err != nil {
	//	c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
	//	return
	//}
	if err := arg.serverCreateBaseCheck(); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	entityArgs, err := arg.toEntityArgs(u)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	resp, err := rc.ci.CreateServer(entityArgs)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.SuccessWith("", resp.ServerID))
}

type DeleteServerArgs struct {
	ServerID *string `json:"server_id"`
	Force    *bool   `json:"force"`
}

func (args *DeleteServerArgs) toEntityArgs(u *entity.User) *entity.ServerDeleteArg {
	return &entity.ServerDeleteArg{
		ProjectID: u.ProjectID,
		ServerID:  *args.ServerID,
		Force:     args.Force,
	}
}
func (rc *RouterCompute) DeleteServer(c *gin.Context) {
	var (
		arg DeleteServerArgs
		u   = common.GetUser(c)
	)
	if err := c.BindJSON(&arg); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	_, err := rc.ci.DeleteServer(arg.toEntityArgs(u))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.Success())
}

type UpdateServerArgs struct {
	ServerID string `json:"server_id"` // 实例ID。

	// 实例名称。长度为2~128个英文或中文字符。必须以大小字母或中文开头，不能以 http:// 和 https:// 开头。可以包含数字、半角冒号（:）、下划线（_）或者连字符（-）。
	ServerName *string `json:"server_name"`

	// 实例描述。长度为2~256个英文或中文字符，不能以 http:// 和 https:// 开头。
	Description *string `json:"description"`

	// 实例的密码。支持长度为8至30个字符，必须同时包含大小写英文字母、数字和特殊符号中的三类字符。特殊符号可以是：
	//
	// ```
	// ()`~!@#$%^&*-_+=|{}[]:;'<>,.?/
	// ````
	// 其中，Windows实例不能以斜线号（/）为密码首字符。
	//
	// @GSD:NOTE 说明 如果传入Password参数，建议您使用HTTPS协议发送请求，避免密码泄露。
	Password *string `json:"password"`

	// 操作系统的主机名。修改主机名后，请调用RebootServer使修改生效。
	// - Windows Server系统：长度为2-15个字符，允许使用大小写字母、数字或连字符（-）。不能以连字符（-）开头或结尾，不能连续使用连字符（-），也不能仅使用数字。
	// - 其他类型实例（Linux等）：长度为2-64个字符，允许使用点号（.）分隔字符成多段，每段允许使用大小写字母、数字或连字符（-），但不能连续使用点号（.）或连字符（-）。不能以点号（.）或连字符（-）开头或结尾。
	HostName *string `json:"host_name"`

	// 实例自定义数据，需要以Base64编码。
	// 编码前，原始数据不能超过16KB。建议不要明文传入敏感信息，例如密码和私钥等。如果必须传入敏感信息，建议您加密后再以Base64编码传入，在实例内部以同样的方式解密。
	UserData *string `json:"user_data"`
}

func (args *UpdateServerArgs) toEntityArgs(u *entity.User) *entity.ServerUpdateArg {
	var entityArg = &entity.ServerUpdateArg{
		ServerID:    args.ServerID,
		ProjectID:   u.ProjectID,
		ServerName:  args.ServerName,
		Description: args.Description,
		Password:    args.Password,
		HostName:    args.HostName,
		UserData:    args.UserData,
	}
	return entityArg
}

func (rc *RouterCompute) UpdateServer(c *gin.Context) {
	var (
		id  = c.Param("id")
		arg UpdateServerArgs
		u   = common.GetUser(c)
	)

	if err := c.BindJSON(&arg); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	arg.ServerID = id
	_, err := rc.ci.UpdateServer(arg.toEntityArgs(u))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.Success())
}

type GetServerArgs struct {
	ProjectID string `json:"project_id"` // 实例ID。
	ServerID  string `json:"server_id"`  // 实例ID。

}

func (args *GetServerArgs) toEntityArgs(u *entity.User) *entity.ServerGetArg {
	var entityArg = &entity.ServerGetArg{
		ServerID:  args.ServerID,
		ProjectID: u.ProjectID,
	}
	return entityArg
}

func (rc *RouterCompute) GetServer(c *gin.Context) {
	var (
		id  = c.Param("id")
		arg GetServerArgs
		u   = common.GetUser(c)
	)

	if err := c.BindJSON(&arg); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	arg.ServerID = id
	_, err := rc.ci.GetServer(arg.toEntityArgs(u))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.Success())
}

type ListServerArgs struct {
	ProjectID string `json:"project_id"` // 实例ID。
	ServerID  string `json:"server_id"`  // 实例ID。

}

func (args *ListServerArgs) toEntityArgs(u *entity.User) *entity.ServerListArg {
	var entityArg = &entity.ServerListArg{
		ProjectID: u.ProjectID,
	}
	return entityArg
}

func (rc *RouterCompute) ListServer(c *gin.Context) {
	var (
		id  = c.Param("id")
		arg ListServerArgs
		u   = common.GetUser(c)
	)

	if err := c.BindJSON(&arg); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	arg.ServerID = id
	_, err := rc.ci.ListServer(arg.toEntityArgs(u))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.Success())
}

type GetServerDisksArgs struct {
}

func (args *GetServerDisksArgs) toEntityArgs(u *entity.User) *entity.ServerDisksGetArg {
	return nil
}

func (rc *RouterCompute) GetServerDisks(c *gin.Context) {

}

func (rc *RouterCompute) StartServer(c *gin.Context) {
	var (
		id = c.Param("id")
		u  = common.GetUser(c)
	)
	_, err := rc.ci.StartServer(&entity.ServerStartArgs{
		ServerID:  id,
		ProjectID: u.ProjectID,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.Success())
}

type StopServerArgs struct {
	ForceStop *bool `json:"force_stop"`
}

func (rc *RouterCompute) StopServer(c *gin.Context) {
	var (
		id   = c.Param("id")
		u    = common.GetUser(c)
		args StopServerArgs
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	_, err := rc.ci.StopServer(&entity.ServerStopArgs{
		ServerID:  id,
		ProjectID: u.ProjectID,
		ForceStop: args.ForceStop,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.Success())
}

type RestartServerArgs struct {
	ForceRestart *bool `json:"force_restart"`
}

func (rc *RouterCompute) RestartServer(c *gin.Context) {
	var (
		id   = c.Param("id")
		u    = common.GetUser(c)
		args RestartServerArgs
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	_, err := rc.ci.RestartServer(&entity.ServerRestartArgs{
		ServerID:     id,
		ProjectID:    u.ProjectID,
		ForceRestart: args.ForceRestart,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.Success())
}
