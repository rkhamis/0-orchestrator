package node

import (
	"gopkg.in/validator.v2"
)

type DHCP struct {
	Domain      string   `json:"domain,omitempty"`
	Hosts       []GWHost `json:"hosts" validate:"nonzero"`
	Interface   string   `json:"interface" validate:"nonzero"`
	Nameservers []string `json:"nameservers,omitempty"`
}

func (s DHCP) Validate() error {

	return validator.Validate(s)
}
