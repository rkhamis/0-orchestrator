package node

import (
	"gopkg.in/validator.v2"
)

type GW struct {
	Dhcps         []DHCP         `json:"dhcps" validate:"nonzero"`
	Httpproxies   []HTTPProxy    `json:"httpproxies" validate:"nonzero"`
	Nics          []ContainerNIC `json:"nics" validate:"nonzero"`
	Portforwards  []PortForward  `json:"portforwards" validate:"nonzero"`
	Portforwards6 []PortForward6 `json:"portforwards6" validate:"nonzero"`
}

func (s GW) Validate() error {

	return validator.Validate(s)
}
