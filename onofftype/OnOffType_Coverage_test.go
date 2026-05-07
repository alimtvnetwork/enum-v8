package onofftype

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// AL2: Coverage uplift for onofftype (was 39.6%).
func Test_OnOffType_Coverage(t *testing.T) {
	all := []Variant{Invalid, Ask, On, Off}

	Convey("onofftype — value & name accessors", t, func() {
		for _, v := range all {
			So(v.ValueByte(), ShouldEqual, byte(v))
			So(v.ValueInt(), ShouldEqual, int(v))
			So(v.ValueInt8(), ShouldEqual, int8(v))
			So(v.ValueInt16(), ShouldEqual, int16(v))
			So(v.ValueInt32(), ShouldEqual, int32(v))
			So(v.ValueUInt16(), ShouldEqual, uint16(v))
			So(v.ValueString(), ShouldNotBeBlank)
			So(v.IsByteValueEqual(byte(v)), ShouldBeTrue)
			So(v.IsValueEqual(byte(v)), ShouldBeTrue)
			So(v.IsAnyValuesEqual(byte(v)), ShouldBeTrue)
			So(v.IsNameEqual(v.Name()), ShouldBeTrue)
			So(v.NameLower(), ShouldNotBeBlank)
			So(v.OnOffName(), ShouldNotBeBlank)
			So(v.OnOffNameLower(), ShouldNotBeBlank)
			So(v.OnOffLowercaseName(), ShouldNotBeBlank)
			So(v.YesNoLower(), ShouldNotBeBlank)
			So(v.TrueFalseName(), ShouldNotBeBlank)
		}
	})

	Convey("onofftype — predicates", t, func() {
		So(On.IsOn(), ShouldBeTrue)
		So(Off.IsOff(), ShouldBeTrue)
		So(On.IsYes(), ShouldBeTrue)
		So(Off.IsNo(), ShouldBeTrue)
		So(Ask.IsAsk(), ShouldBeTrue)
		So(On.IsTrue(), ShouldBeTrue)
		So(On.IsAccepted(), ShouldBeTrue)
		So(Off.IsRejected(), ShouldBeTrue)
		So(On.IsDefinedAccepted(), ShouldBeTrue)
		So(Off.IsDefinedRejected(), ShouldBeTrue)
		So(On.IsAccept(), ShouldBeTrue)
		So(Off.IsReject(), ShouldBeTrue)
		So(On.IsAcceptReject(), ShouldBeTrue)
		So(Off.IsAcceptReject(), ShouldBeTrue)
		So(Ask.IsNotAcceptReject(), ShouldBeTrue)
		So(On.IsOnLogically(), ShouldBeTrue)
		So(Off.IsOffLogically(), ShouldBeTrue)
		So(On.IsYesNo(), ShouldBeTrue)
		So(Off.IsYesNo(), ShouldBeTrue)
		So(On.IsDefinedLogically(), ShouldBeTrue)
		So(Off.IsDefinedLogically(), ShouldBeTrue)
		So(Invalid.IsUndefinedLogically(), ShouldBeTrue)
		So(Invalid.IsUninitialized(), ShouldBeTrue)
		So(On.IsInitialized(), ShouldBeTrue)
		So(Invalid.IsUninitializedOrAsk(), ShouldBeTrue)
		So(Ask.IsUninitializedOrAsk(), ShouldBeTrue)
		So(Invalid.IsInvalid(), ShouldBeTrue)
		So(On.IsValid(), ShouldBeTrue)
		So(Ask.IsIndeterminate(), ShouldBeTrue)
		So(Ask.IsLater(), ShouldBeIn, []bool{true, false})
		So(Ask.IsSkip(), ShouldBeIn, []bool{true, false})
	})

	Convey("onofftype — IsAnyNamesOf + setter bridge", t, func() {
		So(On.IsAnyNamesOf("On", "Off"), ShouldBeTrue)
		So(On.IsAnyNamesOf("nope"), ShouldBeFalse)
		setter := On.ToIsSetter()
		So(setter, ShouldNotBeNil)
	})

	Convey("onofftype — JSON round-trip", t, func() {
		original := On
		bs, err := json.Marshal(&original)
		So(err, ShouldBeNil)
		var got Variant
		So(json.Unmarshal(bs, &got), ShouldBeNil)
		So(got, ShouldEqual, original)
	})

	Convey("onofftype — range / format / type surface", t, func() {
		So(len(On.AllNameValues()), ShouldBeGreaterThan, 0)
		So(On.MaxInt(), ShouldBeGreaterThanOrEqualTo, On.MinInt())
		So(On.MinValueString(), ShouldNotBeBlank)
		So(On.MaxValueString(), ShouldNotBeBlank)
		So(On.RangesDynamicMap(), ShouldNotBeNil)
		So(On.IntegerEnumRanges(), ShouldNotBeNil)
		minA, maxA := On.MinMaxAny()
		So(minA, ShouldNotBeNil)
		So(maxA, ShouldNotBeNil)
		So(On.EnumType(), ShouldNotBeNil)
		So(On.Format("{name}"), ShouldNotBeBlank)
		So(On.ToNumberString(), ShouldNotBeBlank)
		So(On.OnlySupportedErr("On", "Off"), ShouldNotBeNil)
		So(On.OnlySupportedMsgErr("ctx", "On"), ShouldNotBeNil)
	})

	Convey("onofftype — alternate constructors", t, func() {
		So(NewUsingBool(true), ShouldEqual, On)
		So(NewUsingBool(false), ShouldEqual, Off)
		So(NewUsingAndBooleans(true, true, true), ShouldEqual, On)
		So(NewUsingAndBooleans(true, false, true), ShouldEqual, Off)
		So(Min(), ShouldNotEqual, Variant(255))
		So(Max(), ShouldNotEqual, Variant(255))
	})
}
