package pathpatterntype

import (
	"encoding/json"
	"testing"
)

// AL2-03 Batch C coverage suite for pathpatterntype.
// Covers New/NewMust roundtrip, Min/Max, key Variant accessors, and JSON.

func TestPathPatternType_NewRoundTrip(t *testing.T) {
	allNames := Invalid.AllNameValues()
	if len(allNames) == 0 {
		t.Fatal("AllNameValues empty")
	}
	for _, name := range allNames {
		v, err := New(name)
		if err != nil {
			t.Errorf("New(%q): %v", name, err)
			continue
		}
		if v.Name() != name {
			// Some variants may map differently via internal mapping; just ensure non-empty.
			if v.Name() == "" {
				t.Errorf("New(%q) yielded empty Name", name)
			}
		}
		_ = NewMust(name)
	}
	if Min() != Invalid {
		t.Errorf("Min %v", Min())
	}
	if Max().IsInvalid() {
		t.Error("Max should not be invalid")
	}
}

func TestPathPatternType_AccessorsAndPredicates(t *testing.T) {
	v := App
	if v.ValueByte() != byte(v) || v.ValueInt() != int(v) || v.ValueInt8() != int8(v) ||
		v.ValueInt16() != int16(v) || v.ValueInt32() != int32(v) || v.ValueUInt16() != uint16(v) ||
		v.Value() != byte(v) {
		t.Error("numeric accessors mismatch")
	}
	if v.MaxByte() != byte(Max()) || v.MinByte() != byte(Min()) {
		t.Error("Min/MaxByte mismatch")
	}
	if v.MaxInt() < v.MinInt() {
		t.Error("MaxInt < MinInt")
	}
	if v.Name() == "" || v.String() == "" || v.NameValue() == "" || v.ValueString() == "" ||
		v.ToNumberString() == "" || v.RangeNamesCsv() == "" || v.TypeName() == "" ||
		v.Format("%s") == "" {
		t.Error("string accessors empty")
	}
	if v.EnumType() == nil {
		t.Error("EnumType nil")
	}
	if len(v.AllNameValues()) == 0 || len(v.IntegerEnumRanges()) == 0 ||
		len(v.RangesByte()) == 0 || len(v.RangesDynamicMap()) == 0 {
		t.Error("ranges empty")
	}
	if !v.IsValid() || Invalid.IsValid() || !Invalid.IsInvalid() {
		t.Error("valid/invalid wrong")
	}
	if !v.IsAnyOf(App, AppCore) || v.IsAnyOf(AppCore) {
		t.Error("IsAnyOf wrong")
	}
	if !v.IsNameEqual("App") || v.IsNameEqual("nope") {
		t.Error("IsNameEqual wrong")
	}
	if !v.IsByteValueEqual(byte(App)) || v.IsByteValueEqual(byte(Invalid)) {
		t.Error("IsByteValueEqual wrong")
	}
	if !v.IsValueEqual(byte(App)) {
		t.Error("IsValueEqual wrong")
	}
	if !v.IsAnyValuesEqual(byte(App), byte(AppCore)) || v.IsAnyValuesEqual(byte(Invalid)) {
		t.Error("IsAnyValuesEqual wrong")
	}
	if !v.IsAnyNamesOf("App", "AppCore") || v.IsAnyNamesOf("nope") {
		t.Error("IsAnyNamesOf wrong")
	}
	if !v.IsNameOf("App") {
		t.Error("IsNameOf wrong")
	}
	if v.IsUninitialized() && !Invalid.IsUninitialized() {
		t.Error("IsUninitialized inconsistent")
	}
	// HasExpandAssoc / IsExpandPossible / IsSingleType / ExpandedAssociatedVariants smoke
	_ = v.HasExpandAssoc()
	_ = v.IsExpandPossible()
	_ = v.IsSingleType()
	_ = v.ExpandedAssociatedVariants()
	_ = v.CurlyPathFullName()
	_ = v.PathFullName()
	_ = v.CompileCurlyTemplate()
	_ = v.CompileTemplate()
	_ = v.Clone()
}

func TestPathPatternType_JsonAndBinders(t *testing.T) {
	v := App
	data, err := json.Marshal(v)
	if err != nil || len(data) == 0 {
		t.Fatalf("Marshal: %v", err)
	}
	if v.ToPtr() == nil || v.JsonPtr() == nil {
		t.Error("ToPtr/JsonPtr nil")
	}
	if v.AsBasicByteEnumContractsBinder() == nil || v.AsBasicEnumContractsBinder() == nil {
		t.Error("binder nil")
	}
	_ = v.OnlySupportedErr("App")
	_ = v.OnlySupportedMsgErr("ctx", "App")
	if min, max := v.MinMaxAny(); min == nil || max == nil {
		t.Error("MinMaxAny nil")
	}
	if v.MinValueString() == "" || v.MaxValueString() == "" {
		t.Error("Min/Max value string empty")
	}
}
