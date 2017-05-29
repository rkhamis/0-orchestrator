package node

import (
	"gopkg.in/validator.v2"
)

type CloudInit struct {
	MetaData string `yaml:"metadata" json:"metadata" validate:"nonzero"`
	UserData string `yaml:"userdata" json:"userdata" validate:"nonzero"`
}

func (s CloudInit) Validate() error {

	return validator.Validate(s)
}
