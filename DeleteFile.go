package main

import (
	"gopkg.in/validator.v2"
)

type DeleteFile struct {
	Command Command `json:"command,omitempty"`
	Path    string  `json:"path" validate:"nonzero"`
}

func (s DeleteFile) Validate() error {

	return validator.Validate(s)
}
