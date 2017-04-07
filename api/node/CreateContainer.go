package node

import (
	"gopkg.in/validator.v2"
)

type CreateContainer struct {
	Nics           []ContainerNIC            `json:"nics" validate:"nonzero"`
	Filesystems    []string                  `json:"filesystems" validate:"nonzero"`
	Flist          string                    `json:"flist" validate:"nonzero"`
	HostNetworking bool                      `json:"hostNetworking"`
	Hostname       string                    `json:"hostname" validate:"nonzero"`
	Id             string                    `json:"id" validate:"nonzero"`
	InitProcesses  []CoreSystem              `json:"initProcesses" validate:"nonzero"`
	Ports          []string                  `json:"ports" validate:"nonzero"`
	Status         EnumCreateContainerStatus `json:"status" validate:"nonzero"`
	Storage        string                    `json:"storage" validate:"nonzero"`
}

func (s CreateContainer) Validate() error {

	return validator.Validate(s)
}
