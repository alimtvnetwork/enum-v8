package completionstate

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// AL-05: Constructor surface coverage for completionstate.
func Test_CompletionState_Constructors(t *testing.T) {
	knownNames := []string{
		"Invalid", "Initiate", "Running", "Success",
		"SuccessWithWarning", "FailedMiddleWithError", "CompleteWithError",
	}

	Convey("completionstate.New — known names round-trip", t, func() {
		for _, name := range knownNames {
			v, err := New(name)
			So(err, ShouldBeNil)
			So(v.Name(), ShouldEqual, name)
		}
	})

	Convey("completionstate.New — unknown name returns error + Invalid", t, func() {
		v, err := New("__no_such_completion_state__zzz")
		So(err, ShouldNotBeNil)
		So(v, ShouldEqual, Invalid)
	})

	Convey("completionstate.NewMust — known names succeed", t, func() {
		for _, name := range knownNames {
			v := NewMust(name)
			So(v.Name(), ShouldEqual, name)
		}
	})

	Convey("completionstate.Max / Min", t, func() {
		So(Min(), ShouldEqual, Invalid)
		So(int(Max()), ShouldBeGreaterThanOrEqualTo, int(Min()))
	})
}
