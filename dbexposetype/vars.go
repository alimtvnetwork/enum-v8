package dbexposetype

import (
	"github.com/alimtvnetwork/core-v9/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v9/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:    "Invalid",
		AnyIp:      "AnyIp",
		SpecificIp: "SpecificIp",
	}

	RangesMap = map[string]Variant{
		"Invalid":    Invalid,
		"AnyIp":      AnyIp,
		"SpecificIp": SpecificIp,
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.SafeTypeName(Invalid),
		Ranges[:])
)
