package servicescmdnames

import (
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:           "Invalid",
		Help:              "Help",
		Log:               "Log",
		Status:            "Status",
		Upgrade:           "Upgrade",
		Install:           "Install",
		UnInstall:         "UnInstall",
		ApplyKnownFixes:   "ApplyKnownFixes",
		ImportAutoFix:     "ImportAutoFix",
		Info:              "Info",
		Update:            "Update",
		AutoFix:           "AutoFix",
		Nginx:             "Nginx",
		Apache:            "Apache",
		Healthcare:        "Healthcare",
		Logger:            "Logger",
		Cron:              "Cron",
		DbServer:          "DbServer",
		PostgreSql:        "PostgreSql",
		MySql:             "MySql",
		TaskRunner:        "TaskRunner",
		Webserver:         "Webserver",
		AppWebBack:        "AppWebBack",
		AppWebFront:       "AppWebFront",
		AppCommunicator:   "AppCommunicator",
		RestartServices:   "RestartServices",
		SchedulerServices: "SchedulerServices",
		Backup:            "Backup",
		Import:            "Import",
		Export:            "Export",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.DefaultAllCases(
		Invalid,
		Ranges[:])
)
