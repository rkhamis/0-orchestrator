package main

import (
	"gopkg.in/validator.v2"
)

// Arguments for a disk.mount command
type DiskMount struct {
	Command       Command `json:"command,omitempty"`
	Mount_options string  `json:"mount_options" validate:"nonzero"`
	Target        string  `json:"target" validate:"nonzero"`
}

func (s DiskMount) Validate() error {

	return validator.Validate(s)
}
