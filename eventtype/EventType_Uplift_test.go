package eventtype

import (
	"encoding/json"
	"testing"

	"github.com/alimtvnetwork/core-v9/coredata/corejson"
)

// AL2 follow-up coverage uplift for eventtype using only the standard
// `testing` package. Exercises every Variant accessor, predicate, JSON
// path, and BasicEnumImpl-backed helper.
func TestEventType_UpliftAllSurface(t *testing.T) {
	all := []Variant{Invalid, Log, Success, Error, Failure, File, Custom}

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
		_ = v.MaxByte()
		_ = v.MinByte()
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

		// predicates
		_ = v.IsValid()
		_ = v.IsInvalid()
		_ = v.IsLog()
		_ = v.IsError()
		_ = v.IsSuccess()
		_ = v.IsFailure()
		_ = v.IsCustom()
		_ = v.IsFile()
		_ = v.IsSilent()
		_ = v.IsSkip()
		_ = v.IsInfo()
		_ = v.IsTrace()
		_ = v.IsDebug()
		_ = v.IsFatal()
		_ = v.IsPanic()
		_ = v.HasNoLog()
		_ = v.IsErrorLogical()
		_ = v.IsErrorFatalLogical()
		_ = v.IsErrorFatalPanicLogical()
		_ = v.IsErrorLogically()
		_ = v.IsLogLogically()
		_ = v.IsLogNameEqual(v.Name())
		_ = v.IsLogValueEqual(v.ValueByte())
		_ = v.IsValueEqual(v.ValueByte())
		_ = v.IsByteValueEqual(v.ValueByte())
		_ = v.IsNameEqual(v.Name())
		_ = v.IsAnyNamesOf("Log", "Error")
		_ = v.IsAnyValuesEqual(byte(Log), byte(Error))
		_ = v.IsAnyOf(Log, Error, Custom)
		_ = v.IsNameOf("Log", "Custom")

		// JSON round-trip
		data, err := json.Marshal(v)
		if err != nil {
			t.Fatalf("Marshal(%v): %v", v, err)
		}
		var got Variant
		if err := json.Unmarshal(data, &got); err != nil {
			t.Fatalf("Unmarshal(%s): %v", string(data), err)
		}

		// JsonParseSelfInject
		jr := corejson.New(v)
		var sink Variant
		_ = sink.JsonParseSelfInject(&jr)
	}

	// HasPattern panics — recover
	func() {
		defer func() { _ = recover() }()
		_ = Custom.HasPattern("x")
	}()

	// OnlySupported* error helpers
	_ = Log.OnlySupportedErr("Log", "Error")
	_ = Log.OnlySupportedMsgErr("ctx", "Log")

	// Make sure ErrorMap-driven predicates flip on Error/Failure
	if !Error.IsErrorLogical() || !Failure.IsErrorLogical() {
		t.Fatal("Error/Failure should be logical errors")
	}
	if Log.IsErrorLogical() {
		t.Fatal("Log should not be logical error")
	}
	if !Log.IsLogLogically() || Failure.IsLogLogically() {
		t.Fatal("IsLogLogically inverted")
	}

	// MinMaxAny non-nil
	if mn, mx := Custom.MinMaxAny(); mn == nil || mx == nil {
		t.Fatal("MinMaxAny nil")
	}

	// New / NewMust paths
	if _, err := New("Log"); err != nil {
		t.Fatalf("New(Log): %v", err)
	}
	if v := NewMust("Success"); v != Success {
		t.Fatalf("NewMust(Success) = %v", v)
	}
	if _, err := New("__nope__"); err == nil {
		t.Fatal("New(__nope__) should error")
	}

	// Range-invalid err
	if RangesInvalidErr() == nil {
		t.Fatal("RangesInvalidErr nil")
	}
}
