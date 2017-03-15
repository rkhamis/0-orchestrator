package main

import (
	"gopkg.in/validator.v2"
)

// Extended disk information. See https://github.com/karelzak/util-linux/blob/master/misc-utils/lsblk.c#L156
type DiskExtInfo struct {
	Alignment  int    `json:"alignment" validate:"nonzero"`
	DiscAln    int    `json:"discAln" validate:"nonzero"`
	DiscGran   int    `json:"discGran" validate:"nonzero"`
	DiscMax    int    `json:"discMax" validate:"nonzero"`
	DiscZero   int    `json:"discZero" validate:"nonzero"`
	Fstype     string `json:"fstype" validate:"nonzero"`
	Group      string `json:"group" validate:"nonzero"`
	Hctl       string `json:"hctl" validate:"nonzero"`
	Hotplug    int    `json:"hotplug" validate:"nonzero"`
	Kname      string `json:"kname" validate:"nonzero"`
	Label      string `json:"label" validate:"nonzero"`
	LogSec     int    `json:"logSec" validate:"nonzero"`
	Maj_min    string `json:"maj_min" validate:"nonzero"`
	MinIO      int    `json:"minIO" validate:"nonzero"`
	Mode       string `json:"mode" validate:"nonzero"`
	Model      string `json:"model" validate:"nonzero"`
	Mountpoint string `json:"mountpoint" validate:"nonzero"`
	Name       string `json:"name" validate:"nonzero"`
	OptIO      int    `json:"optIO" validate:"nonzero"`
	Owner      string `json:"owner" validate:"nonzero"`
	Partflags  string `json:"partflags" validate:"nonzero"`
	Partlabel  string `json:"partlabel" validate:"nonzero"`
	Parttype   int    `json:"parttype" validate:"nonzero"`
	Partuuid   string `json:"partuuid" validate:"nonzero"`
	PhySec     int    `json:"phySec" validate:"nonzero"`
	Pkname     string `json:"pkname" validate:"nonzero"`
	Ra         int    `json:"ra" validate:"nonzero"`
	Rand       int    `json:"rand" validate:"nonzero"`
	Rev        string `json:"rev" validate:"nonzero"`
	Rm         int    `json:"rm" validate:"nonzero"`
	Ro         int    `json:"ro" validate:"nonzero"`
	Rota       int    `json:"rota" validate:"nonzero"`
	RqSize     int    `json:"rqSize" validate:"nonzero"`
	Sched      string `json:"sched" validate:"nonzero"`
	Serial     string `json:"serial" validate:"nonzero"`
	Size       int    `json:"size" validate:"nonzero"`
	State      string `json:"state" validate:"nonzero"`
	Subsystems string `json:"subsystems" validate:"nonzero"`
	Tran       string `json:"tran" validate:"nonzero"`
	Type       string `json:"type" validate:"nonzero"`
	Uuid       string `json:"uuid" validate:"nonzero"`
	Vendor     string `json:"vendor" validate:"nonzero"`
	Wsame      int    `json:"wsame" validate:"nonzero"`
	Wwn        string `json:"wwn" validate:"nonzero"`
}

func (s DiskExtInfo) Validate() error {

	return validator.Validate(s)
}
