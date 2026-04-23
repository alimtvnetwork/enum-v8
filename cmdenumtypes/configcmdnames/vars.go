package configcmdnames

import (
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:             "Invalid",
		Help:                "Help",
		Log:                 "Log",
		Apply:               "Apply",
		Revert:              "Revert",
		Store:               "Store",
		DockerApply:         "DockerApply",
		ApplyDuringShutdown: "ApplyDuringShutdown",
		ApplyAfterReboot:    "ApplyAfterReboot",
		ApplyAfter:          "ApplyAfter",
		ApplyBefore:         "ApplyBefore",
		Histories:           "Histories",
		Backup:              "Backup",
		Import:              "Import",
		Export:              "Export",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Invalid),
		Ranges[:])
)
