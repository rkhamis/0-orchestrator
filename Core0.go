package main

import (
	"gopkg.in/validator.v2"
)

// Core0 node in the g8os grid
type Core0 struct {
	Hostname string `json:"hostname" validate:"nonzero"`
	Id       string `json:"id" validate:"nonzero"`
}

func (s Core0) Validate() error {

	return validator.Validate(s)
}
