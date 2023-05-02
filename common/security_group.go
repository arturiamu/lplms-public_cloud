package common

import (
	"strings"
)

type NetworkProtocol string

const (
	AllNetworkProtocol  NetworkProtocol = "all"
	TCPNetworkProtocol  NetworkProtocol = "tcp"
	UDPNetworkProtocol  NetworkProtocol = "udp"
	ICMPNetworkProtocol NetworkProtocol = "icmp"
	GRENetworkProtocol  NetworkProtocol = "gre"
)

func (v *NetworkProtocol) Convert() {
	*v = NetworkProtocol(strings.ToLower(string(*v)))
}

func (v *NetworkProtocol) IsValid() bool {
	switch *v {
	case AllNetworkProtocol, TCPNetworkProtocol, UDPNetworkProtocol, ICMPNetworkProtocol, GRENetworkProtocol:
		return true
	}
	return false
}

type SecurityGroupPriority int

func (p SecurityGroupPriority) IsValid() bool {
	if p >= 1 && p <= 100 {
		return true
	}
	return false
}

type Direction string

const (
	IngressDirection Direction = "ingress"
	EgressDirection  Direction = "egress"
	AllDirection     Direction = "all"
)

func (d *Direction) Convert() {
	*d = Direction(strings.ToLower(string(*d)))
}

func (d *Direction) IsValid() bool {
	switch *d {
	case IngressDirection, EgressDirection, AllDirection:
		return true
	}
	return false
}

type SecurityRuleGrantType int

const (
	SecurityGroupType SecurityRuleGrantType = 1
	CIDRType          SecurityRuleGrantType = 2
)

func (t SecurityRuleGrantType) IsValid() bool {
	switch t {
	case SecurityGroupType, CIDRType:
		return true
	}
	return false
}
