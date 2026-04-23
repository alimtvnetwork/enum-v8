package packageinstallmethod

import (
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
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
		coredynamic.TypeName(Invalid),
		Ranges[:])
)
