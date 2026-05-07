package completionstate

import (
	"encoding/json"
	"testing"
)

func Test_CompletionState_Coverage(t *testing.T) {
	all := []Variant{Invalid, Initiate, Running, Success, SuccessWithWarning, FailedMiddleWithError, CompleteWithError}
	for _, v := range all {
		_ = v.ValueByte()
		_ = v.ValueInt()
		_ = v.ValueInt8()
		_ = v.ValueInt16()
		_ = v.ValueInt32()
		_ = v.ValueUInt16()
		_ = v.Value()
		_ = v.ValueString()
		_ = v.Name()
		_ = v.ToNumberString()
		_ = v.RangeNamesCsv()
		_ = v.TypeName()
		_ = v.AllNameValues()
		_ = v.RangesDynamicMap()
		_ = v.IntegerEnumRanges()
		_ = v.MinByte()
		_ = v.MaxByte()
		_ = v.MinInt()
		_ = v.MaxInt()
		_ = v.MinValueString()
		_ = v.MaxValueString()
		_, _ = v.MinMaxAny()
		_ = v.Format("name")
		_ = v.EnumType()
		_ = v.IsByteValueEqual(v.ValueByte())
		_ = v.IsValueEqual(v.ValueByte())
		_ = v.IsNameEqual(v.Name())
		_ = v.IsAnyValuesEqual(v.ValueByte())
		_ = v.IsAnyOf(v)
		_ = v.IsAnyNamesOf(v.Name())
		_ = v.IsStartState()
		_ = v.IsInitiate()
		_ = v.IsRunning()
		_ = v.IsSuccessWithWarning()
		_ = v.IsFailedMiddleWithError()
		_ = v.IsCompleteWithError()
		_ = v.IsEndState()
		_ = v.IsCompletedSuccess()
		_ = v.IsCompletedWithIssues()
		_ = v.IsSuccess()
		_ = v.IsCompletedLogically()
		_ = v.IsSuccessLogically()
		_ = v.IsCompletedWithErrorLogically()
		_ = v.HasErrorLogically()
		_ = v.IsInvalid()
		_ = v.IsValid()
		_ = v.OnlySupportedErr("Success")
		_ = v.OnlySupportedMsgErr("m", "Success")

		raw, err := json.Marshal(v)
		if err != nil {
			t.Fatalf("marshal: %v", err)
		}
		var rt Variant
		_ = json.Unmarshal(raw, &rt)
	}
}
