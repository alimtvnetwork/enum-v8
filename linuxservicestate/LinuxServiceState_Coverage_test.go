package linuxservicestate

import (
	"encoding/json"
	"testing"
)

// AL2-04 Batch D coverage suite for linuxservicestate.

func TestLinuxServiceState_NewAndCodeMapping(t *testing.T) {
	for _, name := range []string{"ActiveRunning", "DeadButPidExists", "DeadButVarLockFileExists",
		"NotRunning", "UnknownService", "InvalidService", "InvalidCode"} {
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
	if Min() != ActiveRunning || Max() != InvalidCode {
		t.Errorf("Min/Max wrong: %v %v", Min(), Max())
	}
	if RangesInvalidErr() == nil {
		t.Error("RangesInvalidErr should be non-nil informational")
	}
	// NewCode raw mappings
	if NewCode(0) != ActiveRunning {
		t.Error("NewCode(0) wrong")
	}
	if NewCode(-1) != InvalidCode || NewCode(999) != InvalidCode {
		t.Error("NewCode out-of-range should be InvalidCode")
	}
	if NewCodeMapping(0) != ActiveRunning {
		t.Error("NewCodeMapping(0) wrong")
	}
	if NewCodeMapping(255) != InvalidCode {
		t.Error("NewCodeMapping out-of-range should be InvalidCode")
	}
}

func TestLinuxServiceState_Accessors(t *testing.T) {
	v := ActiveRunning
	if v.ValueByte() != byte(v) || v.ValueInt() != int(v) || v.ValueInt8() != int8(v) ||
		v.ValueInt16() != int16(v) || v.ValueInt32() != int32(v) || v.ValueUInt16() != uint16(v) {
		t.Error("numeric accessors mismatch")
	}
	if v.MaxByte() != byte(BasicEnumImpl.Max()) || v.MinByte() != byte(BasicEnumImpl.Min()) {
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
	if v.Name() == "" || v.NameValue() == "" || v.ValueString() == "" || v.ToNumberString() == "" ||
		v.RangeNamesCsv() == "" || v.TypeName() == "" || v.Format("%s") == "" {
		t.Error("string accessors empty")
	}
	if v.EnumType() == nil {
		t.Error("EnumType nil")
	}
	if !v.IsByteValueEqual(byte(ActiveRunning)) || v.IsByteValueEqual(byte(NotRunning)) ||
		!v.IsValueEqual(byte(ActiveRunning)) || !v.IsNameEqual("ActiveRunning") {
		t.Error("equality wrong")
	}
	if !v.IsAnyValuesEqual(byte(ActiveRunning), byte(NotRunning)) ||
		!v.IsAnyNamesOf("ActiveRunning", "NotRunning") {
		t.Error("IsAny wrong")
	}
	data, err := json.Marshal(v)
	if err != nil || len(data) == 0 {
		t.Errorf("Marshal: %v", err)
	}
	var got ExitCode
	if err := json.Unmarshal(data, &got); err != nil {
		t.Errorf("Unmarshal: %v", err)
	}
	_ = v.OnlySupportedErr("ActiveRunning")
	_ = v.OnlySupportedMsgErr("ctx", "ActiveRunning")
}
