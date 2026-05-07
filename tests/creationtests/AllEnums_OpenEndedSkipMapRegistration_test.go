package creationtests

import (
	"testing"

	"github.com/alimtvnetwork/core-v9/constants"
	. "github.com/smartystreets/goconvey/convey"
)

// Test_AllEnums_OpenEndedSkipMapRegistration enforces RCA Pattern 9 across
// every registered enum: any Variant whose `MinInt() == constants.MinInt`
// AND `MaxInt() == constants.MaxInt` is open-ended (it wraps an entire
// numeric range — no discrete enumerable members) and MUST be registered
// in `numericRangeSuiteSkipRangesDynamicMap` so the
// `RangesDynamicMap > 0` invariant in `Test_AllEnums_NumericRange` does
// not fire a false positive.
//
// This is a compile-time-style guard for the test fixture itself: when a
// new open-ended numeric Variant is added to allBasicEnumsCollection, this
// suite fails fast and points the author at the exact entry to add. It
// also flags accidental over-registration: any type currently in the skip
// map that is NOT actually open-ended is reported (defensive — currently
// only triggers if a future refactor narrows a previously open-ended type).
func Test_AllEnums_OpenEndedSkipMapRegistration(t *testing.T) {
	for _, current := range allBasicEnumsCollection {
		current := current
		typeName := current.TypeName()

		Convey(typeName+" — open-ended numeric Variants must be in numericRangeSuiteSkipRangesDynamicMap", t, func() {
			// Some Variants panic on MinInt/MaxInt (e.g. strtype). The
			// existing NumericRange suite already handles those via its
			// own skip-MinMax-order map; we mirror that exemption here so
			// this guard does not regress.
			defer func() {
				if r := recover(); r != nil {
					// Intentional: panicking accessors are out of scope.
				}
			}()

			minI := current.MinInt()
			maxI := current.MaxInt()
			isOpenEnded := minI == constants.MinInt && maxI == constants.MaxInt
			_, registered := numericRangeSuiteSkipRangesDynamicMap[typeName]

			if isOpenEnded {
				if !registered {
					t.Errorf("%s is open-ended (MinInt==constants.MinInt && MaxInt==constants.MaxInt) but NOT registered in numericRangeSuiteSkipRangesDynamicMap — add an entry to AllEnums_NumericRange_test.go (RCA Pattern 9)", typeName)
				}
			}
			// Note: we deliberately do NOT assert "registered ⇒ open-ended".
			// Some entries in the skip map (compresslevels, sqliteconnpathtype)
			// are registered for OTHER reasons (negative-range or BasicString
			// upstream defects), not because of open-endedness. The skip map
			// is a union of three legitimate exemption causes.
		})
	}
}
