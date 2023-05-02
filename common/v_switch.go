package common

type VpnSwitch string

const (
	DisableVpnSwitch = VpnSwitch("disable")
	EnableVpnSwitch  = VpnSwitch("enable")
)

func (v VpnSwitch) String() string {
	return string(v)
}

func (v VpnSwitch) Bool() bool {
	switch v {
	case DisableVpnSwitch:
		return false
	case EnableVpnSwitch:
		return true
	}

	return false
}
