package conntrackstate

import (
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:     "Invalid",
		New:         "NEW",
		Established: "ESTABLISHED",
		Related:     "RELATED",
		Untracked:   "UNTRACKED",
		Snat:        "SNAT",
		Dnat:        "DNAT",
	}

	aliasMap = map[string]byte{
		"":            Invalid.ValueByte(),
		"new":         New.ValueByte(),
		"established": Established.ValueByte(),
		"related":     Related.ValueByte(),
		"untracked":   Untracked.ValueByte(),
		"snat":        Snat.ValueByte(),
		"dnat":        Dnat.ValueByte(),
	}

	BasicEnumImpl = enumimpl.New.BasicByte.CreateUsingSlicePlusAliasMapOptions(
		true,
		Invalid,
		Ranges[:],
		aliasMap)
)
