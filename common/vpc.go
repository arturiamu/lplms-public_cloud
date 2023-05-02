package common

type VPCStatus string //enum Pending | Available

const (
	PendingVPCStatus   VPCStatus = "Pending"
	AvailableVPCStatus VPCStatus = "Available"
)

func (v VPCStatus) IsValid() bool {
	switch v {
	case PendingVPCStatus, AvailableVPCStatus:
		return true
	}
	return false
}

func (v VPCStatus) String() string {
	switch v {
	case PendingVPCStatus:
		return "Pending"
	case AvailableVPCStatus:
		return "Available"
	}
	return "N/A"
}

// CenStatusType VPC绑定云企业网的状态
type CenStatusType = string

// 绑定和未绑定
const (
	DetachedCenStatus  CenStatusType = "Detached"
	AvailableCenStatus CenStatusType = "Available"
)

type VSwitchStatus string //enum Pending | Available

const (
	PendingVSwitchStatus   VSwitchStatus = "Pending"
	AvailableVSwitchStatus VSwitchStatus = "Available"
)

func (v VSwitchStatus) IsValid() bool {
	switch v {
	case PendingVSwitchStatus, AvailableVSwitchStatus:
		return true
	}
	return false
}

func (v VSwitchStatus) String() string {
	switch v {
	case PendingVSwitchStatus:
		return "Pending"
	case AvailableVSwitchStatus:
		return "Available"
	}
	return "N/A"
}

type VPNStatus string

const (
	InitVPN         = VPNStatus("init")
	ProvisioningVPN = VPNStatus("provisioning")
	ActiveVPN       = VPNStatus("active")
	UpdatingVPN     = VPNStatus("updating")
	DeletingVPN     = VPNStatus("deleting")
)

type VPNBusinessStatus string

const (
	VPNNormalBusiness          = VPNBusinessStatus("Normal")
	VPNFinancialLockedBusiness = VPNBusinessStatus("FinancialLocked")
)

type IpsecVPNStatus string

const (
	IkeSaNotEstablished   = IpsecVPNStatus("ike_sa_not_established")
	IkeSaEstablished      = IpsecVPNStatus("ike_sa_established")
	IpsecSaNotEstablished = IpsecVPNStatus("ipsec_sa_not_established")
	IpsecSaEstablished    = IpsecVPNStatus("ipsec_sa_established")
)

func (s IpsecVPNStatus) String() string {
	return string(s)
}

type SSLProto string

const (
	UDPSSLProto = SSLProto("UDP")
	TCPSSLProto = SSLProto("TCP")
)

type Cipher string

const (
	AES128CBC = Cipher("AES-128-CBC")
	AES192CBC = Cipher("AES-192-CBC")
	AES256CBC = Cipher("AES-256-CBC")
	NoneEncry = Cipher("none")
)

type SslVpnClientCertStatus string

const (
	ExpiringSslVpnClientCert      = SslVpnClientCertStatus("expiring-soon")
	NormalSslVpnClientCertStatus  = SslVpnClientCertStatus("normal")
	ExpiredSslVpnClientCertStatus = SslVpnClientCertStatus("expired")
)

func (s SslVpnClientCertStatus) String() string {
	return string(s)
}

type VpnInstanceChargeType string

const (
	VpnPREPAY  = VpnInstanceChargeType("PREPAY")
	VpnPOSTPAY = VpnInstanceChargeType("POSTPAY")
)
