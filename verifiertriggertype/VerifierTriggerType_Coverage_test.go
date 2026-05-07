package verifiertriggertype

import (
	"encoding/json"
	"testing"
)

// AL2-03 Batch C coverage suite for verifiertriggertype.

func TestVerifierTriggerType_NewAndRange(t *testing.T) {
	for _, name := range []string{"AllComplete", "AfterRestart", "AfterNetworkReset"} {
		v, err := New(name)
		if err != nil {
			t.Errorf("New(%q): %v", name, err)
		}
		if v.Name() != name {
			t.Errorf("Name mismatch: %q != %q", v.Name(), name)
		}
		_ = NewMust(name)
	}
	if _, err := New("__bogus__"); err == nil {
		t.Error("New bogus should fail")
	}
	if Min() != Invalid || Max() != AfterNetworkReset {
		t.Errorf("Min/Max wrong: %v %v", Min(), Max())
	}
	if RangesInvalidErr() == nil {
		t.Error("RangesInvalidErr should be informational")
	}
}

func TestVerifierTriggerType_Accessors(t *testing.T) {
	v := AllComplete
	if v.ValueByte() != byte(v) || v.ValueInt() != int(v) || v.ValueInt8() != int8(v) ||
		v.ValueInt16() != int16(v) || v.ValueInt32() != int32(v) || v.ValueUInt16() != uint16(v) {
		t.Error("numeric accessor mismatch")
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

func TestVerifierTriggerType_EqualityAndJson(t *testing.T) {
	v := AllComplete
	if !v.IsByteValueEqual(byte(AllComplete)) || v.IsByteValueEqual(byte(AfterRestart)) ||
		!v.IsValueEqual(byte(AllComplete)) || !v.IsNameEqual("AllComplete") {
		t.Error("equality wrong")
	}
	if !v.IsAnyValuesEqual(byte(AllComplete), byte(AfterRestart)) ||
		v.IsAnyValuesEqual(byte(Invalid)) {
		t.Error("IsAnyValuesEqual wrong")
	}
	if !v.IsAnyNamesOf("AllComplete", "AfterRestart") || v.IsAnyNamesOf("nope") {
		t.Error("IsAnyNamesOf wrong")
	}
	if !v.IsValid() || Invalid.IsValid() || !Invalid.IsInvalid() {
		t.Error("valid wrong")
	}
	if !v.IsAnyOf(AllComplete, AfterRestart) || v.IsAnyOf(AfterRestart) {
		t.Error("IsAnyOf wrong")
	}
	for _, x := range []Variant{Invalid, AllComplete, AfterRestart, AfterNetworkReset} {
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
	if v.AsJsonMarshaller() == nil || v.AsBasicByteEnumContractsBinder() == nil ||
		v.AsBasicEnumContractsBinder() == nil {
		t.Error("binder nil")
	}
	_ = v.OnlySupportedErr("AllComplete")
	_ = v.OnlySupportedMsgErr("ctx", "AllComplete")
}
