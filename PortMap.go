package main

import (
	"gopkg.in/validator.v2"
)

// Port mapping
type PortMap struct {
	Core0port int `json:"core0port" validate:"nonzero"`
	CoreXport int `json:"coreXport" validate:"nonzero"`
}

func (s PortMap) Validate() error {

	return validator.Validate(s)
}
