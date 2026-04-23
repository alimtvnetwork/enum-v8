package dockercmdnames

import (
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:           "Invalid",
		Help:              "Help",
		Log:               "Log",
		Status:            "Status",
		Install:           "Install",
		UnInstall:         "UnInstall",
		Upgrade:           "Upgrade",
		Update:            "Update",
		InstallFix:        "InstallFix",
		SudoGroupFix:      "SudoGroupFix",
		Run:               "Run",
		Cmd:               "Cmd",
		Get:               "Get",
		RmIt:              "RmIt",
		Purge:             "Purge",
		Clean:             "Clean",
		RemoveAll:         "RemoveAll",
		AllImages:         "AllImages",
		AllContainer:      "AllContainer",
		AllPorts:          "AllPorts",
		WhichIp:           "WhichIp",
		ChangeIp:          "ChangeIp",
		DockerFile:        "DockerFile",
		CreateDockerFile:  "CreateDockerFile",
		ImagesByPorts:     "ImagesByPorts",
		ImageByNames:      "ImageByNames",
		ContainersByPorts: "ContainersByPorts",
		ContainersByName:  "ContainersByName",
		DockerCompose:     "DockerCompose",
		RemoveImageWhich:  "RemoveImageWhich",
		ListImages:        "ListImages",
		ListContainers:    "ListContainers",
		ListProcesses:     "ListProcesses",
		MacroStart:        "MacroStart",
		MacroEnd:          "MacroEnd",
		StateHistories:    "StateHistories",
		Export:            "Export",
		DumpJson:          "DumpJson",
		ImportJson:        "ImportJson",
		Histories:         "Histories",
		List:              "List",
		ListJson:          "ListJson",
		Backup:            "Backup",
		Import:            "Import",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Invalid),
		Ranges[:])
)
