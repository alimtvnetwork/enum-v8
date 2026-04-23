package verifiertriggertype

import (
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
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
		coredynamic.TypeName(Invalid),
		Ranges[:])
)
