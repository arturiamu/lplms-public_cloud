package entity

type Keypair struct {
	KeyPairName string // SSH密钥对名称。
	FingerPrint string // 密钥对的指纹。
}

type KeypairCreateArg struct {
	KeyPairName   string  // 密钥对名称。长度为2~128个英文或中文字符。必须以大小字母或中文开头，不能以 `http://` 和 `https://` 开头。可以包含数字、半角冒号（:）、下划线（_）或者连字符（-）。
	PublicKeyBody *string // 密钥对的公钥内容。如果非空表示导入密钥对，否则表示新建密钥对。
	ProjectID     string
}

type KeypairDeleteArg struct {
	ProjectID   string
	KeyPairName string // SSH密钥对名称。
}

type KeypairUpdateArg struct {
}

type KeypairGetArg struct {
}

type KeypairListArg struct {
}
