package main

import (
	"gopkg.in/validator.v2"
)

// Arguments for a bridge.delete command
type BridgeDelete struct {
	Command Command `json:"command,omitempty"`
	Name    string  `json:"name" validate:"nonzero"`
}

func (s BridgeDelete) Validate() error {

	return validator.Validate(s)
}
