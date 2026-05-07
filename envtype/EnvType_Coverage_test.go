package envtype

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// AL2: Coverage uplift suite for envtype.
// Targets the surface methods that aren't exercised by the existing
// constructor test (predicates, value accessors, mapping helpers,
// JSON round-trip, range/format surface).
func Test_EnvType_Coverage(t *testing.T) {
	all := []Variant{
		Uninitialized,
		Development, Development1, Development2,
		Test, Test1, Test2,
		Production, Production1, Production2,
	}

	Convey("envtype — value/name accessors round-trip", t, func() {
		for _, v := range all {
			So(v.ValueByte(), ShouldEqual, byte(v))
			So(v.ValueInt(), ShouldEqual, int(v))
			So(v.ValueInt8(), ShouldEqual, int8(v))
			So(v.ValueInt16(), ShouldEqual, int16(v))
			So(v.ValueInt32(), ShouldEqual, int32(v))
			So(v.ValueUInt16(), ShouldEqual, uint16(v))
			So(v.Value(), ShouldEqual, byte(v))
			So(v.IsByteValueEqual(byte(v)), ShouldBeTrue)
			So(v.IsValueEqual(byte(v)), ShouldBeTrue)
			So(v.IsAnyValuesEqual(byte(v)), ShouldBeTrue)
			So(v.IsNameEqual(v.Name()), ShouldBeTrue)
			So(v.IsAnyOf(v), ShouldBeTrue)
			So(v.IsAnyOf(Uninitialized, Production), ShouldBeIn, []bool{true, false})
			So(v.IsNameOf(v.Name()), ShouldBeTrue)
			So(v.IsNameOf("nope"), ShouldBeFalse)
		}
	})

	Convey("envtype — env-class predicates", t, func() {
		So(Development.IsDevelopment(), ShouldBeTrue)
		So(Development1.IsDevelopment1(), ShouldBeTrue)
		So(Development2.IsDevelopment2(), ShouldBeTrue)
		So(Development.IsAnyDevelopment(), ShouldBeTrue)
		So(Development1.IsAnyDevelopment(), ShouldBeTrue)
		So(Development2.IsAnyDevelopment(), ShouldBeTrue)
		So(Test.IsTest(), ShouldBeTrue)
		So(Test1.IsTest1(), ShouldBeTrue)
		So(Test2.IsTest2(), ShouldBeTrue)
		So(Test.IsAnyTestEnv(), ShouldBeTrue)
		So(Production.IsProduction(), ShouldBeTrue)
		So(Production1.IsProduction1(), ShouldBeTrue)
		So(Production2.IsProduction2(), ShouldBeTrue)
		So(Production.IsAnyProduction(), ShouldBeTrue)
		So(Uninitialized.IsAnyProduction(), ShouldBeFalse)

		So(Test.IsTestEnvLogically(), ShouldBeTrue)
		So(Production.IsNotTestEnvLogically(), ShouldBeTrue)
		So(Development.IsDevEnvLogically(), ShouldBeTrue)
		So(Production.IsNotDevEnvLogically(), ShouldBeTrue)
		So(Production.IsProdEnvLogically(), ShouldBeTrue)
		So(Test.IsNotProdEnvLogically(), ShouldBeTrue)

		So(Uninitialized.IsUninitialized(), ShouldBeTrue)
		So(Development.IsInitialized(), ShouldBeTrue)
		So(Uninitialized.IsInvalid(), ShouldBeTrue)
		So(Development.IsValid(), ShouldBeTrue)
	})

	Convey("envtype — mapping helpers", t, func() {
		So(Development.KeyName(), ShouldEqual, "dev")
		So(Production1.KeyName(), ShouldEqual, "prod-1")
		So(Test.CurlyKeyName(), ShouldEqual, "{test}")
		So(Production2.CurlyKeyName(), ShouldEqual, "{prod-2}")
		So(Development1.VersionNumber(), ShouldEqual, 1)
		So(Production2.VersionNumber(), ShouldEqual, 2)
		So(Development.VersionNumber(), ShouldEqual, 0)
	})

	Convey("envtype — IsAnyNamesOf", t, func() {
		So(Development.IsAnyNamesOf("Development", "Production"), ShouldBeTrue)
		So(Development.IsAnyNamesOf("nope"), ShouldBeFalse)
	})

	Convey("envtype — JSON round-trip", t, func() {
		original := Production1
		bs, err := json.Marshal(&original)
		So(err, ShouldBeNil)
		var got Variant
		So(json.Unmarshal(bs, &got), ShouldBeNil)
		So(got, ShouldEqual, original)
	})

	Convey("envtype — range / format / type surface", t, func() {
		So(len(Development.RangesByte()), ShouldBeGreaterThan, 0)
		So(Development.RangeNamesCsv(), ShouldNotBeBlank)
		So(Development.TypeName(), ShouldNotBeBlank)
		So(len(Development.AllNameValues()), ShouldBeGreaterThan, 0)
		So(Development.MaxByte(), ShouldBeGreaterThanOrEqualTo, Development.MinByte())
		So(Development.MaxInt(), ShouldBeGreaterThanOrEqualTo, Development.MinInt())
		So(Development.MinValueString(), ShouldNotBeBlank)
		So(Development.MaxValueString(), ShouldNotBeBlank)
		So(Development.RangesDynamicMap(), ShouldNotBeNil)
		So(Development.IntegerEnumRanges(), ShouldNotBeNil)
		minA, maxA := Development.MinMaxAny()
		So(minA, ShouldNotBeNil)
		So(maxA, ShouldNotBeNil)
		So(Development.EnumType(), ShouldNotBeNil)
		So(Development.NameValue(), ShouldNotBeBlank)
		So(Development.String(), ShouldEqual, "Development")
		So(Development.ValueString(), ShouldNotBeBlank)
		So(Development.Format("{name}={value}"), ShouldNotBeBlank)
	})

	Convey("envtype — Is helper + error paths", t, func() {
		So(Is("Development", Development), ShouldBeTrue)
		So(Is("__nope__", Development), ShouldBeFalse)
		So(Is("Test", Production), ShouldBeFalse)

		// OnlySupportedErr returns non-nil whenever any variant outside the
		// supported list exists; here all 10 envtype variants exist, so the
		// "Development"-only call must report the rest as unsupported.
		err := Development.OnlySupportedErr("Development", "Test")
		So(err, ShouldNotBeNil)
		err = Development.OnlySupportedMsgErr("ctx", "Development")
		So(err, ShouldNotBeNil)
	})

	Convey("envtype — pointer-receiver bindings", t, func() {
		v := Production
		So(v.AsBasicByteEnumContractsBinder(), ShouldNotBeNil)
		So(v.AsBasicEnumContractsBinder(), ShouldNotBeNil)
		So(v.AsEnvironmentTyper(), ShouldNotBeNil)
	})
}
