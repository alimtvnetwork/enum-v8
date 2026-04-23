package osarchs

import (
	"github.com/alimtvnetwork/core-v8/osconsts"
)

func Get(rawGoArch string) Architecture {
	_, isIn32Bit := osconsts.X32ArchitecturesMap[rawGoArch]

	if isIn32Bit {
		return X32
	}

	_, is64Bit := osconsts.X64ArchitecturesMap[rawGoArch]

	if is64Bit {
		return X64
	}

	return Invalid
}
