package onofftype

import (
	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coreinterface"
	"github.com/alimtvnetwork/core-v8/coreinterface/enuminf"
	"github.com/alimtvnetwork/core-v8/issetter"
)

// Variant
//
// Invalid, Ask, On, Off
type Variant byte

const (
	Invalid Variant = iota
	Ask
	On
	Off
)

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

func (it Variant) IsLater() bool {
	return undefinedItems[it]
}

func (it Variant) OnOffLowercaseName() string {
	return onOffNamesLowerMap[it]
}

func (it Variant) IsIndeterminate() bool {
	return undefinedItems[it]
}

func (it Variant) IsAccept() bool {
	return it == On
}

func (it Variant) IsReject() bool {
	return it == Off
}

func (it Variant) IsSkip() bool {
	return undefinedItems[it]
}

func (it Variant) IsAnyNamesOf(names ...string) bool {
	return BasicEnumImpl.IsAnyNamesOf(it.ValueByte(), names...)
}

func (it Variant) ValueInt() int {
	return int(it)
}

func (it Variant) IsAnyValuesEqual(anyByteValues ...byte) bool {
	return BasicEnumImpl.IsAnyOf(it.ValueByte(), anyByteValues...)
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

func (it Variant) IsAcceptReject() bool {
	return it == On || it == Off
}

func (it Variant) IsNotAcceptReject() bool {
	return !it.IsAcceptReject()
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

func (it Variant) IsUninitialized() bool {
	return it == Invalid
}

func (it Variant) IsInitialized() bool {
	return it.IsDefinedLogically()
}

func (it Variant) IsAsk() bool {
	return it == Ask
}

func (it Variant) IsYes() bool {
	return it == On
}

func (it Variant) IsOn() bool {
	return it == On
}

func (it Variant) IsNo() bool {
	return it == Off
}

func (it Variant) IsOff() bool {
	return it == Off
}

func (it Variant) IsOffLogically() bool {
	return it != On
}

func (it Variant) IsOnLogically() bool {
	return it != Off
}

func (it Variant) IsYesNo() bool {
	return it == On || it == Off
}

func (it Variant) IsTrue() bool {
	return it == On
}

func (it Variant) IsAccepted() bool {
	return it == On
}

func (it Variant) IsRejected() bool {
	return it == Off
}

func (it Variant) IsDefinedAccepted() bool {
	return it.IsDefinedLogically() && it == On
}

func (it Variant) IsDefinedRejected() bool {
	return it.IsDefinedLogically() && it == Off
}

// IsDefinedLogically
//
// Not Ask, Invalid
func (it Variant) IsDefinedLogically() bool {
	return !undefinedMap[it]
}

// IsUndefinedLogically
//
// Either Ask, Invalid
func (it Variant) IsUndefinedLogically() bool {
	return undefinedMap[it]
}

// IsUninitializedOrAsk
//
// Not Ask, Invalid
func (it Variant) IsUninitializedOrAsk() bool {
	return undefinedMap[it]
}

func (it Variant) IsInvalid() bool {
	return undefinedMap[it]
}

func (it Variant) IsValid() bool {
	return !undefinedMap[it]
}

// Name
//
// On, Off ... returns from Ranges
func (it Variant) Name() string {
	return BasicEnumImpl.ToEnumString(it.Value())
}

func (it Variant) NameLower() string {
	return lowerCaseNames[it]
}

func (it Variant) YesNoLower() string {
	return lowerCaseNames[it]
}

// OnOffName
//
// Returns from map:

func (it Variant) OnOffName() string {
	return onOffNames[it]
}

func (it Variant) OnOffNameLower() string {
	return onOffNamesLowerMap[it]
}

func (it Variant) TrueFalseName() string {
	return trueFalseNames[it]
}

func (it Variant) ToIsSetter() issetter.Value {
	return variantToIsSetterBooleanMap[it]
}

func (it Variant) ToNumberString() string {
	return BasicEnumImpl.ToNumberString(it.Value())
}

func (it Variant) UnmarshallEnumToValue(
	jsonUnmarshallingValue []byte,
) (byte, error) {
	return BasicEnumImpl.UnmarshallToValue(
		true,
		jsonUnmarshallingValue)
}

func (it Variant) MaxByte() byte {
	return BasicEnumImpl.Max()
}

func (it Variant) MinByte() byte {
	return BasicEnumImpl.Min()
}

func (it Variant) ValueByte() byte {
	return it.Value()
}

func (it Variant) RangesByte() []byte {
	return BasicEnumImpl.Ranges()
}

func (it Variant) Value() byte {
	return byte(it)
}

func (it *Variant) UnmarshalJSON(data []byte) error {
	dataConv, err := it.UnmarshallEnumToValue(data)

	if err == nil {
		*it = Variant(dataConv)
	}

	return err
}

func (it Variant) ToPtr() *Variant {
	return &it
}

func (it *Variant) ToSimple() Variant {
	if it == nil {
		return Invalid
	}

	return *it
}

func (it Variant) IsAnyOf(anyOfItems ...Variant) bool {
	for _, item := range anyOfItems {
		if item == it {
			return true
		}
	}

	return false
}

func (it Variant) IsNameOf(anyNames ...string) bool {
	for _, name := range anyNames {
		if name == it.Name() {
			return true
		}
	}

	return false
}

func (it Variant) MarshalJSON() ([]byte, error) {
	return BasicEnumImpl.ToEnumJsonBytes(it.Value())
}

func (it Variant) RangeNamesCsv() string {
	return BasicEnumImpl.RangeNamesCsv()
}

func (it Variant) TypeName() string {
	return BasicEnumImpl.TypeName()
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

func (it *Variant) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *Variant) AsJsoner() corejson.Jsoner {
	return it
}

func (it *Variant) AsJsonMarshaller() corejson.JsonMarshaller {
	return it
}

func (it *Variant) AsBasicByteEnumContractsBinder() enuminf.BasicByteEnumContractsBinder {
	return it
}

func (it Variant) AsBasicEnumContractsBinder() enuminf.BasicEnumContractsBinder {
	return &it
}

func (it Variant) AsYesNoAcceptRejecter() coreinterface.YesNoAcceptRejecter {
	return &it
}

func (it Variant) AsOnOffLater() enuminf.OnOffLater {
	return &it
}
