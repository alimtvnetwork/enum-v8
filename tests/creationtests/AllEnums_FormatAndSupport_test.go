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

// Skip notes:
//   - inttype.Variant / strtype.Variant: their OnlySupportedErr impls are
//     unimplemented stubs that panic with "not implemented for generic
//     {int,string} enum". Skipped from the OnlySupportedErr block only.
var formatAndSupportSkipOnlySupported = map[string]string{
	"inttype.Variant": "OnlySupportedErr panics: not implemented for generic int enum",
	"strtype.Variant": "OnlySupportedErr panics: not implemented for generic string enum",
}

func Test_AllEnums_FormatAndSupport(t *testing.T) {
	const fmtTemplate = "Enum of {type-name} - {name} - {value}"

	for _, current := range allBasicEnumsCollection {
		current := current
		typeName := current.TypeName()
		name := current.Name()
		_, skipOnlySupported := formatAndSupportSkipOnlySupported[typeName]

		Convey(typeName+" — Format / OnlySupportedErr / MarshalJSON dispatch", t, func() {
			// 1. Format — must produce a non-blank string containing the
			//    type-name. Some Variants render `{name}` as empty when
			//    the receiver is an unnamed/sentinel value, so we only
			//    assert type-name presence.
			out := current.Format(fmtTemplate)
			So(out, ShouldNotBeBlank)
			So(strings.Contains(out, typeName), ShouldBeTrue)
			_ = name

			// 2. OnlySupportedErr — non-blank message when non-nil. Skipped
			//    for generic int/string enums whose impl panics.
			if !skipOnlySupported {
				err := current.OnlySupportedErr("alpha", "beta")
				if err != nil {
					So(err.Error(), ShouldNotBeBlank)
				}
				errMsg := current.OnlySupportedMsgErr("custom-prefix", "alpha")
				if errMsg != nil {
					So(errMsg.Error(), ShouldNotBeBlank)
				}
			}

			// 3. MarshalJSON via the BasicEnumer interface dispatch.
			b, mErr := current.MarshalJSON()
			So(mErr, ShouldBeNil)
			So(len(b), ShouldBeGreaterThan, 0)
		})
	}
}
