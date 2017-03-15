package main

import (
	"gopkg.in/validator.v2"
)

// Virtual NicInfo zerotier configuration
type KVMNicInfoVxLANConfiguration struct {
	Id int `json:"id" validate:"nonzero"`
}

func (s KVMNicInfoVxLANConfiguration) Validate() error {

	return validator.Validate(s)
}
