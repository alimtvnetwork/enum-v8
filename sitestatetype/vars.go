package sitestatetype

import (
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:    "Invalid",
		NewlyAdded: "NewlyAdded",
		Removed:    "Removed",
		Unchanged:  "Unchanged",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.Default(
		Invalid,
		Ranges[:])
)
