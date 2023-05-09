package mock

import (
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/arturiamu/lplms-public_cloud/utils/uuid"
)

var mockKeypairMap map[string]entity.Keypair

var mockKeypairList = []entity.Keypair{
	{
		KeyPairID:   uuid.GenerateUUID(),
		KeyPairName: "测试密钥对1",                           // SSH密钥对名称。
		FingerPrint: "2c38359fb8b9d455b633b8ed35d40f93", // 密钥对的指纹。
	},
	{
		KeyPairID:   uuid.GenerateUUID(),
		KeyPairName: "测试密钥对2",                           // SSH密钥对名称。
		FingerPrint: "3a147005b99e4de5890ae9ff1dfe32c0", // 密钥对的指纹。
	},
	{
		KeyPairID:   uuid.GenerateUUID(),
		KeyPairName: "测试密钥对3",                           // SSH密钥对名称。
		FingerPrint: "9eb03771f79f3c0d972f51797832f571", // 密钥对的指纹。
	},
	{
		KeyPairID:   uuid.GenerateUUID(),
		KeyPairName: "测试密钥对4",                           // SSH密钥对名称。
		FingerPrint: "51a020d4a79618f3237551fac6b72b18", // 密钥对的指纹。
	},
	{
		KeyPairID:   uuid.GenerateUUID(),
		KeyPairName: "测试密钥对5",                           // SSH密钥对名称。
		FingerPrint: "1df0b3ec0f8ab866176e17f0bbd67daa", // 密钥对的指纹。
	},
	{
		KeyPairID:   uuid.GenerateUUID(),
		KeyPairName: "测试密钥对6",                           // SSH密钥对名称。
		FingerPrint: "e37162f32a1a14764f4e9a770b18c005", // 密钥对的指纹。
	},
	{
		KeyPairID:   uuid.GenerateUUID(),
		KeyPairName: "测试密钥对7",                           // SSH密钥对名称。
		FingerPrint: "ff84365c4e58c34afcea4d124f428f57", // 密钥对的指纹。
	},
}

func initKeypair() {
	mockKeypairMap = make(map[string]entity.Keypair)
	for _, k := range mockKeypairList {
		mockKeypairMap[k.KeyPairID] = k
	}
}

func getKeypairList() []entity.Keypair {
	return mockKeypairList
}

func getKeypair(id string) entity.Keypair {
	return mockKeypairMap[id]
}
