package envtype

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// AL-05b: Constructor surface coverage for envtype.
//
// Note: this package's zero-value sentinel is `Uninitialized`, not `Invalid`.
func Test_EnvType_Constructors(t *testing.T) {
	knownNames := []string{
		"Uninitialized", "Development", "Development1", "Development2",
		"Test", "Test1", "Test2",
		"Production", "Production1", "Production2",
	}

	Convey("envtype.New — known names round-trip", t, func() {
		for _, name := range knownNames {
			v, err := New(name)
			So(err, ShouldBeNil)
			So(v.Name(), ShouldEqual, name)
		}
	})

	Convey("envtype.New — unknown name returns error + Uninitialized", t, func() {
		v, err := New("__no_such_env_type__zzz")
		So(err, ShouldNotBeNil)
		So(v, ShouldEqual, Uninitialized)
	})

	Convey("envtype.NewMust — known names succeed", t, func() {
		for _, name := range knownNames {
			v := NewMust(name)
			So(v.Name(), ShouldEqual, name)
		}
	})
}
