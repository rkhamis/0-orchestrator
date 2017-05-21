package node

import (
	"gopkg.in/validator.v2"
)

type GWHost struct {
	Cloudinit  CloudInit `json:"cloudinit,omitempty"`
	Hostname   string    `json:"hostname" validate:"nonzero"`
	Ip6address string    `json:"ip6address,omitempty"`
	Ipaddress  string    `json:"ipaddress" validate:"nonzero"`
	Macaddress string    `json:"macaddress" validate:"nonzero"`
}

func (s GWHost) Validate() error {

	return validator.Validate(s)
}
