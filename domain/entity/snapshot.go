package entity

import "github.com/arturiamu/lplms-public_cloud/common"

type Snapshot struct {
}

type SnapshotInfo struct {
	SnapshotID       string                // 快照ID
	SnapshotName     string                // 快照显示名称。如果创建时指定了快照显示名称，则返回。
	Usage            common.UsageType      // 快照是否被用作创建镜像或云盘。可能值：`image`, `disk`, `image_disk`, `none`
	Status           common.SnapshotStatus // 快照状态。可能值： `progressing`, `accomplished`, `failed`
	Progress         string                // 快照创建进度，单位为百分比。
	RemainTime       int64                 // 正在创建的快照剩余完成时间，单位为秒
	Description      string                // 描述信息
	SourceDiskID     string                // 源云盘ID。如果快照的源云盘已经被释放，该字段仍旧保留
	SourceDiskSize   int64                 // 源云盘容量，单位：GiB
	SourceDiskName   string                // 源云盘名称
	SourceDiskStatus common.DiskStatusType // 硬盘状态
	CreatedAt        int64                 // 创建时间。按照ISO8601标准表示，并使用UTC +0时间，格式为yyyy-MM-ddTHH:mm:ssZ。
}

type SnapshotCreateArg struct {

	// 项目ID
	ProjectID string

	// 云盘ID
	DiskID string

	// 快照的显示名称。长度为2~128个英文或中文字符。必须以大小字母或中文开头，不能以 `http://` 和 `https://` 开头。可以包含数字、半角冒号（:）、下划线（_）或者连字符（-）。
	//
	// 为防止和自动快照的名称冲突，不能以auto开头。
	SnapshotName string

	// 快照的描述。长度为2~256个英文或中文字符，不能以 `http://` 和 `https://` 开头。 默认值：空
	Description *string
}

type SnapshotDeleteArg struct {
	ProjectID  string // 项目ID
	SnapshotID string // 快照ID
	Force      *bool  // 是否强制删除有磁盘关联的快照。说明 删除后该磁盘无法重新初始化。
}

type SnapshotUpdateArg struct {
	// 项目ID
	ProjectID string

	SnapshotID string // 快照ID

	// 快照的显示名称。长度为2~128个英文或中文字符。必须以大小字母或中文开头，不能以 `http://` 和 `https://` 开头。可以包含数字、半角冒号（:）、下划线（_）或者连字符（-）。
	// 为防止和自动快照的名称冲突，不能以auto开头。
	SnapshotName *string

	// 快照的描述。长度为2~256个英文或中文字符，不能以 `http://` 和 `https://` 开头。 默认值：空
	Description *string
}

type SnapshotGetArg struct {
	ProjectID  string // 项目ID
	SnapshotID string // 快照ID
}

type SnapshotListArg struct {
	ProjectID string // 项目ID

	ServerID     *string  // 指定的实例ID
	DiskID       *string  // 指定的云盘设备ID
	SnapshotIDs  []string // 快照标识编码。取值可以由多个快照ID组成一个JSON数组，最多支持100个ID，ID之间用半角逗号（,）隔开。
	SnapshotName *string  // 快照名称

	// 快照状态。取值范围：
	// - progressing：正在创建的快照。
	// - accomplished：创建成功的快照。
	// - failed：创建失败的快照。
	// - all（默认）：所有快照状态。
	Status string

	// 快照源云盘的云盘类型。取值范围：
	// - System：系统盘
	// - Data：数据盘
	// @GSD:NOTE 说明 取值不区分大小写。
	SourceDiskType string

	// 快照是否被用作创建镜像或云盘。取值范围：
	// - image：使用快照创建了自定义镜像。
	// - disk：使用快照创建了云盘。
	// - image_disk：使用快照创建了数据盘和自定义镜像。
	// - none：暂未使用。
	Usage string

	// IsAdmin 标识是不是 admin 查询
	IsAdmin bool
}

type SnapshotCreateResp struct {
	SnapshotID string // 快照ID
}

type SnapshotDeleteResp struct{}

type SnapshotUpdateResp struct{}

type SnapshotGetResp struct {
	Snapshot SnapshotInfo
}

type SnapshotListResp struct {
	Snapshots []SnapshotInfo
}
