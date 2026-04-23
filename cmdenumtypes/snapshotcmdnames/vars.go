package snapshotcmdnames

import (
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:                  "Invalid",
		Help:                     "Help",
		Log:                      "Log",
		Status:                   "Status",
		AddNew:                   "AddNew",
		Checkout:                 "Checkout",
		Push:                     "Push",
		Pull:                     "Pull",
		Revert:                   "Revert",
		Restore:                  "Restore",
		Clone:                    "Clone",
		Deploy:                   "Deploy",
		Remove:                   "Remove",
		Sync:                     "Sync",
		StateOfKey:               "StateOfKey",
		DeployLastState:          "DeployLastState",
		PullLastState:            "PullLastState",
		PushLastState:            "PushLastState",
		ListServices:             "ListServices",
		ListStateByName:          "ListStateByName",
		ListStatesByServiceNames: "ListStatesByServiceNames",
		ListTopStates:            "ListTopStates",
		ListJsonServices:         "ListJsonServices",
		ListByKey:                "ListByKey",
		SearchKeyContains:        "SearchKeyContains",
		ServiceNames:             "ServiceNames",
		VerifyState:              "VerifyState",
		Clean:                    "Clean",
		RemoveAll:                "RemoveAll",
		Search:                   "Search",
		List:                     "List",
		Histories:                "Histories",
		Backup:                   "Backup",
		Import:                   "Import",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.DefaultAllCases(
		Invalid,
		Ranges[:])
)
