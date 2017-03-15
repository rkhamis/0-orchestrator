package main

import (
	"gopkg.in/validator.v2"
)

// Arguments for a core.kill command
type CoreKill struct {
	Command Command `json:"command,omitempty"`
	Id      string  `json:"id" validate:"nonzero"`
}

func (s CoreKill) Validate() error {

	return validator.Validate(s)
}
