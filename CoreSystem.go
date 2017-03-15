package main

import (
	"gopkg.in/validator.v2"
)

// Arguments for a core.system command
type CoreSystem struct {
	Args        []string       `json:"args" validate:"nonzero"`
	Environment []KeyValuePair `json:"environment" validate:"nonzero"`
	Name        string         `json:"name" validate:"nonzero"`
	Pwd         string         `json:"pwd" validate:"nonzero"`
	Stdin       string         `json:"stdin" validate:"nonzero"`
}

func (s CoreSystem) Validate() error {

	return validator.Validate(s)
}
