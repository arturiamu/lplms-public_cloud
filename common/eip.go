package common

import "fmt"

type LockReason string //enum financial | security

const (
	FinancialLockReason LockReason = "financial"
	SecurityLockReason  LockReason = "security"
)

type EIPStatus string // enum Associating | Unassociating | InUse | Available

const (
	CreatingEIPStatus      EIPStatus = "Creating"
	AssociatingEIPStatus   EIPStatus = "Associating"
	UnassociatingEIPStatus EIPStatus = "Unassociating"
	InUseEIPStatus         EIPStatus = "InUse"
	AvailableEIPStatus     EIPStatus = "Available"
)

func (a EIPStatus) IsValid() bool {
	switch a {
	case CreatingEIPStatus, AssociatingEIPStatus, UnassociatingEIPStatus, InUseEIPStatus, AvailableEIPStatus:
		return true
	}
	return false
}

func (a EIPStatus) String() string {
	switch a {
	case CreatingEIPStatus:
		return "Creating"
	case AssociatingEIPStatus:
		return "Associating"
	case UnassociatingEIPStatus:
		return "Unassociating"
	case InUseEIPStatus:
		return "InUse"
	case AvailableEIPStatus:
		return "Available"
	}
	return "N/A"
}

type EIPISP string

const (
	// BGPProEIPISP ALI BGP（多线）精品线路
	BGPProEIPISP EIPISP = "BGP_PRO"
	// BGPEIPISP ALI BGP（多线）线路
	BGPEIPISP EIPISP = "BGP"

	// HWTelcomEIPISP HW 电信
	HWTelcomEIPISP EIPISP = "5_telcom"
	// HWUnionEIPISP HW 联通
	HWUnionEIPISP EIPISP = "5_union"
	// HWBGPEIPISP HW 全动态BGP
	HWBGPEIPISP EIPISP = "5_bgp"
	// HWSBGPEIPISP HW 静态BGP
	HWSBGPEIPISP EIPISP = "5_sbgp"
)

func (isp EIPISP) IsValid() bool {
	switch isp {
	case BGPProEIPISP, BGPEIPISP, HWTelcomEIPISP, HWUnionEIPISP, HWBGPEIPISP, HWSBGPEIPISP:
		return true
	}
	return false
}

func (isp EIPISP) String() string {
	if isp.IsValid() {
		return string(isp)
	}
	return "N/A"
}

type AttachInstanceType string

const (
	AttachInstanceTypeVM  AttachInstanceType = "compute:nova"         // 云主机
	AttachInstanceTypeSLB AttachInstanceType = "network:loadbalancer" // 负载均衡实例
)

func (ait AttachInstanceType) IsValid() bool {
	switch ait {
	case AttachInstanceTypeVM, AttachInstanceTypeSLB:
		return true
	}
	return false
}

func (ait AttachInstanceType) String() string {
	if ait.IsValid() {
		return string(ait)
	}
	return "N/A"
}

type MNOType string

// 线路枚举
const (
	MNOTypeMobile  MNOType = "china_mobile"  // 中国移动
	MNOTypeUnicom  MNOType = "china_unicom"  // 中国联通
	MNOTypeTelecom MNOType = "china_telecom" // 中国电信
)

func (mno MNOType) Valid() error {
	switch mno {
	case MNOTypeMobile, MNOTypeUnicom, MNOTypeTelecom:
		return nil
	}

	return fmt.Errorf("invalid mobile network operator type")
}
