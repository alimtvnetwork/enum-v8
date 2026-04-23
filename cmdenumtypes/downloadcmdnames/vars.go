package downloadcmdnames

import (
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:            "Invalid",
		Help:               "Help",
		Log:                "Log",
		Status:             "Status",
		Install:            "Install",
		Uninstall:          "Uninstall",
		To:                 "To",
		Temp:               "Temp",
		Decompress:         "Decompress",
		TempDecompress:     "TempDecompress",
		Verify:             "Verify",
		DownloadVerify:     "DownloadVerify",
		Schedule:           "Schedule",
		ScheduleTemp:       "ScheduleTemp",
		ScheduleDecompress: "ScheduleDecompress",
		List:               "List",
		ListJson:           "ListJson",
		Backup:             "Backup",
		Import:             "Import",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Invalid),
		Ranges[:])
)
