package node

import (
	"gopkg.in/validator.v2"
)

type GWCreate struct {
	Name         string        `json:"name" validate:"nonzero"`
	Domain       string        `json:"domain" validate:"nonzero"`
	Httpproxies  []HTTPProxy   `json:"httpproxies" validate:"nonzero"`
	Nics         []GWNIC       `json:"nics" validate:"nonzero"`
	Portforwards []PortForward `json:"portforwards,omitempty"`
}

func (s GWCreate) Validate() error {

	return validator.Validate(s)
}
