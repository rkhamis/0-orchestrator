package node

import (
	"gopkg.in/validator.v2"
)

type ContainerNICConfig struct {
	Dhcp        bool     `json:"dhcp"`
	Cidr        string   `json:"cidr"`
	Gateway     string   `json:"gateway"`
	Nameservers []string `json:"nameservers"`
}

type ContainerNIC struct {
	Config ContainerNICConfig   `json:"config,omitempty"`
	Hwaddr string               `json:"hwaddr,omitempty"`
	Id     string               `json:"id" validate:"nonzero"`
	Type   EnumContainerNICType `json:"type" validate:"nonzero"`
}

func (s ContainerNIC) Validate() error {

	return validator.Validate(s)
}
