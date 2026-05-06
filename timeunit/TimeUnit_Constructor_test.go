package timeunit

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// AL-05b: Constructor surface coverage for timeunit.
func Test_TimeUnit_Constructors(t *testing.T) {
	knownNames := []string{
		"Invalid", "Millisecond", "Second", "Minute",
		"Hour", "Day", "Month", "Year",
	}

	Convey("timeunit.New — known names round-trip", t, func() {
		for _, name := range knownNames {
			v, err := New(name)
			So(err, ShouldBeNil)
			So(v.Name(), ShouldEqual, name)
		}
	})

	Convey("timeunit.New — unknown name returns error + Invalid", t, func() {
		v, err := New("__no_such_time_unit__zzz")
		So(err, ShouldNotBeNil)
		So(v, ShouldEqual, Invalid)
	})

	Convey("timeunit.NewMust — known names succeed", t, func() {
		for _, name := range knownNames {
			v := NewMust(name)
			So(v.Name(), ShouldEqual, name)
		}
	})
}
