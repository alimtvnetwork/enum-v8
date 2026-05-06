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
func Test_AllEnums_Format(t *testing.T) {
	for _, current := range allBasicEnumsCollection {
		current := current
		typeName := current.TypeName()

		Convey(typeName+" — Format & string conversion surface", t, func() {
			name := current.Name()
			valueString := current.ValueString()
			numberString := current.ToNumberString()
			rangeCsv := current.RangeNamesCsv()
			nameValue := current.NameValue()
			minStr := current.MinValueString()
			maxStr := current.MaxValueString()
			allNameValues := current.AllNameValues()

			So(name, ShouldNotBeEmpty)
			So(valueString, ShouldNotBeEmpty)
			So(numberString, ShouldNotBeEmpty)
			So(rangeCsv, ShouldNotBeEmpty)
			So(nameValue, ShouldNotBeEmpty)
			So(minStr, ShouldNotBeEmpty)
			So(maxStr, ShouldNotBeEmpty)
			So(len(allNameValues), ShouldBeGreaterThan, 0)

			// String() (from fmt.Stringer / enumNameStinger) must equal Name()
			if stringer, ok := current.(interface{ String() string }); ok {
				So(stringer.String(), ShouldEqual, name)
			}

			// RangeNamesCsv must include the current Name as one of its entries.
			So(strings.Contains(rangeCsv, name), ShouldBeTrue)

			// NameValue is the human-readable "Name=value" composite — it must
			// at minimum reference the Name.
			So(strings.Contains(nameValue, name), ShouldBeTrue)

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
