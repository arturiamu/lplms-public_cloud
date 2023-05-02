package common

type DiskCategory string

const (
	CloudDiskCategory   DiskCategory = "cloud"     // 普通云盘
	CloudSSDiskCategory DiskCategory = "cloud_ssd" // SSD云盘
)

func (d DiskCategory) IsValid() bool {
	switch d {
	case CloudDiskCategory, CloudSSDiskCategory:
		return true
	}

	return false
}

func (d DiskCategory) String() string {
	if d.IsValid() {
		return string(d)
	}

	return "N/A"
}

func ParseDiskCategory(c string) DiskCategory {
	switch c {
	case "cloud":
		return CloudDiskCategory
	case "cloud_ssd":
		return CloudSSDiskCategory
	}
	return "N/A"
}

func (d DiskCategory) Humanize() string {
	switch d {
	case CloudDiskCategory:
		return "普通云盘"
	case CloudSSDiskCategory:
		return "SSD云盘"
	default:
		return "N/A"
	}
}

type DiskType string

const (
	SystemDiskType DiskType = "system"
	DataDiskType   DiskType = "data"
)

func (d DiskType) IsValid() bool {
	switch d {
	case SystemDiskType, DataDiskType:
		return true
	}
	return false
}

func (d DiskType) String() string {
	if d.IsValid() {
		return string(d)
	}

	return "N/A"
}

func ParseDiskType(s string) *DiskType {
	var diskType DiskType
	switch s {
	case "system":
		diskType = SystemDiskType
	case "data":
		diskType = DataDiskType
	}
	return &diskType
}

type DiskStatusType string

const (
	DiskInUse         DiskStatusType = "In_use"
	DiskAvailable     DiskStatusType = "Available"
	DiskAttaching     DiskStatusType = "Attaching"
	DiskDetaching     DiskStatusType = "Detaching"
	DiskCreating      DiskStatusType = "Creating"
	DiskInitializing  DiskStatusType = "Initializing"
	DiskReIniting     DiskStatusType = "ReIniting"
	DiskProcessing    DiskStatusType = "Processing"
	DiskDeleting      DiskStatusType = "Deleting"
	DiskDownloading   DiskStatusType = "Downloading"
	DiskError         DiskStatusType = "Error"
	AllDiskStatusType DiskStatusType = "All"
)

func (d DiskStatusType) IsValid() bool {
	switch d {
	case DiskInUse, DiskAvailable, DiskAttaching, DiskDetaching, DiskCreating,
		DiskProcessing, DiskInitializing, DiskReIniting, DiskDeleting, DiskDownloading, DiskError, AllDiskStatusType:
		return true
	}

	return false
}

func (d DiskStatusType) String() string {
	if d.IsValid() {
		return string(d)
	}

	return "N/A"
}
