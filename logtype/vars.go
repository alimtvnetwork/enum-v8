package logtype

import (
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Silent:  "Silent",
		Success: "Success",
		Info:    "Info",
		Trace:   "Trace",
		Debug:   "Debug",
		Warning: "Warning",
		Error:   "Error",
		Fatal:   "Fatal",
		Panic:   "Panic",
		Custom:  "Custom",
		File:    "File",
		Pattern: "Pattern",
		Invalid: "Invalid",
	}

	TraceMap = map[Variant]bool{
		Trace: true,
		Debug: true,
	}

	ErrorMap = map[Variant]bool{
		Error: true,
		Fatal: true,
		Panic: true,
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Trace),
		Ranges[:])
)
