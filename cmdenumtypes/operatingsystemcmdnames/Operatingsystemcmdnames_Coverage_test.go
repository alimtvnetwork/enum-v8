package operatingsystemcmdnames

import (
	"encoding/json"
	"testing"
)

// AL2-09 cmdenumtypes coverage sweep — generated uniform test.

func TestOperatingsystemcmdnames_NewAndAccessors(t *testing.T) {
	v, err := New("Help")
	if err != nil {
		t.Fatalf("New(%q): %v", "Help", err)
	}
	if v.Name() != "Help" {
		t.Errorf("Name mismatch: %q != %q", v.Name(), "Help")
	}
	_ = NewMust("Help")
	if _, err := New("__bogus__"); err == nil {
		t.Error("New bogus should fail")
	}
	if Min() != Invalid {
		t.Errorf("Min wrong: %v", Min())
	}
	if Max() == Invalid {
		t.Errorf("Max wrong: %v", Max())
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
	_ = v.OnlySupportedErr("Help")
	_ = v.OnlySupportedMsgErr("ctx", "Help")
}
