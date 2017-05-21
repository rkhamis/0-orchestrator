package node

import (
	"gopkg.in/validator.v2"
)

type CloudInit struct {
	MetaData string `json:"meta-data" validate:"nonzero"`
	UserData string `json:"user-data" validate:"nonzero"`
}

func (s CloudInit) Validate() error {

	return validator.Validate(s)
}
