package common

type SourceDiskType string

const (
	SystemSourceDiskType SourceDiskType = "system"
	DataSourceDiskType   SourceDiskType = "data"
)

func (sourceDisk SourceDiskType) String() string {
	switch sourceDisk {
	case SystemSourceDiskType:
		return "system"
	case DataSourceDiskType:
		return "data"
	}
	return "N/A"
}

type UsageType string

const (
	ImageUsageType     UsageType = "image"
	DiskUsageType      UsageType = "disk"
	ImageDiskUsageType UsageType = "image_disk"
	NoneUsageType      UsageType = "none"
)

func (u UsageType) String() string {
	switch u {
	case ImageUsageType:
		return "image"
	case DiskUsageType:
		return "disk"
	case ImageDiskUsageType:
		return "image_disk"
	case NoneUsageType:
		return "none"
	}
	return "N/A"
}

type SnapshotStatus string

const (
	ProgressingSnapshotStatus  SnapshotStatus = "progressing"
	AccomplishedSnapshotStatus SnapshotStatus = "accomplished"
	FailedSnapshotStatus       SnapshotStatus = "failed"
	RestoringSnapshotStatus    SnapshotStatus = "restoring"
	ProcessingSnapshotStatus   SnapshotStatus = "processing"
)

func (s SnapshotStatus) String() string {
	switch s {
	case ProgressingSnapshotStatus:
		return "progressing"
	case AccomplishedSnapshotStatus:
		return "accomplished"
	case FailedSnapshotStatus:
		return "failed"
	case RestoringSnapshotStatus:
		return "restoring"
	case ProcessingSnapshotStatus:
		return "processing"
	}
	return "N/A"
}

type SnapshotType string

const (
	AutoSnapshotType SnapshotType = "auto"
	UserSnapshotType SnapshotType = "user"
	AllSnapshotType  SnapshotType = "all"
)

func (s SnapshotType) IsValid() bool {
	switch s {
	case AutoSnapshotType, UserSnapshotType, AllSnapshotType:
		return true
	}
	return false
}

func (s SnapshotType) String() string {
	if s.IsValid() {
		return string(s)
	}
	return "N/A"
}
