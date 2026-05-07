package dbexposetype

import (
	"encoding/json"
	"testing"
)

// AL2-02 Batch B coverage suite for dbexposetype.

func TestDbExposeType_NewRoundTrip(t *testing.T) {
	for _, n := range []string{"Invalid", "AnyIp", "SpecificIp"} {
		v, err := New(n)
		if err != nil {
			t.Errorf("New(%q): %v", n, err)
		}
		if v.Name() != n {
			t.Errorf("Name() = %q, want %q", v.Name(), n)
		}
	}
}

func TestDbExposeType_NewInvalid(t *testing.T) {
	if v, err := New("__bogus__"); err == nil || v != Invalid {
		t.Errorf("New bogus: v=%v err=%v", v, err)
	}
}

func TestDbExposeType_NewMust(t *testing.T) {
	if NewMust("AnyIp") != AnyIp {
		t.Error("NewMust mismatch")
	}
}

func TestDbExposeType_Predicates(t *testing.T) {
	if !AnyIp.IsAnyIP() || !SpecificIp.IsSpecificIP() || !Invalid.IsInvalid() {
		t.Error("predicate mismatch")
	}
	if !AnyIp.IsValid() || Invalid.IsValid() {
		t.Error("IsValid mismatch")
	}
	if !AnyIp.IsAnyOf(AnyIp, SpecificIp) || AnyIp.IsAnyOf(SpecificIp) {
		t.Error("IsAnyOf mismatch")
	}
	if !AnyIp.IsNameOf("AnyIp") || AnyIp.IsNameOf("zzz") {
		t.Error("IsNameOf mismatch")
	}
	if !AnyIp.IsAnyNamesOf("AnyIp", "SpecificIp") {
		t.Error("IsAnyNamesOf mismatch")
	}
	if !AnyIp.IsAnyValuesEqual(byte(AnyIp)) || !AnyIp.IsByteValueEqual(byte(AnyIp)) ||
		!AnyIp.IsValueEqual(byte(AnyIp)) || !AnyIp.IsNameEqual("AnyIp") {
		t.Error("value-equal mismatch")
	}
	if !Invalid.IsUninitialized() {
		t.Error("IsUninitialized mismatch")
	}
}

func TestDbExposeType_NumericAccessors(t *testing.T) {
	v := AnyIp
	_ = v.ValueInt()
	_ = v.ValueInt8()
	_ = v.ValueInt16()
	_ = v.ValueInt32()
	_ = v.ValueUInt16()
	_ = v.ValueByte()
	_ = v.Value()
	_ = v.MinByte()
	_ = v.MaxByte()
	if v.MinInt() > v.MaxInt() {
		t.Error("MinInt > MaxInt")
	}
	if v.MinValueString() == "" || v.MaxValueString() == "" {
		t.Error("min/max value string empty")
	}
	min, max := v.MinMaxAny()
	if min == nil || max == nil {
		t.Error("MinMaxAny nil")
	}
	if len(v.RangesByte()) == 0 || len(v.IntegerEnumRanges()) == 0 ||
		len(v.AllNameValues()) == 0 || len(v.RangesDynamicMap()) == 0 {
		t.Error("ranges empty")
	}
	if v.RangeNamesCsv() == "" || v.NameValue() == "" || v.TypeName() == "" {
		t.Error("string accessor empty")
	}
}

func TestDbExposeType_JsonRoundTrip(t *testing.T) {
	for _, x := range []Variant{Invalid, AnyIp, SpecificIp} {
		data, err := json.Marshal(x)
		if err != nil {
			t.Fatalf("Marshal: %v", err)
		}
		var got Variant
		if err := json.Unmarshal(data, &got); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if got != x {
			t.Errorf("round-trip: got %v want %v", got, x)
		}
	}
}

func TestDbExposeType_FormatAndPtrAndBinders(t *testing.T) {
	v := SpecificIp
	if v.Format("v=%v") == "" || v.String() == "" || v.ValueString() == "" || v.ToNumberString() == "" {
		t.Error("string method empty")
	}
	if v.ToPtr() == nil || (*v.ToPtr()).ToSimple() != v {
		t.Error("ToPtr mismatch")
	}
	var nilP *Variant
	if nilP.ToSimple() != Invalid {
		t.Error("ToSimple(nil) should be Invalid")
	}
	if v.JsonPtr().HasError() {
		t.Errorf("Json error: %v", v.JsonPtr().Error)
	}
	if v.JsonPtr() == nil {
		t.Error("JsonPtr nil")
	}
	if v.AsJsoner() == nil || v.AsJsonContractsBinder() == nil ||
		v.AsJsonMarshaller() == nil || v.AsBasicByteEnumContractsBinder() == nil ||
		v.AsBasicEnumContractsBinder() == nil {
		t.Error("As* binder nil")
	}
	if v.EnumType() == nil {
		t.Error("EnumType nil")
	}
	// Diagnostic descriptors — informational, always non-nil errors.
	_ = v.OnlySupportedErr("AnyIp")
	_ = v.OnlySupportedMsgErr("ctx", "AnyIp")
	_ = RangesInvalidErr()
	_ = Min()
	_ = Max()
}
