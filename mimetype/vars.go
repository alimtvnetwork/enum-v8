package mimetype

import (
	"github.com/alimtvnetwork/core-v9/coreimpl/enumimpl"
)

var (
	ranges = [...]string{
		Application: "Application",
		Audio:       "Audio",
		Font:        "Font",
		Image:       "Image",
		Message:     "Message",
		Model:       "Model",
		Multipart:   "Multipart",
		Text:        "Text",
		Video:       "Video",
		Invalid:     "Invalid",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingFirstItemSliceAllCases(
		Application,
		ranges[:])
)
