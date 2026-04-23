package revokereason

import (
	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coreinterface/enuminf"
)

// Variant
//
// RevokeCertification reasoning
//   - Unspecified          Variant = iota
//   - KeyCompromise        Variant = 1
//   - CaCompromise         Variant = 2
//   - AffiliationChanged   Variant = 3
//   - Superseded           Variant = 4
//   - CessationOfOperation Variant = 5
//   - CertificateHold      Variant = 6
//   - _Unused              Variant = 7
//   - RemoveFromCRL        Variant = 8
//   - PrivilegeWithdrawn   Variant = 9
//   - AaCompromise         Variant = 10
//
// References:
//   - Reason Code                      : https://tools.ietf.org/html/rfc5280#section-5.3.1
//   - Reasoning Numbers                : https://prnt.sc/26gwwsm
//   - PKIX Certificate and CRL Profile : https://prnt.sc/26gwxgi
type Variant byte

// Don't modify, order matters
// Constants used for certificate revocation, used for RevokeCertificate
// See https://tools.ietf.org/html/rfc5280#section-5.3.1
const (
	Unspecified          Variant = iota // 0 - needs to be as it is, cannot be changed
	KeyCompromise        Variant = 1    // 1
	CaCompromise         Variant = 2    // 2
	AffiliationChanged   Variant = 3    // 3
	Superseded           Variant = 4    // 4
	CessationOfOperation Variant = 5    // 5
	CertificateHold      Variant = 6    // 6
	_Unused              Variant = 7    // 7 - Unused
	RemoveFromCRL        Variant = 8    // 8
	PrivilegeWithdrawn   Variant = 9    // 9
	AaCompromise         Variant = 10   // 10
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

func (it Variant) IsUnspecified() bool {
	return it == Unspecified
}

func (it Variant) IsKeyCompromise() bool {
	return it == KeyCompromise
}

func (it Variant) IsCaCompromise() bool {
	return it == CaCompromise
}

func (it Variant) IsAffiliationChanged() bool {
	return it == AffiliationChanged
}

func (it Variant) IsSuperseded() bool {
	return it == Superseded
}

func (it Variant) IsCessationOfOperation() bool {
	return it == CessationOfOperation
}

func (it Variant) IsCertificateHold() bool {
	return it == CertificateHold
}

func (it Variant) IsRemoveFromCRL() bool {
	return it == RemoveFromCRL
}

func (it Variant) IsPrivilegeWithdrawn() bool {
	return it == PrivilegeWithdrawn
}

func (it Variant) IsAaCompromise() bool {
	return it == AaCompromise
}

func (it Variant) IsUnspecifiedLogically() bool {
	return it == Unspecified || it == _Unused
}

func (it Variant) Name() string {
	return BasicEnumImpl.ToEnumString(it.Value())
}

func (it Variant) IsInvalid() bool {
	return it.IsUnspecifiedLogically()
}

func (it Variant) IsValid() bool {
	return !it.IsInvalid()
}

func (it Variant) IsAnyOf(logTypes ...Variant) bool {
	for _, logType := range logTypes {
		if logType == it {
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
		return Unspecified
	}

	return *it
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
