package sqliteconnpathtype

import (
	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coreinterface/enuminf"
)

type Variant string

const (
	Invalid                           Variant = "Invalid"
	AllSqlitePath                     Variant = "All"
	AllWithTypeSqlitePath             Variant = "AllWithType"
	AllWithTypeAndDynamicSqlitePath   Variant = "AllWithTypeAndDynamicSqlitePath"
	AllWithTypeAndSequenceSqlitePath  Variant = "AllWithTypeAndSequenceSqlitePath"
	PrefixSqlitePath                  Variant = "Prefix"
	PrefixTypeSqlitePath              Variant = "PrefixType"
	SpecificSqlitePath                Variant = "Specific"
	DynamicSpecificSqlitePath         Variant = "DynamicSpecific"
	SequenceSpecificSqlitePath        Variant = "SequenceSpecific"
	DynamicSequenceSpecificSqlitePath Variant = "DynamicSequenceSpecific"
)

func (it Variant) ValueUInt16() uint16 {
	return 0
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

func (it Variant) Name() string {
	return string(it)
}

func (it Variant) String() string {
	return string(it)
}

func (it Variant) PathFormat() string {
	return sqliteConnectionFormats[it]
}

func (it Variant) ValueByte() byte {
	return 0
}

func (it Variant) IsValid() bool {
	return it != Invalid
}

func (it Variant) IsInvalid() bool {
	return it == Invalid
}

func (it Variant) NameValue() string {
	return BasicEnumImpl.NameWithValue(it.String())
}

func (it Variant) IsNameEqual(name string) bool {
	return it.Name() == name
}

func (it Variant) IsAnyNamesOf(names ...string) bool {
	return BasicEnumImpl.IsAnyOf(it.Name(), names...)
}

func (it Variant) TypeName() string {
	return BasicEnumImpl.TypeName()
}

func (it Variant) ToNumberString() string {
	return "0"
}

func (it Variant) ValueInt() int {
	return 0
}

func (it Variant) ValueInt8() int8 {
	return 0
}

func (it Variant) ValueInt16() int16 {
	return 0
}

func (it Variant) ValueInt32() int32 {
	return 0
}

func (it Variant) ValueString() string {
	return it.String()
}

func (it Variant) RangeNamesCsv() string {
	return BasicEnumImpl.RangeNamesCsv()
}

func (it Variant) Format(format string) (compiled string) {
	return BasicEnumImpl.Format(format, it.ValueByte())
}

func (it Variant) EnumType() enuminf.EnumTyper {
	return BasicEnumImpl.EnumType()
}

func (it Variant) UnmarshallEnumToValue(
	jsonUnmarshallingValue []byte,
) (string, error) {
	return BasicEnumImpl.UnmarshallToValue(
		true,
		jsonUnmarshallingValue)
}

func (it Variant) MarshalJSON() ([]byte, error) {
	return BasicEnumImpl.ToEnumJsonBytes(it.String())
}

func (it *Variant) UnmarshalJSON(data []byte) error {
	dataConv, err := it.UnmarshallEnumToValue(data)

	if err == nil {
		*it = Variant(dataConv)
	}

	return err
}

func (it Variant) StringRangesPtr() *[]string {
	return BasicEnumImpl.StringRangesPtr()
}

func (it Variant) StringRanges() []string {
	return BasicEnumImpl.StringRanges()
}

func (it Variant) RangesInvalidMessage() string {
	return BasicEnumImpl.RangesInvalidMessage()
}

func (it Variant) RangesInvalidErr() error {
	return BasicEnumImpl.RangesInvalidErr()
}

func (it Variant) IsValidRange() bool {
	return BasicEnumImpl.IsValidRange(it.Name())
}

func (it Variant) IsInvalidRange() bool {
	return !BasicEnumImpl.IsValidRange(it.Name())
}

func (it Variant) Json() corejson.Result {
	return corejson.New(it)
}

func (it Variant) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it Variant) JsonParseSelfInject(jsonResult *corejson.Result) error {
	err := jsonResult.Unmarshal(it)

	return err
}

func (it Variant) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return &it
}

func (it Variant) AsBasicEnumContractsBinder() enuminf.BasicEnumContractsBinder {
	return &it
}

func (it Variant) AsStandardEnumerContractsBinder() enuminf.StandardEnumerContractsBinder {
	return &it
}

func (it Variant) ToPtr() *Variant {
	return &it
}
