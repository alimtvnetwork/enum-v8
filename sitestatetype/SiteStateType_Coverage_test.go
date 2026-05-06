package sitestatetype

import "testing"

// AL2-01 Batch A coverage suite for sitestatetype.
// Covers New/NewMust/Max/RangesInvalidErr + Variant accessors.

func TestSiteStateType_NewRoundTrip(t *testing.T) {
	names := []string{"NewlyAdded", "Removed", "Unchanged"}
	for _, n := range names {
		v, err := New(n)
		if err != nil {
			t.Errorf("New(%q) error: %v", n, err)
			continue
		}
		if v == Invalid {
			t.Errorf("New(%q) returned Invalid", n)
		}
	}
}

func TestSiteStateType_NewInvalid(t *testing.T) {
	v, err := New("__bogus__")
	if err == nil || v != Invalid {
		t.Errorf("New bogus: v=%v err=%v", v, err)
	}
}

func TestSiteStateType_NewMust(t *testing.T) {
	if NewMust("NewlyAdded") != NewlyAdded {
		t.Error("NewMust(NewlyAdded) mismatch")
	}
}

func TestSiteStateType_Max(t *testing.T) {
	if Max() == Invalid {
		t.Error("Max() should not be Invalid")
	}
}

func TestSiteStateType_RangesInvalidErr(t *testing.T) {
	if err := RangesInvalidErr(); err != nil {
		t.Errorf("RangesInvalidErr unexpected: %v", err)
	}
}

func TestSiteStateType_Ranges(t *testing.T) {
	if Ranges[NewlyAdded] != "NewlyAdded" || Ranges[Removed] != "Removed" || Ranges[Unchanged] != "Unchanged" {
		t.Error("Ranges mapping mismatch")
	}
}
