package dbaction

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// AL-05b: Constructor surface coverage for dbaction.
func Test_DbAction_Constructors(t *testing.T) {
	knownNames := []string{
		"Invalid", "Create", "Update", "Delete", "Read",
		"CreateOrUpdate", "SkipOnExist", "SkipOnNonExist",
		"DropOnExist", "UpdateOnExist",
	}

	Convey("dbaction.New — known names round-trip", t, func() {
		for _, name := range knownNames {
			v, err := New(name)
			So(err, ShouldBeNil)
			So(v.Name(), ShouldEqual, name)
		}
	})

	Convey("dbaction.New — unknown name returns error + Invalid", t, func() {
		v, err := New("__no_such_db_action__zzz")
		So(err, ShouldNotBeNil)
		So(v, ShouldEqual, Invalid)
	})

	Convey("dbaction.NewMust — known names succeed", t, func() {
		for _, name := range knownNames {
			v := NewMust(name)
			So(v.Name(), ShouldEqual, name)
		}
	})
}
