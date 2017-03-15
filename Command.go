package main

import (
	"gopkg.in/validator.v2"
)

// Command that can be executed in the Core0/CoreX. More information can be found here https://github.com/g8os/core0/blob/master/docs/commands.md
type Command struct {
	Id              string `json:"id,omitempty"`
	LogLevels       []int  `json:"logLevels,omitempty"`
	MaxRestart      int    `json:"maxRestart,omitempty"`
	MaxTime         int    `json:"maxTime,omitempty"`
	Queue           string `json:"queue,omitempty"`
	RecurringPeriod int    `json:"recurringPeriod,omitempty"`
	StatsInterval   int    `json:"statsInterval,omitempty"`
	Tags            string `json:"tags,omitempty"`
}

func (s Command) Validate() error {

	return validator.Validate(s)
}
