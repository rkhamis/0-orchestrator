package main

import (
	"gopkg.in/validator.v2"
)

// Zerotier details
type ZeroTier struct {
	AllowDefault      bool             `json:"allowDefault" validate:"nonzero"`
	AllowGlobal       bool             `json:"allowGlobal" validate:"nonzero"`
	AllowManaged      bool             `json:"allowManaged" validate:"nonzero"`
	AssignedAddresses []string         `json:"assignedAddresses" validate:"nonzero"`
	Bridge            bool             `json:"bridge" validate:"nonzero"`
	BroadcastEnabled  bool             `json:"broadcastEnabled" validate:"nonzero"`
	Dhcp              bool             `json:"dhcp" validate:"nonzero"`
	Mac               string           `json:"mac" validate:"nonzero"`
	Mtu               int              `json:"mtu" validate:"nonzero"`
	Name              string           `json:"name" validate:"nonzero"`
	NetconfRevision   int              `json:"netconfRevision" validate:"nonzero"`
	Nwid              string           `json:"nwid" validate:"nonzero"`
	PortDeviceName    string           `json:"portDeviceName" validate:"nonzero"`
	PortError         int              `json:"portError" validate:"nonzero"`
	Routes            []ZeroTierRoute  `json:"routes" validate:"nonzero"`
	Status            string           `json:"status" validate:"nonzero"`
	Type              EnumZeroTierType `json:"type" validate:"nonzero"`
}

func (s ZeroTier) Validate() error {

	return validator.Validate(s)
}
