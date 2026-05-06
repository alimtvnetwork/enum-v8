package iptype

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// AL-05b: Constructor surface coverage for iptype.
//
// Note: the registered enum names in iptype/vars.go use the prefixed form
// "IpV4"/"IpV6" (the Go identifiers V4/V6 are aliases). Aliases like "v4",
// "ver4", "ipv4" are also accepted via aliasMap.
func Test_IpType_Constructors(t *testing.T) {
	knownNames := []string{"Invalid", "IpV4", "IpV6"}

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
