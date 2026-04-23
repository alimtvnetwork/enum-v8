package sqljointype

import (
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Default:   "Default",
		Invalid:   "Invalid",
		Join:      "Join",
		Inner:     "Inner",
		Left:      "Left",
		Right:     "Right",
		FullOuter: "FullOuter",
		Cross:     "Cross",
	}

	InnerJoinMap = map[Variant]bool{
		Default: true, // inner join
		Join:    true,
		Inner:   true,
	}

	OuterJoinMap = map[Variant]bool{
		Left:      true, // inner join
		Right:     true,
		FullOuter: true,
	}

	SqlSyntax = [...]string{
		Default:   "JOIN", // inner join
		Invalid:   "",
		Join:      "JOIN",
		Inner:     "INNER JOIN",
		Left:      "LEFT JOIN",
		Right:     "RIGHT JOIN",
		FullOuter: "FULL OUTER JOIN", // Reference : https://www.w3schools.com/sql/sql_join_full.asp
		Cross:     "CROSS JOIN",
	}

	RangesMap = map[string]Variant{
		"Default":   Default, // inner join
		"Invalid":   Invalid,
		"":          Invalid,
		"Join":      Join,
		"Inner":     Inner,
		"Left":      Left,
		"Right":     Right,
		"FullOuter": FullOuter,
		"Cross":     Cross,
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Left),
		Ranges[:])
)
