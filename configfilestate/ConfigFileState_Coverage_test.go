package configfilestate

import "testing"

// AL2-01 Batch A coverage suite for configfilestate.
// Covers all IsX predicates + Min/Max + Variant logical helpers.
// No New constructor — Ranges + BasicEnumImpl driven.

func TestConfigFileState_AllVariantPredicates(t *testing.T) {
	all := []Variant{
		Invalid, Unchanged, Permission, Added, Modified, Deleted,
		SymbolicLinkAdded, SymbolicLinkDelete,
		ChmodChanged, ChownChanged, LastModifiedDateChanged, SizeChanged,
		ChmodChownBothChanged,
		ChmodOrChownOrLastModifiedDateChanged,
		SizeOrChmodOrChownOrLastModifiedDateChanged,
		MismatchFileOrDir,
	}

	// Spot-check unique predicates per variant.
	if !Invalid.IsUnknown() {
		t.Error("Invalid.IsUnknown() should be true")
	}
	if !Unchanged.IsUnchanged() {
		t.Error("Unchanged.IsUnchanged() should be true")
	}
	if !Added.IsAdded() {
		t.Error("Added.IsAdded() should be true")
	}
	if !Deleted.IsDeleted() {
		t.Error("Deleted.IsDeleted() should be true")
	}
	if !Permission.IsPermission() {
		t.Error("Permission.IsPermission() should be true")
	}
	if !Permission.IsUnknownOrPermission() || !Invalid.IsUnknownOrPermission() {
		t.Error("IsUnknownOrPermission failed")
	}
	if !Invalid.IsUnsafeCase() || !Permission.IsUnsafeCase() || !MismatchFileOrDir.IsUnsafeCase() {
		t.Error("IsUnsafeCase failed for unsafe variants")
	}
	if Unchanged.IsUnsafeCase() {
		t.Error("Unchanged.IsUnsafeCase() should be false")
	}
	if !Modified.IsModified() {
		t.Error("Modified.IsModified() should be true")
	}

	// HasChangeLogically: anything not Unchanged.
	if !Added.HasChangeLogically() || !Modified.HasChangeLogically() {
		t.Error("HasChangeLogically failed for Added/Modified")
	}
	if Unchanged.HasChangeLogically() {
		t.Error("Unchanged.HasChangeLogically() should be false")
	}
	if !Unchanged.HasNoChangeLogically() {
		t.Error("Unchanged.HasNoChangeLogically() should be true")
	}

	// Sanity: no panic when iterating all.
	for _, v := range all {
		_ = v.IsUnknown()
		_ = v.IsUnsafeCase()
		_ = v.HasChangeLogically()
	}
}

func TestConfigFileState_MinMaxRanges(t *testing.T) {
	if Min() != Invalid {
		t.Errorf("Min() = %v, want Invalid", Min())
	}
	max := Max()
	if max == Invalid {
		t.Errorf("Max() should not be Invalid, got %v", max)
	}
	if int(max) != len(Ranges)-1 {
		// Max returns last enumerated entry; verify it indexes into Ranges.
		if int(max) >= len(Ranges) {
			t.Errorf("Max() = %d out of range len=%d", max, len(Ranges))
		}
	}
}

// Pattern-7: AllNameValues coverage — exercise the accessor for coverage and
// verify it produces one entry per non-blank Ranges slot. AllNameValues() emits
// upstream's "Name(value)" format (e.g. "Invalid(0)"), so we count entries
// rather than checking raw-name presence.
func TestConfigFileState_AllNameValuesCoverage(t *testing.T) {
	names := Invalid.AllNameValues()
	if len(names) == 0 {
		t.Fatal("AllNameValues empty")
	}
	wantCount := 0
	for _, r := range Ranges {
		if r != "" {
			wantCount++
		}
	}
	if len(names) != wantCount {
		t.Errorf("AllNameValues count=%d, want %d (one per non-blank Ranges entry)", len(names), wantCount)
	}
}
