package httpmethodtype

import (
	"github.com/alimtvnetwork/core-v9/coreimpl/enumimpl"
)

var (
	ranges = [...]string{
		Get:     "Get",
		Post:    "Post",
		Put:     "Put",
		Patch:   "Patch",
		Delete:  "Delete",
		Head:    "Head",
		Options: "Options",
		Invalid: "Invalid",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingFirstItemSliceAllCases(
		Get,
		ranges[:])
)
