package main

import (
	"gopkg.in/validator.v2"
)

// Devices that are assigned to a btrfs filesystem
type BtrfsDevice struct {
	Id   int    `json:"id" validate:"nonzero"`
	Path string `json:"path" validate:"nonzero"`
	Size int    `json:"size" validate:"nonzero"`
	Used int    `json:"used" validate:"nonzero"`
}

func (s BtrfsDevice) Validate() error {

	return validator.Validate(s)
}
