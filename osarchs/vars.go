package osarchs

import (
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core-v8/osconsts"
)

var (
	architectures = [...]string{
		X32:     "x32",
		X64:     "x64",
		Invalid: "Invalid",
	}

	// https://t.ly/XHVi
	aliasMap = map[string]byte{
		"":      Invalid.ValueByte(),
		"32":    X32.ValueByte(),
		"x32":   X32.ValueByte(),
		"X32":   X32.ValueByte(),
		"386":   X32.ValueByte(),
		"arm":   X32.ValueByte(),
		"i386":  X32.ValueByte(),
		"x86":   X32.ValueByte(),
		"86":    X32.ValueByte(),
		"64":    X64.ValueByte(),
		"x64":   X64.ValueByte(),
		"X64":   X64.ValueByte(),
		"amd64": X64.ValueByte(),
		"arm64": X64.ValueByte(),
	}

	CurrentArch   = Get(osconsts.CurrentSystemArchitecture)
	BasicEnumImpl = enumimpl.New.BasicByte.DefaultWithAliasMap(
		Invalid,
		architectures[:],
		aliasMap)
)
