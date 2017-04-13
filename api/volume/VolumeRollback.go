package volume

import (
	"gopkg.in/validator.v2"
)

type VolumeRollback struct {
	Epoch int `yaml:"epoch" json:"epoch" validate:"nonzero"`
}

func (s VolumeRollback) Validate() error {

	return validator.Validate(s)
}
