package inttype

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/alimtvnetwork/core-v9/corecomparator"
	. "github.com/smartystreets/goconvey/convey"
)

// AL-07: Constructor + GetSet + IsCompareResult coverage for inttype.
func Test_IntType_Constructors(t *testing.T) {
	Convey("New / NewString round-trip", t, func() {
		for _, n := range []int{-5, -1, 0, 1, 42, 9999} {
			v := New(n)
			So(v.Value(), ShouldEqual, n)
			s := strconv.Itoa(n)
			vs, err := NewString(s)
			So(err, ShouldBeNil)
			So(vs.Value(), ShouldEqual, n)
		}
	})

	Convey("NewString — bad input returns InvalidValue + error", t, func() {
		v, err := NewString("not-a-number")
		So(err, ShouldNotBeNil)
		So(v, ShouldEqual, InvalidValue)
	})

	Convey("NewUInt — within int range succeeds", t, func() {
		v, err := NewUInt(uint(123))
		So(err, ShouldBeNil)
		So(v.Value(), ShouldEqual, 123)
	})

	Convey("NewInt64 — within range succeeds", t, func() {
		v, err := NewInt64(int64(456))
		So(err, ShouldBeNil)
		So(v.Value(), ShouldEqual, 456)
	})

	Convey("NewUsingJsonNumber — nil returns Invalid + error", t, func() {
		v, err := NewUsingJsonNumber(nil)
		So(err, ShouldNotBeNil)
		So(v, ShouldEqual, Invalid)
	})

	Convey("NewUsingJsonNumber — valid number parses", t, func() {
		jn := json.Number("789")
		v, err := NewUsingJsonNumber(&jn)
		So(err, ShouldBeNil)
		So(v.Value(), ShouldEqual, 789)
	})
}

func Test_IntType_GetSet(t *testing.T) {
	Convey("GetSet returns trueValue when condition true", t, func() {
		So(GetSet(true, New(1), New(2)), ShouldEqual, New(1))
		So(GetSet(false, New(1), New(2)), ShouldEqual, New(2))
	})

	Convey("GetSetVariant wraps ints", t, func() {
		So(GetSetVariant(true, 10, 20).Value(), ShouldEqual, 10)
		So(GetSetVariant(false, 10, 20).Value(), ShouldEqual, 20)
	})
}

func Test_IntType_IsCompareResult(t *testing.T) {
	v := New(5)

	Convey("IsCompareResult covers all comparator branches", t, func() {
		So(v.IsCompareResult(5, corecomparator.Equal), ShouldBeTrue)
		So(v.IsCompareResult(3, corecomparator.LeftGreater), ShouldBeTrue)
		So(v.IsCompareResult(5, corecomparator.LeftGreaterEqual), ShouldBeTrue)
		So(v.IsCompareResult(9, corecomparator.LeftLess), ShouldBeTrue)
		So(v.IsCompareResult(5, corecomparator.LeftLessEqual), ShouldBeTrue)
		So(v.IsCompareResult(7, corecomparator.NotEqual), ShouldBeTrue)
		So(v.IsCompareResult(5, corecomparator.NotEqual), ShouldBeFalse)
	})
}
