package inttype

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"

	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coreinterface/enuminf"
	"github.com/alimtvnetwork/core-v8/errcore"
)

func New(val int) Variant {
	return Variant(val)
}

func NewString(valueString string) (Variant, error) {
	val, err := strconv.Atoi(valueString)

	if err == nil {
		return Variant(val), nil
	}

	return InvalidValue, errcore.
		FailedToConvertType.
		Error(
			"cannot convert string to inttype!"+err.Error(), valueString)
}

func NewUInt(val uint) (Variant, error) {
	if val <= math.MaxInt {
		return Variant(val), nil
	}

	return InvalidValue, errcore.
		FailedToConvertType.
		Error(
			"cannot convert uint to inttype!", val)
}

func NewInt64(val int64) (Variant, error) {
	if val <= math.MaxInt {
		return Variant(val), nil
	}

	return InvalidValue, errcore.
		FailedToConvertType.
		Error(
			"cannot convert int64 to inttype!", val)
}

func NewUsingEnum(valueEnum enuminf.BasicEnumer) (Variant, error) {
	return NewString(valueEnum.Name())
}

func NewUsingStringer(valueStringer fmt.Stringer) (Variant, error) {
	return NewString(valueStringer.String())
}

func NewUsingJsoner(jsoner corejson.Jsoner) (Variant, error) {
	jsonPtr := jsoner.JsonPtr()

	return NewUsingJsonResult(jsonPtr)
}

func NewUsingJsonResult(jsonResult *corejson.Result) (Variant, error) {
	if jsonResult.HasError() {
		return Invalid, jsonResult.MeaningfulError()
	}

	return NewString(jsonResult.JsonString())
}

func NewUsingJsonNumber(jsonNumber *json.Number) (Variant, error) {
	if jsonNumber == nil {
		return Invalid,
			errcore.
				InvalidNullPointerType.
				ErrorNoRefs("nil jsonNumber")
	}

	return NewString(jsonNumber.String())
}
