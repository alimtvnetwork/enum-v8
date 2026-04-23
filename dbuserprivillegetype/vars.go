package dbuserprivilegetype

import (
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:    "Invalid",
		All:        "All",
		Select:     "Select",
		Insert:     "Insert",
		Create:     "Create",
		Update:     "Update",
		Alter:      "Alter",
		Delete:     "Delete",
		Drop:       "Drop",
		Execute:    "Execute",
		Event:      "Event",
		CreateView: "CreateView",
		Index:      "Index",
		LockTables: "LockTables",
		References: "References",
		ShowView:   "ShowView",
		Trigger:    "Trigger",
	}

	editLogically = map[Variant]bool{
		Update: true,
		Insert: true,
		Alter:  true,
		Create: true,
	}

	createLogicallyMap = map[Variant]bool{
		Create: true,
		Insert: true,
	}

	crudOnlyLogically = map[Variant]bool{
		Create: true,
		Update: true,
		Delete: true,
		Insert: true,
		Alter:  true,
		Drop:   true,
	}

	readEditLogically = map[Variant]bool{
		Select: true,
		Update: true,
		Insert: true,
	}

	updateOrRemoveLogicallyMap = map[Variant]bool{
		Update: true,
		Delete: true,
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Invalid),
		Ranges[:])
)
