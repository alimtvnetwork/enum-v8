package macrocmdnames

import (
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:                    "Invalid",
		Help:                       "Help",
		Log:                        "Log",
		Status:                     "Status",
		RemoveMarcoStates:          "RemoveMarcoStates",
		InjectMacroAt:              "InjectMacroAt",
		Create:                     "Create",
		Session:                    "Session",
		CreateSession:              "CreateSession",
		SessionStart:               "SessionStart",
		EndSession:                 "EndSession",
		SessionStatus:              "SessionStatus",
		CreateOrUpdate:             "CreateOrUpdate",
		Remove:                     "Remove",
		RemoveOnExist:              "RemoveOnExist",
		StateChange:                "StateChange",
		SearchMacroNames:           "SearchMacroNames",
		SearchMacroByKey:           "SearchMacroByKey",
		SearchMacroByService:       "SearchMacroByService",
		SearchMacroByUserService:   "SearchMacroByUserSer",
		ServiceNames:               "ServiceNames",
		HasUserMacro:               "HasUserMacro",
		HasServiceMarcos:           "HasServiceMarcos",
		ListMacros:                 "ListMacros",
		ListMarcoStates:            "ListMarcoStates",
		ListMacroByServiceName:     "ListMacroByServiceNa",
		ListMacroByUserServiceName: "ListMacroByUserServiceName",
		ListMacroByKey:             "ListMacroByKey",
		LastMacro:                  "LastMacro",
		MacroByUser:                "MacroByUser",
		CompileMacro:               "CompileMacro",
		CompileMacroToInstruction:  "CompileMacroToInstruction",
		CompileMacroToPkg:          "CompileMacroToPkg",
		DumpMacro:                  "DumpMacro",
		ClearAll:                   "ClearAll",
		List:                       "List",
		ListJson:                   "ListJson",
		Histories:                  "Histories",
		Backup:                     "Backup",
		ExportSpecific:             "ExportSpecific",
		ImportSpecific:             "ImportSpecific",
		Export:                     "Export",
		Import:                     "Import",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Invalid),
		Ranges[:])
)
