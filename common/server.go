package common

type ServerStatus string

const (
	NormalServerStatus    ServerStatus = "Normal"
	CreatingServerStatus  ServerStatus = "Creating"
	ChangingServerStatus  ServerStatus = "Changing"
	RunningServerStatus   ServerStatus = "Running"
	StartingServerStatus  ServerStatus = "Starting"
	StoppingServerStatus  ServerStatus = "Stopping"
	StoppedServerStatus   ServerStatus = "Stopped"
	InactiveServerStatus  ServerStatus = "Inactive"
	RebootingServerStatus ServerStatus = "Rebooting"
	PendingServerStatus   ServerStatus = "Pending"
	DeletedServerStatus   ServerStatus = "Deleted"
	DeletingServerStatus  ServerStatus = "Deleting"
	UnknownServerStatus   ServerStatus = "Unknown"
)

func (s ServerStatus) IsValid() bool {
	switch s {
	case NormalServerStatus, CreatingServerStatus, ChangingServerStatus, RunningServerStatus,
		StartingServerStatus, StoppingServerStatus, StoppedServerStatus, InactiveServerStatus,
		RebootingServerStatus, PendingServerStatus, DeletedServerStatus, DeletingServerStatus,
		UnknownServerStatus:
		return true
	}
	return false
}

func (s ServerStatus) String() string {
	if s.IsValid() {
		return string(s)
	}
	return ""
}
