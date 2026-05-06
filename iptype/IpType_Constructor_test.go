package iptype

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// AL-05b: Constructor surface coverage for iptype.
func Test_IpType_Constructors(t *testing.T) {
	knownNames := []string{"Invalid", "V4", "V6"}

	Convey("iptype.New — known names round-trip", t, func() {
		for _, name := range knownNames {
			v, err := New(name)
			So(err, ShouldBeNil)
			So(v.Name(), ShouldEqual, name)
		}
	})

	Convey("iptype.New — unknown name returns error + Invalid", t, func() {
		v, err := New("__no_such_ip_type__zzz")
		So(err, ShouldNotBeNil)
		So(v, ShouldEqual, Invalid)
	})

	Convey("iptype.NewMust — known names succeed", t, func() {
		for _, name := range knownNames {
			v := NewMust(name)
			So(v.Name(), ShouldEqual, name)
		}
	})
}
