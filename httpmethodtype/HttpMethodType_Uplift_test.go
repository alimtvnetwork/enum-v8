package httpmethodtype

import (
	"encoding/json"
	"testing"

	"github.com/alimtvnetwork/core-v9/coredata/corejson"
)

// AL2 follow-up coverage uplift for httpmethodtype using only the standard
// `testing` package.
func TestHttpMethodType_UpliftAllSurface(t *testing.T) {
	all := []Variant{Get, Post, Put, Patch, Delete, Head, Options, Invalid}

	for _, v := range all {
		_ = v.Value()
		_ = v.ValueByte()
		_ = v.ValueInt()
		_ = v.ValueInt8()
		_ = v.ValueInt16()
		_ = v.ValueInt32()
		_ = v.ValueUInt16()
		_ = v.ValueString()
		_ = v.Name()
		_ = v.String()
		_ = v.NameValue()
		_ = v.ToNumberString()
		_ = v.RangeNamesCsv()
		_ = v.TypeName()
		_ = v.Format("%s")
		_ = v.MinByte()
		_ = v.MaxByte()
		_ = v.MaxInt()
		_ = v.MinInt()
		_ = v.MinValueString()
		_ = v.MaxValueString()
		_ = v.AllNameValues()
		_ = v.IntegerEnumRanges()
		_ = v.RangesDynamicMap()
		_ = v.RangesByte()
		_, _ = v.MinMaxAny()
		_ = v.EnumType()
		_ = v.ToPtr()

		_ = v.IsGet()
		_ = v.IsPost()
		_ = v.IsPut()
		_ = v.IsPatch()
		_ = v.IsDelete()
		_ = v.IsHead()
		_ = v.IsOptions()
		_ = v.IsBodyAllowed()
		_ = v.IsSafe()
		_ = v.IsIdempotent()
		_ = v.IsValid()
		_ = v.IsInvalid()
		_ = v.IsEqual(v)
		_ = v.IsValueEqual(v.ValueByte())
		_ = v.IsByteValueEqual(v.ValueByte())
		_ = v.IsNameEqual(v.Name())
		_ = v.IsAboveOrEqual(Get)
		_ = v.IsLowerOrEqual(Invalid)
		_ = v.IsAnyOf(Get, Post)
		_ = v.IsAnyValuesEqual(byte(Get), byte(Post))
		_ = v.IsAnyNamesOf("Get", "Post")

		_ = v.OnlySupportedErr("Get")
		_ = v.OnlySupportedMsgErr("ctx", "Get")

		// JSON round-trip
		data, err := json.Marshal(v)
		if err != nil {
			t.Fatalf("Marshal(%v): %v", v, err)
		}
		var got Variant
		if err := json.Unmarshal(data, &got); err != nil {
			t.Fatalf("Unmarshal(%s): %v", string(data), err)
		}

		// corejson helpers
		jr := v.Json()
		_ = v.JsonPtr()
		var sink Variant
		_ = sink.JsonParseSelfInject(&jr)

		_ = v.AsJsonContractsBinder()
		_ = v.AsJsoner()
		_ = v.AsJsonMarshaller()
		_ = v.AsBasicByteEnumContractsBinder()
		_ = v.AsBasicEnumContractsBinder()
	}

	// Min/Max constructor functions
	if Min() != Get {
		t.Fatalf("Min() = %v", Min())
	}
	if Max() != Options {
		t.Fatalf("Max() = %v", Max())
	}
	if RangesInvalidErr() == nil {
		t.Fatal("RangesInvalidErr nil")
	}
	if v := NewMust("Get"); v != Get {
		t.Fatalf("NewMust(Get) = %v", v)
	}

	// Bounce against corejson directly
	_ = corejson.New(Get)
}
