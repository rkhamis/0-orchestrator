package node

import (
	"gopkg.in/validator.v2"
)

type GWNIC struct {
	Config     string        `json:"config,omitempty"`
	Dhcpserver DHCP          `json:"dhcpserver,omitempty"`
	Id         string        `json:"id" validate:"nonzero"`
	Name       string        `json:"name" validate:"nonzero"`
	Type       EnumGWNICType `json:"type" validate:"nonzero"`
}

func (s GWNIC) Validate() error {

	return validator.Validate(s)
}
