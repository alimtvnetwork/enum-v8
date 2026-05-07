package httpstatusfamily

import (
	"github.com/alimtvnetwork/core-v9/coreimpl/enumimpl"
)

var (
	ranges = [...]string{
		Informational: "Informational",
		Successful:    "Successful",
		Redirection:   "Redirection",
		ClientError:   "ClientError",
		ServerError:   "ServerError",
		Invalid:       "Invalid",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingFirstItemSliceAllCases(
		Informational,
		ranges[:])
)
