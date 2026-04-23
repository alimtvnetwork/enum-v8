package strtype

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/alimtvnetwork/core-v8/codestack"
	"github.com/alimtvnetwork/core-v8/constants"
	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coredata/corestr"
	"github.com/alimtvnetwork/core-v8/coredata/stringslice"
	"github.com/alimtvnetwork/core-v8/errcore"
)

type fileReader struct {
	filePath string
}

func (it fileReader) OpenFile() (*os.File, error) {
	file, err := os.Open(it.filePath)

	if err != nil {
		return nil, errcore.PathInvalidErrorType.Error(
			"cannot open file!"+err.Error(),
			errcore.MessageWithRef(
				"StackTraces",
				codestack.StacksString(codestack.Skip1)))
	}

	return file, nil
}

func (it fileReader) OpenFileLock() (*os.File, error) {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.OpenFile()
}

func (it fileReader) TrimText() (string, error) {
	line, err := it.String()

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(line), err
}

func (it fileReader) TrimTextLock() (string, error) {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.TrimText()
}

func (it fileReader) TrimLine() (string, error) {
	line, err := it.Line()

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(line), err
}

func (it fileReader) TrimLineLock() (string, error) {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.TrimLine()
}

func (it fileReader) FilePath() string {
	return it.filePath
}

func (it fileReader) Type() (Variant, error) {
	toString, err := it.String()

	return Variant(toString), err
}

func (it fileReader) TypeLock() (Variant, error) {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.Type()
}

func (it fileReader) Line() (string, error) {
	allBytes, err := it.Bytes()

	if err != nil {
		return "", err
	}

	return string(allBytes), err
}

func (it fileReader) LineLock() (string, error) {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.Line()
}

func (it fileReader) Text() (string, error) {
	allBytes, err := it.Bytes()

	if err != nil {
		return "", err
	}

	return string(allBytes), err
}

func (it fileReader) TextLock() (string, error) {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.Text()
}

func (it fileReader) String() (string, error) {
	allBytes, err := it.Bytes()

	if err != nil {
		return "", err
	}

	return string(allBytes), err
}

func (it fileReader) StringLock() (string, error) {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.String()
}

func (it fileReader) Strings() ([]string, error) {
	return it.Lines()
}

func (it fileReader) StringsLock() ([]string, error) {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.Lines()
}

func (it fileReader) SimpleSlice() (*corestr.SimpleSlice, error) {
	lines, err := it.Lines()

	if err != nil {
		return corestr.New.SimpleSlice.Empty(), err
	}

	return corestr.New.SimpleSlice.Create(lines), err
}

func (it fileReader) SimpleSliceLock() (*corestr.SimpleSlice, error) {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.SimpleSlice()
}

func (it fileReader) Lines() ([]string, error) {
	text, err := it.Text()

	if err != nil {
		return nil, err
	}

	return strings.Split(text, constants.DefaultLine), err
}

func (it fileReader) LinesLock() ([]string, error) {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.Lines()
}

func (it fileReader) NonEmptyLines() ([]string, error) {
	lines, err := it.Lines()

	if err != nil {
		return nil, err
	}

	return stringslice.NonEmptySlice(lines), err
}

func (it fileReader) NonEmptyLinesLock() ([]string, error) {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.NonEmptyLines()
}

func (it fileReader) TrimNonEmptyLines() ([]string, error) {
	text, err := it.Text()

	if err != nil {
		return nil, err
	}

	return stringslice.SplitTrimmedNonEmptyAll(
		text,
		constants.DefaultLine), err
}

func (it fileReader) TrimNonEmptyLinesLock() ([]string, error) {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.TrimNonEmptyLines()
}

func (it fileReader) TrimNonWhitespaceLines() ([]string, error) {
	text, err := it.Text()

	if err != nil {
		return nil, err
	}

	return stringslice.SplitTrimmedNonEmptyAll(
		text,
		constants.DefaultLine), err
}

func (it fileReader) TrimNonWhitespaceLinesLock() ([]string, error) {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.TrimNonWhitespaceLines()
}

func (it fileReader) Bytes() ([]byte, error) {
	allBytes, err := ioutil.ReadFile(it.filePath)

	if err != nil {
		return nil, errcore.
			InvalidAnyPathEmptyType.
			Error("cannot read the file", it.filePath)
	}

	return allBytes, err
}

func (it fileReader) BytesLock() ([]byte, error) {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.Bytes()
}

func (it fileReader) Raw() ([]byte, error) {
	allBytes, err := ioutil.ReadFile(it.filePath)

	if err != nil {
		return nil, errcore.
			InvalidAnyPathEmptyType.
			Error("cannot read the file", it.filePath)
	}

	return allBytes, err
}

func (it fileReader) RawLock() ([]byte, error) {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.Raw()
}

func (it fileReader) JsonResult() (*corejson.Result, error) {
	rawBytes, err := it.Raw()

	if err != nil {
		return nil, err
	}

	return corejson.
		Deserialize.
		ResultPtr(rawBytes)
}

func (it fileReader) JsonResultLock() (*corejson.Result, error) {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.JsonResult()
}

func (it fileReader) RawAsJsonResult() *corejson.Result {
	rawBytes, err := it.Raw()

	return corejson.NewResult.Ptr(
		rawBytes,
		err,
		"File : "+it.filePath)
}

func (it fileReader) RawAsJsonResultLock() *corejson.Result {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.RawAsJsonResult()
}
