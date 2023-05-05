package common

import kmetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

func GetBizName(o *kmetav1.ObjectMeta) string {
	return o.Annotations[AnnotationName]
}

func GetDesc(o *kmetav1.ObjectMeta) string {
	return o.Annotations[AnnotationDescription]
}
