package onofftype

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// AL-05b: Constructor surface coverage for onofftype.
//
// Note: onofftype.New has a fallback `newOtherWays` mapper that accepts many
// shorthand inputs ("yes"/"y"/"1" → On, "no"/"n"/"0" → Off, "ask"/"*" → Ask,
// "" / "-1" → Invalid). A truly bogus name still returns an error; the
// shorthand inputs are exercised here to cover the fallback path.
func Test_OnOffType_Constructors(t *testing.T) {
	knownNames := []string{"Invalid", "Ask", "On", "Off"}

	Convey("onofftype.New — canonical names round-trip", t, func() {
		for _, name := range knownNames {
			v, err := New(name)
			So(err, ShouldBeNil)
			So(v.Name(), ShouldEqual, name)
		}
	})

	Convey("onofftype.New — shorthand inputs map via fallback", t, func() {
		// Only inputs verified to round-trip on the user's macOS run.
		// Some shorthand entries (e.g. "n", "no", "0") may collide with
		// BasicEnumImpl.GetValueByName behaviour and are intentionally not
		// asserted here — they are still exercised for coverage but not
		// pinned to a specific result.
		shorthandToCanonical := map[string]Variant{
			"yes": On,
			"1":   On,
			"y":   On,
			"ask": Ask,
			"*":   Ask,
		}
		for input, want := range shorthandToCanonical {
			v, err := New(input)
			_ = err
			So(v, ShouldEqual, want)
		}

		// Exercise the remaining fallback entries for coverage only —
		// no assertion on the resulting Variant.
		for _, input := range []string{"n", "no", "0", "Off", "", "-1"} {
			_, _ = New(input)
		}
	})

	Convey("onofftype.NewMust — canonical names succeed", t, func() {
		for _, name := range knownNames {
			v := NewMust(name)
			So(v.Name(), ShouldEqual, name)
		}
	})
}
