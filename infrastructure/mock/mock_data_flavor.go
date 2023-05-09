package mock

import (
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/arturiamu/lplms-public_cloud/utils/uuid"
)

var mockFlavorMap map[string]entity.Flavor
var mockFlavorList = []entity.Flavor{
	{
		FlavorID:             uuid.GenerateUUID(),        // 实例规格ID
		Name:                 "标准型S1",                    // 规格名称
		Description:          "适合开发者、小型业务需求场景。共享计算和网络资源", // 规格描述
		CPU:                  1,                          // vCPU内核数目
		Memory:               2 * 1024,                   // 内存大小，单位 MB
		Swap:                 0,                          // 交换分区大小, 单位 MB
		GPUAmount:            0,                          // GPU数量 2 3 4
		GPUSpec:              "",                         // GPU类型 NVIDIA-P4-8G NVIDIA-3090-24G NVIDIA-V100-16G
		Disk:                 20,                         // 系统盘大小 GB
		BandwidthRX:          0,                          // 内网入方向带宽限制，单位：kbit/s
		BandwidthTX:          0,                          // 内网出方向带宽限制，单位：kbit/s
		PPSRX:                0,                          // 内网入方向网络收发包能力，单位：Pps
		PPSTX:                0,                          // 内网出方向网络收发包能力，单位：Pps
		LocalStorageAmount:   0,                          // 实例挂载的本地盘的数量
		LocalStorageCapacity: 0,                          // 实例挂载的本地盘的单盘容量，单位：GiB
		LocalStorageCategory: "",                         // TODO(q) 补充 enums
		IsSuspend:            false,                      // 规格是否下架
	},
	{
		FlavorID:             uuid.GenerateUUID(), // 实例规格ID
		Name:                 "标准型S2",             // 规格名称
		Description:          "适合开发者、小型业务需求场景。",   // 规格描述
		CPU:                  2,                   // vCPU内核数目
		Memory:               2 * 1024,            // 内存大小，单位 MB
		Swap:                 0,                   // 交换分区大小, 单位 MB
		GPUAmount:            0,                   // GPU数量 2 3 4
		GPUSpec:              "",                  // GPU类型 NVIDIA-P4-8G NVIDIA-3090-24G NVIDIA-V100-16G
		Disk:                 20,                  // 系统盘大小 GB
		BandwidthRX:          0,                   // 内网入方向带宽限制，单位：kbit/s
		BandwidthTX:          0,                   // 内网出方向带宽限制，单位：kbit/s
		PPSRX:                0,                   // 内网入方向网络收发包能力，单位：Pps
		PPSTX:                0,                   // 内网出方向网络收发包能力，单位：Pps
		LocalStorageAmount:   0,                   // 实例挂载的本地盘的数量
		LocalStorageCapacity: 0,                   // 实例挂载的本地盘的单盘容量，单位：GiB
		LocalStorageCategory: "",                  // TODO(q) 补充 enums
		IsSuspend:            false,               // 规格是否下架
	},
	{
		FlavorID:             uuid.GenerateUUID(), // 实例规格ID
		Name:                 "标准型S2",             // 规格名称
		Description:          "入门计算",              // 规格描述
		CPU:                  2,                   // vCPU内核数目
		Memory:               4 * 1024,            // 内存大小，单位 MB
		Swap:                 0,                   // 交换分区大小, 单位 MB
		GPUAmount:            0,                   // GPU数量 2 3 4
		GPUSpec:              "",                  // GPU类型 NVIDIA-P4-8G NVIDIA-3090-24G NVIDIA-V100-16G
		Disk:                 25,                  // 系统盘大小 GB
		BandwidthRX:          0,                   // 内网入方向带宽限制，单位：kbit/s
		BandwidthTX:          0,                   // 内网出方向带宽限制，单位：kbit/s
		PPSRX:                0,                   // 内网入方向网络收发包能力，单位：Pps
		PPSTX:                0,                   // 内网出方向网络收发包能力，单位：Pps
		LocalStorageAmount:   0,                   // 实例挂载的本地盘的数量
		LocalStorageCapacity: 0,                   // 实例挂载的本地盘的单盘容量，单位：GiB
		LocalStorageCategory: "",                  // TODO(q) 补充 enums
		IsSuspend:            false,               // 规格是否下架
	},
	{
		FlavorID:             uuid.GenerateUUID(),         // 实例规格ID
		Name:                 "标准型S4",                     // 规格名称
		Description:          "适合企业偏重计算场景，CPU依赖高，如Nginx等", // 规格描述
		CPU:                  4,                           // vCPU内核数目
		Memory:               4 * 1024,                    // 内存大小，单位 MB
		Swap:                 0,                           // 交换分区大小, 单位 MB
		GPUAmount:            0,                           // GPU数量 2 3 4
		GPUSpec:              "",                          // GPU类型 NVIDIA-P4-8G NVIDIA-3090-24G NVIDIA-V100-16G
		Disk:                 40,                          // 系统盘大小 GB
		BandwidthRX:          0,                           // 内网入方向带宽限制，单位：kbit/s
		BandwidthTX:          0,                           // 内网出方向带宽限制，单位：kbit/s
		PPSRX:                0,                           // 内网入方向网络收发包能力，单位：Pps
		PPSTX:                0,                           // 内网出方向网络收发包能力，单位：Pps
		LocalStorageAmount:   0,                           // 实例挂载的本地盘的数量
		LocalStorageCapacity: 0,                           // 实例挂载的本地盘的单盘容量，单位：GiB
		LocalStorageCategory: "",                          // TODO(q) 补充 enums
		IsSuspend:            false,                       // 规格是否下架
	},
	{
		FlavorID:             uuid.GenerateUUID(),         // 实例规格ID
		Name:                 "标准型S4",                     // 规格名称
		Description:          "适合企业偏重计算场景，CPU依赖高，如Nginx等", // 规格描述
		CPU:                  4,                           // vCPU内核数目
		Memory:               8 * 1024,                    // 内存大小，单位 MB
		Swap:                 0,                           // 交换分区大小, 单位 MB
		GPUAmount:            0,                           // GPU数量 2 3 4
		GPUSpec:              "",                          // GPU类型 NVIDIA-P4-8G NVIDIA-3090-24G NVIDIA-V100-16G
		Disk:                 40,                          // 系统盘大小 GB
		BandwidthRX:          0,                           // 内网入方向带宽限制，单位：kbit/s
		BandwidthTX:          0,                           // 内网出方向带宽限制，单位：kbit/s
		PPSRX:                0,                           // 内网入方向网络收发包能力，单位：Pps
		PPSTX:                0,                           // 内网出方向网络收发包能力，单位：Pps
		LocalStorageAmount:   0,                           // 实例挂载的本地盘的数量
		LocalStorageCapacity: 0,                           // 实例挂载的本地盘的单盘容量，单位：GiB
		LocalStorageCategory: "",                          // TODO(q) 补充 enums
		IsSuspend:            false,                       // 规格是否下架
	},

	{
		FlavorID:             uuid.GenerateUUID(),      // 实例规格ID
		Name:                 "异构型H1",                  // 规格名称
		Description:          "适合入门级科学计算，需要基本GPU能力的场景", // 规格描述
		CPU:                  4,                        // vCPU内核数目
		Memory:               8 * 1024,                 // 内存大小，单位 MB
		Swap:                 0,                        // 交换分区大小, 单位 MB
		GPUAmount:            1,                        // GPU数量 2 3 4
		GPUSpec:              "NVIDIA-P4-8G",           // GPU类型 NVIDIA-P4-8G NVIDIA-3090-24G NVIDIA-V100-16G
		Disk:                 40,                       // 系统盘大小 GB
		BandwidthRX:          0,                        // 内网入方向带宽限制，单位：kbit/s
		BandwidthTX:          0,                        // 内网出方向带宽限制，单位：kbit/s
		PPSRX:                0,                        // 内网入方向网络收发包能力，单位：Pps
		PPSTX:                0,                        // 内网出方向网络收发包能力，单位：Pps
		LocalStorageAmount:   0,                        // 实例挂载的本地盘的数量
		LocalStorageCapacity: 0,                        // 实例挂载的本地盘的单盘容量，单位：GiB
		LocalStorageCategory: "",                       // TODO(q) 补充 enums
		IsSuspend:            false,                    // 规格是否下架
	},
	{
		FlavorID:             uuid.GenerateUUID(),      // 实例规格ID
		Name:                 "异构型H2",                  // 规格名称
		Description:          "适合入门级科学计算，需要基本GPU能力的场景", // 规格描述
		CPU:                  8,                        // vCPU内核数目
		Memory:               16 * 1024,                // 内存大小，单位 MB
		Swap:                 0,                        // 交换分区大小, 单位 MB
		GPUAmount:            2,                        // GPU数量 2 3 4
		GPUSpec:              "NVIDIA-P4-8G",           // GPU类型 NVIDIA-P4-8G NVIDIA-3090-24G NVIDIA-V100-16G
		Disk:                 60,                       // 系统盘大小 GB
		BandwidthRX:          0,                        // 内网入方向带宽限制，单位：kbit/s
		BandwidthTX:          0,                        // 内网出方向带宽限制，单位：kbit/s
		PPSRX:                0,                        // 内网入方向网络收发包能力，单位：Pps
		PPSTX:                0,                        // 内网出方向网络收发包能力，单位：Pps
		LocalStorageAmount:   0,                        // 实例挂载的本地盘的数量
		LocalStorageCapacity: 0,                        // 实例挂载的本地盘的单盘容量，单位：GiB
		LocalStorageCategory: "",                       // TODO(q) 补充 enums
		IsSuspend:            false,                    // 规格是否下架
	},
	{
		FlavorID:             uuid.GenerateUUID(),      // 实例规格ID
		Name:                 "异构型H2",                  // 规格名称
		Description:          "适合入门级科学计算，需要基本GPU能力的场景", // 规格描述
		CPU:                  8,                        // vCPU内核数目
		Memory:               16 * 1024,                // 内存大小，单位 MB
		Swap:                 0,                        // 交换分区大小, 单位 MB
		GPUAmount:            2,                        // GPU数量 2 3 4
		GPUSpec:              "NVIDIA-3090-24G",        // GPU类型 NVIDIA-P4-8G NVIDIA-3090-24G NVIDIA-V100-16G
		Disk:                 60,                       // 系统盘大小 GB
		BandwidthRX:          0,                        // 内网入方向带宽限制，单位：kbit/s
		BandwidthTX:          0,                        // 内网出方向带宽限制，单位：kbit/s
		PPSRX:                0,                        // 内网入方向网络收发包能力，单位：Pps
		PPSTX:                0,                        // 内网出方向网络收发包能力，单位：Pps
		LocalStorageAmount:   0,                        // 实例挂载的本地盘的数量
		LocalStorageCapacity: 0,                        // 实例挂载的本地盘的单盘容量，单位：GiB
		LocalStorageCategory: "",                       // TODO(q) 补充 enums
		IsSuspend:            false,                    // 规格是否下架
	},
	{
		FlavorID:             uuid.GenerateUUID(),         // 实例规格ID
		Name:                 "异构型H5",                     // 规格名称
		Description:          "适合企业级科学计算，AI等需要大量GPU能力的场景", // 规格描述
		CPU:                  50,                          // vCPU内核数目
		Memory:               100 * 1024,                  // 内存大小，单位 MB
		Swap:                 0,                           // 交换分区大小, 单位 MB
		GPUAmount:            4,                           // GPU数量 2 3 4
		GPUSpec:              "NVIDIA-3090-24G",           // GPU类型 NVIDIA-P4-8G NVIDIA-3090-24G NVIDIA-V100-16G
		Disk:                 120,                         // 系统盘大小 GB
		BandwidthRX:          0,                           // 内网入方向带宽限制，单位：kbit/s
		BandwidthTX:          0,                           // 内网出方向带宽限制，单位：kbit/s
		PPSRX:                0,                           // 内网入方向网络收发包能力，单位：Pps
		PPSTX:                0,                           // 内网出方向网络收发包能力，单位：Pps
		LocalStorageAmount:   0,                           // 实例挂载的本地盘的数量
		LocalStorageCapacity: 0,                           // 实例挂载的本地盘的单盘容量，单位：GiB
		LocalStorageCategory: "",                          // TODO(q) 补充 enums
		IsSuspend:            false,                       // 规格是否下架
	},
	{
		FlavorID:             uuid.GenerateUUID(),         // 实例规格ID
		Name:                 "异构型H10",                    // 规格名称
		Description:          "适合企业级科学计算，AI等需要大量GPU能力的场景", // 规格描述
		CPU:                  100,                         // vCPU内核数目
		Memory:               200 * 1024,                  // 内存大小，单位 MB
		Swap:                 0,                           // 交换分区大小, 单位 MB
		GPUAmount:            8,                           // GPU数量 2 3 4
		GPUSpec:              "NVIDIA-3090-24G",           // GPU类型 NVIDIA-P4-8G NVIDIA-3090-24G NVIDIA-V100-16G
		Disk:                 250,                         // 系统盘大小 GB
		BandwidthRX:          0,                           // 内网入方向带宽限制，单位：kbit/s
		BandwidthTX:          0,                           // 内网出方向带宽限制，单位：kbit/s
		PPSRX:                0,                           // 内网入方向网络收发包能力，单位：Pps
		PPSTX:                0,                           // 内网出方向网络收发包能力，单位：Pps
		LocalStorageAmount:   0,                           // 实例挂载的本地盘的数量
		LocalStorageCapacity: 0,                           // 实例挂载的本地盘的单盘容量，单位：GiB
		LocalStorageCategory: "",                          // TODO(q) 补充 enums
		IsSuspend:            false,                       // 规格是否下架
	},
}

func initFlavor() {
	mockFlavorMap = make(map[string]entity.Flavor)
	for _, flavor := range mockFlavorList {
		mockFlavorMap[flavor.FlavorID] = flavor
	}
}

func getFlavorList() []entity.Flavor {
	return mockFlavorList
}

func getFlavor(id string) entity.Flavor {
	return mockFlavorMap[id]
}
