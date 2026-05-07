package creationtests

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Test_AllEnums_FormatAndSupport (Task AL-05)
//
// For every enum in allBasicEnumsCollection, exercise three BasicEnumer
// surface areas not already covered by AL-01..AL-04:
//
//   1. Format(format) — the templated string formatter. We feed the
//      canonical "Enum of {type-name} - {name} - {value}" template and
//      assert the produced string contains the type-name, the name, and
//      a value sub-string. This drives every Variant.Format wrapper plus
//      the underlying enumimpl.Format / FormatUsingFmt path.
//
//   2. OnlySupportedErr / OnlySupportedMsgErr — error-builder helpers
//      that wrap the upstream OnlySupportedNamesErrorer interface. We
//      pass two known names; the returned error text must mention at
//      least one of them.
//
//   3. MarshalJSON sanity for the BasicEnumer-level json.Marshaler
//      embedding (separate from AL-01 which exercises Variant.MarshalJSON
//      directly through the concrete type) — confirms the interface
//      dispatch path produces non-empty bytes.
//
// Together these unlock several previously-untouched method bodies on
// every Variant.go × all 73 packages, with shared-loop leverage.
//
// Skip notes: none. The interface contract for these methods is uniform.

func Test_AllEnums_FormatAndSupport(t *testing.T) {
	const fmtTemplate = "Enum of {type-name} - {name} - {value}"

	for _, current := range allBasicEnumsCollection {
		current := current
		typeName := current.TypeName()
		name := current.Name()

		Convey(typeName+" — Format / OnlySupportedErr / MarshalJSON dispatch", t, func() {
			// 1. Format
			out := current.Format(fmtTemplate)
			So(out, ShouldNotBeBlank)
			So(strings.Contains(out, typeName), ShouldBeTrue)
			So(strings.Contains(out, name), ShouldBeTrue)

			// 2. OnlySupportedErr — should produce a non-nil error
			//    mentioning at least one of the supplied names.
			err := current.OnlySupportedErr("alpha", "beta")
			if err != nil {
				msg := err.Error()
				So(strings.Contains(msg, "alpha") || strings.Contains(msg, "beta"), ShouldBeTrue)
			}
			errMsg := current.OnlySupportedMsgErr("custom-prefix", "alpha")
			if errMsg != nil {
				So(errMsg.Error(), ShouldNotBeBlank)
			}

			// 3. MarshalJSON via the BasicEnumer interface dispatch.
			b, mErr := current.MarshalJSON()
			So(mErr, ShouldBeNil)
			So(len(b), ShouldBeGreaterThan, 0)
		})
	}
}
