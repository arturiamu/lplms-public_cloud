package entity

import "github.com/arturiamu/lplms-public_cloud/common"

type Disk struct {
	//ZoneID           enums.Zone
	ImageID          string
	Device           string
	DetachedAt       int64
	DiskType         *common.DiskType
	Bootable         bool
	ServerID         string
	ServerName       string
	AttachedTime     int64
	SourceSnapshotID string
	Size             int64
	Description      string
	Portable         bool
	DiskName         string
	CreatedAt        int64
	Status           common.DiskStatusType
	Category         common.DiskCategory
	DeleteWithServer bool
	DiskID           string
	OSType           common.OSType
	OSName           string
	ImageName        string
}

type DiskInfo struct {
	DiskID           string                `json:"disk_id"`
	DiskName         string                `json:"disk_name"`
	DiskCategory     common.DiskCategory   `json:"disk_category"`
	DiskType         common.DiskType       `json:"disk_type"`
	Device           string                `json:"device"`
	Size             int64                 `json:"size"`
	Status           common.DiskStatusType `json:"status"`
	SourceSnapshotID string                `json:"source_snapshot_id"`
	Server           DiskServerInfo        `json:"server"`
	ImageID          string                `json:"image_id"`
	OSType           common.OSType         `json:"os_type"`
	OSName           string                `json:"os_name"`
	ImageName        string                `json:"image_name"`
	Portable         bool                  `json:"portable"`
	Bootable         bool                  `json:"bootable"`
	AttachedTime     int64                 `json:"attached_time"`
	DeleteWithServer bool                  `json:"delete_with_server"`
	Description      string                `json:"description"`
	DetachedAt       int64                 `json:"detached_at"`
	CreatedAt        int64                 `json:"created_at"`
}

type DiskServerInfo struct {
	ServerID   string `json:"server_id"`
	ServerName string `json:"server_name"`
}

func BuildDiskInfo(disk Disk) *DiskInfo {
	diskInfo := &DiskInfo{
		DiskID:           disk.DiskID,
		DiskName:         disk.DiskName,
		DiskCategory:     disk.Category,
		Device:           disk.Device,
		Size:             disk.Size,
		Status:           disk.Status,
		SourceSnapshotID: disk.SourceSnapshotID,
		Server: DiskServerInfo{
			ServerID:   disk.ServerID,
			ServerName: disk.ServerName,
		},
		ImageID:          disk.ImageID,
		Portable:         disk.Portable,
		Bootable:         disk.Bootable,
		DeleteWithServer: disk.DeleteWithServer,
		AttachedTime:     disk.AttachedTime,
		Description:      disk.Description,
		DetachedAt:       disk.DetachedAt,
		CreatedAt:        disk.CreatedAt,
		OSType:           disk.OSType,
		OSName:           disk.OSName,
		ImageName:        disk.ImageName,
	}

	if disk.DiskType != nil {
		diskInfo.DiskType = *disk.DiskType
	}

	return diskInfo
}

type DiskCreateArg struct {

	// 在指定可用区内创建一块按量付费云盘。
	//ZoneID enums.Zone

	// 项目ID
	ProjectID string

	// 容量大小，以GiB为单位。指定该参数后，其取值必须≥指定快照ID的容量大小
	Size *int64

	// 云盘名称。长度为2~128个英文或中文字符。必须以大小字母或中文开头，不能以 http:// 和 https:// 开头。可以包含数字、半角冒号（:）、下划线（_）或者连字符（-）。
	//
	// 默认值：空
	DiskName *string

	// 云盘种类
	Category string

	// 云盘描述。长度为2~256个英文或中文字符，不能以 http:// 和 https:// 开头。
	//
	// 默认值：空
	Description *string

	// 创建云盘使用的快照。
	SnapshotID *string

	// 创建云盘使用的镜像
	ImageID *string

	// 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。
	ClientToken *string
}

type DiskDeleteArg struct {
	ProjectID string // 项目ID
	DiskID    string // 需要释放的云盘设备ID。
}

type DiskUpdateArg struct {
	// 所属的地域ID。您可以调用DescribeRegions查看最新的七牛云地域列表。
	//ZoneID enums.Zone

	// 项目ID
	ProjectID string

	// 待修改明细的磁盘ID。
	// @GSD:NOTE 说明 DiskId和DiskIds.N两个参数不能同时被调用，请您根据需求任选其一传值。
	DiskID string

	// 磁盘名称。长度为2~128个英文或中文字符。必须以大小字母或中文开头，不能以 `http://` 和 `https://` 开头。可以包含数字、半角冒号（:）、下划线（_）或者连字符（-）。
	DiskName *string

	// 磁盘描述。 长度为2~256个英文或中文字符，不能以 `http://` 和 `https://` 开头。
	Description *string

	// 磁盘是否随实例释放。默认值：无，无表示不改变当前的值。
	//
	// 在下列两种情况下，将参数DeleteWithServer设置成false时会报错。
	// - 磁盘的种类（category）为本地盘（ephemeral）时。
	DeleteWithServer *bool
	// - 磁盘的种类（category）为普通云盘（cloud），且不可以卸载（Portable=false）时。
}

type DiskGetArg struct {
	// 可用区ID。
	//ZoneID enums.Zone

	// 项目ID
	ProjectID string

	// 云盘ID
	DiskID string
}

type DiskListArg struct {

	// 可用区ID。
	//ZoneID enums.Zone

	// 项目ID
	ProjectID string

	// 云盘或本地盘ID。一个带有格式的JSON数组，最多支持100个ID，用半角逗号（,）隔开。
	DiskIDs []string

	// 云盘或本地盘挂载的实例ID。
	ServerID *string

	// 要查询的云盘或本地盘类型。取值范围：
	// - all：同时查询系统盘与数据盘。
	// - system：只查询系统盘。
	// - data：只查询数据盘。
	// 默认值：all

	DiskType *common.DiskType

	// 云盘或本地盘种类。取值范围：
	// - all：所有云盘以及本地盘
	// - cloud：普通云盘
	// - cloud_efficiency：高效云盘
	// - cloud_ssd：SSD盘
	// - cloud_essd：ESSD云盘
	// - local_ssd_pro：I/O密集型本地盘
	// - local_hdd_pro：吞吐密集型本地盘
	// - ephemeral：（已停售）本地盘
	// - ephemeral_ssd：（已停售）本地SSD盘
	// 默认值：all
	Category *common.DiskCategory

	// 云盘状态，详情参见云盘状态。取值范围：
	// - In_use
	// - Available
	// - Attaching
	// - Detaching
	// - Creating
	// - ReIniting
	// - All
	// 默认值：All
	Status *common.DiskStatusType

	// 创建云盘时使用的快照ID。
	SnapshotID *string

	// 云盘或本地盘名称。
	DiskName *string

	// 云盘是否设置了随实例释放。取值范围：
	// - true：云盘随实例一起释放。
	// - false：云盘保留不释放，转为按量付费数据盘而保留下来。
	// 默认值：false
	DeleteWithServer *bool

	// IsAdmin 标识是否是 admin 查询
	IsAdmin bool

	common.Pagination // 分页信息

}

type DiskCreateResp struct {
	DiskID string // 云盘ID。
}

type DiskDeleteResp struct{}

type DiskUpdateResp struct{}

type DiskGetResp struct {
	Disk Disk
}

type DiskListResp struct {
	Disks []Disk
}

type DiskAttachArg struct {
	// 项目 ID
	ProjectID string

	// 待挂载的云盘 ID。云盘（DiskID）和实例（ServerID）必须在同一个可用区。
	DiskID string

	// 目标ECS实例的ID。
	ServerID string

	// 释放实例时，该云盘是否随实例一起释放。
	// - true：释放。
	// - false：不释放。云盘会转换成按量付费数据盘而被保留下来。
	// 默认值：false
	DeleteWithServer *bool

	// 是否作为系统盘挂载。 默认值：false
	//
	// @GSD:NOTE 说明 设置为 `Bootable=true` 时，目标ECS实例必须处于无系统盘状态。
	Bootable *bool

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
	Password *string

	// 密钥对名称。
	// 挂载系统盘时，为Linux系统ECS实例绑定的SSH密钥对的名称。
	//
	// @GSD:NOTE 说明 Windows Server系统：不支持SSH密钥对。即使填写了该参数，只执行Password的配置。Linux系统：密码登录方式会被初始化成禁止。
	KeyPairName *string
}

type DiskAttachResp struct{}

type DiskDetachArg struct {
	ProjectID string // 项目 ID
	ServerID  string // 待卸载的ECS实例ID。
	DiskID    string // 待卸载的云盘ID。
}

type DiskDetachResp struct{}

type ResizeDiskType string

var (
	ResizeDiskTypeOffline ResizeDiskType = "offline"
	ResizeDiskTypeOnline  ResizeDiskType = "online"
)

type DiskResizeArg struct {
	// 项目ID
	ProjectID string

	// 云盘ID。
	DiskID string

	// 希望扩容到的云盘容量大小。单位为GiB。取值范围：
	// - 高效云盘（cloud_efficiency）：20~32768
	// - SSD云盘（cloud_ssd）：20~32768
	// - ESSD云盘（cloud_essd）：20~32768
	// - 普通云盘（cloud）：5~2000
	// 指定的新云盘容量必须比原云盘容量大。
	NewSize int64

	// 扩容云盘的方式。取值范围：
	// - **offline**（默认）：离线扩容。扩容后，您必须在控制台重启实例或者调用API RebootInstance使操作生效。
	// - **online**：在线扩容，无需重启实例即可完成扩容。云盘类型支持高效云盘、SSD云盘和ESSD云盘。
	Type *ResizeDiskType
}

type DiskResizeResp struct{}

type DiskResetArg struct {
	ProjectID  string // 项目ID
	DiskID     string // 指定的磁盘设备ID。
	SnapshotID string // 需要恢复到某一磁盘阶段的历史快照ID。
}

type DiskResetResp struct{}
