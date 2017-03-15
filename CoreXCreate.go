package main

import (
	"gopkg.in/validator.v2"
)

// Arguments for a corex.create command
type CoreXCreate struct {
	Command          Command        `json:"command,omitempty"`
	Mounts           []KeyValuePair `json:"mounts" validate:"nonzero"`
	Name             string         `json:"name" validate:"nonzero"`
	NetworkBridges   []Bridge       `json:"networkBridges,omitempty"`
	Plist            string         `json:"plist" validate:"nonzero"`
	PortMaps         []PortMap      `json:"portMaps" validate:"nonzero"`
	ZerotierNetworks []string       `json:"zerotierNetworks,omitempty"`
}

func (s CoreXCreate) Validate() error {

	return validator.Validate(s)
}
