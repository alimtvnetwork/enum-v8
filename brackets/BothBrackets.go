package brackets

import (
	"fmt"
	"strings"

	"github.com/alimtvnetwork/core-v8/constants"
	"github.com/alimtvnetwork/core-v8/converters"
)

type BothBrackets struct {
	Start, End Bracket
	Category   Category
	IsInvalid  bool
}

func (it *BothBrackets) IsValid() bool {
	return it != nil && !it.IsInvalid
}

func (it *BothBrackets) IsSafeInvalid() bool {
	return !it.IsValid()
}

func (it BothBrackets) WrapAny(source interface{}) string {
	isInvalid := it.IsSafeInvalid()

	if isInvalid && source == nil {
		return ""
	}

	toString := fmt.Sprintf(
		constants.SprintValueFormat,
		source)

	if isInvalid {
		return toString
	}

	return it.Start.String() + toString + it.End.String()
}

func (it BothBrackets) WrapString(sourceString string) string {
	isInvalid := it.IsSafeInvalid()

	if isInvalid {
		return sourceString
	}

	return it.Start.String() + sourceString + it.End.String()
}

// WrapFmtString
//
//	{wrapped} will be replaced in the
//	format by the wrapped string.
func (it BothBrackets) WrapFmtString(
	format, sourceString string,
) string {
	wrapped := it.WrapString(sourceString)

	return strings.ReplaceAll(
		format,
		"{wrapped}",
		wrapped)
}

func (it BothBrackets) IsWrapped(source string) bool {
	return HasBothWrappedWith(
		source,
		it.Category)
}

func (it BothBrackets) UnWrap(source string) string {
	return UnWrapWith(
		source,
		it.Category)
}

func (it BothBrackets) WrapWithOptions(
	isSkipOnExist bool,
	source string,
) string {
	return WrapWith(
		source,
		it.Category,
		isSkipOnExist)
}

func (it BothBrackets) WrapSkipOnExist(
	source string,
) string {
	return WrapWith(
		source,
		it.Category,
		true)
}

func (it BothBrackets) String() string {
	return converters.AnyToValueString(it)
}
