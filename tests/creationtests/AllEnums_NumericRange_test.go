package creationtests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Test_AllEnums_NumericRange
//
// For every enum registered in allBasicEnumsCollection, exercise the
// numeric range / width surface guaranteed to exist on BasicEnumer:
//
//   - MinInt() <= MaxInt() (range is well-formed for numeric-backed enums).
//   - MinMaxAny() returns a (min, max) pair that is non-nil for every type.
//   - MinValueString() and MaxValueString() are produced without panic and
//     are non-empty for numeric-backed enums.
//   - RangesDynamicMap() returns a non-nil map covering at least one entry
//     for numeric-backed enums.
//   - AllNameValues() returns a non-empty slice.
//   - IntegerEnumRanges() returns a non-nil structure (where applicable).
//
// Skip notes:
//   - sqliteconnpathtype.Variant has known-broken MinValueString (PI-006);
//     non-empty MinValueString assertion is skipped for that type.
//   - strtype.Variant is a string-backed enum: numeric MinInt/MaxInt are
//     present on the interface but their semantic meaning differs; only the
//     MinInt <= MaxInt invariant is skipped (not asserted) for that type.
var numericRangeSuiteSkipMinValueString = map[string]string{
	"sqliteconnpathtype.Variant": "PI-006 — MinValueString returns empty",
}

var numericRangeSuiteSkipMinMaxIntOrder = map[string]string{
	"strtype.Variant": "string-backed enum; MinInt/MaxInt semantics differ",
}

func Test_AllEnums_NumericRange(t *testing.T) {
	for _, current := range allBasicEnumsCollection {
		current := current
		typeName := current.TypeName()
		_, skipMinValString := numericRangeSuiteSkipMinValueString[typeName]
		_, skipMinMaxOrder := numericRangeSuiteSkipMinMaxIntOrder[typeName]

		Convey(typeName+" — numeric range / width surface", t, func() {
			// MinInt <= MaxInt invariant for numeric-backed enums.
			if !skipMinMaxOrder {
				minI := current.MinInt()
				maxI := current.MaxInt()
				So(minI, ShouldBeLessThanOrEqualTo, maxI)
			}

			// MinMaxAny returns non-nil pair.
			minAny, maxAny := current.MinMaxAny()
			So(minAny, ShouldNotBeNil)
			So(maxAny, ShouldNotBeNil)

			// MaxValueString must be non-empty for every type.
			So(current.MaxValueString(), ShouldNotBeBlank)

			// MinValueString must be non-empty for every type
			// except known-broken sqliteconnpathtype (PI-006).
			if !skipMinValString {
				So(current.MinValueString(), ShouldNotBeBlank)
			}

			// RangesDynamicMap returns a non-nil map.
			rangesMap := current.RangesDynamicMap()
			So(rangesMap, ShouldNotBeNil)
			So(len(rangesMap), ShouldBeGreaterThan, 0)

			// AllNameValues is non-empty.
			allNV := current.AllNameValues()
			So(allNV, ShouldNotBeNil)
			So(len(allNV), ShouldBeGreaterThan, 0)

			// IntegerEnumRanges returns a non-nil structure.
			integerRanges := current.IntegerEnumRanges()
			So(integerRanges, ShouldNotBeNil)
		})
	}
}
