package node

import (
	"gopkg.in/validator.v2"
)

type IPProtocol struct {
}

func (s IPProtocol) Validate() error {

	return validator.Validate(s)
}
