package verifiertriggertype

import (
	"github.com/alimtvnetwork/core-v9/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v9/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:           "Invalid",
		AllComplete:       "AllComplete",
		AfterRestart:      "AfterRestart",
		AfterNetworkReset: "AfterNetworkReset",
	}

	RangesMap = map[string]Variant{
		"Invalid":           Invalid,
		"AllComplete":       AllComplete,
		"AfterRestart":      AfterRestart,
		"AfterNetworkReset": AfterNetworkReset,
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.SafeTypeName(Invalid),
		Ranges[:])
)
