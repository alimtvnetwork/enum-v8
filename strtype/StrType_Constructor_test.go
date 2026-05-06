package strtype

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// AL-07: Constructor + GetSet coverage for strtype.
func Test_StrType_Constructors(t *testing.T) {
	Convey("New wraps string", t, func() {
		v := New("hello")
		So(string(v), ShouldEqual, "hello")
	})

	Convey("NewUsingInteger converts int to string-backed Variant", t, func() {
		v := NewUsingInteger(42)
		So(string(v), ShouldEqual, "42")
	})

	Convey("NewFileReader returns a FileReader bound to the path", t, func() {
		fr := NewFileReader("/tmp/__strtype_does_not_exist__.txt")
		So(fr, ShouldNotBeNil)
	})
}

func Test_StrType_GetSet(t *testing.T) {
	Convey("GetSet picks based on condition", t, func() {
		So(string(GetSet(true, New("a"), New("b"))), ShouldEqual, "a")
		So(string(GetSet(false, New("a"), New("b"))), ShouldEqual, "b")
	})

	Convey("GetSetVariant wraps a single byte", t, func() {
		So(string(GetSetVariant(true, 'x', 'y')), ShouldEqual, "x")
		So(string(GetSetVariant(false, 'x', 'y')), ShouldEqual, "y")
	})
}
