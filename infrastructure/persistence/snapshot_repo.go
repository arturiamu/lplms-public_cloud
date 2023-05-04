package persistence

import (
	"context"
	"errors"
	"github.com/arturiamu/lplms-public_cloud/common"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/arturiamu/lplms-public_cloud/utils/uuid"
	sv1 "github.com/kubernetes-csi/external-snapshotter/client/v4/apis/volumesnapshot/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
	"time"
)

// CreateSnapshot
// 调用CreateSnapshot为一块云盘创建一份快照。
func (s *StorageRepo) CreateSnapshot(args *entity.SnapshotCreateArg) (*entity.SnapshotCreateResp, error) {
	var (
		ns = args.ProjectID
		k  = s.k8Virt.KubernetesSnapshotClient().SnapshotV1().VolumeSnapshots(ns)
		id = uuid.GenerateUUID()
	)
	annotations := map[string]string{}

	annotations[common.AnnotationName] = args.SnapshotName
	if args.Description != nil {
		annotations[common.AnnotationDescription] = *args.Description
	}

	labels := make(map[string]string)
	labels[common.LabelDiskID] = args.DiskID
	createVolSnapshotArgs := &sv1.VolumeSnapshot{

		ObjectMeta: metav1.ObjectMeta{
			Name:        id,
			Annotations: annotations,
			Labels:      labels,
		},
		Spec: sv1.VolumeSnapshotSpec{
			Source: sv1.VolumeSnapshotSource{
				PersistentVolumeClaimName: &args.DiskID,
			},
		},
	}
	_, err := k.Create(context.Background(), createVolSnapshotArgs, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return &entity.SnapshotCreateResp{
		SnapshotID: id,
	}, nil
}

// DeleteSnapshot
// 调用DeleteSnapshot删除指定的快照。如果需要取消正在创建的快照，也可以调用该接口删除快照，即取消创建快照任务。
func (s *StorageRepo) DeleteSnapshot(args *entity.SnapshotDeleteArg) (*entity.SnapshotDeleteResp, error) {
	var (
		ns = args.ProjectID
		k  = s.k8Virt.KubernetesSnapshotClient().SnapshotV1().VolumeSnapshots(ns)
		vs = s.k8Virt.VirtualMachineSnapshot(ns)
	)

	snapshot, err := k.Get(context.Background(), args.SnapshotID, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	//获取vmsnapshot
	name := snapshot.GetAnnotations()[common.AnnotationVMSnapshotName]
	//name为空代表用户创建的快照
	if name == "" {
		err = k.Delete(context.Background(), args.SnapshotID, metav1.DeleteOptions{})
		return nil, err
	}

	vmsnapshot, er := vs.Get(context.Background(), name, metav1.GetOptions{})
	if er != nil {
		//找不到直接删除快照
		if k8serrors.IsNotFound(er) {
			err = k.Delete(context.Background(), args.SnapshotID, metav1.DeleteOptions{})
			return nil, err
		}
		err = er
		return nil, err
	}

	if vmsnapshot.GetAnnotations()[common.AnnotationHide] == "" {
		//镜像没有被删除了 不可以删除快照
		return nil, errors.New("delete snapshot error,can't delete Referenced Snapshots by image")
	}

	err = k.Delete(context.Background(), args.SnapshotID, metav1.DeleteOptions{})
	if err != nil {
		return nil, err
	}
	str := vmsnapshot.GetAnnotations()[common.AnnotationImageReference]
	cnt, err := strconv.Atoi(str)
	if err != nil {
		return nil, err
	}
	cnt--
	// 引用计数为0代表该镜像快照全部被删除可以释放该镜像
	if cnt == 0 {
		err = vs.Delete(context.Background(), name, metav1.DeleteOptions{})
		if err != nil {
			return nil, err
		}
	} else {
		//todo 并发不安全 写个定时任务删除未被释放的隐藏镜像
		vmsnapshot.GetAnnotations()[common.AnnotationImageReference] = strconv.Itoa(cnt)
		for i := 0; i < 5; i++ {
			_, err = vs.Update(context.Background(), vmsnapshot, metav1.UpdateOptions{})
			if err == nil {
				break
			}
			time.Sleep(time.Second << i)
		}

	}
	return nil, err
}

func (s *StorageRepo) UpdateSnapshot(args *entity.SnapshotUpdateArg) (*entity.SnapshotUpdateResp, error) {
	var (
		ns = args.ProjectID
		k  = s.k8Virt.KubernetesSnapshotClient().SnapshotV1().VolumeSnapshots(ns)
	)
	snapshot, err := k.Get(context.Background(), args.SnapshotID, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	modifyVolSnapshotArgs := snapshot.DeepCopy()
	annotations := modifyVolSnapshotArgs.ObjectMeta.Annotations

	if args.SnapshotName != nil {
		annotations[common.AnnotationName] = *args.SnapshotName
	}

	if args.Description != nil {
		annotations[common.AnnotationDescription] = *args.Description
	}
	_, err = k.Update(context.Background(), modifyVolSnapshotArgs, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// GetSnapshot
// 调用DescribeSnapshots查询一台ECS实例或一块云盘所有的快照列表。
func (s *StorageRepo) GetSnapshot(args *entity.SnapshotGetArg) (*entity.SnapshotGetResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *StorageRepo) ListSnapshot(args *entity.SnapshotListArg) (*entity.SnapshotListResp, error) {
	//TODO implement me
	panic("implement me")
}

///////////////////////////// help functions /////////////
