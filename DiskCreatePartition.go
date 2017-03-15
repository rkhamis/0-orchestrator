package main

import (
	"gopkg.in/validator.v2"
)

// Arguments for a disk.mkpart command
// part_type, start & end values must be supported by the parted mkpart command
type DiskCreatePartition struct {
	Command   Command `json:"command,omitempty"`
	End       string  `json:"end" validate:"nonzero"`
	Part_type string  `json:"part_type" validate:"nonzero"`
	Start     int     `json:"start" validate:"nonzero"`
}

func (s DiskCreatePartition) Validate() error {

	return validator.Validate(s)
}
