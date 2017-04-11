package volume

import (
	"gopkg.in/validator.v2"
)

type VolumeListItem struct {
	ID             string                     `json:"id" validate:"nonzero"`
	Status         EnumVolumeStatus           `json:"status,omitempty"`
	Storagecluster string                     `json:"storageCluster" validate:"nonzero"`
	Volumetype     EnumVolumeCreateVolumetype `json:"type" validate:"nonzero"`
}

func (s VolumeListItem) Validate() error {

	return validator.Validate(s)
}
