package main

import (
	"gopkg.in/validator.v2"
)

// Zerotier route
type ZeroTierRoute struct {
	Flags  int    `json:"flags" validate:"nonzero"`
	Metric int    `json:"metric" validate:"nonzero"`
	Target string `json:"target" validate:"nonzero"`
	Via    string `json:"via" validate:"nonzero"`
}

func (s ZeroTierRoute) Validate() error {

	return validator.Validate(s)
}
