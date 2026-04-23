package userrolecmdnames

import (
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:                 "Invalid",
		Help:                    "Help",
		AddOrUpdatePolicy:       "AddOrUpdatePolicy",
		AddPolicy:               "AddPolicy",
		RemovePolicy:            "RemovePolicy",
		RemovePolicyOnExist:     "RemovePolicyOnExist",
		HasPolicy:               "HasPolicy",
		HasRole:                 "HasRole",
		AssignRole:              "AssignRole",
		AssignRoleOrSkipOnExist: "AssignRoleOrSkipOnExist",
		RemoveRole:              "RemoveRole",
		DetachRole:              "DetachRole",
		List:                    "List",
		Search:                  "Search",
		Histories:               "Histories",
		StateChange:             "StateChange",
		Backup:                  "Backup",
		Import:                  "Import",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.DefaultAllCases(
		Invalid,
		Ranges[:])
)
