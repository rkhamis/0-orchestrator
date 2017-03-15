package main

import (
	"gopkg.in/validator.v2"
)

type CoreX struct {
	Bridges  []Bridge       `json:"bridges" validate:"nonzero"`
	Hostname string         `json:"hostname" validate:"nonzero"`
	Id       string         `json:"id" validate:"nonzero"`
	Mounts   []MountMapping `json:"mounts" validate:"nonzero"`
	Ports    []PortMap      `json:"ports" validate:"nonzero"`
	Root_url string         `json:"root_url" validate:"nonzero"`
	Zerotier string         `json:"zerotier,omitempty"`
}

func (s CoreX) Validate() error {

	return validator.Validate(s)
}
