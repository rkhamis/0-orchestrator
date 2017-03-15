package main

import (
	"gopkg.in/validator.v2"
)

type KVMCreate struct {
	Command Command   `json:"command,omitempty"`
	Domain  KVMDomain `json:"domain" validate:"nonzero"`
}

func (s KVMCreate) Validate() error {

	return validator.Validate(s)
}
