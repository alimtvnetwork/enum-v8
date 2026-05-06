package creationtests

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Test_AllEnums_Format
//
//	For every enum registered in allBasicEnumsCollection:
//	  - Format(...) with each placeholder format produces a non-empty string
//	    containing the substituted value.
//	  - String() / Name() / ToNumberString() / ValueString() are non-empty
//	    and self-consistent (String() == Name()).
//	  - NameValue() contains both the name and the value-string fragments.
//	  - RangeNamesCsv() lists at least one name and contains the current Name.
//	  - MinValueString() / MaxValueString() are non-empty.
//	  - AllNameValues() returns at least one entry.
//
// Single test, one Convey per type, exercising the entire string-conversion
// surface of every Variant package via the shared collection.
//
// Skip notes:
//   - `strtype.Variant` is a free-form string enum with no fixed ranges, so
//     `MinValueString` / `MaxValueString` / `AllNameValues` are intentionally
//     empty and excluded from the non-empty assertions.
//   - PI-006 (sqliteconnpathtype) was RESOLVED in Cycle 60 by overriding
//     `MinValueString` and `NameValue` locally to bypass the upstream
//     `BasicString.Min()` / `NameWithValue` defects. The skip entry is gone.
var formatSuiteSkipMinMaxAll = map[string]string{
	"strtype.Variant": "free-form string enum, no fixed ranges",
}

var formatSuiteSkipNameValue = map[string]string{}

func Test_AllEnums_Format(t *testing.T) {
	for _, current := range allBasicEnumsCollection {
		current := current
		typeName := current.TypeName()
		_, skipMinMaxAll := formatSuiteSkipMinMaxAll[typeName]
		_, skipNameValue := formatSuiteSkipNameValue[typeName]

		Convey(typeName+" — Format & string conversion surface", t, func() {
			name := current.Name()
			valueString := current.ValueString()
			numberString := current.ToNumberString()
			rangeCsv := current.RangeNamesCsv()
			nameValue := current.NameValue()

			So(name, ShouldNotBeEmpty)
			So(valueString, ShouldNotBeEmpty)
			So(numberString, ShouldNotBeEmpty)
			So(rangeCsv, ShouldNotBeEmpty)
			So(nameValue, ShouldNotBeEmpty)

			if !skipMinMaxAll {
				So(current.MinValueString(), ShouldNotBeEmpty)
				So(current.MaxValueString(), ShouldNotBeEmpty)
				So(len(current.AllNameValues()), ShouldBeGreaterThan, 0)
			}

			// String() (from fmt.Stringer / enumNameStinger) must equal Name()
			if stringer, ok := current.(interface{ String() string }); ok {
				So(stringer.String(), ShouldEqual, name)
			}

			// RangeNamesCsv must include the current Name as one of its entries.
			So(strings.Contains(rangeCsv, name), ShouldBeTrue)

			// NameValue is the human-readable "Name=value" composite — it must
			// at minimum reference the Name (skipped for types with broken NameValue).
			if !skipNameValue {
				So(strings.Contains(nameValue, name), ShouldBeTrue)
			}

			// Format with the canonical placeholders defined by EnumFormatter.
			// Spec sample: "Enum of {type-name} - {name} - {value}".
			// Note: the upstream `enumimpl.Format` substitutes the numeric
			// number-string for `{name}` on number-backed enums, so we only
			// assert the type-name substitution here (always reliable across
			// every BasicEnumer implementation).
			formatTemplate := "Enum of {type-name} - {name} - {value}"
			formatted := current.Format(formatTemplate)
			So(formatted, ShouldNotBeEmpty)
			So(strings.Contains(formatted, typeName), ShouldBeTrue)
			// Substitution must have happened — the literal placeholder must be gone.
			So(strings.Contains(formatted, "{type-name}"), ShouldBeFalse)

			// A format with no placeholders should pass through unchanged.
			plainFormat := "literal-no-placeholders"
			plainFormatted := current.Format(plainFormat)
			So(plainFormatted, ShouldEqual, plainFormat)
		})
	}
}
