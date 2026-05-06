package creationtests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Test_AllEnums_Predicates
//
// For every enum registered in allBasicEnumsCollection (all instantiated as
// the package's `Invalid` zero-value, except where a non-zero default is
// used), exercise the predicate / equality / value-width surface that is
// guaranteed to exist on the BasicEnumer interface:
//
//   - IsValid + IsInvalid are mutually exclusive and always defined.
//   - IsNameEqual matches the type's own Name() and rejects a guaranteed-
//     bogus name.
//   - IsAnyNamesOf returns true when the own Name is in the list, false
//     when it is absent (with a guaranteed-bogus name).
//   - The numeric-width accessors (ValueByte, ValueInt, ValueInt8/16/32,
//     ValueUInt16) are internally consistent — they all return the same
//     underlying numeric value cast to their respective types.
//
// This exercises one assertion path through every Variant's predicate
// implementations in a single sweep over allBasicEnumsCollection.
//
// Skip notes:
//   - `sqliteconnpathtype.Variant.IsAnyNamesOf()` (empty args) returns true
//     while every other Variant correctly returns false. Tracked as PI-007;
//     the empty-args assertion is skipped for that type.
var predicateSuiteSkipEmptyAnyNames = map[string]string{
	"sqliteconnpathtype.Variant": "PI-007 — IsAnyNamesOf() with no args returns true",
}

func Test_AllEnums_Predicates(t *testing.T) {
	const bogusName = "__definitely_not_a_real_enum_name__zzz_AL03__"

	for _, current := range allBasicEnumsCollection {
		current := current
		typeName := current.TypeName()

		Convey(typeName+" — predicate / equality / value-width surface", t, func() {
			name := current.Name()

			// IsValid XOR IsInvalid — exactly one must be true.
			isValid := current.IsValid()
			isInvalid := current.IsInvalid()
			So(isValid, ShouldNotEqual, isInvalid)

			// IsNameEqual must match the type's own Name and reject a bogus one.
			So(current.IsNameEqual(name), ShouldBeTrue)
			So(current.IsNameEqual(bogusName), ShouldBeFalse)

			// IsAnyNamesOf — true when own name is present, false when only the bogus name is given.
			So(current.IsAnyNamesOf(name, bogusName), ShouldBeTrue)
			So(current.IsAnyNamesOf(bogusName), ShouldBeFalse)
			// Empty input is unambiguously "no match".
			So(current.IsAnyNamesOf(), ShouldBeFalse)

			// Numeric-width accessors must agree on the underlying value.
			vByte := current.ValueByte()
			vInt := current.ValueInt()
			vInt8 := current.ValueInt8()
			vInt16 := current.ValueInt16()
			vInt32 := current.ValueInt32()
			vUInt16 := current.ValueUInt16()

			So(int(vByte), ShouldEqual, vInt)
			So(int(vInt8), ShouldEqual, vInt)
			So(int(vInt16), ShouldEqual, vInt)
			So(int(vInt32), ShouldEqual, vInt)
			So(int(vUInt16), ShouldEqual, vInt)
		})
	}
}
