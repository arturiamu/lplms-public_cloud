package entity

type Disk struct {
	ImageID          string
	Device           string
	DetachedAt       int64
	DiskType         string
	Bootable         bool
	ServerID         string
	ServerName       string
	ZoneID           string
	AttachedTime     int64
	SourceSnapshotID string
	Size             int64
	Description      string
	Portable         bool
	DiskName         string
	CreatedAt        int64
	Status           string
	Category         string
	DeleteWithServer bool
	DiskID           string
	OSType           string
	OSName           string
	ImageName        string
}

type DiskCreateArg struct {

	// 项目ID
	ProjectID string

	// 容量大小，以GiB为单位。指定该参数后，其取值必须≥指定快照ID的容量大小
	Size *int64

	// 云盘名称。长度为2~128个英文或中文字符。必须以大小字母或中文开头，不能以 http:// 和 https:// 开头。可以包含数字、半角冒号（:）、下划线（_）或者连字符（-）。
	// 默认值：空
	DiskName *string

	// 云盘种类
	Category string

	// 云盘描述。长度为2~256个英文或中文字符，不能以 http:// 和 https:// 开头。
	// 默认值：空
	Description *string

	// 创建云盘使用的快照。
	SnapshotID *string

	// 创建云盘使用的镜像
	ImageID *string
}

type DiskDeleteArg struct {
	ProjectID string // 项目ID
	DiskID    string // 需要释放的云盘设备ID。
}

type DiskUpdateArg struct {
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
	// 在下列两种情况下，将参数DeleteWithServer设置成false时会报错。
	// - 磁盘的种类（category）为本地盘（ephemeral）时。
	// - 磁盘的种类（category）为普通云盘（cloud），且不可以卸载（Portable=false）时。
	DeleteWithServer *bool
}

type DiskGetArg struct {

	// 项目ID
	ProjectID string

	// 云盘ID
	DiskID string
}

type DiskListArg struct {

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
	DiskType string

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
	Category string

	// 云盘状态，详情参见云盘状态。取值范围：
	// - In_use
	// - Available
	// - Attaching
	// - Detaching
	// - Creating
	// - ReIniting
	// - All
	// 默认值：All
	Status string

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
}
