package linescomparetype

import (
	"github.com/alimtvnetwork/core-v9/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v9/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:      "Invalid",
		Equal:        "Equal",
		Less:         "Less",
		LessEqual:    "LessEqual",
		Greater:      "Greater",
		GreaterEqual: "GreaterEqual",
		NotEqual:     "NotEqual",
	}

	RangesMap = map[string]Variant{
		"Invalid":      Invalid,
		"Equal":        Equal,
		"Less":         Less,
		"LessEqual":    LessEqual,
		"Greater":      Greater,
		"GreaterEqual": GreaterEqual,
		"NotEqual":     NotEqual,
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.SafeTypeName(Invalid),
		Ranges[:])
)
