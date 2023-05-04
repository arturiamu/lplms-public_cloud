package common

const (
	AnnotationName             = "virt.lplms.com/name"
	AnnotationDescription      = "virt.lplms.com/description"
	AnnotationPassword         = "virt.lplms.com/password"
	AnnotationVPCID            = "virt.lplms.com/vpc_id"
	AnnotationVSwitchID        = "virt.lplms.com/vswitch_id"
	AnnotationDiskType         = "virt.lplms.com/disk_type"
	AnnotationFingerPrint      = "virt.lplms.com/finger_print"
	AnnotationFlavor           = "virt.lplms.com/flavor"
	AnnotationPrivateIPAddress = "virt.lplms.com/private_ip_addr"
	AnnotationMacAddress       = "virt.lplms.com/mac_address"
	AnnotationImageOwnerAlias  = "virt.lplms.com/image_owner_alias"
	AnnotationImageStatus      = "virt.lplms.com/image_status"
	AnnotationHide             = "virt.lplms.com/hide"
	AnnotationImageVersion     = "virt.lplms.com/image_version"
	AnnotationVMSnapshotName   = "virt.lplms.com/vm_snapshot_name"
	AnnotationImageReference   = "virt.lplms.com/image_reference"
	LabelProject               = "virt.lplms.com/project" // labels to gather non-namespaced resource under same project

	AnnotationOvnLogicalSwitch = "ovn.kubernetes.io/logical_switch"  // subnet
	AnnotationOvnLogicalRouter = "ovn.kubernetes.io/logical_router"  // vpc
	AnnotationOvnSecurityGroup = "ovn.kubernetes.io/security_groups" // security_group
)

const (
	LabelName   = "name"
	LabelDiskID = "label/disk"
)
