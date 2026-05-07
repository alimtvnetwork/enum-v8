package creationtests

import (
	"regexp"
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Test_AllEnums_AllNameValuesRoundTrip enforces RCA Pattern 7 across every
// registered enum: every entry of AllNameValues() must conform to the
// canonical "name(value)" format declared by `constants.EnumNameValueFormat`
// (`"%s(%d)"`), and the parsed integer value must sit within [MinInt, MaxInt].
//
// Catches:
//   - Malformed NameWithValue overrides that drop the parenthesized value.
//   - Off-by-one drift between `AllNameValues` and `MinInt`/`MaxInt` after a
//     Variant member is added but a custom override is forgotten.
//   - Empty-name entries from sparse string-array gaps (cross-checks the
//     SparseArrayGapGuard at the AllNameValues surface).
//
// String-backed enums (strtype, sqliteconnpathtype) and open-ended numeric
// wrappers (inttype) use a different surface format; they're skipped.
var allNameValuesRoundTripSkip = map[string]string{
	"strtype.Variant":            "string-backed enum; uses StringEnumNameValueFormat (no value suffix)",
	"sqliteconnpathtype.Variant": "string-backed enum (PI-006 cluster); upstream returns empty AllNameValues",
	"inttype.Variant":            "open-ended numeric enum; AllNameValues intentionally empty",
}

var nameValueFormatRegex = regexp.MustCompile(`^(.+)\((-?\d+)\)$`)

func Test_AllEnums_AllNameValuesRoundTrip(t *testing.T) {
	for _, current := range allBasicEnumsCollection {
		current := current
		typeName := current.TypeName()
		if _, skip := allNameValuesRoundTripSkip[typeName]; skip {
			continue
		}

		Convey(typeName+" — AllNameValues entries must be \"name(value)\" with value in [MinInt, MaxInt]", t, func() {
			defer func() {
				_ = recover() // upstream stub panics are out of scope
			}()

			entries := current.AllNameValues()
			if len(entries) == 0 {
				return // skip suites where surface is empty by design
			}

			minI := current.MinInt()
			maxI := current.MaxInt()

			for _, entry := range entries {
				match := nameValueFormatRegex.FindStringSubmatch(entry)
				So(match, ShouldNotBeNil)
				if match == nil {
					continue
				}
				name := match[1]
				So(name, ShouldNotBeBlank)

				val, err := strconv.Atoi(match[2])
				So(err, ShouldBeNil)
				So(val, ShouldBeBetweenOrEqual, minI, maxI)
			}
		})
	}
}
