package mock

// implement check
//var _ application.ComputeAppInterface = &BizMock{}
//var _ application.NetworkAppInterface = &BizMock{}
//var _ application.StorageAppInterface = &BizMock{}

type BizMock struct{}

func init() {
	initMock()
	setID()
}

func initMock() {
	initFlavor()        // 无依赖
	initImage()         // 无依赖
	initKeypair()       // 无依赖
	initSecurityGroup() // 无依赖

	initVSwitch() // 循环依赖vpc 指定v-switch参数
	initVpc()     // 循环依赖v-switch 指定vpc参数

	initDisk()     // 依赖server, image 指定server信息
	initSnapshot() // 依赖disk
	initEip()      // 依赖server
	initServer()   // 依赖disk,eip, SecurityGroup,image,flavor,vpc
}

func setID() {

}
