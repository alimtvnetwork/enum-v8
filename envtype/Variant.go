package envtype

import "github.com/alimtvnetwork/core-v8/coreinterface/enuminf"

type Variant byte

const (
	Uninitialized Variant = iota
	Development
	Development1
	Development2
	Test
	Test1
	Test2
	Production
	Production1
	Production2
)

func (it Variant) IsAnyTestEnv() bool {
	return testEnvMap[it]
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

func (it Variant) Value() byte {
	return byte(it)
}

func (it Variant) KeyName() string {
	return keyNameMap[it]
}

func (it Variant) CurlyKeyName() string {
	return curlyKeyNameMap[it]
}

func (it Variant) VersionNumber() int {
	return envVersionNumber[it]
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

func (it *Variant) Name() string {
	return BasicEnumImpl.ToEnumString(it.ValueByte())
}

func (it Variant) NameValue() string {
	return BasicEnumImpl.NameWithValue(it)
}

func (it *Variant) ToNumberString() string {
	return BasicEnumImpl.ToNumberString(it.ValueByte())
}

func (it Variant) IsInvalid() bool {
	return it == Uninitialized
}

func (it Variant) IsValid() bool {
	return it != Uninitialized
}

func (it Variant) IsUninitialized() bool {
	return it == Uninitialized
}

func (it Variant) IsInitialized() bool {
	return it != Uninitialized
}

func (it Variant) IsDevelopment() bool {
	return it == Development
}

func (it Variant) IsDevelopment1() bool {
	return it == Development1
}

func (it Variant) IsDevelopment2() bool {
	return it == Development2
}

func (it Variant) IsAnyDevelopment() bool {
	return it == Development ||
		it == Development1 ||
		it == Development2
}

func (it Variant) IsTest() bool {
	return it == Test
}

func (it Variant) IsTest1() bool {
	return it == Test1
}

func (it Variant) IsTest2() bool {
	return it == Test2
}

func (it Variant) IsTestEnvLogically() bool {
	return testEnvMap[it]
}

func (it Variant) IsNotTestEnvLogically() bool {
	return !testEnvMap[it]
}

func (it Variant) IsDevEnvLogically() bool {
	return devEnvMap[it]
}

func (it Variant) IsNotDevEnvLogically() bool {
	return !devEnvMap[it]
}

func (it Variant) IsProdEnvLogically() bool {
	return productionEnvMap[it]
}

func (it Variant) IsNotProdEnvLogically() bool {
	return !productionEnvMap[it]
}

func (it Variant) IsProduction() bool {
	return it == Production
}

func (it Variant) IsProduction1() bool {
	return it == Production1
}

func (it Variant) IsProduction2() bool {
	return it == Production2
}

func (it Variant) IsAnyProduction() bool {
	return it == Production ||
		it == Production1 ||
		it == Production2
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

func (it *Variant) UnmarshallEnumToValue(
	jsonUnmarshallingValue []byte,
) (byte, error) {
	return BasicEnumImpl.UnmarshallToValue(
		isMappedToDefault,
		jsonUnmarshallingValue)
}

func (it Variant) String() string {
	return BasicEnumImpl.ToEnumString(it.ValueByte())
}

func (it Variant) RangeNamesCsv() string {
	return BasicEnumImpl.RangeNamesCsv()
}

func (it Variant) TypeName() string {
	return BasicEnumImpl.TypeName()
}

func (it *Variant) MarshalJSON() ([]byte, error) {
	return BasicEnumImpl.ToEnumJsonBytes(it.ValueByte())
}

func (it *Variant) UnmarshalJSON(data []byte) error {
	rawScriptType, err := it.UnmarshallEnumToValue(
		data)

	if err == nil {
		*it = Variant(rawScriptType)
	}

	return err
}

func (it *Variant) AsBasicEnumContractsBinder() enuminf.BasicEnumContractsBinder {
	return it
}

func (it Variant) MaxByte() byte {
	return BasicEnumImpl.Max()
}

func (it Variant) MinByte() byte {
	return BasicEnumImpl.Min()
}

func (it Variant) ValueByte() byte {
	return byte(it)
}

func (it Variant) RangesByte() []byte {
	return BasicEnumImpl.Ranges()
}

func (it Variant) AsEnvironmentTyper() enuminf.EnvironmentTyper {
	return &it
}

func (it Variant) AsBasicByteEnumContractsBinder() enuminf.BasicByteEnumContractsBinder {
	return &it
}
