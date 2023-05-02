package common

type OSType string

const (
	LinuxOSType   OSType = "linux"
	WindowsOSType OSType = "windows"
)

func (a OSType) IsValid() bool {
	switch a {
	case LinuxOSType, WindowsOSType:
		return true
	}
	return false
}

func (a OSType) String() string {
	if a.IsValid() {
		return string(a)
	}
	return "N/A"
}

type PlatformType string

const (
	PlatformTypeWindows  PlatformType = "Windows"
	PlatformTypeCentOS   PlatformType = "CentOS"
	PlatformTypeRedhat   PlatformType = "Redhat"
	PlatformTypeUbuntu   PlatformType = "Ubuntu"
	PlatformTypeSUSE     PlatformType = "SUSE"
	PlatformTypeDebian   PlatformType = "Debian"
	PlatformTypeFedora   PlatformType = "Fedora"
	PlatformTypeNeoshine PlatformType = "Neoshine Linux Server"
	PlatformTypeOracle   PlatformType = "Oracle Linux"
	PlatformTypeCustom   PlatformType = "凝思Linux"
	PlatformTypeAsianux  PlatformType = "红旗Linux"
	PlatformTypeKylin    PlatformType = "Kylin Linux"
	PlatformTypeUOS      PlatformType = "UOS"
	PlatformTypeCirrOS   PlatformType = "CirrOS"
	PlatformTypeOther    PlatformType = "Other Linux"
)

func (a PlatformType) IsValid() bool {
	switch a {
	case PlatformTypeWindows, PlatformTypeCentOS, PlatformTypeRedhat, PlatformTypeUbuntu, PlatformTypeSUSE, PlatformTypeDebian, PlatformTypeFedora, PlatformTypeNeoshine,
		PlatformTypeOracle, PlatformTypeCustom, PlatformTypeAsianux, PlatformTypeKylin, PlatformTypeUOS, PlatformTypeCirrOS, PlatformTypeOther:
		return true
	}
	return false
}

func (a PlatformType) String() string {
	if a.IsValid() {
		return string(a)
	}
	return ""
}

type ArchitectureType string

const (
	ArchitectureType86    ArchitectureType = "x86_64"
	ArchitectureTypeARM   ArchitectureType = "arm"
	ArchitectureTypePPC64 ArchitectureType = "pcc64"
	ArchitectureTypeI386  ArchitectureType = "i386"
)

func (a ArchitectureType) IsValid() bool {
	switch a {
	case ArchitectureType86, ArchitectureTypeARM, ArchitectureTypePPC64:
		return true
	}
	return false
}

func (a ArchitectureType) String() string {
	if a.IsValid() {
		return string(a)
	}
	return ""
}

type ImageStatus string

const (
	ImageStatusUnAvailable  ImageStatus = "UnAvailable"  // 不可用
	ImageStatusAvailable    ImageStatus = "Available"    // 可用
	ImageStatusCreating     ImageStatus = "Creating"     // 创建中
	ImageStatusCreateFailed ImageStatus = "CreateFailed" // 创建失败
	ImageStatusWaiting      ImageStatus = "Waiting"      // 排队中
	ImageStatusDeprecated   ImageStatus = "Deprecated"   // 已弃用
)

func (a ImageStatus) IsValid() bool {
	switch a {
	case ImageStatusUnAvailable, ImageStatusAvailable, ImageStatusCreating, ImageStatusCreateFailed, ImageStatusWaiting, ImageStatusDeprecated:
		return true
	}
	return false
}

func (a ImageStatus) String() string {
	if a.IsValid() {
		return string(a)
	}
	return "N/A"
}

type ImageDisk string

const (
	ImageDiskAMI   ImageDisk = "ami"
	ImageDiskARI   ImageDisk = "ari"
	ImageDiskAKI   ImageDisk = "aki"
	ImageDiskVHD   ImageDisk = "vhd"
	ImageDiskVHDX  ImageDisk = "vhdx"
	ImageDiskVMDK  ImageDisk = "vmdk"
	ImageDiskRAW   ImageDisk = "raw"
	ImageDiskQCOW2 ImageDisk = "qcow2"
	ImageDiskVDI   ImageDisk = "vdi"
	ImageDiskPLOOP ImageDisk = "ploop"
	ImageDiskISO   ImageDisk = "iso"
)

func (a ImageDisk) IsValid() bool {
	switch a {
	case ImageDiskAMI, ImageDiskARI, ImageDiskAKI, ImageDiskVHD, ImageDiskVHDX, ImageDiskVMDK, ImageDiskRAW, ImageDiskQCOW2, ImageDiskVDI, ImageDiskPLOOP, ImageDiskISO:
		return true
	}
	return false
}

func (a ImageDisk) String() string {
	if a.IsValid() {
		return string(a)
	}
	return "N/A"
}

type ImageOwnerAlias string

const (
	ImageOwnerAliasSystem      ImageOwnerAlias = "system"      // 公共镜像
	ImageOwnerAliasSelf        ImageOwnerAlias = "self"        // 自定义镜像
	ImageOwnerAliasOthers      ImageOwnerAlias = "others"      // 共享镜像
	ImageOwnerAliasMarketplace ImageOwnerAlias = "marketplace" // 市场镜像
)

func (a ImageOwnerAlias) IsValid() bool {
	switch a {
	case ImageOwnerAliasSystem, ImageOwnerAliasSelf, ImageOwnerAliasOthers, ImageOwnerAliasMarketplace:
		return true
	}
	return false
}

func (a ImageOwnerAlias) String() string {
	if a.IsValid() {
		return string(a)
	}
	return "N/A"
}
