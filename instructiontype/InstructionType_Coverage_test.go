package instructiontype

import (
	"encoding/json"
	"testing"
)

// AL2-05 Batch E coverage suite for instructiontype.

func TestInstructionType_NewAndAccessors(t *testing.T) {
	for _, name := range []string{"Scoping", "DependsOn", "InstallPackages", "Nginx", "Apache", "Verify"} {
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
	if Min() != Invalid || Max() != Verify {
		t.Errorf("Min/Max wrong: %v %v", Min(), Max())
	}
	if RangesInvalidErr() == nil {
		t.Error("RangesInvalidErr should be non-nil")
	}
	if !Invalid.IsUninitialized() || Scoping.IsUninitialized() {
		t.Error("IsUninitialized wrong")
	}
	v := Nginx
	if len(v.AllNameValues()) == 0 || len(v.IntegerEnumRanges()) == 0 {
		t.Error("ranges empty")
	}
	data, err := json.Marshal(v)
	if err != nil || len(data) == 0 {
		t.Errorf("Marshal: %v", err)
	}
	var got Variant
	if err := json.Unmarshal(data, &got); err != nil {
		t.Errorf("Unmarshal: %v", err)
	}
	_ = v.OnlySupportedErr("Nginx")
	_ = v.OnlySupportedMsgErr("ctx", "Nginx")
}
