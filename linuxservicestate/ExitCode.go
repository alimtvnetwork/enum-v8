package linuxservicestate

import (
	"github.com/alimtvnetwork/core-v8/constants"
	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coreinterface/enuminf"
)

// ExitCode
//
//	ActiveRunning, DeadButPidExists,
//	DeadButVarLockFileExists, NotRunning,
//	UnknownService
//
//	What code means?
//	- Invalid (0):
//	    Created by us, actually its value 0
//	    but here we are considered for parsing issue
//	- ActiveRunning(1):
//	    We will map this to 0,
//	    consider as program is running and service is okay.
//	- DeadButPidExists(2):
//	    We will map this to 1,
//	    consider as program is dead and /var/run pid exists.
//	- DeadButVarLockFileExists(3):
//	    We will map this to 2,
//	    consider as program is dead and /var/lock-lock file exists.
//	- NotRunning(4):
//	    We will map this to 3,
//	    consider as program is not running but service exists in the system.
//	- UnknownService(5) / InvalidService(7):
//	    We will map this to 4, 5 respectively and
//	    consider as program doesn't exist in the system or invalid exited.
//
// Reference :
//   - LSB Returns Codes Screenshot : https://t.ly/3jkY
//   - LSB Returns Codes Screenshot : https://prnt.sc/26gunol
//   - Mapping (RawMapping)         : https://prnt.sc/26gwnxw
type ExitCode byte

const (
	Invalid ExitCode = iota
	ActiveRunning
	DeadButPidExists         // unit not failed
	DeadButVarLockFileExists // Unused
	NotRunning
	UnknownService // there is no service exist with that name.
	InvalidService
	InvalidCode // represents that not listed code
)

func (it ExitCode) ValueUInt16() uint16 {
	return uint16(it)
}

func (it ExitCode) AllNameValues() []string {
	return BasicEnumImpl.AllNameValues()
}

func (it ExitCode) OnlySupportedErr(names ...string) error {
	return BasicEnumImpl.OnlySupportedErr(names...)
}

func (it ExitCode) OnlySupportedMsgErr(message string, names ...string) error {
	return BasicEnumImpl.OnlySupportedMsgErr(message, names...)
}

func (it ExitCode) IntegerEnumRanges() []int {
	return BasicEnumImpl.IntegerEnumRanges()
}

func (it ExitCode) MinMaxAny() (min, max interface{}) {
	return BasicEnumImpl.MinMaxAny()
}

func (it ExitCode) MinValueString() string {
	return BasicEnumImpl.MinValueString()
}

func (it ExitCode) MaxValueString() string {
	return BasicEnumImpl.MaxValueString()
}

func (it ExitCode) MaxInt() int {
	return BasicEnumImpl.MaxInt()
}

func (it ExitCode) MinInt() int {
	return BasicEnumImpl.MinInt()
}

func (it ExitCode) RangesDynamicMap() map[string]interface{} {
	return BasicEnumImpl.RangesDynamicMap()
}

func (it ExitCode) IsAnyNamesOf(names ...string) bool {
	return BasicEnumImpl.IsAnyNamesOf(it.ValueByte(), names...)
}

func (it ExitCode) ValueInt() int {
	return int(it)
}

func (it ExitCode) IsAnyValuesEqual(anyByteValues ...byte) bool {
	return BasicEnumImpl.IsAnyOf(it.ValueByte(), anyByteValues...)
}

func (it ExitCode) IsByteValueEqual(value byte) bool {
	return it.ValueByte() == value
}

func (it ExitCode) IsNameEqual(name string) bool {
	return it.Name() == name
}

func (it ExitCode) IsValueEqual(value byte) bool {
	return it.ValueByte() == value
}

func (it ExitCode) ValueInt8() int8 {
	return int8(it)
}

func (it ExitCode) ValueInt16() int16 {
	return int16(it)
}

func (it ExitCode) ValueInt32() int32 {
	return int32(it)
}

func (it ExitCode) ValueString() string {
	return it.ToNumberString()
}

func (it ExitCode) Format(format string) (compiled string) {
	return BasicEnumImpl.Format(format, it.ValueByte())
}

func (it ExitCode) EnumType() enuminf.EnumTyper {
	return BasicEnumImpl.EnumType()
}

func (it ExitCode) IsUndefined() bool {
	return it == Invalid
}

func (it ExitCode) IsDefined() bool {
	return it != Invalid
}

func (it ExitCode) IsActiveRunning() bool {
	return it == ActiveRunning
}

func (it ExitCode) IsDeadButPidExists() bool {
	return it == DeadButPidExists
}

func (it ExitCode) IsDeadButVarLockFileExists() bool {
	return it == DeadButVarLockFileExists
}

func (it ExitCode) IsNotRunning() bool {
	return it == NotRunning
}

func (it ExitCode) IsUnknownService() bool {
	return it == UnknownService
}

func (it ExitCode) IsInvalidService() bool {
	return it == InvalidService
}

func (it ExitCode) IsInvalidCode() bool {
	return it == InvalidCode
}

func (it ExitCode) IsSuccess() bool {
	return it == ActiveRunning
}

func (it ExitCode) IsFailed() bool {
	return it != ActiveRunning
}

func (it ExitCode) IsAllOf(codes ...int) bool {
	if len(codes) == 0 {
		return false
	}

	for _, code := range codes {
		if !it.IsEqual(code) {
			return false
		}
	}

	return true
}

func (it ExitCode) IsAnyOf(codes ...int) bool {
	if len(codes) == 0 {
		return false
	}

	for _, code := range codes {
		if it.IsEqual(code) {
			return true
		}
	}

	return false
}

func (it ExitCode) IsEqual(code int) bool {
	if code == InvalidExitCode || code < 0 {
		return false
	}

	if code > constants.MaxUnit8AsInt {
		return false
	}

	return byte(code) == it.ValueByte()
}

func (it ExitCode) IsInvalid() bool {
	return it == Invalid
}

func (it ExitCode) IsValid() bool {
	return it != Invalid
}

func (it ExitCode) IsAnyOfExitCode(anyOfItems ...ExitCode) bool {
	for _, item := range anyOfItems {
		if item == it {
			return true
		}
	}

	return false
}

func (it ExitCode) IsNameOf(anyNames ...string) bool {
	for _, name := range anyNames {
		if name == it.Name() {
			return true
		}
	}

	return false
}

func (it ExitCode) Name() string {
	return BasicEnumImpl.ToEnumString(
		it.Value())
}

func (it ExitCode) RangeNamesCsv() string {
	return BasicEnumImpl.RangeNamesCsv()
}

func (it ExitCode) TypeName() string {
	return BasicEnumImpl.TypeName()
}

func (it ExitCode) ToNumberString() string {
	return BasicEnumImpl.ToNumberString(it.Value())
}

func (it ExitCode) Value() byte {
	return byte(it)
}

func (it ExitCode) UnmarshallEnumToValue(jsonUnmarshallingValue []byte) (byte, error) {
	return BasicEnumImpl.UnmarshallToValue(
		false,
		jsonUnmarshallingValue)
}

func (it ExitCode) MaxByte() byte {
	return BasicEnumImpl.Max()
}

func (it ExitCode) MinByte() byte {
	return BasicEnumImpl.Min()
}

func (it ExitCode) ValueByte() byte {
	return byte(it)
}

func (it ExitCode) RangesByte() []byte {
	return Ranges[:]
}

func (it ExitCode) ToByteEnumString(input byte) string {
	return BasicEnumImpl.ToNumberString(input)
}

func (it ExitCode) NameValue() string {
	return BasicEnumImpl.NameWithValue(it)
}

func (it ExitCode) String() string {
	return BasicEnumImpl.ToEnumString(it.Value())
}

func (it ExitCode) MarshalJSON() ([]byte, error) {
	return BasicEnumImpl.ToEnumJsonBytes(it.Value())
}

func (it *ExitCode) UnmarshalJSON(data []byte) error {
	dataConv, err := it.UnmarshallEnumToValue(data)

	if err == nil {
		*it = ExitCode(dataConv)
	}

	return err
}

func (it ExitCode) AsBasicEnumContractsBinder() enuminf.BasicEnumContractsBinder {
	return &it
}

func (it ExitCode) AsJsonMarshaller() corejson.JsonMarshaller {
	return &it
}

func (it ExitCode) AsBasicByteEnumContractsBinder() enuminf.BasicByteEnumContractsBinder {
	return &it
}

func (it ExitCode) AsBasicByteEnumContractsDelegateBinder() enuminf.BasicByteEnumContractsDelegateBinder {
	return &it
}

func (it ExitCode) ToPtr() *ExitCode {
	return &it
}
