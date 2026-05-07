package eventtype

import (
	"encoding/json"
	"testing"
)

// AL2-05 Batch E coverage suite for eventtype.

func TestEventType_NewAndAccessors(t *testing.T) {
	for _, name := range []string{"Log", "Success", "Error", "Failure", "File", "Custom"} {
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
	if Min() != Invalid || Max() != Custom {
		t.Errorf("Min/Max wrong: %v %v", Min(), Max())
	}
	if RangesInvalidErr() == nil {
		t.Error("RangesInvalidErr should be non-nil")
	}
	v := Failure
	if !v.IsFailure() || Custom.IsCustom() != true || File.IsFile() != true {
		t.Error("predicate wrong")
	}
	if v.ValueUInt16() != uint16(v) {
		t.Error("ValueUInt16 mismatch")
	}
	if len(v.AllNameValues()) == 0 || len(v.IntegerEnumRanges()) == 0 ||
		len(v.RangesDynamicMap()) == 0 {
		t.Error("ranges empty")
	}
	if v.Name() == "" || v.MinValueString() == "" || v.MaxValueString() == "" {
		t.Error("string accessors empty")
	}
	if min, max := v.MinMaxAny(); min == nil || max == nil {
		t.Error("MinMaxAny nil")
	}
	data, err := json.Marshal(v)
	if err != nil || len(data) == 0 {
		t.Errorf("Marshal: %v", err)
	}
	var got Variant
	if err := json.Unmarshal(data, &got); err != nil {
		t.Errorf("Unmarshal: %v", err)
	}
	_ = v.OnlySupportedErr("Log")
	_ = v.OnlySupportedMsgErr("ctx", "Log")

	defer func() { _ = recover() }()
	_ = v.HasPattern("x")
}
