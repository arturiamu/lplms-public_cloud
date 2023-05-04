package common

type ResourceType string

const (
	ZoneResourceType        ResourceType = "Zone"
	IoOptimizedResourceType ResourceType = "IoOptimized"
	ServerTypeResourceType  ResourceType = "ServerType"
	SystemDiskResourceType  ResourceType = "SystemDisk"
	DataDiskResourceType    ResourceType = "DataDisk"
	NetworkResourceType     ResourceType = "Network"
)

func (r ResourceType) IsValid() bool {
	switch r {
	case ZoneResourceType, IoOptimizedResourceType, ServerTypeResourceType, SystemDiskResourceType,
		DataDiskResourceType, NetworkResourceType:
		return true
	}
	return false
}

func (r ResourceType) String() string {
	if r.IsValid() {
		return string(r)
	}
	return "N/A"
}

type ResourceStatusType string

const (
	AvailableResourceStatusType ResourceStatusType = "Available"
	SoldOutResourceStatusType   ResourceStatusType = "SoldOut"
)

func (r ResourceStatusType) IsValid() bool {
	switch r {
	case AvailableResourceStatusType, SoldOutResourceStatusType:
		return true
	}
	return false
}

func (r ResourceStatusType) String() string {
	if r.IsValid() {
		return string(r)
	}
	return "N/A"
}
