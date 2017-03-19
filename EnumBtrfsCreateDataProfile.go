package main

type EnumBtrfsCreateDataProfile string

const (
	EnumBtrfsCreateDataProfileraid0  EnumBtrfsCreateDataProfile = "raid0"
	EnumBtrfsCreateDataProfileraid1  EnumBtrfsCreateDataProfile = "raid1"
	EnumBtrfsCreateDataProfileraid5  EnumBtrfsCreateDataProfile = "raid5"
	EnumBtrfsCreateDataProfileraid6  EnumBtrfsCreateDataProfile = "raid6"
	EnumBtrfsCreateDataProfileraid10 EnumBtrfsCreateDataProfile = "raid10"
	EnumBtrfsCreateDataProfiledup    EnumBtrfsCreateDataProfile = "dup"
	EnumBtrfsCreateDataProfilesingle EnumBtrfsCreateDataProfile = "single"
)
