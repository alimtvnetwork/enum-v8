package runtype

import (
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:         "Invalid",
		Now:             "Now",
		OnReboot:        "OnReboot",
		OnShutdown:      "OnShutdown",
		OnEveryReboot:   "OnEveryReboot",
		OnEveryShutdown: "OnEveryShutdown",
		OnFailRetry:     "OnFailRetry",
		EveryMinute:     "EveryMinute",
		EveryHour:       "EveryHour",
		EveryDay:        "EveryDay",
		EveryMonth:      "EveryMonth",
		EveryYear:       "EveryYear",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Invalid),
		Ranges[:])
)
