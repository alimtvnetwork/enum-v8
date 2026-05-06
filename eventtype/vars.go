package eventtype

import (
	"github.com/alimtvnetwork/core-v9/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v9/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid: "Invalid",
		Log:     "Log",
		Success: "Success",
		Error:   "Error",
		Failure: "Failure",
		File:    "File",
		Custom:  "Custom",
	}

	ErrorMap = map[Variant]bool{
		Failure: true,
		Error:   true,
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.SafeTypeName(Success),
		Ranges[:])
)
