package taskpriority

import (
	"encoding/json"
	"testing"
)

// AL2-06 Batch F coverage suite for taskpriority.

func TestTaskPriority_NewAndAccessors(t *testing.T) {
	for _, name := range []string{"Default", "DefaultLock", "Reminder", "Notification", "SystemUpdate", "LowerPriority"} {
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
	if RangesInvalidErr := (func() error { _, err := New("__bogus__"); return err })(); RangesInvalidErr == nil {
		t.Error("expected error")
	}
	pm := GetPriorityMap()
	if pm["Default"] != 40 {
		t.Errorf("priorityMap[Default] = %d", pm["Default"])
	}
	if PriorityMapString() == "" {
		t.Error("PriorityMapString empty")
	}
	v := Default
	if v.ValueUInt16() != uint16(v) || len(v.AllNameValues()) == 0 ||
		len(v.IntegerEnumRanges()) == 0 || len(v.RangesDynamicMap()) == 0 {
		t.Error("accessors broken")
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
	_ = v.OnlySupportedErr("Default")
	_ = v.OnlySupportedMsgErr("ctx", "Default")
}
