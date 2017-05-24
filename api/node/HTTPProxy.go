package node

import (
	"gopkg.in/validator.v2"
)

type HTTPProxy struct {
	Destinations []string       `json:"destinations" validate:"nonzero"`
	Host         string         `json:"host" validate:"nonzero"`
	Types        []EnumHTTPType `json:"types" validate:"nonzero"`
}

func (s HTTPProxy) Validate() error {

	return validator.Validate(s)
}
