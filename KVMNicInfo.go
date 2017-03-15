package main

import (
	"gopkg.in/validator.v2"
)

// Definition of a virtual NicInfo
type KVMNicInfo struct {
	Configuration interface{}        `json:"configuration" validate:"nonzero"`
	Type          EnumKVMNicInfoType `json:"type" validate:"nonzero"`
}

func (s KVMNicInfo) Validate() error {

	return validator.Validate(s)
}
