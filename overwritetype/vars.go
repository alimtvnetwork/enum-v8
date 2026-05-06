package overwritetype

import (
	"github.com/alimtvnetwork/core-v9/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v9/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:                            "Invalid",
		ForceWrite:                         "ForceWrite",
		SkipOnExistFiles:                   "SkipOnExistFiles",
		IgnoreRepeatInFolderNameExtraction: "IgnoreRepeatInFolderNameExtraction",
		Yes:                                "Yes",
		No:                                 "No",
	}

	RangesMap = map[string]Variant{
		"Invalid":                            Invalid,
		"ForceWrite":                         ForceWrite,
		"SkipOnExistFiles":                   SkipOnExistFiles,
		"IgnoreRepeatInFolderNameExtraction": IgnoreRepeatInFolderNameExtraction,
		"Yes":                                Yes,
		"No":                                 No,
	}

	overwriteMap = map[Variant]bool{
		ForceWrite:       true,
		ForceWriteRepeat: true,
		Yes:              true,
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.SafeTypeName(Invalid),
		Ranges[:])
)
