package timeunit

import (
	"github.com/alimtvnetwork/core-v9/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v9/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:     "Invalid",
		Millisecond: "Millisecond",
		Second:      "Second",
		Minute:      "Minute",
		Hour:        "Hour",
		Day:         "Day",
		Month:       "Month",
		Year:        "Year",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.SafeTypeName(Invalid),
		Ranges[:])
)
