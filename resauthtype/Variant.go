package resauthtype

import (
	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coreinterface"
	"github.com/alimtvnetwork/core-v8/coreinterface/enuminf"
)

// Variant
//
//	authorization response type
//
// Types:
//   - Invalid
//   - AllAccess
//   - Error
//   - Warning
//   - Restricted
//   - UnAuthorized
//   - PermissionIssue
//   - Forbidden
//   - ReadAccess
//   - WriteAccess
//   - CreateAccess
//   - EditAccess
type Variant byte

const (
	Invalid Variant = iota
	AllAccess
	Error
	Warning
	Restricted
	UnAuthorized
	PermissionIssue
	Forbidden
	ReadAccess
	WriteAccess
	CreateAccess
	EditAccess
	AccessGranted
	AccessRejected
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

func (it Variant) IsUninitialized() bool {
	return it == Invalid
}

func (it Variant) IsAccess() bool {
	return it == AllAccess
}

func (it Variant) IsError() bool {
	return it == Error
}

func (it Variant) IsWarning() bool {
	return it == Warning
}

func (it Variant) IsRestricted() bool {
	return it == Restricted
}

func (it Variant) IsUnAuthorized() bool {
	return it == UnAuthorized
}

func (it Variant) IsPermissionIssue() bool {
	return it == PermissionIssue
}

func (it Variant) IsForbidden() bool {
	return it == Forbidden
}

func (it Variant) IsReadAccess() bool {
	return it == ReadAccess
}

func (it Variant) IsWriteAccess() bool {
	return it == WriteAccess
}

func (it Variant) IsCreateAccess() bool {
	return it == CreateAccess
}

func (it Variant) IsEditAccess() bool {
	return it == EditAccess
}

func (it Variant) IsReadLogically() bool {
	return it.IsAccess() || it.IsReadAccess()
}

func (it Variant) IsWriteLogically() bool {
	return it.IsAccess() || it.IsWriteAccess() || it.IsCreateAccess() || it.IsEditAccess()
}

func (it Variant) IsCreateLogically() bool {
	return it.IsAccess() || it.IsCreateAccess()
}

func (it Variant) IsEditLogically() bool {
	return it.IsAccess() || it.IsWriteAccess() || it.IsEditAccess()
}

func (it Variant) IsAnyErrorLogically() bool {
	return errorMap[it]
}

func (it Variant) IsReadOrUpdateLogically() bool {
	return it.IsReadLogically() || it.IsEditLogically()
}

func (it Variant) IsUnAuthorizedLogically() bool {
	return errorMap[it]
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

func (it Variant) IsNameOf(anyNames ...string) bool {
	for _, name := range anyNames {
		if name == it.Name() {
			return true
		}
	}

	return false
}

func (it Variant) Name() string {
	return BasicEnumImpl.ToEnumString(it.Value())
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

func (it Variant) IsYes() bool {
	return allAccessMap[it]
}

func (it Variant) IsNo() bool {
	return errorMap[it]
}

func (it Variant) IsAsk() bool {
	return it == Invalid || it == Warning
}

func (it Variant) IsIndeterminate() bool {
	return it == Invalid || it == Warning
}

func (it Variant) IsAccept() bool {
	return it.IsYes()
}

func (it Variant) IsReject() bool {
	return it.IsNo()
}

func (it Variant) IsSkip() bool {
	return it.IsAsk()
}

func (it Variant) IsSuccess() bool {
	return it.IsYes()
}

func (it Variant) IsFailed() bool {
	return it.IsNo()
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

func (it Variant) AsIsSuccessValidator() coreinterface.IsSuccessValidator {
	return &it
}

func (it Variant) AsYesNoAcceptRejecter() coreinterface.YesNoAcceptRejecter {
	return &it
}

func (it Variant) AsJsoner() corejson.Jsoner {
	return &it
}

func (it Variant) AsJsonMarshaller() corejson.JsonMarshaller {
	return &it
}

func (it Variant) AsBasicByteEnumContractsBinder() enuminf.BasicByteEnumContractsBinder {
	return &it
}

func (it Variant) AsBasicEnumContractsBinder() enuminf.BasicEnumContractsBinder {
	return &it
}
