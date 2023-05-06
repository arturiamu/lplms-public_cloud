package entity

import "github.com/arturiamu/lplms-public_cloud/common"

type Image struct {
	//ZoneID             enums.Zone
	ImageID            string                  // 镜像 ID
	SnapshotID         string                  // 快照 ID
	Description        string                  // 镜像描述
	OSType             common.OSType           // 操作系统类型
	Architecture       common.ArchitectureType // 镜像系统架构类型
	OSName             string                  // 操作系统的中文显示名称
	Progress           string                  // 对于导入中的镜像，返回导入任务的进度
	ImageVersion       string                  // 镜像版本
	Status             common.ImageStatus      // 镜像的状态
	ImageName          string                  // 镜像名称
	ImageOwnerAlias    common.ImageOwnerAlias  // 镜像来源
	Platform           common.PlatformType     // 操作系统平台
	Size               int64                   // 镜像大小，单位：Bytes
	IsSupportCloudInit bool                    // 镜像是否支持cloud-init
	IsCopied           bool                    // 是否是拷贝的镜像
	CreatedAt          int64                   // 镜像的创建时间
	//ImageDeviceMappings []*ImageDeviceMapping   // 镜像下包含磁盘和快照的映射关系
	MinImageSize int64 // 镜像要求的最小系统盘容量，单位：Bytes
}

type ImageCreateArg struct {
	// 项目ID
	ProjectID string

	// 根据指定的快照创建自定义镜像
	SnapshotID string

	// 实例 ID
	ServerID string

	// 镜像名称。长度为2~128个英文或中文字符。必须以大小字母或中文开头，不能以 `http://` 和 `https://` 开头。可以包含数字、半角冒号（:）、下划线（_）或者连字符（-）。
	// 默认值：空
	ImageName *string

	// 镜像的描述信息。长度为2~256个英文或中文字符，不能以 `http://` 和 `https://` 开头。
	// 默认值：空
	Description *string

	// 镜像格式
	ImageFormat string
}

type ImageDeleteArg struct {

	// 项目 ID
	ProjectID string

	// 镜像ID。如果指定的自定义镜像不存在，则请求将被忽略。
	ImageID string

	// 是否执行强制删除。取值范围：
	// - true：强制删除自定义镜像，忽略当前镜像是否被其他实例使用。
	// - false：正常删除自定义镜像，删除前检查当前镜像是否被其他实例使用。
	// 默认值：false
	Force *bool

	// 是否删除相关快照，默认不删除
	DeleteRelatedSnapshot *bool
}

type ImageUpdateArg struct {
	ImageID string // 自定义镜像的 ID。

	// 项目 ID
	ProjectID string

	ImageName *string // 自定义镜像的名称。能包含2~128个字符。必须以大小字母或中文开头，可包含数字、半角冒号（:）、下划线（_）或者连字符（-）。不能以 `http://` 和 `https://` 开头。 默认值：空，表示保持原有名称不变。

	// 自定义镜像的描述信息。能包含2~256个字符。不能以 `http://` 和 `https://` 开头。默认值：空，表示保持原有描述信息不变。
	Description *string

	// 镜像版本。必须使用镜像名称开头，大版本号使用点（.）连接小版本号。例如：CentOS6.9
	Version *string
}

type ImageGetArg struct {
}

type ImageListArg struct {
	// 项目 ID
	ProjectID string

	// 镜像ID。
	ImageID string

	// 根据某一快照ID创建的自定义镜像。
	SnapshotID string

	// 状态
	Status []common.ImageStatus

	// 镜像名称。
	ImageName string

	// 镜像来源
	ImageOwnerAlias common.ImageOwnerAlias

	// 指定实例类型可以使用的镜像。
	ServerType string

	// 镜像的操作系统类型
	OSType common.OSType

	// 镜像的体系架构
	Architecture common.ArchitectureType
}

type ImageCreateResp struct{}

type ImageDeleteResp struct{}

type ImageUpdateResp struct{}

type ImageGetResp struct{}

type ImageListResp struct{}
