package main

import (
	"gopkg.in/validator.v2"
)

type WriteFile struct {
	Command  Command `json:"command,omitempty"`
	Contents string  `json:"contents" validate:"nonzero"`
	Path     string  `json:"path" validate:"nonzero"`
}

func (s WriteFile) Validate() error {

	return validator.Validate(s)
}
