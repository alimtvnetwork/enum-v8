package strtype

import (
	"os"

	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coredata/corestr"
)

type FileReader interface {
	FilePath() string

	OpenFile() (*os.File, error)
	OpenFileLock() (*os.File, error)

	Type() (Variant, error)
	TypeLock() (Variant, error)

	TrimText() (string, error)
	TrimTextLock() (string, error)

	TrimLine() (string, error)
	TrimLineLock() (string, error)

	Line() (string, error)
	LineLock() (string, error)

	Text() (string, error)
	TextLock() (string, error)

	String() (string, error)
	StringLock() (string, error)

	Strings() ([]string, error)
	StringsLock() ([]string, error)

	SimpleSlice() (*corestr.SimpleSlice, error)
	SimpleSliceLock() (*corestr.SimpleSlice, error)

	Lines() ([]string, error)
	LinesLock() ([]string, error)

	NonEmptyLines() ([]string, error)
	NonEmptyLinesLock() ([]string, error)

	TrimNonEmptyLines() ([]string, error)
	TrimNonEmptyLinesLock() ([]string, error)

	TrimNonWhitespaceLines() ([]string, error)
	TrimNonWhitespaceLinesLock() ([]string, error)

	Bytes() ([]byte, error)
	BytesLock() ([]byte, error)

	Raw() ([]byte, error)
	RawLock() ([]byte, error)

	JsonResult() (*corejson.Result, error)
	JsonResultLock() (*corejson.Result, error)

	RawAsJsonResult() *corejson.Result
	RawAsJsonResultLock() *corejson.Result
}
