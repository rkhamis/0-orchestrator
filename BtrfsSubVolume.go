package main

import (
	"gopkg.in/validator.v2"
)

// Btrfs subvolume details
type BtrfsSubVolume struct {
	Gen      int    `json:"gen" validate:"nonzero"`
	Id       int    `json:"id" validate:"nonzero"`
	Path     string `json:"path" validate:"nonzero"`
	TopLevel int    `json:"topLevel" validate:"nonzero"`
}

func (s BtrfsSubVolume) Validate() error {

	return validator.Validate(s)
}
