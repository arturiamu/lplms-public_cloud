package mock

import "github.com/arturiamu/lplms-public_cloud/domain/entity"

var mockKeypairList = []entity.Keypair{
	{
		KeyPairName: "", // SSH密钥对名称。
		FingerPrint: "", // 密钥对的指纹。
	},
}
