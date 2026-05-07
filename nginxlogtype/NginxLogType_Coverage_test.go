package nginxlogtype

import (
	"encoding/json"
	"testing"
)

// AL2-03 Batch C coverage suite for nginxlogtype.

func TestNginxLogType_NewTypeAndPredicates(t *testing.T) {
	for name, want := range RangesMap {
		got := NewType(name)
		if got != want {
			t.Errorf("NewType(%q) = %v want %v", name, got, want)
		}
	}
	if Min() != Invalid || Max() != DuplicateDefaultError {
		t.Errorf("Min/Max wrong: %v %v", Min(), Max())
	}
	if RangesInvalidErr() == nil {
		t.Error("RangesInvalidErr should be informational")
	}
	if !Invalid.IsUnknown() || !Error.IsError() || !Warning.IsWarning() ||
		!Notice.IsNotice() || !AlertError.IsAlert() || !FileIssueError.IsFileIssue() ||
		!SyntaxIssueError.IsSyntaxIssue() || !DuplicateDomainWarningError.IsDuplicateDomain() ||
		!DuplicateDefaultError.IsDuplicateDefault() {
		t.Error("Is* predicate wrong")
	}
	if !Error.IsAnyKindOfError() || Notice.IsAnyKindOfError() {
		t.Error("IsAnyKindOfError wrong")
	}
	if !Notice.IsNotError() || Error.IsNotError() {
		t.Error("IsNotError wrong")
	}
}

func TestNginxLogType_LevelComparisons(t *testing.T) {
	if !Error.IsEqual(Error) || Error.IsEqual(Notice) {
		t.Error("IsEqual wrong")
	}
	if !Error.IsAboveOrEqual(Notice) || !Notice.IsLowerOrEqual(Error) {
		t.Error("level comparisons wrong")
	}
}

func TestNginxLogType_Accessors(t *testing.T) {
	v := Error
	if v.ValueByte() != byte(v) || v.ValueInt() != int(v) || v.ValueInt8() != int8(v) ||
		v.ValueInt16() != int16(v) || v.ValueInt32() != int32(v) || v.ValueUInt16() != uint16(v) ||
		v.Value() != byte(v) {
		t.Error("numeric accessors mismatch")
	}
	if v.MaxByte() != byte(Max()) || v.MinByte() != byte(Min()) {
		t.Error("Min/MaxByte mismatch")
	}
	if v.MaxInt() < v.MinInt() || v.MinValueString() == "" || v.MaxValueString() == "" {
		t.Error("Min/Max wrong")
	}
	if min, max := v.MinMaxAny(); min == nil || max == nil {
		t.Error("MinMaxAny nil")
	}
	if len(v.AllNameValues()) == 0 || len(v.IntegerEnumRanges()) == 0 ||
		len(v.RangesByte()) == 0 || len(v.RangesDynamicMap()) == 0 {
		t.Error("ranges empty")
	}
	if v.Name() == "" || v.String() == "" || v.NameValue() == "" || v.ValueString() == "" ||
		v.ToNumberString() == "" || v.RangeNamesCsv() == "" || v.TypeName() == "" ||
		v.Format("%s") == "" {
		t.Error("string accessors empty")
	}
	if v.EnumType() == nil {
		t.Error("EnumType nil")
	}
}

func TestNginxLogType_EqualityAndJson(t *testing.T) {
	v := Error
	if !v.IsByteValueEqual(byte(Error)) || v.IsByteValueEqual(byte(Notice)) ||
		!v.IsValueEqual(byte(Error)) || !v.IsNameEqual("Error") {
		t.Error("equality wrong")
	}
	if !v.IsAnyValuesEqual(byte(Error), byte(Notice)) || v.IsAnyValuesEqual(byte(Invalid)) {
		t.Error("IsAnyValuesEqual wrong")
	}
	if !v.IsAnyNamesOf("Error", "Notice") || v.IsAnyNamesOf("nope") {
		t.Error("IsAnyNamesOf wrong")
	}
	if !v.IsValid() || Invalid.IsValid() || !Invalid.IsInvalid() {
		t.Error("valid wrong")
	}
	if !v.IsAnyOf(Error, Warning) || v.IsAnyOf(Notice) {
		t.Error("IsAnyOf wrong")
	}
	for _, x := range []Variant{Invalid, Notice, Warning, Error, AlertError} {
		data, err := json.Marshal(x)
		if err != nil || len(data) == 0 {
			t.Errorf("Marshal(%v): %v", x, err)
		}
		var got Variant
		if err := json.Unmarshal(data, &got); err != nil {
			t.Errorf("Unmarshal(%v): %v", x, err)
		}
	}
	if v.ToPtr() == nil || v.JsonPtr() == nil {
		t.Error("ToPtr/JsonPtr nil")
	}
	if v.AsJsonMarshaller() == nil || v.AsJsonContractsBinder() == nil ||
		v.AsJsoner() == nil || v.AsBasicByteEnumContractsBinder() == nil ||
		v.AsBasicEnumContractsBinder() == nil {
		t.Error("binder nil")
	}
	_ = v.OnlySupportedErr("Error")
	_ = v.OnlySupportedMsgErr("ctx", "Error")
}
