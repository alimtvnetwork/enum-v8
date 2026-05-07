package revokereason

import (
	"encoding/json"
	"testing"
)

// AL2-05 Batch E coverage suite for revokereason.

func TestRevokeReason_NewAndAccessors(t *testing.T) {
	for _, name := range []string{"Unspecified", "KeyCompromise", "CaCompromise", "AffiliationChanged", "Superseded", "CessationOfOperation", "CertificateHold", "RemoveFromCRL", "PrivilegeWithdrawn", "AaCompromise"} {
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
	if Min() != Unspecified || Max() != AaCompromise {
		t.Errorf("Min/Max wrong: %v %v", Min(), Max())
	}
	if RangesInvalidErr() == nil {
		t.Error("RangesInvalidErr should be non-nil")
	}
	v := KeyCompromise
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
	_ = v.OnlySupportedErr("KeyCompromise")
	_ = v.OnlySupportedMsgErr("ctx", "KeyCompromise")
}
