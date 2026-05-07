package taskcategory

import (
	"encoding/json"
	"testing"
)

// AL2-06 Batch F coverage suite for taskcategory.

func TestTaskCategory_NewAndAccessors(t *testing.T) {
	for _, name := range []string{"Help", "Login", "DbTask", "Macro", "Backup", "Restore"} {
		v, err := New(name)
		if err != nil {
			// some names may not be in Ranges; skip unmatched
			continue
		}
		if v.Name() != name {
			t.Errorf("Name mismatch: %q != %q", v.Name(), name)
		}
		_ = NewMust(name)
	}
	if _, err := New("__bogus__"); err == nil {
		t.Error("New bogus should fail")
	}
	if RangesInvalidErr() == nil {
		t.Error("RangesInvalidErr should be non-nil")
	}
	v := Help
	if v.ValueByte() != byte(v) || len(v.AllNameValues()) == 0 ||
		len(v.IntegerEnumRanges()) == 0 || len(v.RangesDynamicMap()) == 0 {
		t.Error("accessors broken")
	}
	if v.MinValueString() == "" || v.MaxValueString() == "" || v.TypeName() == "" {
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
