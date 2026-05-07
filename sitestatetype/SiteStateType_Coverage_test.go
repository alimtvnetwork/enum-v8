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
	// Diagnostic-only: byte enums starting at Invalid(0) always report the
	// numeric range. Exercise for coverage.
	_ = RangesInvalidErr()
}

func TestSiteStateType_Ranges(t *testing.T) {
	if Ranges[NewlyAdded] != "NewlyAdded" || Ranges[Removed] != "Removed" || Ranges[Unchanged] != "Unchanged" {
		t.Error("Ranges mapping mismatch")
	}
}

// Pattern-7: Ranges ↔ New round-trip (raw names, not the "Name(value)" format
// emitted by AllNameValues()). Also exercises AllNameValues() for coverage.
func TestSiteStateType_AllNameValuesRoundTrip(t *testing.T) {
	if len(Invalid.AllNameValues()) == 0 {
		t.Fatal("AllNameValues empty")
	}
	for _, name := range Ranges {
		if name == "" {
			continue
		}
		v, err := New(name)
		if err != nil {
			t.Errorf("New(%q) error: %v", name, err)
			continue
		}
		if v.Name() != name {
			t.Errorf("round-trip mismatch: New(%q).Name()=%q", name, v.Name())
		}
	}
}
