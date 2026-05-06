package envpathcmdnames

import (
	"github.com/alimtvnetwork/core-v9/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v9/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:    "Invalid",
		Help:       "Help",
		Append:     "Append",
		Remove:     "Remove",
		TempAppend: "TempAppend",
		TempRemove: "TempRemove",
		Source:     "Source",
		Fix:        "Fix",
		OrderFix:   "OrderFix",
		HasIssues:  "HasIssues",
		List:       "List",
		ListJson:   "ListJson",
		Backup:     "Backup",
		Import:     "Import",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.SafeTypeName(Invalid),
		Ranges[:])
)
