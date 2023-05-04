package persistence

import (
	"fmt"
	"github.com/arturiamu/lplms-public_cloud/common"
	"github.com/arturiamu/lplms-public_cloud/config"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/arturiamu/lplms-public_cloud/utils/uuid"
	"golang.org/x/net/context"
	"k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "kubevirt.io/api/core/v1"
	v1alpha2 "kubevirt.io/api/instancetype/v1alpha2"
	"strings"
)

var gpuFlavorDriverMap = map[string]string{
	"NVIDIA-V100-32G": "nvidia.com/Tesla_V100_32G",
	"NVIDIA-V100-16G": "nvidia.com/Tesla_V100",
	"NVIDIA-P4-8G":    "nvidia.com/P4",
	"NVIDIA-3090-24G": "nvidia.com/GA102_GEFORCE_RTX_3090",
}

func (c *ComputeRepo) CreateFlavor(arg *entity.FlavorCreateArg) (*entity.FlavorCreateResp, error) {
	var (
		ns = config.AppConfig.Cluster.BaseNamespace
		k  = c.k8Virt.VirtualMachineInstancetype(ns)
		id = uuid.GenerateUUID()
	)
	var gpus []corev1.GPU
	for i := 0; i < arg.GPUAmount; i++ {
		var gpu = corev1.GPU{
			Name:       fmt.Sprintf("%s_%d", arg.GPUSpec, i+1),
			DeviceName: arg.DeviceName,
		}
		gpus = append(gpus, gpu)
	}
	quantity, _ := resource.ParseQuantity(fmt.Sprintf("%dMi", arg.RAM))
	_, err := k.Create(context.Background(), &v1alpha2.VirtualMachineInstancetype{
		ObjectMeta: v1.ObjectMeta{
			Name: id,
			Annotations: map[string]string{
				common.AnnotationName: arg.Name,
			},
		},
		Spec: v1alpha2.VirtualMachineInstancetypeSpec{
			CPU:    v1alpha2.CPUInstancetype{Guest: uint32(arg.VCPUs)},
			Memory: v1alpha2.MemoryInstancetype{Guest: quantity},
			GPUs:   gpus,
		},
	}, v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return &entity.FlavorCreateResp{
		Flavor: entity.Flavor{
			FlavorID: id,
			Name:     arg.Name,
		},
	}, nil
}

func (c *ComputeRepo) DeleteFlavor(arg *entity.FlavorDeleteArg) (*entity.FlavorDeleteResp, error) {
	var (
		ns = config.AppConfig.Cluster.BaseNamespace
		k  = c.k8Virt.VirtualMachineInstancetype(ns)
	)
	err := k.Delete(context.Background(), arg.FlavorID, v1.DeleteOptions{})
	return nil, err
}

func (c *ComputeRepo) UpdateFlavor(arg *entity.FlavorUpdateArg) (*entity.FlavorUpdateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) GetFlavor(arg *entity.FlavorGetArg) (*entity.FlavorGetResp, error) {
	var (
		ns     = config.AppConfig.Cluster.BaseNamespace
		k      = c.k8Virt.VirtualMachineInstancetype(ns)
		flavor entity.Flavor
	)
	resp, err := k.List(context.Background(), v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	if resp != nil {
		for _, v := range resp.Items {
			if v.Name == arg.FlavorID {
				flavor = entity.Flavor{
					FlavorID:  v.ObjectMeta.GetName(),
					Name:      v.ObjectMeta.GetAnnotations()[common.AnnotationName],
					CPU:       uint(v.Spec.CPU.Guest),
					Memory:    uint(v.Spec.Memory.Guest.AsDec().UnscaledBig().Uint64() / 1024 / 1024),
					GPUAmount: uint(len(v.Spec.GPUs)),
					GPUSpec:   getGUPSpec(v.Spec.GPUs),
				}
				break
			}
		}
	}
	return &entity.FlavorGetResp{Flavor: flavor}, nil
}

func (c *ComputeRepo) ListFlavor(arg *entity.FlavorListArg) (*entity.FlavorListResp, error) {
	//TODO implement me
	panic("implement me")
}

///////////////////////////// help functions /////////////

func getGUPSpec(gpus []corev1.GPU) string {
	if gpus == nil || len(gpus) == 0 {
		return ""
	}
	//NVIDIA-P4-8G_x
	//NVIDIA-3090-24G_x
	//NVIDIA-V100-16G_x
	//NVIDIA-V100-32G_x
	name := strings.Split(gpus[0].Name, "_")
	return name[0]
}

func GetDriver(name string) string {
	return gpuFlavorDriverMap[name]
}
