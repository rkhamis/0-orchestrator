package node

import (
	"gopkg.in/validator.v2"
)

type PortForward struct {
	Dstip     string       `json:"dstip" validate:"nonzero"`
	Dstport   int          `json:"dstport" validate:"nonzero"`
	Protocols []EnumIPProtocol `json:"protocols" validate:"nonzero"`
	Srcip     string       `json:"srcip" validate:"nonzero"`
	Srcport   int          `json:"srcport" validate:"nonzero"`
}

func (s PortForward) Validate() error {

	return validator.Validate(s)
}
