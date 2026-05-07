package osarchs

import (
	"encoding/json"
	"testing"
)

// AL2-04 Batch D coverage suite for osarchs.

func TestOsArchs_NewAndGet(t *testing.T) {
	for _, name := range []string{"x32", "x64"} {
		v, err := New(name)
		if err != nil {
			t.Errorf("New(%q): %v", name, err)
		}
		if v == Invalid {
			t.Errorf("New(%q) yielded Invalid", name)
		}
		_ = NewMust(name)
	}
	if _, err := New("__bogus__"); err == nil {
		t.Error("New bogus should fail")
	}
	// Get via aliases
	if Get("amd64") != X64 || Get("386") != X32 || Get("__nope__") != Invalid {
		t.Error("Get alias mapping wrong")
	}
	if Min() != Invalid {
		t.Errorf("Min %v", Min())
	}
	if Max() == Invalid {
		t.Error("Max should not be Invalid")
	}
	if RangesInvalidErr() == nil {
		t.Error("RangesInvalidErr should be non-nil informational")
	}
}

func TestOsArchs_Predicates(t *testing.T) {
	if !X32.IsX32() || X32.IsX64() {
		t.Error("X32 predicates wrong")
	}
	if !X64.IsX64() || X64.IsX32() {
		t.Error("X64 predicates wrong")
	}
}

func TestOsArchs_Accessors(t *testing.T) {
	v := X64
	if v.ValueByte() != byte(v) || v.ValueInt() != int(v) || v.ValueInt8() != int8(v) ||
		v.ValueInt16() != int16(v) || v.ValueInt32() != int32(v) || v.ValueUInt16() != uint16(v) {
		t.Error("numeric accessors mismatch")
	}
	if v.MaxByte() != byte(Max()) {
		t.Error("MaxByte mismatch")
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
	if v.Name() == "" || v.NameValue() == "" || v.ValueString() == "" ||
		v.RangeNamesCsv() == "" || v.TypeName() == "" || v.Format("%s") == "" {
		t.Error("string accessors empty")
	}
	if v.EnumType() == nil {
		t.Error("EnumType nil")
	}
	if !v.IsByteValueEqual(byte(X64)) || v.IsByteValueEqual(byte(X32)) ||
		!v.IsValueEqual(byte(X64)) || !v.IsNameEqual("x64") {
		t.Error("equality wrong")
	}
	if !v.IsAnyValuesEqual(byte(X64), byte(X32)) ||
		!v.IsAnyNamesOf("x64", "x32") {
		t.Error("IsAny wrong")
	}
	data, err := json.Marshal(v)
	if err != nil || len(data) == 0 {
		t.Errorf("Marshal: %v", err)
	}
	var got Architecture
	if err := json.Unmarshal(data, &got); err != nil {
		t.Errorf("Unmarshal: %v", err)
	}
	_ = v.OnlySupportedErr("x64")
	_ = v.OnlySupportedMsgErr("ctx", "x64")
	_ = CurrentArch
}
