package accesstype

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// AL-05: Constructor surface coverage for accesstype.
//   - New(name) returns (Variant, nil) for every known name.
//   - New(unknown) returns (Invalid, non-nil error).
//   - NewMust(name) returns the matching Variant for known names.
//   - RangesInvalidErr() is non-nil and stable.
func Test_AccessType_Constructors(t *testing.T) {
	knownNames := []string{
		"Invalid", "Create", "Update", "Delete", "Read",
		"CreateOrUpdate", "SkipOnExist", "SkipOnNonExist",
		"DropOnExist", "UpdateOnExist",
	}

	Convey("accesstype.New — known names round-trip", t, func() {
		for _, name := range knownNames {
			v, err := New(name)
			So(err, ShouldBeNil)
			So(v.Name(), ShouldEqual, name)
		}
	})

	Convey("accesstype.New — unknown name returns error + Invalid", t, func() {
		v, err := New("__no_such_access_type__zzz")
		So(err, ShouldNotBeNil)
		So(v, ShouldEqual, Invalid)
	})

	Convey("accesstype.NewMust — known names succeed", t, func() {
		for _, name := range knownNames {
			v := NewMust(name)
			So(v.Name(), ShouldEqual, name)
		}
	})

	Convey("accesstype.RangesInvalidErr — non-nil", t, func() {
		So(RangesInvalidErr(), ShouldNotBeNil)
	})
}
