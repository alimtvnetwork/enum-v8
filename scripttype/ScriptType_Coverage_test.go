package scripttype

import (
	"encoding/json"
	"testing"
)

// AL2-06 Batch F coverage suite for scripttype.

func TestScriptType_NewAndAccessors(t *testing.T) {
	for _, name := range []string{"Shell", "Bash", "Perl", "Python", "Python2", "Python3", "CLang", "MakeScript", "Powershell", "Cmd"} {
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
	if Min() != Default || Max() != Invalid {
		t.Errorf("Min/Max: %v %v", Min(), Max())
	}
	if RangesInvalidErr() == nil {
		t.Error("RangesInvalidErr nil")
	}
	v := Bash
	if v.ValueUInt16() != uint16(v) || len(v.AllNameValues()) == 0 ||
		len(v.IntegerEnumRanges()) == 0 {
		t.Error("accessors broken")
	}
	if min, max := v.MinMaxAny(); min == nil || max == nil {
		t.Error("MinMaxAny nil")
	}
	if v.MinValueString() == "" || v.MaxValueString() == "" {
		t.Error("string accessors empty")
	}
	data, err := json.Marshal(v)
	if err != nil || len(data) == 0 {
		t.Errorf("Marshal: %v", err)
	}
	var got Variant
	if err := json.Unmarshal(data, &got); err != nil {
		t.Errorf("Unmarshal: %v", err)
	}
	_ = v.OnlySupportedErr("Bash")
	_ = v.OnlySupportedMsgErr("ctx", "Bash")

	// helpers
	if CondBool(true, Bash, Cmd) != Bash || CondBool(false, Bash, Cmd) != Cmd {
		t.Error("CondBool wrong")
	}
	if CondFunc(true, func() Variant { return Powershell }) != Powershell {
		t.Error("CondFunc wrong (true)")
	}
	if CondFunc(false, func() Variant { return Powershell }) != Invalid {
		t.Error("CondFunc wrong (false)")
	}
	if OsDefaultScriptType() == Invalid {
		t.Error("OsDefaultScriptType Invalid")
	}
	if DefaultOsScript() == nil {
		t.Error("DefaultOsScript nil")
	}
}
