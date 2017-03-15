package main

import (
	"gopkg.in/validator.v2"
)

// Arguments for a corex.terminate command
type CoreXTerminate struct {
	Command Command `json:"command,omitempty"`
	CoreX   int     `json:"coreX" validate:"nonzero"`
}

func (s CoreXTerminate) Validate() error {

	return validator.Validate(s)
}
