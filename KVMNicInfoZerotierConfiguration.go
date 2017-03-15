package main

import (
	"gopkg.in/validator.v2"
)

// Virtual NicInfo zerotier configuration
type KVMNicInfoZerotierConfiguration struct {
	Id string `json:"id" validate:"nonzero"`
}

func (s KVMNicInfoZerotierConfiguration) Validate() error {

	return validator.Validate(s)
}
