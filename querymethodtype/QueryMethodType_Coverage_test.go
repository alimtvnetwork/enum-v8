package querymethodtype

import (
	"encoding/json"
	"testing"
)

// AL2-02 Batch B coverage suite for querymethodtype.
// Covers Variant predicates, accessors, JSON round-trip, and Min/Max range
// helpers. Pattern follows resauthtype/dbexposetype Batch-B suites.

func TestQueryMethod_MinMaxAndPredicates(t *testing.T) {
	if Min() != Invalid {
		t.Errorf("Min() = %v want Invalid", Min())
	}
	if Max() != ByServerName {
		t.Errorf("Max() = %v want ByServerName", Max())
	}
	if !Invalid.IsInvalid() || Invalid.IsValid() {
		t.Error("Invalid predicates wrong")
	}
	if !ByFile.IsByFile() || ByFile.IsByServerName() {
		t.Error("ByFile predicate wrong")
	}
	if !ByServerName.IsByServerName() || ByServerName.IsByFile() {
		t.Error("ByServerName predicate wrong")
	}
	if !ByFile.IsValid() {
		t.Error("ByFile should be valid")
	}
}

func TestQueryMethod_Accessors(t *testing.T) {
	v := ByServerName
	if v.Value() != byte(v) || v.ValueByte() != byte(v) ||
		v.ValueInt() != int(v) || v.ValueInt8() != int8(v) ||
		v.ValueInt16() != int16(v) || v.ValueInt32() != int32(v) ||
		v.ValueUInt16() != uint16(v) {
		t.Error("numeric accessors mismatch")
	}
	if v.Name() == "" || v.String() == "" || v.NameValue() == "" ||
		v.ToNumberString() == "" || v.ValueString() == "" {
		t.Error("string accessors empty")
	}
	if v.MaxByte() != byte(Max()) || v.MinByte() != byte(Min()) {
		t.Error("Min/MaxByte mismatch")
	}
	if v.MaxInt() < v.MinInt() {
		t.Error("MaxInt < MinInt")
	}
	if v.MinValueString() == "" || v.MaxValueString() == "" {
		t.Error("Min/MaxValueString empty")
	}
	if min, max := v.MinMaxAny(); min == nil || max == nil {
		t.Error("MinMaxAny nil")
	}
	if len(v.AllNameValues()) == 0 || len(v.IntegerEnumRanges()) == 0 ||
		len(v.RangesByte()) == 0 || len(v.RangesDynamicMap()) == 0 {
		t.Error("range accessors empty")
	}
	if v.RangeNamesCsv() == "" || v.TypeName() == "" {
		t.Error("csv/typename empty")
	}
	if v.EnumType() == nil {
		t.Error("EnumType nil")
	}
	if v.Format("%s") == "" {
		t.Error("Format empty")
	}
}

func TestQueryMethod_EqualityAndIsAny(t *testing.T) {
	v := ByFile
	if !v.IsByteValueEqual(byte(ByFile)) || v.IsByteValueEqual(byte(ByServerName)) {
		t.Error("IsByteValueEqual wrong")
	}
	if !v.IsValueEqual(byte(ByFile)) || !v.IsNameEqual("ByFile") {
		t.Error("IsValueEqual/IsNameEqual wrong")
	}
	if !v.IsAnyValuesEqual(byte(ByFile), byte(ByServerName)) {
		t.Error("IsAnyValuesEqual should match")
	}
	if v.IsAnyValuesEqual(byte(Invalid)) {
		t.Error("IsAnyValuesEqual should not match Invalid")
	}
	if !v.IsAnyNamesOf("ByFile", "ByServerName") {
		t.Error("IsAnyNamesOf should match")
	}
	if v.IsAnyNamesOf("nope") {
		t.Error("IsAnyNamesOf should not match unknown")
	}
}

func TestQueryMethod_JsonRoundTrip(t *testing.T) {
	for _, v := range []Variant{Invalid, ByFile, ByServerName} {
		data, err := json.Marshal(v)
		if err != nil {
			t.Fatalf("Marshal %v: %v", v, err)
		}
		if len(data) == 0 {
			t.Errorf("%v: empty json", v)
		}
		var got Variant
		if err := json.Unmarshal(data, &got); err != nil {
			t.Errorf("Unmarshal %v: %v", v, err)
		}
	}
	if err := RangesInvalidErr(); err == nil {
		t.Error("RangesInvalidErr should be non-nil informational error")
	}
}

func TestQueryMethod_PtrAndBinders(t *testing.T) {
	v := ByFile
	p := v.ToPtr()
	if p == nil || p.ToSimple() != ByFile {
		t.Error("ToPtr/ToSimple wrong")
	}
	var nilP *Variant
	if nilP.ToSimple() != Invalid {
		t.Error("nil ToSimple should be Invalid")
	}
	if v.AsJsonMarshaller() == nil || v.AsBasicByteEnumContractsBinder() == nil ||
		v.AsBasicEnumContractsBinder() == nil {
		t.Error("binder nil")
	}
	if p.AsJsonContractsBinder() == nil || p.AsJsoner() == nil {
		t.Error("ptr binders nil")
	}
	if v.Json().IsEmpty() && v.JsonPtr() == nil {
		t.Error("Json/JsonPtr broken")
	}
	jr := v.Json()
	var dst Variant
	if err := dst.JsonParseSelfInject(&jr); err != nil {
		t.Errorf("JsonParseSelfInject: %v", err)
	}
}
