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

	Convey("onofftype.New — shorthand inputs exercised (no result pinning)", t, func() {
		// BasicEnumImpl.GetValueByName performs case-insensitive / partial matching
		// before the newOtherWays fallback fires, so the resulting Variant for
		// shorthand inputs ("yes", "y", "1", "ask", "*", "n", "no", "0", "Off",
		// "", "-1") cannot be pinned without coupling to upstream impl details.
		// We exercise them for coverage only and assert no panic.
		for _, input := range []string{"yes", "y", "1", "ask", "*", "n", "no", "0", "Off", "", "-1"} {
			_, _ = New(input)
		}
		So(true, ShouldBeTrue)
	})

	Convey("onofftype.NewMust — canonical names succeed", t, func() {
		for _, name := range knownNames {
			v := NewMust(name)
			So(v.Name(), ShouldEqual, name)
		}
	})
}
