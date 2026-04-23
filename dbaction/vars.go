package dbaction

import (
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:        "Invalid",
		Create:         "Create",
		Update:         "Update",
		Delete:         "Delete",
		Read:           "Read",
		CreateOrUpdate: "CreateOrUpdate",
		SkipOnExist:    "SkipOnExist",
		SkipOnNonExist: "SkipOnNonExist",
		DropOnExist:    "DropOnExist",
		UpdateOnExist:  "UpdateOnExist",
	}

	onExistCheck = map[Variant]bool{
		SkipOnExist:   true,
		DropOnExist:   true,
		UpdateOnExist: true,
	}

	editLogically = map[Variant]bool{
		Update:         true,
		CreateOrUpdate: true,
		UpdateOnExist:  true,
	}

	crudOnlyLogically = map[Variant]bool{
		Create: true,
		Update: true,
		Delete: true,
		Read:   true,
	}

	readEditLogically = map[Variant]bool{
		Read:           true,
		Update:         true,
		CreateOrUpdate: true,
	}

	updateOrRemoveLogicallyMap = map[Variant]bool{
		Update:         true,
		Delete:         true,
		DropOnExist:    true,
		CreateOrUpdate: true,
	}

	dropMap = map[Variant]bool{
		Delete:      true,
		DropOnExist: true,
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Invalid),
		Ranges[:])
)
