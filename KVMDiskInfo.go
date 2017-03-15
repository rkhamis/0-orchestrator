package main

import (
	"gopkg.in/validator.v2"
)

// Definition of a virtual DiskInfo
type KVMDiskInfo struct {
	MaxIOps int    `json:"maxIOps" validate:"nonzero"`
	Url     string `json:"url" validate:"nonzero"`
}

func (s KVMDiskInfo) Validate() error {

	return validator.Validate(s)
}
