package envvarscmdnames

import (
	"github.com/alimtvnetwork/core-v9/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v9/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:       "Invalid",
		Help:          "Help",
		AppendPath:    "AppendPath",
		Set:           "Set",
		Remove:        "Remove",
		TempSet:       "TempSet",
		TempRemove:    "TempRemove",
		Source:        "Source",
		HasIssues:     "HasIssues",
		List:          "List",
		ListJson:      "ListJson",
		ListPaths:     "ListPaths",
		ListPathsJson: "ListPathsJson",
		Backup:        "Backup",
		Import:        "Import",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.SafeTypeName(Invalid),
		Ranges[:])
)
