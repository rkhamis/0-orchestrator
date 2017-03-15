package main

import (
	"gopkg.in/validator.v2"
)

// Key value pair
type KeyValuePair struct {
	Name  string `json:"name" validate:"nonzero"`
	Value string `json:"value" validate:"nonzero"`
}

func (s KeyValuePair) Validate() error {

	return validator.Validate(s)
}
