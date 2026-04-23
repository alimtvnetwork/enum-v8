package brackets

import (
	"reflect"
	"strings"

	"github.com/alimtvnetwork/core-v8/coreinterface/enuminf"
)

type Category byte

const (
	UnknownCategory Category = iota
	Parenthesis
	Curly
	Square
)

func (it Category) ValueByte() byte {
	return byte(it)
}

func (it Category) TypeName() string {
	return reflect.TypeOf(UnknownCategory).String()
}

func (it Category) SelfWrap() string {
	return categorySelfWrapMap[it]
}

func (it Category) Pair() Pair {
	pair, _ := pairMaps[it]

	return pair
}

func (it Category) IsInvalid() bool {
	return it == UnknownCategory
}

func (it Category) IsValid() bool {
	return it != UnknownCategory
}

func (it Category) Start() Bracket {
	return startCategoryMap[it]
}

func (it Category) End() Bracket {
	return endCategoryMap[it]
}

func (it Category) IsParenthesis() bool {
	return it == Parenthesis
}

func (it Category) IsCurly() bool {
	return it == Curly
}

func (it Category) IsSquare() bool {
	return it == Square
}

func (it Category) String() string {
	return categoriesRange[it]
}

func (it Category) WrapAny(source interface{}) string {
	pair := it.Pair()

	return pair.WrapAny(source)
}

func (it Category) WrapString(sourceString string) string {
	pair := it.Pair()

	return pair.Wrap(sourceString)
}

// WrapFmtString
//
//	{wrapped} will be replaced in the
//	format by the wrapped string.
func (it Category) WrapFmtString(
	format, sourceString string,
) string {
	wrapped := it.WrapString(sourceString)

	return strings.ReplaceAll(
		format,
		"{wrapped}",
		wrapped)
}

func (it Category) IsWrapped(source string) bool {
	return HasBothWrappedWith(
		source,
		it)
}

func (it Category) UnWrap(source string) string {
	return UnWrapWith(
		source,
		it)
}

func (it Category) WrapWithOptions(
	isSkipOnExist bool,
	source string,
) string {
	return WrapWith(
		source,
		it,
		isSkipOnExist)
}

func (it Category) Name() string {
	return it.String()
}

func (it Category) AsSimpleEnumer() enuminf.SimpleEnumer {
	return &it
}
