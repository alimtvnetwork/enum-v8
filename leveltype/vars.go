package leveltype

import (
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid: "Invalid",
		Level1:  "Level1",
		Level2:  "Level2",
		Level3:  "Level3",
		Level4:  "Level4",
		Level5:  "Level5",
		Level6:  "Level6",
		Level7:  "Level7",
		Level8:  "Level8",
		Level9:  "Level9",
		Level10: "Level10",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Level1),
		Ranges[:])
)
