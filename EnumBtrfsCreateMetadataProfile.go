package main

type EnumBtrfsCreateMetadataProfile string

const (
	EnumBtrfsCreateMetadataProfileraid0  EnumBtrfsCreateMetadataProfile = "raid0"
	EnumBtrfsCreateMetadataProfileraid1  EnumBtrfsCreateMetadataProfile = "raid1"
	EnumBtrfsCreateMetadataProfileraid5  EnumBtrfsCreateMetadataProfile = "raid5"
	EnumBtrfsCreateMetadataProfileraid6  EnumBtrfsCreateMetadataProfile = "raid6"
	EnumBtrfsCreateMetadataProfileraid10 EnumBtrfsCreateMetadataProfile = "raid10"
	EnumBtrfsCreateMetadataProfiledup    EnumBtrfsCreateMetadataProfile = "dup"
	EnumBtrfsCreateMetadataProfilesingle EnumBtrfsCreateMetadataProfile = "single"
)
