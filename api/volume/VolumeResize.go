package volume

import (
	"gopkg.in/validator.v2"
)

type VolumeResize struct {
	NewSize int `yaml:"newSize" json:"newSize" validate:"nonzero"`
}

func (s VolumeResize) Validate() error {

	return validator.Validate(s)
}
