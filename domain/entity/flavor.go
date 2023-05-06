package entity

type Flavor struct {
	FlavorID             string // 实例规格ID
	Name                 string // 规格名称
	Description          string // 规格描述
	CPU                  uint   // vCPU内核数目
	Memory               uint   // 内存大小，单位 MB
	Swap                 uint   // 交换分区大小, 单位 MB
	GPUAmount            uint   // GPU数量 2 3 4
	GPUSpec              string // GPU类型 NVIDIA-P4-8G NVIDIA-3090-24G NVIDIA-V100-16G
	Disk                 uint   // 系统盘大小 GB
	BandwidthRX          uint64 // 内网入方向带宽限制，单位：kbit/s
	BandwidthTX          uint64 // 内网出方向带宽限制，单位：kbit/s
	PPSRX                uint   // 内网入方向网络收发包能力，单位：Pps
	PPSTX                uint   // 内网出方向网络收发包能力，单位：Pps
	LocalStorageAmount   uint   // 实例挂载的本地盘的数量
	LocalStorageCapacity uint   // 实例挂载的本地盘的单盘容量，单位：GiB
	LocalStorageCategory string // TODO(q) 补充 enums
	IsSuspend            bool   // 规格是否下架

	// 本地盘类型，详情请参见本地盘。可能值：
	// - local_hdd_pro：实例规格族d1ne和d1搭载的SATA HDD本地盘。
	// - local_ssd_pro：实例规格族i2、i2g、i1、ga1和gn5等搭载的NVMe SSD本地盘。
}

type FlavorCreateArg struct {
	Name       string // 名称
	VCPUs      int    // 虚拟 CPU 核数
	RAM        int    // 内存，单位：MiB
	Disk       int    // 磁盘，单位：GiB
	GPUAmount  int    // GPU数量 2 3 4
	GPUSpec    string // GPU类型 NVIDIA-P4-8G NVIDIA-3090-24G NVIDIA-V100-16G
	DeviceName string // GPU DeviceName nvidia.com/p4
}

type FlavorDeleteArg struct {
	FlavorID string // 实例规格ID
}

type FlavorUpdateArg struct {
	FlavorID string
}

type FlavorGetArg struct {
	FlavorID string
}

type FlavorListArg struct {
}

type FlavorCreateResp struct {
	Flavor Flavor
}

type FlavorDeleteResp struct{}

type FlavorUpdateResp struct{}

type FlavorGetResp struct {
	Flavor Flavor
}

type FlavorListResp struct {
	Flavors []Flavor
}
