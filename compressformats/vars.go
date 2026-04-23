package compressformats

import (
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	ranges = [...]string{
		Zip:     "Zip",
		Tar:     "Tar",
		TarGZ:   "TarGZ",
		TarXZ:   "TarXZ",
		TarBz2:  "TarBz2",
		Invalid: "Invalid",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingFirstItemSliceAllCases(
		Zip,
		ranges[:])
)
