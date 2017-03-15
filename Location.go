package main

import (
	"gopkg.in/validator.v2"
)

// Url to object details
type Location struct {
	Id   string `json:"id" validate:"nonzero"`
	Name string `json:"name" validate:"nonzero"`
	Url  string `json:"url" validate:"nonzero"`
}

func (s Location) Validate() error {

	return validator.Validate(s)
}
