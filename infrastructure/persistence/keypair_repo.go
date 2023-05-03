package persistence

import (
	"context"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *ComputeRepo) CreateKeypair(arg *entity.KeypairCreateArg) (*entity.KeypairCreateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) DeleteKeypair(arg *entity.KeypairDeleteArg) (*entity.KeypairDeleteResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) UpdateKeypair(arg *entity.KeypairUpdateArg) (*entity.KeypairUpdateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) GetKeypair(arg *entity.KeypairGetArg) (*entity.KeypairGetResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) ListKeypair(arg *entity.KeypairListArg) (*entity.KeypairListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ComputeRepo) getPubkey(keypairname string, ns string) (pubKey string, err error) {
	resp, err := c.k8Virt.CoreV1().Secrets(ns).Get(context.Background(), keypairname, v1.GetOptions{})
	if err != nil {
		return
	}

	if resp != nil {
		pubKey = string(resp.Data["key1"])
	}

	return
}
