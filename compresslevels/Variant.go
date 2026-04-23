package compresslevels

import (
	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coreinterface/enuminf"
)

type Variant int8

const (
	Default Variant = iota
	Best
	Fast
	NoCompression
	Invalid
)

func (it Variant) IsBest() bool {
	return it == Best
}

func (it Variant) IsFast() bool {
	return it == Fast
}

func (it Variant) IsDefault() bool {
	return it == Default
}

func (it Variant) IsNoCompression() bool {
	return it == NoCompression
}

func (it Variant) Flate() int8 {
	return flateRanges[it]
}

func (it Variant) Value() int8 {
	return int8(it)
}

func (it Variant) ValueUInt16() uint16 {
	return uint16(it)
}

func (it Variant) AllNameValues() []string {
	return BasicEnumImpl.AllNameValues()
}

func (it Variant) OnlySupportedErr(names ...string) error {
	return BasicEnumImpl.OnlySupportedErr(names...)
}

func (it Variant) OnlySupportedMsgErr(message string, names ...string) error {
	return BasicEnumImpl.OnlySupportedMsgErr(message, names...)
}

func (it Variant) IntegerEnumRanges() []int {
	return BasicEnumImpl.IntegerEnumRanges()
}

func (it Variant) MinMaxAny() (min, max interface{}) {
	return BasicEnumImpl.MinMaxAny()
}

func (it Variant) MinValueString() string {
	return BasicEnumImpl.MinValueString()
}

func (it Variant) MaxValueString() string {
	return BasicEnumImpl.MaxValueString()
}

func (it Variant) MaxInt() int {
	return BasicEnumImpl.MaxInt()
}

func (it Variant) MinInt() int {
	return BasicEnumImpl.MinInt()
}

func (it Variant) RangesDynamicMap() map[string]interface{} {
	return BasicEnumImpl.RangesDynamicMap()
}

func (it Variant) IsAnyNamesOf(names ...string) bool {
	return BasicEnumImpl.IsAnyNamesOf(it.Value(), names...)
}

func (it Variant) ValueInt() int {
	return int(it)
}

func (it Variant) IsAnyValuesEqual(anyByteValues ...int8) bool {
	return BasicEnumImpl.IsAnyOf(it.Value(), anyByteValues...)
}

func (it Variant) IsByteValueEqual(value byte) bool {
	return it.ValueByte() == value
}

func (it Variant) IsNameEqual(name string) bool {
	return it.Name() == name
}

func (it Variant) IsValueEqual(value byte) bool {
	return it.ValueByte() == value
}

func (it Variant) ValueInt8() int8 {
	return int8(it)
}

func (it Variant) ValueInt16() int16 {
	return int16(it)
}

func (it Variant) ValueInt32() int32 {
	return int32(it)
}

func (it Variant) ValueString() string {
	return it.ToNumberString()
}

func (it Variant) Format(format string) (compiled string) {
	return BasicEnumImpl.Format(format, it.ValueByte())
}

func (it Variant) EnumType() enuminf.EnumTyper {
	return BasicEnumImpl.EnumType()
}

func (it Variant) Name() string {
	return BasicEnumImpl.ToEnumString(it.ValueInt8())
}

func (it Variant) ToNumberString() string {
	return BasicEnumImpl.ToNumberString(it.ValueByte())
}

func (it Variant) MarshalJSON() ([]byte, error) {
	return BasicEnumImpl.ToEnumJsonBytes(it.ValueInt8())
}

func (it *Variant) UnmarshalJSON(data []byte) error {
	dataConv, err := it.UnmarshallEnumToValue(
		data)

	if err == nil {
		*it = Variant(dataConv)
	}

	return err
}

func (it Variant) RangeNamesCsv() string {
	return BasicEnumImpl.RangeNamesCsv()
}

func (it Variant) TypeName() string {
	return BasicEnumImpl.TypeName()
}

func (it Variant) IsEqual(level Variant) bool {
	return level == it
}

func (it Variant) IsAboveOrEqual(level Variant) bool {
	return level.ValueByte() >= it.ValueByte()
}

func (it Variant) IsLowerOrEqual(level Variant) bool {
	return level.ValueByte() <= it.ValueByte()
}

func (it Variant) IsInvalid() bool {
	return it == Invalid
}

func (it Variant) IsValid() bool {
	return it != Invalid
}

func (it Variant) IsAnyOf(anyOfItems ...Variant) bool {
	for _, item := range anyOfItems {
		if item == it {
			return true
		}
	}

	return false
}

func (it Variant) UnmarshallEnumToValue(
	jsonUnmarshallingValue []byte,
) (int8, error) {
	return BasicEnumImpl.UnmarshallToValue(
		true,
		jsonUnmarshallingValue)
}

func (it Variant) MaxByte() int8 {
	return BasicEnumImpl.Max()
}

func (it Variant) MinByte() int8 {
	return BasicEnumImpl.Min()
}

func (it Variant) ValueByte() byte {
	return byte(it)
}

func (it Variant) RangesInt8() []int8 {
	return BasicEnumImpl.Ranges()
}

func (it Variant) NameValue() string {
	return BasicEnumImpl.NameWithValue(it)
}

func (it Variant) String() string {
	return BasicEnumImpl.ToEnumString(it.Value())
}

func (it *Variant) JsonParseSelfInject(jsonResult *corejson.Result) error {
	err := jsonResult.Unmarshal(it)

	return err
}

func (it Variant) Json() corejson.Result {
	return corejson.New(it)
}

func (it Variant) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it Variant) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return &it
}

func (it Variant) AsJsoner() corejson.Jsoner {
	return it
}

func (it Variant) AsJsonMarshaller() corejson.JsonMarshaller {
	return &it
}

func (it Variant) UnmarshallEnumToValueInt8(
	jsonUnmarshallingValue []byte,
) (int8, error) {
	return it.UnmarshallEnumToValue(jsonUnmarshallingValue)
}

func (it Variant) MaxInt8() int8 {
	return BasicEnumImpl.Max()
}

func (it Variant) MinInt8() int8 {
	return BasicEnumImpl.Min()
}

func (it Variant) ToEnumString(input int8) string {
	return BasicEnumImpl.ToEnumString(input)
}

func (it Variant) ToInt8EnumString(input int8) string {
	return BasicEnumImpl.ToNumberString(input)
}

func (it Variant) IsInteger8ValueEqual(value int8) bool {
	return it.Value() == value
}

func (it Variant) AsBasicInt8EnumContractsBinder() enuminf.BasicInt8EnumContractsBinder {
	return &it
}

func (it Variant) AsBasicByteEnumContractsBinder() enuminf.BasicInt8EnumContractsBinder {
	return &it
}

func (it Variant) AsBasicEnumContractsBinder() enuminf.BasicEnumContractsBinder {
	return &it
}

func (it Variant) ToPtr() *Variant {
	return &it
}
