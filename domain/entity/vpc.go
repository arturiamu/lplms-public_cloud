package entity

type Vpc struct {
	PrivateIPAddress string
	Mac              string
	VSwitchId        string
	VPCId            string
	VPCName          string
	VSwitchName      string
	Type             string
	EipAddress       *EipAddress `json:"eip_address"`
}

type EipAddress struct {
	Name         string `json:"name"`
	AllocationID string `json:"allocation_id"`
	IPAddress    string `json:"ip_address"`
	Bandwidth    int64  `json:"bandwidth"`
}

type VpcCreateArg struct {
}

type VpcDeleteArg struct {
}

type VpcUpdateArg struct {
}

type VpcGetArg struct {
}

type VpcListArg struct {
	VPCIDs    []string
	ProjectID string
}

type VpcCreateResp struct{}

type VpcDeleteResp struct{}

type VpcUpdateResp struct{}

type VpcGetResp struct{}

type VpcListResp struct {
	VPCs []Vpc
}
