package main

import (
	"gopkg.in/validator.v2"
)

type MountMapping struct {
	Destination string `json:"destination" validate:"nonzero"`
	Source      string `json:"source" validate:"nonzero"`
}

func (s MountMapping) Validate() error {

	return validator.Validate(s)
}
