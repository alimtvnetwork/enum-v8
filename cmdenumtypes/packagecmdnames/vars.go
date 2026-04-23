package packagecmdnames

import (
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:        "Invalid",
		Help:           "Help",
		Install:        "Install",
		Cleanup:        "Cleanup",
		Reinstall:      "Reinstall",
		Lock:           "Lock",
		Uninstall:      "Uninstall",
		Sync:           "Sync",
		Download:       "Download",
		GitClone:       "GitClone",
		Clone:          "Clone",
		Revert:         "Revert",
		Macro:          "Macro",
		Create:         "Create",
		Push:           "Push",
		List:           "List",
		Search:         "Search",
		Histories:      "Histories",
		StateChange:    "StateChange",
		MacroHistories: "MacroHistories",
		RemoveMacro:    "RemoveMacro",
		ExportMacro:    "ExportMacro",
		ImportMacro:    "ImportMacro",
		Backup:         "Backup",
		Import:         "Import",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.DefaultAllCases(
		Invalid,
		Ranges[:])
)
