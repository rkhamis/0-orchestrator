package volume

import (
	"gopkg.in/validator.v2"
)

type Volume struct {
	Blocksize          int                  `yaml:"blocksize" json:"blocksize" validate:"nonzero"`
	ID                 string               `yaml:"id" json:"id" validate:"nonzero"`
	ReadOnly           bool                 `yaml:"readOnly" json:"readOnly,omitempty"`
	Size               int                  `yaml:"size" json:"size" validate:"nonzero"`
	Status             EnumVolumeStatus     `yaml:"status" json:"status" validate:"nonzero"`
	Storagecluster     string               `yaml:"storagecluster" json:"storagecluster" validate:"nonzero"`
	TlogStoragecluster string               `yaml:"tlogStoragecluster" json:"tlogStoragecluster" validate:"nonzero"`
	Volumetype         EnumVolumeVolumetype `yaml:"type" json:"type" validate:"nonzero"`
}

func (s Volume) Validate() error {

	return validator.Validate(s)
}
