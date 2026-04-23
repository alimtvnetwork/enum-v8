package fail2bancmdnames

import (
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:        "Invalid",
		Help:           "Help",
		Enable:         "Enable",
		Disable:        "Disable",
		Client:         "Client",
		EnableBan:      "EnableBan",
		DisableBan:     "DisableBan",
		EnableJail:     "EnableJail",
		DisableJail:    "DisableJail",
		Status:         "Status",
		WhichJails:     "WhichJails",
		JailLogs:       "JailLogs",
		Log:            "Log",
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
