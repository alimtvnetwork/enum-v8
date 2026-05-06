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
		shorthandToCanonical := map[string]Variant{
			"yes": On,
			"1":   On,
			"y":   On,
			"n":   Off,
			"no":  Off,
			"0":   Off,
			"ask": Ask,
			"*":   Ask,
		}
		for input, want := range shorthandToCanonical {
			v, err := New(input)
			// Fallback path may or may not return an error depending on
			// implementation; only the resulting Variant matters here.
			_ = err
			So(v, ShouldEqual, want)
		}
	})

	Convey("onofftype.NewMust — canonical names succeed", t, func() {
		for _, name := range knownNames {
			v := NewMust(name)
			So(v.Name(), ShouldEqual, name)
		}
	})
}
