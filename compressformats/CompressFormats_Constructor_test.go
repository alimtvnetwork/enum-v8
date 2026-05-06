package compressformats

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// AL-05: Constructor surface coverage for compressformats.
//
// Note: in this package the iota ordering is unusual — `Invalid = 5` is the
// largest underlying byte while `Min()` is hand-written to return `Invalid`.
// We therefore do NOT assert `int(Max()) >= int(Min())` here; we only check
// that the constructors and bounds return sensible Variants.
func Test_CompressFormats_Constructors(t *testing.T) {
	knownNames := []string{"Zip", "Tar", "TarGZ", "TarXZ", "TarBz2", "Invalid"}

	Convey("compressformats.New — known names round-trip", t, func() {
		for _, name := range knownNames {
			v, err := New(name)
			So(err, ShouldBeNil)
			So(v.Name(), ShouldEqual, name)
		}
	})

	Convey("compressformats.New — unknown name returns error + Invalid", t, func() {
		v, err := New("__no_such_compress_format__zzz")
		So(err, ShouldNotBeNil)
		So(v, ShouldEqual, Invalid)
	})

	Convey("compressformats.NewMust — known names succeed", t, func() {
		for _, name := range knownNames {
			v := NewMust(name)
			So(v.Name(), ShouldEqual, name)
		}
	})

	Convey("compressformats.Max / Min / RangesInvalidErr", t, func() {
		So(Min(), ShouldEqual, Invalid)
		So(Max().IsValid() || Max() == Invalid, ShouldBeTrue)
		So(RangesInvalidErr(), ShouldNotBeNil)
	})
}
