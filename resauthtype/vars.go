package resauthtype

import (
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:         "Invalid",
		AllAccess:       "AllAccess",
		Error:           "Error",
		Warning:         "Warning",
		Restricted:      "Restricted",
		UnAuthorized:    "UnAuthorized",
		PermissionIssue: "PermissionIssue",
		Forbidden:       "Forbidden",
		ReadAccess:      "ReadAccess",
		WriteAccess:     "WriteAccess",
		CreateAccess:    "CreateAccess",
		EditAccess:      "EditAccess",
		AccessGranted:   "AccessGranted",
		AccessRejected:  "AccessRejected",
	}

	errorMap = map[Variant]bool{
		Error:           true,
		Restricted:      true,
		UnAuthorized:    true,
		PermissionIssue: true,
		Forbidden:       true,
		AccessRejected:  true,
	}

	allAccessMap = map[Variant]bool{
		AllAccess:     true,
		AccessGranted: true,
	}

	accessMap = map[Variant]bool{
		AllAccess:     true,
		AccessGranted: true,
		ReadAccess:    true,
		WriteAccess:   true,
		CreateAccess:  true,
		EditAccess:    true,
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Invalid),
		Ranges[:])
)
