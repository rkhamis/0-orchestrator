package main

import (
	"gopkg.in/validator.v2"
)

// Command arguments for creating btrfs filesystems
type BtrfsCreate struct {
	Command         Command                        `json:"command,omitempty"`
	DataProfile     EnumBtrfsCreateDataProfile     `json:"dataProfile" validate:"nonzero"`
	Devices         []string                       `json:"devices" validate:"nonzero"`
	Label           string                         `json:"label" validate:"nonzero"`
	MetadataProfile EnumBtrfsCreateMetadataProfile `json:"metadataProfile" validate:"nonzero"`
}

func (s BtrfsCreate) Validate() error {

	return validator.Validate(s)
}
