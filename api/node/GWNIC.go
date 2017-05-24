package node

import (
	"gopkg.in/validator.v2"
)

type GWNIC struct {
	Config     *GWNICConfig  `json:"config,omitempty" yaml:"config,omitempty"`
	Dhcpserver *DHCP         `json:"dhcpserver,omitempty" yaml:"dhcpserver,omitempty"`
	Id         string        `json:"id"   yaml:"id"   validate:"nonzero"`
	Name       string        `json:"name" yaml:"name" validate:"nonzero"`
	Type       EnumGWNICType `json:"type" yaml:"type" validate:"nonzero"`
}

func (s GWNIC) Validate() error {

	return validator.Validate(s)
}
