package brackets

import "github.com/alimtvnetwork/core-v8/coreinterface/enuminf"

type Bracket byte

const (
	Invalid          Bracket = iota
	ParenthesisStart Bracket = '('
	ParenthesisEnd   Bracket = ')'
	CurlyStart       Bracket = '{'
	CurlyEnd         Bracket = '}'
	SquareStart      Bracket = '['
	SquareEnd        Bracket = ']'
)

func (it Bracket) ValueUInt16() uint16 {
	return uint16(it)
}

func (it Bracket) AllNameValues() []string {
	return BasicEnumImpl.AllNameValues()
}

func (it Bracket) OnlySupportedErr(names ...string) error {
	return BasicEnumImpl.OnlySupportedErr(names...)
}

func (it Bracket) OnlySupportedMsgErr(message string, names ...string) error {
	return BasicEnumImpl.OnlySupportedMsgErr(message, names...)
}

func (it Bracket) IntegerEnumRanges() []int {
	return BasicEnumImpl.IntegerEnumRanges()
}

func (it Bracket) MinMaxAny() (min, max interface{}) {
	return BasicEnumImpl.MinMaxAny()
}

func (it Bracket) MinValueString() string {
	return BasicEnumImpl.MinValueString()
}

func (it Bracket) MaxValueString() string {
	return BasicEnumImpl.MaxValueString()
}

func (it Bracket) MaxInt() int {
	return BasicEnumImpl.MaxInt()
}

func (it Bracket) MinInt() int {
	return BasicEnumImpl.MinInt()
}

func (it Bracket) RangesDynamicMap() map[string]interface{} {
	return BasicEnumImpl.RangesDynamicMap()
}

func (it Bracket) SelfWrap() string {
	return bracketsSelfWrapMap[it]
}

func (it Bracket) IsStart() bool {
	return it == ParenthesisStart ||
		it == CurlyStart ||
		it == SquareStart
}

func (it Bracket) IsEnd() bool {
	return it == ParenthesisEnd ||
		it == CurlyEnd ||
		it == SquareEnd
}

func (it Bracket) IsParenthesis() bool {
	return it == ParenthesisStart || it == ParenthesisEnd
}

func (it Bracket) IsCurly() bool {
	return it == CurlyStart || it == CurlyEnd
}

func (it Bracket) IsSquare() bool {
	return it == SquareStart || it == SquareEnd
}

func (it Bracket) IsParenthesisStart() bool {
	return it == ParenthesisStart
}

func (it Bracket) IsParenthesisEnd() bool {
	return it == ParenthesisEnd
}

func (it Bracket) IsCurlyStart() bool {
	return it == CurlyStart
}

func (it Bracket) IsCurlyEnd() bool {
	return it == CurlyEnd
}

func (it Bracket) IsSquareStart() bool {
	return it == SquareStart
}

func (it Bracket) IsSquareEnd() bool {
	return it == SquareEnd
}

func (it Bracket) OtherBracket() Bracket {
	other, _ := otherBracketMaps[it]

	return other
}

func (it Bracket) BothBrackets() BothBrackets {
	bothBracket, has := bothBracketsMap[it]

	if has {
		return *bothBracket
	}

	return BothBrackets{
		IsInvalid: true,
	}
}

func (it Bracket) Category() Category {
	brackets := it.BothBrackets()

	return brackets.Category
}

func (it Bracket) Pair() Pair {
	brackets := it.BothBrackets()

	return brackets.Category.Pair()
}

func (it Bracket) WrapAny(source interface{}) string {
	brackets := it.BothBrackets()

	return brackets.WrapAny(source)
}

func (it Bracket) WrapString(
	sourceString string,
) string {
	brackets := it.BothBrackets()

	return brackets.WrapString(sourceString)
}

// WrapFmtString
//
//	{wrapped} will be replaced in the
//	format by the wrapped string.
func (it Bracket) WrapFmtString(
	format, sourceString string,
) string {
	brackets := it.BothBrackets()

	return brackets.WrapFmtString(format, sourceString)
}

func (it Bracket) IsWrapped(
	source string,
) bool {
	brackets := it.BothBrackets()

	return brackets.IsWrapped(source)
}

func (it Bracket) UnWrap(
	source string,
) string {
	brackets := it.BothBrackets()

	return brackets.UnWrap(source)
}

func (it Bracket) WrapWithOptions(
	isSkipOnExist bool,
	source string,
) string {
	brackets := it.BothBrackets()

	return brackets.WrapWithOptions(
		isSkipOnExist,
		source)
}

func (it Bracket) WrapSkipOnExist(
	source string,
) string {
	brackets := it.BothBrackets()

	return brackets.WrapWithOptions(
		true,
		source)
}

func (it Bracket) WrapRegardless(
	source string,
) string {
	brackets := it.BothBrackets()

	return brackets.WrapWithOptions(
		false,
		source)
}

func (it Bracket) IsEqual(char uint8) bool {
	return it.Value() == char
}

func (it Bracket) Value() byte {
	return byte(it)
}

func (it Bracket) IsAnyNamesOf(names ...string) bool {
	return BasicEnumImpl.IsAnyNamesOf(it.ValueByte(), names...)
}

func (it Bracket) ValueInt() int {
	return int(it)
}

func (it Bracket) IsAnyValuesEqual(anyByteValues ...byte) bool {
	return BasicEnumImpl.IsAnyOf(it.ValueByte(), anyByteValues...)
}

func (it Bracket) IsByteValueEqual(value byte) bool {
	return it.ValueByte() == value
}

func (it Bracket) IsNameEqual(name string) bool {
	return it.Name() == name
}

func (it Bracket) IsValueEqual(value byte) bool {
	return it.ValueByte() == value
}

func (it Bracket) ValueInt8() int8 {
	return int8(it)
}

func (it Bracket) ValueInt16() int16 {
	return int16(it)
}

func (it Bracket) ValueInt32() int32 {
	return int32(it)
}

func (it Bracket) ValueString() string {
	return it.ToNumberString()
}

func (it Bracket) Format(format string) (compiled string) {
	return BasicEnumImpl.Format(format, it.ValueByte())
}

func (it Bracket) EnumType() enuminf.EnumTyper {
	return BasicEnumImpl.EnumType()
}

func (it Bracket) Name() string {
	return BasicEnumImpl.ToEnumString(it.ValueByte())
}

func (it Bracket) NameValue() string {
	return BasicEnumImpl.NameWithValue(it)
}

func (it *Bracket) ToNumberString() string {
	return BasicEnumImpl.ToNumberString(it.ValueByte())
}

func (it Bracket) IsInvalid() bool {
	return it == Invalid
}

func (it Bracket) IsValid() bool {
	return it != Invalid
}

func (it Bracket) IsAnyOf(anyOfItems ...Bracket) bool {
	for _, item := range anyOfItems {
		if item == it {
			return true
		}
	}

	return false
}

func (it Bracket) IsNameOf(anyNames ...string) bool {
	for _, name := range anyNames {
		if name == it.Name() {
			return true
		}
	}

	return false
}

func (it *Bracket) UnmarshallEnumToValue(
	jsonUnmarshallingValue []byte,
) (byte, error) {
	return BasicEnumImpl.UnmarshallToValue(
		isMappedToDefault,
		jsonUnmarshallingValue)
}

func (it Bracket) String() string {
	return BasicEnumImpl.ToEnumString(it.ValueByte())
}

func (it Bracket) RangeNamesCsv() string {
	return BasicEnumImpl.RangeNamesCsv()
}

func (it Bracket) TypeName() string {
	return BasicEnumImpl.TypeName()
}

func (it Bracket) MarshalJSON() ([]byte, error) {
	return BasicEnumImpl.ToEnumJsonBytes(it.ValueByte())
}

func (it *Bracket) UnmarshalJSON(data []byte) error {
	rawScriptType, err := it.UnmarshallEnumToValue(
		data)

	if err == nil {
		*it = Bracket(rawScriptType)
	}

	return err
}

func (it *Bracket) AsBasicEnumContractsBinder() enuminf.BasicEnumContractsBinder {
	return it
}

func (it Bracket) MaxByte() byte {
	return BasicEnumImpl.Max()
}

func (it Bracket) MinByte() byte {
	return BasicEnumImpl.Min()
}

func (it Bracket) ValueByte() byte {
	return byte(it)
}

func (it Bracket) RangesByte() []byte {
	return BasicEnumImpl.Ranges()
}

func (it Bracket) AsBasicByteEnumContractsBinder() enuminf.BasicByteEnumContractsBinder {
	return &it
}

func (it Bracket) ToPtr() *Bracket {
	return &it
}
