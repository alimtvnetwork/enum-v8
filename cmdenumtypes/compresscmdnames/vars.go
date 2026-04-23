package compresscmdnames

import (
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:                  "Invalid",
		Help:                     "Help",
		Install:                  "Install",
		Compress:                 "Compress",
		Tar:                      "Tar",
		Gz:                       "Gz",
		Zip:                      "Zip",
		Unzip:                    "Unzip",
		Extract:                  "Extract",
		Decompress:               "Decompress",
		DecompressInstall:        "DecompressInstall",
		DownloadDecompress:       "DownloadDecompress",
		DownloadDecompressRemove: "DownloadDecompressRemove",
		SneakList:                "SneakList",
		SneakSearch:              "SneakSearch",
		ListJson:                 "ListJson",
		List:                     "List",
		Histories:                "Histories",
		StateChange:              "StateChange",
		MacroHistories:           "MacroHistories",
		RemoveMacro:              "RemoveMacro",
		ExportMacro:              "ExportMacro",
		ImportMacro:              "ImportMacro",
		Backup:                   "Backup",
		Import:                   "Import",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Invalid),
		Ranges[:])
)
