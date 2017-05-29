package node

type EnumGWNICType string

const (
	EnumGWNICTypezerotier EnumGWNICType = "zerotier"
	EnumGWNICTypevxlan    EnumGWNICType = "vxlan"
	EnumGWNICTypevlan     EnumGWNICType = "vlan"
	EnumGWNICTypedefault  EnumGWNICType = "default"
	EnumGWNICTypebridge   EnumGWNICType = "bridge"
)
