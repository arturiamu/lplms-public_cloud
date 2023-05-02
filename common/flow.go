package common

// FlowDirection 流量方向
type FlowDirection string

// FlowDirection enums
const (
	FlowDirectionIngress FlowDirection = "ingress" // 入口方向流量
	FlowDirectionEgress  FlowDirection = "egress"  // 出口方向流量
)
