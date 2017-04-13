package volume

import (
	"gopkg.in/validator.v2"
)

type VolumeListItem struct {
	ID             string                     `yaml:"id" json:"id" validate:"nonzero"`
	Status         EnumVolumeStatus           `yaml:"status" json:"status,omitempty"`
	Storagecluster string                     `yaml:"storageCluster" json:"storageCluster" validate:"nonzero"`
	Volumetype     EnumVolumeCreateVolumetype `yaml:"type" json:"type" validate:"nonzero"`
}

func (s VolumeListItem) Validate() error {

	return validator.Validate(s)
}
