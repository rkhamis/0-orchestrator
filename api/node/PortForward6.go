package node

import (
	"gopkg.in/validator.v2"
)

type PortForward6 struct {
	Dstip     string       `json:"dstip" validate:"nonzero"`
	Dstport   int          `json:"dstport,omitempty"`
	Protocols []IPProtocol `json:"protocols" validate:"nonzero"`
	Srcip     string       `json:"srcip,omitempty"`
	Srcport   int          `json:"srcport" validate:"nonzero"`
}

func (s PortForward6) Validate() error {

	return validator.Validate(s)
}
