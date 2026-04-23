package osgroupexecution

import (
	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coreinterface/enuminf"
)

type Precedence byte

const (
	Invalid Precedence = iota
	Create
	Delete
	Update
	ManageByUsers
	AddGroupsToSudoers
	GroupManage
)

func (it Precedence) ValueUInt16() uint16 {
	return uint16(it)
}

func (it Precedence) AllNameValues() []string {
	return BasicEnumImpl.AllNameValues()
}

func (it Precedence) OnlySupportedErr(names ...string) error {
	return BasicEnumImpl.OnlySupportedErr(names...)
}

func (it Precedence) OnlySupportedMsgErr(message string, names ...string) error {
	return BasicEnumImpl.OnlySupportedMsgErr(message, names...)
}

func (it Precedence) IntegerEnumRanges() []int {
	return BasicEnumImpl.IntegerEnumRanges()
}

func (it Precedence) MinMaxAny() (min, max interface{}) {
	return BasicEnumImpl.MinMaxAny()
}

func (it Precedence) MinValueString() string {
	return BasicEnumImpl.MinValueString()
}

func (it Precedence) MaxValueString() string {
	return BasicEnumImpl.MaxValueString()
}

func (it Precedence) MaxInt() int {
	return BasicEnumImpl.MaxInt()
}

func (it Precedence) MinInt() int {
	return BasicEnumImpl.MinInt()
}

func (it Precedence) RangesDynamicMap() map[string]interface{} {
	return BasicEnumImpl.RangesDynamicMap()
}

func (it Precedence) IsAnyNamesOf(names ...string) bool {
	return BasicEnumImpl.IsAnyNamesOf(it.ValueByte(), names...)
}

func (it Precedence) ValueInt() int {
	return int(it)
}

func (it Precedence) IsAnyValuesEqual(anyByteValues ...byte) bool {
	return BasicEnumImpl.IsAnyOf(it.ValueByte(), anyByteValues...)
}

func (it Precedence) IsByteValueEqual(value byte) bool {
	return it.ValueByte() == value
}

func (it Precedence) IsNameEqual(name string) bool {
	return it.Name() == name
}

func (it Precedence) IsValueEqual(value byte) bool {
	return it.ValueByte() == value
}

func (it Precedence) ValueInt8() int8 {
	return int8(it)
}

func (it Precedence) ValueInt16() int16 {
	return int16(it)
}

func (it Precedence) ValueInt32() int32 {
	return int32(it)
}

func (it Precedence) ValueString() string {
	return it.ToNumberString()
}

func (it Precedence) Format(format string) (compiled string) {
	return BasicEnumImpl.Format(format, it.ValueByte())
}

func (it Precedence) EnumType() enuminf.EnumTyper {
	return BasicEnumImpl.EnumType()
}

func (it Precedence) IsInvalid() bool {
	return it == Invalid
}

func (it Precedence) IsValid() bool {
	return it != Invalid
}

func (it Precedence) IsCreate() bool {
	return it == Create
}

func (it Precedence) IsDelete() bool {
	return it == Delete
}

func (it Precedence) IsUpdate() bool {
	return it == Update
}

func (it Precedence) IsAnyOf(anyOfItems ...Precedence) bool {
	for _, item := range anyOfItems {
		if item == it {
			return true
		}
	}

	return false
}

func (it Precedence) IsManageByUsers() bool {
	return it == ManageByUsers
}

func (it Precedence) Name() string {
	return BasicEnumImpl.ToEnumString(it.ValueByte())
}

func (it Precedence) ToNumberString() string {
	return BasicEnumImpl.ToNumberString(it.ValueByte())
}

func (it Precedence) UnmarshallEnumToValue(jsonUnmarshallingValue []byte) (byte, error) {
	return BasicEnumImpl.UnmarshallToValue(
		true,
		jsonUnmarshallingValue)
}

func (it Precedence) MaxByte() byte {
	return BasicEnumImpl.Max()
}

func (it Precedence) MinByte() byte {
	return BasicEnumImpl.Min()
}

func (it Precedence) ValueByte() byte {
	return it.Value()
}

func (it Precedence) Value() byte {
	return byte(it)
}

func (it Precedence) RangesByte() []byte {
	return BasicEnumImpl.Ranges()
}

func (it Precedence) MarshalJSON() ([]byte, error) {
	return BasicEnumImpl.ToEnumJsonBytes(it.ValueByte())
}

func (it *Precedence) UnmarshalJSON(data []byte) error {
	dataConv, err := it.UnmarshallEnumToValue(data)

	if err == nil {
		*it = Precedence(dataConv)
	}

	return err
}

func (it Precedence) RangeNamesCsv() string {
	return BasicEnumImpl.RangeNamesCsv()
}

func (it Precedence) TypeName() string {
	return BasicEnumImpl.TypeName()
}

func (it Precedence) NameValue() string {
	return BasicEnumImpl.NameWithValue(it)
}

func (it Precedence) String() string {
	return BasicEnumImpl.ToEnumString(it.Value())
}

func (it *Precedence) JsonParseSelfInject(jsonResult *corejson.Result) error {
	err := jsonResult.Unmarshal(it)

	return err
}

func (it Precedence) Json() corejson.Result {
	return corejson.New(it)
}

func (it Precedence) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it Precedence) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return &it
}

func (it Precedence) AsJsoner() corejson.Jsoner {
	return it
}

func (it Precedence) AsJsonMarshaller() corejson.JsonMarshaller {
	return &it
}

func (it Precedence) AsBasicByteEnumContractsBinder() enuminf.BasicByteEnumContractsBinder {
	return &it
}

func (it Precedence) AsBasicEnumContractsBinder() enuminf.BasicEnumContractsBinder {
	return &it
}

func (it Precedence) ToPtr() *Precedence {
	return &it
}
