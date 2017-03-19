package main

import (
	"gopkg.in/validator.v2"
)

// Result object of a command
type CommandResult struct {
	Data      string                 `json:"data" validate:"nonzero"`
	Id        string                 `json:"id" validate:"nonzero"`
	Level     string                 `json:"level" validate:"nonzero"`
	Name      EnumCommandResultName  `json:"name" validate:"nonzero"`
	Starttime int                    `json:"starttime" validate:"nonzero"`
	State     EnumCommandResultState `json:"state" validate:"nonzero"`
	Stderr    string                 `json:"stderr" validate:"nonzero"`
	Stdout    string                 `json:"stdout" validate:"nonzero"`
}

func (s CommandResult) Validate() error {

	return validator.Validate(s)
}
