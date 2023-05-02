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

	VPCAttributes      []*Vpc              // 专有网络VPC属性。
	DeviceAvailable    bool                // 实例是否可以挂载数据盘。
	SecurityGroupInfos []*SecurityGroup    // 实例所属安全组集合。
	HostName           string              // 实例主机名。
	FlavorID           string              // 规格 ID
	ServerType         string              // 实例规格。
	Status             common.ServerStatus // 实例状态。
	Disks              []*Disk             // 磁盘信息
	Recyclable         bool                // 实例是否可以回收。
	KeyPairName        string              // 密钥对名称。
	Tags               []string            // 实例的标签集合。
	CreatedAt          int64
}

type ServerCreateArg struct {
}

type ServerDeleteArg struct {
}

type ServerUpdateArg struct {
}

type ServerGetArg struct {
}

type ServerListArg struct {
}

type ServerCreateResp struct{}

type ServerDeleteResp struct{}

type ServerUpdateResp struct{}

type ServerGetResp struct{}

type ServerListResp struct{}
