package main

import (
	"gopkg.in/validator.v2"
)

// Btrfs details
type Btrfs struct {
	DeviceCount int           `json:"deviceCount" validate:"nonzero"`
	Devices     []BtrfsDevice `json:"devices" validate:"nonzero"`
	Label       string        `json:"label" validate:"nonzero"`
	Used        int           `json:"used" validate:"nonzero"`
	Uuid        string        `json:"uuid" validate:"nonzero"`
}

func (s Btrfs) Validate() error {

	return validator.Validate(s)
}
