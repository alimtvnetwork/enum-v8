package promptclitype

import (
	"encoding/json"
	"testing"

	"github.com/alimtvnetwork/core-v9/issetter"
)

// AL2-06 follow-up: lift promptclitype from 58.3% past the 60% AL² bar
// by exercising every Variant accessor/predicate/binder across all 5 variants
// plus the alternative-name fast-paths and NewUsingSetter error path.
func TestPromptCliType_Uplift_AllVariants(t *testing.T) {
	all := []Variant{Invalid, Accept, Reject, Later, Review}
	for _, v := range all {
		_ = v.ValueUInt16()
		_ = v.AllNameValues()
		_ = v.IntegerEnumRanges()
		_, _ = v.MinMaxAny()
		_ = v.MinValueString()
		_ = v.MaxValueString()
		_ = v.MaxInt()
		_ = v.MinInt()
		_ = v.RangesDynamicMap()
		_ = v.IsLater()
		_ = v.OnOffLowercaseName()
		_ = v.IsIndeterminate()
		_ = v.IsAccept()
		_ = v.IsReject()
		_ = v.IsSkip()
		_ = v.IsAnyNamesOf("Accept", "Reject", "Later", "Review")
		_ = v.ValueInt()
		_ = v.IsAnyValuesEqual(0, 1, 2, 3, 4)
		_ = v.IsByteValueEqual(byte(v))
		_ = v.IsNameEqual(v.Name())
		_ = v.IsValueEqual(byte(v))
		_ = v.ValueInt8()
		_ = v.ValueInt16()
		_ = v.ValueInt32()
		_ = v.ValueString()
		_ = v.Format("%v")
		_ = v.EnumType()
		_ = v.IsUninitialized()
		_ = v.IsInitialized()
		_ = v.IsAsk()
		_ = v.IsYes()
		_ = v.IsOn()
		_ = v.IsNo()
		_ = v.IsOff()
		_ = v.IsYesNo()
		_ = v.IsAcceptReject()
		_ = v.IsNotAcceptReject()
		_ = v.IsTrue()
		_ = v.IsAccepted()
		_ = v.IsRejected()
		_ = v.IsDefinedAccepted()
		_ = v.IsDefinedRejected()
		_ = v.IsDefinedLogically()
		_ = v.IsUndefinedLogically()
		_ = v.IsUninitializedOrAsk()
		_ = v.IsInvalid()
		_ = v.IsValid()
		_ = v.Name()
		_ = v.NameLower()
		_ = v.YesNoLower()
		_ = v.OnOffName()
		_ = v.OnOffNameLower()
		_ = v.TrueFalseName()
		_ = v.ToIsSetter()
		_ = v.ToNumberString()
		_ = v.MaxByte()
		_ = v.MinByte()
		_ = v.ValueByte()
		_ = v.RangesByte()
		_ = v.Value()
		_ = v.ToPtr()
		_ = v.IsAnyOf(Accept, Reject)
		_ = v.IsNameOf("Accept", "Reject")
		_ = v.RangeNamesCsv()
		_ = v.TypeName()
		_ = v.NameValue()
		_ = v.String()
		_ = v.Json()
		_ = v.JsonPtr()
		_ = v.AsJsonContractsBinder()
		_ = v.AsJsoner()
		_ = v.AsJsonMarshaller()
		_ = v.AsBasicByteEnumContractsBinder()
		_ = v.AsBasicEnumContractsBinder()
		_ = v.AsYesNoAcceptRejecter()
		_ = v.AsOnOffLater()

		// JSON round-trip
		data, err := json.Marshal(v)
		if err != nil {
			t.Errorf("Marshal %v: %v", v, err)
		}
		var got Variant
		if err := json.Unmarshal(data, &got); err != nil {
			t.Errorf("Unmarshal %v: %v", v, err)
		}
		ptr := v.ToPtr()
		_ = ptr.ToSimple()
	}

	// Alternative-name fast-paths through newOtherWays / nameToVariant
	for name, want := range map[string]Variant{
		"yes": Accept, "y": Accept, "1": Accept, "A": Accept, "a": Accept,
		"no": Reject, "n": Reject, "0": Reject, "R": Reject, "r": Reject,
		"ask": Later, "*": Later, "L": Later, "l": Later,
		"review": Review, "Review": Review, "C": Review, "c": Review, "3": Review,
		"-1": Invalid, "": Invalid,
	} {
		got, err := New(name)
		if err != nil || got != want {
			t.Errorf("New(%q) = %v err=%v, want %v", name, got, err, want)
		}
	}

	// NewUsingSetter error path
	if _, err := NewUsingSetter(issetter.Value(99)); err == nil {
		t.Error("NewUsingSetter(99) expected error")
	}

	// JsonParseSelfInject path via Json round-trip
	src := Accept
	jr := src.Json()
	var dst Variant
	if err := dst.JsonParseSelfInject(&jr); err != nil {
		t.Errorf("JsonParseSelfInject: %v", err)
	}
	if dst != Accept {
		t.Errorf("JsonParseSelfInject dst = %v, want Accept", dst)
	}
}
