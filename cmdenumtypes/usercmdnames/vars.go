package usercmdnames

import (
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:            "Invalid",
		Help:               "Help",
		Add:                "Add",
		Remove:             "Remove",
		RemoveOnExist:      "RemoveOnExist",
		AddOrUpdate:        "AddOrUpdate",
		Update:             "Update",
		UpdateOnExist:      "UpdateOnExist",
		HasUser:            "HasUser",
		AttachUserWithRole: "AttachUserWithRole",
		AddUserWithPolicy:  "AddUserWithPolicy",
		AssignRole:         "AssignRole",
		DetachRole:         "DetachRole",
		RemoveRole:         "RemoveRole",
		List:               "List",
		Search:             "Search",
		Histories:          "Histories",
		StateChange:        "StateChange",
		Backup:             "Backup",
		Import:             "Import",
	}

	aliasMap = map[string]byte{
		"delete":                Remove.Value(),
		"remove-on-exist":       RemoveOnExist.Value(),
		"add-or-update":         AddOrUpdate.Value(),
		"update-on-exist":       UpdateOnExist.Value(),
		"has-user":              HasUser.Value(),
		"attach-user-with-role": AttachUserWithRole.Value(),
		"attach-user-role":      AttachUserWithRole.Value(),
		"assign-role":           AttachUserWithRole.Value(),
		"detach-role":           AttachUserWithRole.Value(),
		"remove-role":           AttachUserWithRole.Value(),
		"delete-role":           AttachUserWithRole.Value(),
		"list-all":              List.Value(),
		"state-change":          StateChange.Value(),
	}

	BasicEnumImpl = enumimpl.New.BasicByte.DefaultWithAliasMapAllCases(
		Invalid,
		Ranges[:],
		aliasMap)
)
