package servicestate

import (
	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coreinterface/enuminf"
)

// Action
//
// Status, Start, Restart, Reload, Disable, Enable, Stop
type Action byte

const (
	Invalid Action = iota
	Status
	Start
	Restart
	Reload
	Enable
	Disable
	Stop
	StopEnableStart
)

func (it Action) IsStopEnableStart() bool {
	return it == StopEnableStart
}

func (it Action) IsStopDisable() bool {
	return it == Invalid
}

func (it Action) IsUndefined() bool {
	return it == Invalid
}

func (it Action) ValueUInt16() uint16 {
	return uint16(it)
}

func (it Action) AllNameValues() []string {
	return BasicEnumImpl.AllNameValues()
}

func (it Action) OnlySupportedErr(names ...string) error {
	return BasicEnumImpl.OnlySupportedErr(names...)
}

func (it Action) OnlySupportedMsgErr(message string, names ...string) error {
	return BasicEnumImpl.OnlySupportedMsgErr(message, names...)
}

func (it Action) IntegerEnumRanges() []int {
	return BasicEnumImpl.IntegerEnumRanges()
}

func (it Action) MinMaxAny() (min, max interface{}) {
	return BasicEnumImpl.MinMaxAny()
}

func (it Action) MinValueString() string {
	return BasicEnumImpl.MinValueString()
}

func (it Action) MaxValueString() string {
	return BasicEnumImpl.MaxValueString()
}

func (it Action) MaxInt() int {
	return BasicEnumImpl.MaxInt()
}

func (it Action) MinInt() int {
	return BasicEnumImpl.MinInt()
}

func (it Action) RangesDynamicMap() map[string]interface{} {
	return BasicEnumImpl.RangesDynamicMap()
}

func (it Action) IsStopSleepStart() bool {
	return false
}

func (it Action) IsSuspend() bool {
	return it == Stop
}

func (it Action) IsPause() bool {
	return it == Stop
}

func (it Action) IsResumed() bool {
	return it == Start
}

func (it Action) IsAnyAction() bool {
	return true
}

func (it Action) IsNotAnyAction() bool {
	return false
}

func (it Action) IsAnyNamesOf(names ...string) bool {
	return BasicEnumImpl.IsAnyNamesOf(it.ValueByte(), names...)
}

func (it Action) ValueInt() int {
	return int(it)
}

func (it Action) IsAnyValuesEqual(anyByteValues ...byte) bool {
	return BasicEnumImpl.IsAnyOf(it.ValueByte(), anyByteValues...)
}

func (it Action) IsByteValueEqual(value byte) bool {
	return it.ValueByte() == value
}

func (it Action) IsNameEqual(name string) bool {
	return it.Name() == name
}

func (it Action) IsValueEqual(value byte) bool {
	return it.ValueByte() == value
}

func (it Action) ValueInt8() int8 {
	return int8(it)
}

func (it Action) ValueInt16() int16 {
	return int16(it)
}

func (it Action) ValueInt32() int32 {
	return int32(it)
}

func (it Action) ValueString() string {
	return it.ToNumberString()
}

func (it Action) Format(format string) (compiled string) {
	return BasicEnumImpl.Format(format, it.ValueByte())
}

func (it Action) EnumType() enuminf.EnumTyper {
	return BasicEnumImpl.EnumType()
}

func (it Action) IsStart() bool {
	return it == Start
}

func (it Action) IsRestart() bool {
	return it == Restart
}

func (it Action) IsReload() bool {
	return it == Reload
}

func (it Action) IsEnable() bool {
	return it == Enable
}

func (it Action) IsDisable() bool {
	return it == Disable
}

func (it Action) IsStop() bool {
	return it == Stop
}

func (it Action) IsStatus() bool {
	return it == Status
}

func (it Action) Name() string {
	return BasicEnumImpl.ToEnumString(it.Value())
}

func (it Action) Names() string {
	return BasicEnumImpl.RangeNamesCsv()
}

func (it Action) IsInvalid() bool {
	return it == Invalid
}

func (it Action) IsValid() bool {
	return it != Invalid
}

func (it Action) IsAnyOf(anyOfItems ...Action) bool {
	for _, item := range anyOfItems {
		if item == it {
			return true
		}
	}

	return false
}

func (it Action) ToNumberString() string {
	return BasicEnumImpl.ToNumberString(it.Value())
}

func (it Action) UnmarshallEnumToValue(
	jsonUnmarshallingValue []byte,
) (byte, error) {
	return BasicEnumImpl.UnmarshallToValue(
		true,
		jsonUnmarshallingValue)
}

func (it Action) MaxByte() byte {
	return BasicEnumImpl.Max()
}

func (it Action) MinByte() byte {
	return BasicEnumImpl.Min()
}

func (it Action) ValueByte() byte {
	return it.Value()
}

func (it Action) RangesByte() []byte {
	return BasicEnumImpl.Ranges()
}

func (it Action) Value() byte {
	return byte(it)
}

func (it Action) NameValue() string {
	return BasicEnumImpl.NameWithValue(it)
}

func (it Action) NameCapital() string {
	return capitalNameMap[it]
}

func (it Action) CommandName() string {
	return Ranges[it]
}

func (it Action) String() string {
	return BasicEnumImpl.ToEnumString(it.Value())
}

func (it *Action) UnmarshalJSON(data []byte) error {
	dataConv, err := it.UnmarshallEnumToValue(data)

	if err == nil {
		*it = Action(dataConv)
	}

	return err
}

func (it Action) ToPtr() *Action {
	return &it
}

func (it Action) RangeNamesCsv() string {
	return BasicEnumImpl.RangeNamesCsv()
}

func (it Action) TypeName() string {
	return BasicEnumImpl.TypeName()
}

func (it Action) MarshalJSON() ([]byte, error) {
	return BasicEnumImpl.ToEnumJsonBytes(it.Value())
}

func (it Action) AsBasicEnumContractsBinder() enuminf.BasicEnumContractsBinder {
	return &it
}

func (it *Action) AsJsonMarshaller() corejson.JsonMarshaller {
	return it
}

func (it Action) AsActionTyper() enuminf.ActionTyper {
	return &it
}

func (it Action) AsBasicByteEnumContractsBinder() enuminf.BasicByteEnumContractsBinder {
	return &it
}
