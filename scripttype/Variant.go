package scripttype

import "github.com/alimtvnetwork/core-v8/coreinterface/enuminf"

type Variant byte

const (
	Default Variant = iota
	Shell
	Bash
	Perl
	Python
	Python2
	Python3
	CLang
	MakeScript
	Powershell
	Cmd
	Invalid
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

func (it Variant) OsDefaultScriptType() Variant {
	return osDefaultScriptType
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

func (it Variant) Name() string {
	return BasicEnumImpl.ToEnumString(it.ValueByte())
}

func (it Variant) NameValue() string {
	return BasicEnumImpl.NameWithValue(it)
}

func (it *Variant) ToNumberString() string {
	return BasicEnumImpl.ToNumberString(it.ValueByte())
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

func (it *Variant) MaxByte() byte {
	return BasicEnumImpl.Max()
}

func (it *Variant) MinByte() byte {
	return BasicEnumImpl.Min()
}

func (it Variant) ValueByte() byte {
	return byte(it)
}

func (it *Variant) RangesByte() []byte {
	return BasicEnumImpl.Ranges()
}

func (it Variant) IsUninitialized() bool {
	return it == Invalid
}

func (it Variant) IsDefault() bool {
	return it == Default
}

func (it Variant) IsDefaultOsScriptType() bool {
	return it == osDefaultScriptType
}

func (it Variant) IsShellOrBash() bool {
	return it == Shell || it == Bash
}

func (it Variant) IsCmdOrPowershell() bool {
	return it == Cmd || it == Powershell
}

func (it Variant) IsPowershellOrShell() bool {
	return it == Shell || it == Powershell
}

func (it Variant) IsPowershellOrShellOrBash() bool {
	return it == Shell || it == Powershell || it == Bash
}

func (it Variant) IsShell() bool {
	return it == Shell
}

func (it Variant) IsBash() bool {
	return it == Bash
}

func (it Variant) IsPerl() bool {
	return it == Perl
}

func (it Variant) IsPython() bool {
	return it == Python
}

func (it Variant) IsPython2() bool {
	return it == Python2
}

func (it Variant) IsPython3() bool {
	return it == Python3
}

func (it Variant) IsCLang() bool {
	return it == CLang
}

func (it Variant) IsMakeScript() bool {
	return it == MakeScript
}

func (it Variant) IsPowershell() bool {
	return it == Powershell
}

func (it Variant) IsCmd() bool {
	return it == Cmd
}

func (it Variant) IsAnyPython() bool {
	return it.IsPython() ||
		it.IsPython2() ||
		it.IsPython3()
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

func (it *Variant) RangesVariants() []Variant {
	return scriptTypeRanges[:]
}

func (it *Variant) ScriptDefault() *ScriptDefault {
	if it.IsDefault() {
		return DefaultOsScript()
	}

	return RangesMap[*it]
}

func (it *Variant) ProcessName() string {
	scriptDefault := it.ScriptDefault()

	if scriptDefault == nil {
		return ""
	}

	return scriptDefault.ProcessName
}

func (it *Variant) DefaultArgs() []string {
	scriptDefault := it.ScriptDefault()

	if scriptDefault == nil {
		return []string{}
	}

	return scriptDefault.DefaultArguments
}

func (it *Variant) IsImplemented() bool {
	scriptDefault := it.ScriptDefault()

	if scriptDefault == nil {
		return false
	}

	return scriptDefault.IsImplemented
}

func (it *Variant) IsInvalidImplement() bool {
	scriptDefault := it.ScriptDefault()

	if scriptDefault == nil {
		return true
	}

	return !scriptDefault.IsImplemented
}

func (it Variant) AsBasicByteEnumContractsBinder() enuminf.BasicByteEnumContractsBinder {
	return &it
}

func (it Variant) AsBasicEnumContractsBinder() enuminf.BasicEnumContractsBinder {
	return &it
}

func (it Variant) ToPtr() *Variant {
	return &it
}
