package main

type EnumCommandResultState string

const (
	EnumCommandResultStateSCHEDULED      EnumCommandResultState = "SCHEDULED"
	EnumCommandResultStateRUNNING        EnumCommandResultState = "RUNNING"
	EnumCommandResultStateSUCCESS        EnumCommandResultState = "SUCCESS"
	EnumCommandResultStateKILLED         EnumCommandResultState = "KILLED"
	EnumCommandResultStateKILLED_TIMEOUT EnumCommandResultState = "KILLED_TIMEOUT"
	EnumCommandResultStateFAILED         EnumCommandResultState = "FAILED"
)
