package compresslevels

import (
	"compress/flate"
	"encoding/json"
	"testing"
)

// AL2-01 Batch A coverage suite for compresslevels.
// No New/NewMust constructors — covers Variant methods + flate mapping + JSON round-trip.

func TestCompressLevels_VariantPredicatesAndFlate(t *testing.T) {
	cases := []struct {
		v          Variant
		isDefault  bool
		isBest     bool
		isFast     bool
		isNoCompr  bool
		flateValue int8
	}{
		{Default, true, false, false, false, flate.DefaultCompression},
		{Best, false, true, false, false, flate.BestCompression},
		{Fast, false, false, true, false, flate.BestSpeed},
		{NoCompression, false, false, false, true, flate.NoCompression},
	}
	for _, c := range cases {
		if c.v.IsDefault() != c.isDefault ||
			c.v.IsBest() != c.isBest ||
			c.v.IsFast() != c.isFast ||
			c.v.IsNoCompression() != c.isNoCompr {
			t.Errorf("%v: predicate mismatch", c.v)
		}
		if c.v.Flate() != c.flateValue {
			t.Errorf("%v: Flate() = %d, want %d", c.v, c.v.Flate(), c.flateValue)
		}
		if c.v.Value() != int8(c.v) {
			t.Errorf("%v: Value() mismatch", c.v)
		}
		if c.v.ValueUInt16() != uint16(c.v) {
			t.Errorf("%v: ValueUInt16() mismatch", c.v)
		}
	}
}

func TestCompressLevels_RangesAndAccessors(t *testing.T) {
	v := Default
	if len(v.AllNameValues()) == 0 {
		t.Error("AllNameValues empty")
	}
	if v.MinInt() > v.MaxInt() {
		t.Errorf("MinInt %d > MaxInt %d", v.MinInt(), v.MaxInt())
	}
	if v.MinValueString() == "" || v.MaxValueString() == "" {
		t.Error("Min/Max value string empty")
	}
	min, max := v.MinMaxAny()
	if min == nil || max == nil {
		t.Error("MinMaxAny returned nil")
	}
	if len(v.IntegerEnumRanges()) == 0 {
		t.Error("IntegerEnumRanges empty")
	}
	if err := v.OnlySupportedErr("Default", "Best"); err != nil {
		t.Errorf("OnlySupportedErr unexpected: %v", err)
	}
	if err := v.OnlySupportedMsgErr("ctx", "Default"); err != nil {
		t.Errorf("OnlySupportedMsgErr unexpected: %v", err)
	}
}

func TestCompressLevels_JsonRoundTrip(t *testing.T) {
	for _, v := range []Variant{Default, Best, Fast, NoCompression} {
		data, err := json.Marshal(v)
		if err != nil {
			t.Fatalf("Marshal %v: %v", v, err)
		}
		// Variant has no UnmarshalJSON; test that BasicEnumImpl serialises non-empty.
		if len(data) == 0 {
			t.Errorf("%v: empty JSON", v)
		}
	}
}

func TestCompressLevels_FlateRangesMap(t *testing.T) {
	if flateRangesMap[flate.BestCompression] != Best {
		t.Errorf("flateRangesMap mapping wrong for BestCompression")
	}
	if flateRangesMap[flate.NoCompression] != NoCompression {
		t.Errorf("flateRangesMap mapping wrong for NoCompression")
	}
	if rangesMap[9] != Best || rangesMap[1] != Fast || rangesMap[0] != NoCompression || rangesMap[-1] != Default {
		t.Error("rangesMap mismatch")
	}
	if stringRangesMap["Default"] != Default {
		t.Error("stringRangesMap mismatch")
	}
}
