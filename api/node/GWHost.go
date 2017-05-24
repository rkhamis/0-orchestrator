package node

import (
	"gopkg.in/validator.v2"
)

type GWHost struct {
	Cloudinit  *CloudInit `json:"cloudinit,omitempty" yaml:"cloudinit,omitempty"`
	Ip6address string     `json:"ip6address,omitempty" yaml:"ip6address,omitempty"`
	Hostname   string     `json:"hostname"  yaml:"hostname" validate:"nonzero"`
	Ipaddress  string     `json:"ipaddress" yaml:"ipaddress" validate:"nonzero,ip"`
	Macaddress string     `json:"macaddress" yaml:"macaddress" validate:"nonzero,macaddress"`
}

func (s GWHost) Validate() error {

	return validator.Validate(s)
}
