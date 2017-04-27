package node

import (
	"gopkg.in/validator.v2"
)

// VMInfo Statistical information about a vm
type VMInfo struct {
	CPU  []float64    `yaml:"cpu" json:"cpu" validate:"nonzero"`
	Disk []VMDiskInfo `yaml:"disk" json:"disk" validate:"nonzero"`
	Net  []VMNetInfo  `yaml:"net" json:"net" validate:"nonzero"`
}

func (s VMInfo) Validate() error {

	return validator.Validate(s)
}
