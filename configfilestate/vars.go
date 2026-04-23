package configfilestate

import (
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:                               "Invalid",
		Unchanged:                             "Unchanged",
		Permission:                            "Permission",
		Added:                                 "Added",
		Deleted:                               "Deleted",
		Modified:                              "Modified",
		SymbolicLinkAdded:                     "SymbolicLinkAdded",
		SymbolicLinkDelete:                    "SymbolicLinkDelete",
		ChmodChanged:                          "ChmodChanged",
		ChownChanged:                          "ChownChanged",
		LastModifiedDateChanged:               "LastModifiedDateChanged",
		SizeChanged:                           "SizeChanged",
		ChmodChownBothChanged:                 "ChmodChownBothChanged",
		ChmodOrChownOrLastModifiedDateChanged: "ChmodOrChownOrLastModifiedDateChanged",
		SizeOrChmodOrChownOrLastModifiedDateChanged: "SizeOrChmodOrChownOrLastModifiedDateChanged",
		MismatchFileOrDir: "MismatchFileOrDir",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.CreateUsingSlicePlusAliasMapOptions(
		true,
		Invalid,
		Ranges[:],
		nil)
)
