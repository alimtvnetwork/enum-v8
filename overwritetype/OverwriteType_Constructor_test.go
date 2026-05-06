package overwritetype

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// AL-05b: Constructor surface coverage for overwritetype.
//
// Note: the Go consts include ForceWriteRepeat (4) and SkipFilesRepeat (5),
// but the enum-map registration in vars.go skips those two (Ranges leaves
// indices 4 and 5 with empty names). Only the 6 registered names round-trip
// through New / NewMust.
func Test_OverwriteType_Constructors(t *testing.T) {
	knownNames := []string{
		"Invalid", "ForceWrite", "SkipOnExistFiles",
		"IgnoreRepeatInFolderNameExtraction", "Yes", "No",
	}

	Convey("overwritetype.New — known names round-trip", t, func() {
		for _, name := range knownNames {
			v, err := New(name)
			So(err, ShouldBeNil)
			So(v.Name(), ShouldEqual, name)
		}
	})

	Convey("overwritetype.New — unknown name returns error + Invalid", t, func() {
		v, err := New("__no_such_overwrite_type__zzz")
		So(err, ShouldNotBeNil)
		So(v, ShouldEqual, Invalid)
	})

	Convey("overwritetype.NewMust — known names succeed", t, func() {
		for _, name := range knownNames {
			v := NewMust(name)
			So(v.Name(), ShouldEqual, name)
		}
	})
}
