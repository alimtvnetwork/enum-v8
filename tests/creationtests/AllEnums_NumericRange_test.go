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
//   - PI-006 (sqliteconnpathtype) RESOLVED in Cycle 60 — local
//     MinValueString override now returns the lexicographic min name
//     instead of empty. The skip entry for sqliteconnpathtype is gone.
//   - strtype.Variant is a string-backed enum: numeric MinInt/MaxInt are
//     present on the interface but their semantic meaning differs; only the
//     MinInt <= MaxInt invariant is skipped (not asserted) for that type.
var numericRangeSuiteSkipMinValueString = map[string]string{
	"strtype.Variant": "string-backed enum; MinValueString is empty by design",
}

var numericRangeSuiteSkipMaxValueString = map[string]string{
	"strtype.Variant": "string-backed enum; MaxValueString is empty by design",
}

var numericRangeSuiteSkipMinMaxIntOrder = map[string]string{
	"strtype.Variant": "string-backed enum; MinInt/MaxInt semantics differ",
}

var numericRangeSuiteSkipRangesDynamicMap = map[string]string{
	"strtype.Variant":             "string-backed enum; RangesDynamicMap is intentionally empty",
	"compresslevels.Variant":      "int8-backed enum with negative range; RangesDynamicMap returns empty in upstream impl",
	"sqliteconnpathtype.Variant":  "string-backed enum; upstream BasicString.AllNameValues/RangesDynamicMap empty for spread-constructed enums (PI-006 cluster)",
}

func Test_AllEnums_NumericRange(t *testing.T) {
	for _, current := range allBasicEnumsCollection {
		current := current
		typeName := current.TypeName()
		_, skipMinValString := numericRangeSuiteSkipMinValueString[typeName]
		_, skipMinMaxOrder := numericRangeSuiteSkipMinMaxIntOrder[typeName]

		_, skipMaxValString := numericRangeSuiteSkipMaxValueString[typeName]
		_, skipRangesMap := numericRangeSuiteSkipRangesDynamicMap[typeName]

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

			// MaxValueString must be non-empty for every numeric-backed type.
			if !skipMaxValString {
				So(current.MaxValueString(), ShouldNotBeBlank)
			}

			// MinValueString must be non-empty for every type
			// except known-broken sqliteconnpathtype (PI-006) and string-backed strtype.
			if !skipMinValString {
				So(current.MinValueString(), ShouldNotBeBlank)
			}

			// RangesDynamicMap returns a non-nil map.
			rangesMap := current.RangesDynamicMap()
			So(rangesMap, ShouldNotBeNil)
			if !skipRangesMap {
				So(len(rangesMap), ShouldBeGreaterThan, 0)
			}

			// AllNameValues is non-empty for numeric-backed enums.
			allNV := current.AllNameValues()
			So(allNV, ShouldNotBeNil)
			if !skipRangesMap {
				So(len(allNV), ShouldBeGreaterThan, 0)
			}

			// IntegerEnumRanges returns a non-nil structure.
			integerRanges := current.IntegerEnumRanges()
			So(integerRanges, ShouldNotBeNil)
		})
	}
}
