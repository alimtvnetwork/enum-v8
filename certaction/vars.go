package certaction

import (
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid: "Invalid",
		Create:  "Create",
		Renew:   "Renew",
		Revoke:  "Revoke",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.Default(
		Invalid,
		Ranges[:])
)
