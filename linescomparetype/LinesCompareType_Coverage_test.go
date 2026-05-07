package linescomparetype

import (
	"encoding/json"
	"testing"
)

// AL2-05 Batch E coverage suite for linescomparetype.

func TestLinesCompareType_NewAndAccessors(t *testing.T) {
	for _, name := range []string{"Equal", "Less", "LessEqual", "Greater", "GreaterEqual", "NotEqual"} {
		v, err := New(name)
		if err != nil {
			t.Errorf("New(%q): %v", name, err)
		}
		if v.Name() != name {
			t.Errorf("Name mismatch: %q != %q", v.Name(), name)
		}
		_ = NewMust(name)
		if RangesMap[name] != v {
			t.Errorf("RangesMap mismatch for %q", name)
		}
	}
	if _, err := New("__bogus__"); err == nil {
		t.Error("New bogus should fail")
	}
	if Min() != Invalid || Max() != NotEqual {
		t.Errorf("Min/Max wrong: %v %v", Min(), Max())
	}
	if RangesInvalidErr() == nil {
		t.Error("RangesInvalidErr should be non-nil")
	}
	v := Equal
	if v.ValueUInt16() != uint16(v) {
		t.Error("ValueUInt16 mismatch")
	}
	if len(v.AllNameValues()) == 0 || len(v.IntegerEnumRanges()) == 0 {
		t.Error("ranges empty")
	}
	if v.MinValueString() == "" || v.MaxValueString() == "" {
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
	_ = v.OnlySupportedErr("Equal")
	_ = v.OnlySupportedMsgErr("ctx", "Equal")
}
