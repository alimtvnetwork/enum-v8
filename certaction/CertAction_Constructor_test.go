package certaction

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// AL-05: Constructor surface coverage for certaction.
func Test_CertAction_Constructors(t *testing.T) {
	knownNames := []string{"Invalid", "Create", "Renew", "Revoke"}

	Convey("certaction.New — known names round-trip", t, func() {
		for _, name := range knownNames {
			v, err := New(name)
			So(err, ShouldBeNil)
			So(v.Name(), ShouldEqual, name)
		}
	})

	Convey("certaction.New — unknown name returns error + Invalid", t, func() {
		v, err := New("__no_such_cert_action__zzz")
		So(err, ShouldNotBeNil)
		So(v, ShouldEqual, Invalid)
	})

	Convey("certaction.NewMust — known names succeed", t, func() {
		for _, name := range knownNames {
			v := NewMust(name)
			So(v.Name(), ShouldEqual, name)
		}
	})

	Convey("certaction.Max / Min / RangesInvalidErr", t, func() {
		So(Max(), ShouldEqual, Revoke)
		So(Min(), ShouldEqual, Invalid)
		So(int(Max()), ShouldBeGreaterThanOrEqualTo, int(Min()))
		So(RangesInvalidErr(), ShouldNotBeNil)
	})
}
