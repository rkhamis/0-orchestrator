package node

import (
	"gopkg.in/validator.v2"
)

type HTTPType struct {
}

func (s HTTPType) Validate() error {

	return validator.Validate(s)
}
