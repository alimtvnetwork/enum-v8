package osgroupexecution

import (
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	ranges = [...]string{
		Invalid:            "Invalid",
		Create:             "Create",
		Delete:             "Delete",
		Update:             "Update",
		ManageByUsers:      "ManageByUsers",
		AddGroupsToSudoers: "AddGroupsToSudoers",
		GroupManage:        "GroupManage",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Create),
		ranges[:])
)
