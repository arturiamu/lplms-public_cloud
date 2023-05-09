package entity

type Keypair struct {
	KeyPairID   string
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
	ProjectID string
	// 密钥对名称。支持正则表达式模糊搜索，使用*匹配子表达式，示例：
	// - *SshKey：查询以SshKey结尾的密钥对名称，包括SshKey。
	// - SshKey*：查询以SshKey开头的密钥对名称，包括SshKey。
	// - *SshKey*：查询名称中间有SshKey的密钥对，包括SshKey。
	// - SshKey：精确匹配SshKey。
	KeyPairID   string
	KeyPairName *string
}

type KeypairListArg struct {
	ProjectID string
	// 密钥对名称。支持正则表达式模糊搜索，使用*匹配子表达式，示例：
	// - *SshKey：查询以SshKey结尾的密钥对名称，包括SshKey。
	// - SshKey*：查询以SshKey开头的密钥对名称，包括SshKey。
	// - *SshKey*：查询名称中间有SshKey的密钥对，包括SshKey。
	// - SshKey：精确匹配SshKey。
	KeyPairName *string
}

type KeypairCreateResp struct {
	KeyPairName string

	// 密钥对的指纹。根据 `RFC4716` 定义的公钥指纹格式，采用MD5信息摘要算法。更多详情，请参见RFC4716。
	FingerPrint string

	// 密钥对的私钥。PEM编码的PKCS#8格式的私钥部分。如果是导入密钥对则不返回该参数。
	PrivateKeyBody *string
}

type KeypairDeleteResp struct{}

type KeypairUpdateResp struct{}

type KeypairGetResp struct {
	Keypair Keypair
}

type KeypairListResp struct {
	KeyPairs []Keypair
}

type KeypairAttachArg struct {
	ProjectID   string
	KeyPairName string // SSH密钥对名称。
	ServerID    string // 绑定SSH密钥对的实例ID。
}

type KeypairAttachResp struct{}

type KeypairDetachArg struct {
	KeyPairName string // SSH密钥对名称。
	ServerID    string // 解绑SSH密钥对的实例ID。
	ProjectID   string
}

type KeypairDetachResp struct{}
