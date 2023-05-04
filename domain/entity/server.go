package entity

import "github.com/arturiamu/lplms-public_cloud/common"

type Server struct {
	//ZoneID                  common.Zone    // 实例所属可用区。
	ServerID                string        // 实例ID。
	ServerName              string        // 实例名称。
	Description             string        // 实例描述。
	CPU                     int64         // vCPU数。
	Memory                  int64         // 内存大小，单位MiB。
	GPUSpec                 string        // 实例规格附带的GPU类型。
	GPUAmount               int64         // 实例规格附带的GPU数量。
	OSType                  common.OSType // 实例的操作系统类型，分为Windows Server和Linux两种。可能值：  - windows - linux
	OSName                  string        // 实例的操作系统名称。
	ImageID                 string        // 实例运行的镜像ID。
	ImageName               string        //  镜像名称
	VlanID                  string        // 实例的VLAN ID。 @GSD:NOTE 说明 该参数即将被弃用，为提高兼容性，请尽量使用其他参数。
	InternetMaxBandwidthIn  int           // 公网入带宽最大值，单位为Mbit/s。
	InternetMaxBandwidthOut int           // 公网出带宽最大值，单位为Mbit/s。

	VPCAttributes      []*VPCAttribute      // 专有网络VPC属性。
	DeviceAvailable    bool                 // 实例是否可以挂载数据盘。
	SecurityGroupInfos []*SecurityGroupInfo // 实例所属安全组集合。
	HostName           string               // 实例主机名。
	FlavorID           string               // 规格 ID
	ServerType         string               // 实例规格。
	Status             common.ServerStatus  // 实例状态。
	Disks              []*DiskInfo          // 磁盘信息
	Recyclable         bool                 // 实例是否可以回收。
	KeyPairName        string               // 密钥对名称。
	Tags               []string             // 实例的标签集合。
	CreatedAt          int64
}

type ServerCreateArg struct {
	//ZoneID          enums.Zone // 实例所属的可用区 ID，您可以调用 DescribeZones 获取可用区列表。默认值：系统随机选择。
	ProjectID       string // 所属项目 ID
	ImageID         string // 镜像 ID，启动实例时选择的镜像资源。您可以通过 DescribeImages 查询您可以使用的镜像资源。
	FlavorID        string // 实例的资源规格ID。  Flavor ID 和 Flavor Name 二选一必传 如果两个都传优先使用 Flavor ID
	SecurityGroupID string // 指定新创建实例所属于的安全组 ID。同一个安全组内的实例之间可以互相访问，一个安全组能容纳的实例数量视安全组类型而定，具体请参见使用限制的安全组章节。

	// 虚拟交换机ID。如果您创建的是VPC类型ECS实例，必须指定虚拟交换机ID，且安全组和虚拟交换机在同一个专有网络VPC中。
	VSwitchID *string

	// 实例的名称。长度为2~128个英文或中文字符。必须以大小字母或中文开头，不能以 http:// 和 https:// 开头。
	// 可以包含数字、半角冒号（:）、下划线（_）、英文句号（.）或者连字符（-）。
	// 如果没有指定该参数，默认值为实例的ServerId。
	ServerName *string

	// 云服务器的主机名。
	// - 英文句号（.）和短横线（-）不能作为首尾字符，更不能连续使用。
	// - Windows实例：字符长度为2~15，不支持英文句号（.），不能全是数字。允许大小写英文字母、数字和短横线（-）。
	// - 其他类型实例（Linux等）：字符长度为2~64，支持多个英文句号（.），英文句号之间为一段，每段允许大小写英文字母、数字和短横线（-）。
	HostName *string

	// 实例的描述。长度为 2~256 个英文或中文字符，不能以 http:// 和 https:// 开头。
	Description *string

	// 实例的密码。长度为 8 至 30 个字符，必须同时包含大小写英文字母、数字和特殊符号中的三类字符。特殊符号可以是：
	//
	// ```
	// ()`~!@#$%^&*-_+=|{}[]:;'<>,.?/
	// ```
	// 其中，Windows实例不能以斜线号（/）为密码首字符。
	//
	// @GSE:NOTE: 如果传入 Password 参数，建议您使用 HTTPS 协议发送请求，避免密码泄露。
	Password *string

	// 实例自定义数据，需要以 Base64 方式编码，原始数据最多为16KB。
	UserData *string

	// 密钥对名称。
	// - Windows实例，忽略该参数。默认为空。即使填写了该参数，仍旧只执行Password的内容。
	// - Linux实例的密码登录方式会被初始化成禁止。为提高实例安全性，强烈建议您使用密钥对的连接方式。
	KeyPairName *string

	VPCID *string // VPC ID

	Host *string // 指定云主机创建在哪台物理机 这个 host 必须是浪潮平台上的物理节点

	SystemDisk *DiskArgs   // 系统盘
	DataDisks  []*DiskArgs // 数据盘

	//ClientToken *string // 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。
}

type DiskInfo struct {
	ID                  string
	Size                int    // GB
	IsBoot              bool   // 是否是系统盘
	DeleteOnTermination bool   // 是否随主机释放
	DiskCategory        string // 磁盘类型
	DiskType            *common.DiskType
}

type SecurityGroupInfo struct {
	ID   string
	Name string
}

type DiskArgs struct {
	Size             int
	Category         string
	SnapshotID       *string
	DiskName         *string
	Description      *string
	DeleteWithServer *bool
}

type ServerDeleteArg struct {
	ServerID  string // 实例ID。N的取值范围：1~100
	ProjectID string

	// 是否强制释放运行中（Running）的ECS实例。
	//
	// - true：强制释放运行中（Running）的实例。强制释放相当于断电，实例内存以及存储中的临时数据都会被擦除，无法恢复。
	// - false（默认值）：正常释放实例，此时实例必须处于已停止（Stopped）状态。
	Force *bool
}

type ServerUpdateArg struct {
	ProjectID   string
	ServerID    string  // 指定的实例ID。
	FlavorID    string  // 实例的目标规格。更多详情，请参见实例规格族，也可以调用DescribeServerTypes接口获得最新的规格表。
	ClientToken *string // 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。
}

type ServerGetArg struct {
	//ZoneID    enums.Zone // 实例所属可用区。
	ProjectID string // 项目 ID
	ServerID  string // 一个或ECS实例ID。N的取值范围：1~100，多个取值使用重复列表的形式。
}

type ServerListArg struct {
	ServerIDs          []string             // 实例 ID 列表。最多支持100个ID
	VPCID              *string              // 专有网络VPC ID。
	VSwitchID          *string              // 交换机 ID。
	FlavorID           *string              // 实例的规格 ID。
	PrivateIPAddresses []string             // VPC网络类型实例的私有IP。当ServerNetworkType=vpc时生效，取值可以由多个IP组成一个JSON数组，最多支持100个IP，IP之间用半角逗号（,）隔开。
	PublicIPAddresses  []string             // 实例的公网IP列表。取值可以由多个IP组成一个JSON数组，最多支持100个IP，IP之间用半角逗号（,）隔开。
	SecurityGroupID    *string              // 实例所属的安全组。
	ServerName         *string              // 实例名称，支持使用通配符*进行模糊搜索。
	ImageID            *string              // 镜像 ID。
	Status             *common.ServerStatus // 实例状态。取值范围： - Pending：创建中 - Running：运行中 - Starting：启动中 - Stopping：停止中 - Stopped：已停止
	SortKey            *string              // 排序字段, 前面+ -号表示排序顺序
	ProjectID          string
}

type ServerCreateResp struct {
	ServerID string
}

type ServerDeleteResp struct{}

type ServerUpdateResp struct{}

type ServerGetResp struct {
	Server Server
}

type ServerListResp struct {
	Servers []*Server // 由 Servers 组成的数组格式，返回实例的信息。
}

type ServerDisksGetArg struct {
	ProjectID string
	ServerID  string // 实例 ID
}

type ServerDisksGetResp struct {
	Disks []*DiskInfo
}
