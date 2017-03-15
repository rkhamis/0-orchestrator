package main

import (
	"gopkg.in/validator.v2"
)

// Definition of a virtual CPUInfo
type KVMCPUInfo struct {
}

func (s KVMCPUInfo) Validate() error {

	return validator.Validate(s)
}
