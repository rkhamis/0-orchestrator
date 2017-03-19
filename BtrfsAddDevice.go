package main

import (
	"gopkg.in/validator.v2"
)

// Command to add a device to an existing Btrfs filesystem
type BtrfsAddDevice struct {
	Command Command `json:"command,omitempty"`
	Path    string  `json:"path" validate:"nonzero"`
}

func (s BtrfsAddDevice) Validate() error {

	return validator.Validate(s)
}
