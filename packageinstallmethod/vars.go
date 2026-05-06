package packageinstallmethod

import (
	"github.com/alimtvnetwork/core-v9/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v9/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:       "Invalid",
		Url:           "Url",
		OsPackages:    "OsPackages",
		AdvanceScript: "AdvanceScript",
	}

	RangesMap = map[string]Variant{
		"Invalid":       Invalid,
		"Url":           Url,
		"OsPackages":    OsPackages,
		"AdvanceScript": AdvanceScript,
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.SafeTypeName(Invalid),
		Ranges[:])
)
