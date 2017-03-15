package main

import (
	"gopkg.in/validator.v2"
)

type KVMDomain struct {
	Cpu    []KVMCPUInfo  `json:"cpu" validate:"nonzero"`
	Disk   []KVMDiskInfo `json:"disk" validate:"nonzero"`
	Memory int           `json:"memory" validate:"nonzero"`
	Name   string        `json:"name" validate:"nonzero"`
	Nic    []KVMNicInfo  `json:"nic" validate:"nonzero"`
}

func (s KVMDomain) Validate() error {

	return validator.Validate(s)
}
