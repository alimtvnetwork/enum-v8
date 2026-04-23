package timeunit

import (
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
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
		coredynamic.TypeName(Invalid),
		Ranges[:])
)
