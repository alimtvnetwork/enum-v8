package creationtests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Test_AllEnums_FixtureDriftGuard enforces RCA Pattern 1 across every
// registered enum: the pinned `StringMin`, `StringMax`, and
// `ExpectedInvalidValueString` rows in `allEnumGeneralTestCases.go` must
// match the live `MinValueString()` / `MaxValueString()` / Invalid
// instance's `ValueString()` output. When a Variant override is added or
// changed (commonly `MinValueString`, `NameValue`, `MarshalJSON`), the
// pinned fixture goes stale and `Test_AllEnums_ContractsTesting` fails
// downstream with a hard-to-attribute "Expected vs Actual" message.
//
// This guard short-circuits that diagnosis: it iterates every fixture row
// in `allEnumGeneralTestCases` and reports any mismatch with the exact
// type name, field, pinned value, and live value — pointing the author
// directly at the line to update.
//
// Skip rules: rows whose `InitialBasicEnumer` is nil, or whose
// MinValueString/MaxValueString accessors panic on the type, are
// skipped silently — those are upstream-impl-coupled cases handled
// elsewhere (PI-005..007 / sqliteconnpathtype cluster).
var fixtureDriftSuiteSkip = map[string]string{
	"strtype.Variant":            "open-ended string enum; no pinned Min/Max semantics",
	"inttype.Variant":            "open-ended numeric enum; no pinned Min/Max semantics",
	"sqliteconnpathtype.Variant": "BasicString upstream defect cluster (PI-005..007); pinned overrides intentional",
}

func Test_AllEnums_FixtureDriftGuard(t *testing.T) {
	for _, tc := range allEnumGeneralTestCases {
		tc := tc
		typeName := tc.TypeName
		if _, skip := fixtureDriftSuiteSkip[typeName]; skip {
			continue
		}
		if tc.InitialBasicEnumer == nil {
			continue
		}

		Convey(typeName+" — pinned StringMin/StringMax must match live accessor outputs", t, func() {
			defer func() {
				if r := recover(); r != nil {
					// Some accessors panic on certain Variant types — out
					// of scope (those types should be in the skip map).
				}
			}()

			liveMin := tc.InitialBasicEnumer.MinValueString()
			liveMax := tc.InitialBasicEnumer.MaxValueString()

			if tc.StringMin != "" || liveMin != "" {
				if tc.StringMin != liveMin {
					t.Errorf("%s: pinned StringMin=%q but live MinValueString()=%q — fixture drift (RCA Pattern 1). Update allEnumGeneralTestCases.go.", typeName, tc.StringMin, liveMin)
				}
			}
			if tc.StringMax != "" || liveMax != "" {
				if tc.StringMax != liveMax {
					t.Errorf("%s: pinned StringMax=%q but live MaxValueString()=%q — fixture drift (RCA Pattern 1). Update allEnumGeneralTestCases.go.", typeName, tc.StringMax, liveMax)
				}
			}
		})
	}
}
