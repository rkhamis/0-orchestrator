package main

import (
	"gopkg.in/validator.v2"
)

// Arguments for a disk.mktable command
type DiskMKTable struct {
	Command   Command `json:"command,omitempty"`
	Device    string  `json:"device" validate:"nonzero"`
	TableType string  `json:"tableType" validate:"nonzero"`
}

func (s DiskMKTable) Validate() error {

	return validator.Validate(s)
}
