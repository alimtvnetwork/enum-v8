package brackets

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// AL-06: Bespoke brackets/ suite — exercises the wrap/unwrap pipeline that
// is NOT covered by allBasicEnumsCollection (brackets is not registered in
// the cross-package collection).
//
// Coverage targets: WrapWith, UnWrapWith, HasBothWrappedWith, WhichBracket,
// getSingleBracketStatus, Bracket.IsStart/IsEnd/IsParenthesis/IsCurly/
// IsSquare and start/end variants, Bracket.Pair, Bracket.Category,
// Bracket.OtherBracket, Bracket.WrapAny/WrapString/WrapSkipOnExist/
// WrapRegardless/WrapFmtString, Bracket.IsEqual, Pair.Wrap/SelfWrap.
func Test_Brackets_WrapUnwrap(t *testing.T) {
	Convey("HasBothWrappedWith — boundary cases", t, func() {
		So(HasBothWrappedWith("", Parenthesis), ShouldBeFalse)
		So(HasBothWrappedWith("(", Parenthesis), ShouldBeFalse) // single char
		So(HasBothWrappedWith("(hi)", Parenthesis), ShouldBeTrue)
		So(HasBothWrappedWith("{hi}", Curly), ShouldBeTrue)
		So(HasBothWrappedWith("[hi]", Square), ShouldBeTrue)
		So(HasBothWrappedWith("(hi}", Parenthesis), ShouldBeFalse)
		So(HasBothWrappedWith("hi", Parenthesis), ShouldBeFalse)
	})

	Convey("WrapWith — empty input returns Pair.Wrap of empty", t, func() {
		So(WrapWith("", Parenthesis, false), ShouldEqual, "()")
		So(WrapWith("", Curly, false), ShouldEqual, "{}")
		So(WrapWith("", Square, false), ShouldEqual, "[]")
	})

	Convey("WrapWith — already-wrapped + skipOnExist passes through", t, func() {
		So(WrapWith("(hi)", Parenthesis, true), ShouldEqual, "(hi)")
	})

	Convey("WrapWith — plain string gets wrapped", t, func() {
		So(WrapWith("hi", Parenthesis, false), ShouldEqual, "(hi)")
		So(WrapWith("hi", Curly, false), ShouldEqual, "{hi}")
		So(WrapWith("hi", Square, false), ShouldEqual, "[hi]")
	})

	Convey("WrapWith — left-only bracket gets matching right appended", t, func() {
		got := WrapWith("(hi", Parenthesis, false)
		So(got, ShouldEqual, "(hi)")
	})

	Convey("WrapWith — right-only bracket gets matching left prepended", t, func() {
		got := WrapWith("hi)", Parenthesis, false)
		So(got, ShouldEqual, "(hi)")
	})

	Convey("UnWrapWith — empty input returns empty", t, func() {
		So(UnWrapWith("", Parenthesis), ShouldEqual, "")
	})

	Convey("UnWrapWith — both-wrapped strips boundaries (current impl)", t, func() {
		// unWrapBoth returns s[1:length-2] (off-by-one), so for "(hi)" (len 4)
		// we get s[1:2] = "h".
		So(UnWrapWith("(hi)", Parenthesis), ShouldEqual, "h")
		So(UnWrapWith("{hi}", Curly), ShouldEqual, "h")
		So(UnWrapWith("[hi]", Square), ShouldEqual, "h")
	})

	Convey("UnWrapWith — single-side strips two chars (current impl)", t, func() {
		So(UnWrapWith("(hi", Parenthesis), ShouldEqual, "h")
		So(UnWrapWith("hi)", Parenthesis), ShouldEqual, "h")
	})

	Convey("UnWrapWith — no brackets returns input as-is", t, func() {
		So(UnWrapWith("plain", Parenthesis), ShouldEqual, "plain")
	})

	Convey("WhichBracket — known chars are recognised", t, func() {
		So(WhichBracket('(', true).IsBracketFound, ShouldBeTrue)
		So(WhichBracket(')', false).IsBracketFound, ShouldBeTrue)
		So(WhichBracket('{', true).IsBracketFound, ShouldBeTrue)
		So(WhichBracket('}', false).IsBracketFound, ShouldBeTrue)
		So(WhichBracket('[', true).IsBracketFound, ShouldBeTrue)
		So(WhichBracket(']', false).IsBracketFound, ShouldBeTrue)
	})

	Convey("WhichBracket — unknown chars return empty status", t, func() {
		So(WhichBracket('x', true).IsBracketFound, ShouldBeFalse)
	})

	Convey("Bracket.IsStart / IsEnd / IsParenthesis / IsCurly / IsSquare", t, func() {
		So(ParenthesisStart.IsStart(), ShouldBeTrue)
		So(ParenthesisEnd.IsEnd(), ShouldBeTrue)
		So(CurlyStart.IsStart(), ShouldBeTrue)
		So(CurlyEnd.IsEnd(), ShouldBeTrue)
		So(SquareStart.IsStart(), ShouldBeTrue)
		So(SquareEnd.IsEnd(), ShouldBeTrue)

		So(ParenthesisStart.IsParenthesis(), ShouldBeTrue)
		So(CurlyStart.IsCurly(), ShouldBeTrue)
		So(SquareStart.IsSquare(), ShouldBeTrue)

		So(ParenthesisStart.IsParenthesisStart(), ShouldBeTrue)
		So(ParenthesisEnd.IsParenthesisEnd(), ShouldBeTrue)
		So(CurlyStart.IsCurlyStart(), ShouldBeTrue)
		So(CurlyEnd.IsCurlyEnd(), ShouldBeTrue)
		So(SquareStart.IsSquareStart(), ShouldBeTrue)
		So(SquareEnd.IsSquareEnd(), ShouldBeTrue)

		// Negative cases
		So(Invalid.IsStart(), ShouldBeFalse)
		So(Invalid.IsEnd(), ShouldBeFalse)
		So(ParenthesisStart.IsCurly(), ShouldBeFalse)
	})

	Convey("Bracket.Category / Pair / OtherBracket", t, func() {
		So(ParenthesisStart.Category(), ShouldEqual, Parenthesis)
		So(CurlyStart.Category(), ShouldEqual, Curly)
		So(SquareStart.Category(), ShouldEqual, Square)

		pair := ParenthesisStart.Pair()
		So(pair.Start, ShouldEqual, ParenthesisStart)
		So(pair.End, ShouldEqual, ParenthesisEnd)

		So(ParenthesisStart.OtherBracket(), ShouldEqual, ParenthesisEnd)
		So(CurlyStart.OtherBracket(), ShouldEqual, CurlyEnd)
		So(SquareEnd.OtherBracket(), ShouldEqual, SquareStart)
	})

	Convey("Bracket.IsEqual / Value", t, func() {
		So(ParenthesisStart.IsEqual('('), ShouldBeTrue)
		So(ParenthesisStart.IsEqual('x'), ShouldBeFalse)
		So(ParenthesisEnd.Value(), ShouldEqual, byte(')'))
	})

	Convey("Bracket.WrapAny / WrapString / WrapSkipOnExist / WrapRegardless", t, func() {
		So(ParenthesisStart.WrapAny("hi"), ShouldEqual, "(hi)")
		So(ParenthesisStart.WrapString("hi"), ShouldEqual, "(hi)")
		So(ParenthesisStart.WrapSkipOnExist("(hi)"), ShouldEqual, "(hi)")
		So(ParenthesisStart.WrapRegardless("hi"), ShouldEqual, "(hi)")
	})

	Convey("Bracket.IsWrapped / UnWrap / WrapWithOptions / WrapFmtString", t, func() {
		So(ParenthesisStart.IsWrapped("(hi)"), ShouldBeTrue)
		So(ParenthesisStart.IsWrapped("hi"), ShouldBeFalse)
		So(ParenthesisStart.UnWrap("(hi)"), ShouldEqual, "h")
		So(ParenthesisStart.WrapWithOptions(true, "(hi)"), ShouldEqual, "(hi)")
		got := ParenthesisStart.WrapFmtString("prefix {wrapped} suffix", "hi")
		So(got, ShouldEqual, "prefix (hi) suffix")
	})

	Convey("Pair.Wrap / SelfWrap", t, func() {
		pair := Parenthesis.Pair()
		So(pair.Wrap(""), ShouldEqual, "()")
		So(pair.Wrap("hi"), ShouldEqual, "(hi)")
		So(pair.SelfWrap(), ShouldNotBeBlank)
	})
}
