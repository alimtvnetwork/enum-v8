package sysgroupcmdnames

import (
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:                         "Invalid",
		Help:                            "Help",
		Create:                          "Create",
		Remove:                          "Remove",
		Rename:                          "Rename",
		AddGroupMembers:                 "AddGroupMembers",
		DeleteGroupMembers:              "DeleteGroupMembers",
		AddToSudoers:                    "AddToSudoers",
		RemoveSudoers:                   "RemoveSudoers",
		CreateOrUpdate:                  "CreateOrUpdate",
		OnlyKeepUsersToGroup:            "OnlyKeepUsersToGroup",
		OnlyKeepUsersToSudoers:          "OnlyKeepUsersToSudoers",
		RemoveAllUsers:                  "RemoveAllUsers",
		AddUsersToGroupUsingFile:        "AddUsersToGroupUsingFile",
		OnlyKeepUsersToGroupUsingFile:   "OnlyKeepUsersToGroupUsingFile",
		OnlyKeepUsersToSudoersUsingFile: "OnlyKeepUsersToSudoersUsingFile",
		Profile:                         "Profile",
		Details:                         "Details",
		ListJson:                        "ListJson",
		List:                            "List",
		Search:                          "Search",
		Histories:                       "Histories",
		StateChange:                     "StateChange",
		Backup:                          "Backup",
		Import:                          "Import",
	}

	aliasMap = map[string]byte{
		"create":                               Create.Value(),
		"remove":                               Remove.Value(),
		"rename-group":                         Rename.Value(),
		"add-members":                          AddGroupMembers.Value(),
		"delete-members":                       DeleteGroupMembers.Value(),
		"add-to-sudoers":                       AddToSudoers.Value(),
		"remove-sudoers":                       RemoveSudoers.Value(),
		"create-update":                        CreateOrUpdate.Value(),
		"only-members-keep":                    OnlyKeepUsersToGroup.Value(),
		"only-members-keep-sudoers":            OnlyKeepUsersToSudoers.Value(),
		"remove-all":                           RemoveAllUsers.Value(),
		"add-members-using-file":               AddUsersToGroupUsingFile.Value(),
		"only-members-keep-using-file":         OnlyKeepUsersToGroupUsingFile.Value(),
		"only-members-keep-sudoers-using-file": OnlyKeepUsersToSudoersUsingFile.Value(),
		"profile":                              Profile.Value(),
		"details":                              Details.Value(),
		"list-json":                            ListJson.Value(),
	}

	BasicEnumImpl = enumimpl.New.BasicByte.DefaultWithAliasMap(
		Invalid,
		Ranges[:],
		aliasMap)
)
