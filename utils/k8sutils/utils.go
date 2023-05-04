package k8sutils

import (
	"github.com/arturiamu/lplms-public_cloud/common"
	kmetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetBizName(o *kmetav1.ObjectMeta) string {
	return o.Annotations[common.AnnotationName]
}

func GetDesc(o *kmetav1.ObjectMeta) string {
	return o.Annotations[common.AnnotationDescription]
}
