package brackets

import (
	"fmt"
	"strings"

	"github.com/alimtvnetwork/core-v9/constants"
)

type Pair struct {
	Start    Bracket
	End      Bracket
	Category Category
}

func (it Pair) SelfWrap() string {
	return it.Category.SelfWrap()
}

func (it Pair) Wrap(str string) string {
	if str == "" {
		return it.Start.String() + it.End.String()
	}

	return it.Start.String() +
		str +
		it.End.String()
}

func (it *Pair) IsValid() bool {
	return it != nil && !it.Start.IsInvalid()
}

func (it *Pair) IsSafeInvalid() bool {
	return !it.IsValid()
}

func (it Pair) WrapAny(source interface{}) string {
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

func (it Pair) WrapString(sourceString string) string {
	isInvalid := it.IsSafeInvalid()

	if isInvalid {
		return sourceString
	}

	return it.Start.String() + sourceString + it.End.String()
}

// WrapFmtString
//
//	{wrapped} will be replaced in
//	the format by the wrapped string.
func (it Pair) WrapFmtString(
	format, sourceString string,
) string {
	wrapped := it.WrapString(sourceString)

	return strings.ReplaceAll(
		format,
		"{wrapped}",
		wrapped)
}

func (it Pair) IsWrapped(source string) bool {
	return HasBothWrappedWith(
		source,
		it.Category)
}

func (it Pair) UnWrap(source string) string {
	return UnWrapWith(
		source,
		it.Category)
}

func (it Pair) WrapWithOptions(
	isSkipOnExist bool,
	source string,
) string {
	return WrapWith(
		source,
		it.Category,
		isSkipOnExist)
}

func (it Pair) WrapSkipOnExist(
	source string,
) string {
	return WrapWith(
		source,
		it.Category,
		true)
}

// String returns a deterministic representation of the Pair without
// re-entering converters.AnyTo.ValueString — that path falls back to
// fmt.Sprintf("%v", it) which calls back into String() (infinite recursion
// → stack overflow). RCA pattern 9 (2026-05-07).
func (it Pair) String() string {
	return fmt.Sprintf(
		"{Start:%s End:%s Category:%s}",
		it.Start.String(),
		it.End.String(),
		it.Category.String())
}
