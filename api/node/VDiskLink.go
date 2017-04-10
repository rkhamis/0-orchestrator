package node

import (
	"gopkg.in/validator.v2"
)

// Definition of a virtual disk
type VDiskLink struct {
	MaxIOps          int    `json:"maxIOps" validate:"nonzero"`
	Volumeid         string `json:"volumeid" validate:"nonzero"`
}

func (s VDiskLink) Validate() error {

	return validator.Validate(s)
}
