package compressformats

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// AL-05: Constructor surface coverage for compressformats.
//
// Note: iota ordering is unusual — `Invalid = 5` is the largest underlying
// byte. As of v0.64.0 (Pattern-8 fix) `Min()` returns `Zip` (the first real
// member, iota 0) and `Max()` returns `TarBz2` (the last real member),
// bypassing the trailing-Invalid sentinel defect in BasicEnumImpl.
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
		So(Min(), ShouldEqual, Zip)
		So(Max(), ShouldEqual, TarBz2)
		So(RangesInvalidErr(), ShouldNotBeNil)
	})
}
