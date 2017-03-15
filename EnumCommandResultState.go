package main

type EnumCommandResultState string

const (
	EnumCommandResultStateSCHEDULED      EnumCommandResultState = "SCHEDULED"
	EnumCommandResultStateRUNNING        EnumCommandResultState = "RUNNING"
	EnumCommandResultStateSUCCES         EnumCommandResultState = "SUCCES"
	EnumCommandResultStateKILLED         EnumCommandResultState = "KILLED"
	EnumCommandResultStateKILLED_TIMEOUT EnumCommandResultState = "KILLED_TIMEOUT"
	EnumCommandResultStateFAILED         EnumCommandResultState = "FAILED"
)
