package main

import (
	"gopkg.in/validator.v2"
)

// Arguments for a bridge.create command
type BridgeCreate struct {
	Command Command `json:"command,omitempty"`
	Hwaddr  string  `json:"hwaddr,omitempty"`
	Name    string  `json:"name" validate:"nonzero"`
}

func (s BridgeCreate) Validate() error {

	return validator.Validate(s)
}
