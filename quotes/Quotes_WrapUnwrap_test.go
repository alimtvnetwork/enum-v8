package quotes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// AL-06: Bespoke quotes/ suite — exercises the wrap/unwrap pipeline that
// is NOT covered by allBasicEnumsCollection (quotes is not registered in
// the cross-package collection).
//
// Coverage targets: WrapWith, UnWrapWith, HasBothWrappedWith, WhichQuote,
// getQuoteStatus, Quote.Wrap, Quote.SelfWrap, Quote.UnWrap, Quote.IsWrapped,
// Quote.GetOther, Quote.WrapAny, Quote.WrapString, Quote.WrapSkipOnExist,
// Quote.WrapRegardless, Quote.WrapFmtString, Quote.IsEqual.
func Test_Quotes_WrapUnwrap(t *testing.T) {
	Convey("HasBothWrappedWith — boundary cases", t, func() {
		So(HasBothWrappedWith("", Double), ShouldBeFalse)
		So(HasBothWrappedWith(`"`, Double), ShouldBeFalse) // single char
		So(HasBothWrappedWith(`"hi"`, Double), ShouldBeTrue)
		So(HasBothWrappedWith(`'hi'`, Single), ShouldBeTrue)
		So(HasBothWrappedWith("`hi`", Backtick), ShouldBeTrue)
		So(HasBothWrappedWith(`"hi'`, Double), ShouldBeFalse) // mismatched
		So(HasBothWrappedWith(`hi`, Double), ShouldBeFalse)
	})

	Convey("WrapWith — empty input returns SelfWrap", t, func() {
		So(WrapWith("", Double, false), ShouldEqual, Double.SelfWrap())
		So(WrapWith("", Single, false), ShouldEqual, Single.SelfWrap())
		So(WrapWith("", Backtick, false), ShouldEqual, Backtick.SelfWrap())
	})

	Convey("WrapWith — already-wrapped + skipOnExist passes through", t, func() {
		So(WrapWith(`"hi"`, Double, true), ShouldEqual, `"hi"`)
	})

	Convey("WrapWith — already-wrapped + !skipOnExist re-wraps via Quote.Wrap", t, func() {
		// HasBothWrappedWith returns true but skipOnExists is false; falls through to
		// getQuoteStatus which detects the leading quote (left=true) and appends it.
		got := WrapWith(`"hi"`, Double, false)
		So(got, ShouldNotEqual, `"hi"`)
		So(got, ShouldNotBeBlank)
	})

	Convey("WrapWith — plain string gets wrapped", t, func() {
		So(WrapWith("hi", Double, false), ShouldEqual, `"hi"`)
		So(WrapWith("hi", Single, false), ShouldEqual, `'hi'`)
		So(WrapWith("hi", Backtick, false), ShouldEqual, "`hi`")
	})

	Convey("WrapWith — left-only quote gets right side appended", t, func() {
		got := WrapWith(`"hi`, Double, false)
		So(got, ShouldEqual, `"hi"`)
	})

	Convey("WrapWith — right-only quote gets left side prepended", t, func() {
		got := WrapWith(`hi"`, Double, false)
		So(got, ShouldEqual, `"hi"`)
	})

	Convey("UnWrapWith — empty input returns empty", t, func() {
		So(UnWrapWith("", Double), ShouldEqual, "")
	})

	Convey("UnWrapWith — both-wrapped strips one char from each side (PI-008 fix)", t, func() {
		// PI-008 fix: unWrapBoth now correctly returns s[1:length-1].
		So(UnWrapWith(`"hi"`, Double), ShouldEqual, "hi")
		So(UnWrapWith(`'hi'`, Single), ShouldEqual, "hi")
		So(UnWrapWith("`hi`", Backtick), ShouldEqual, "hi")
	})

	Convey("UnWrapWith — single-side strips exactly the wrapped boundary (PI-008 fix)", t, func() {
		// PI-008 fix: unWrapSingle now strips exactly one char on the wrapped side.
		// Left-only: status.IsLeft = true → s[1:length].
		So(UnWrapWith(`"hi`, Double), ShouldEqual, "hi")
		// Right-only: status.IsLeft = false → s[0:length-1].
		So(UnWrapWith(`hi"`, Double), ShouldEqual, "hi")
	})

	Convey("UnWrapWith — no quotes returns input as-is", t, func() {
		So(UnWrapWith("plain", Double), ShouldEqual, "plain")
	})

	Convey("WhichQuote — known chars are recognised", t, func() {
		So(WhichQuote('"', true).IsQuoteFound, ShouldBeTrue)
		So(WhichQuote('"', true).IsLeft, ShouldBeTrue)
		So(WhichQuote('\'', false).IsQuoteFound, ShouldBeTrue)
		So(WhichQuote('\'', false).IsLeft, ShouldBeFalse)
		So(WhichQuote('`', true).IsQuoteFound, ShouldBeTrue)
	})

	Convey("WhichQuote — unknown chars return empty status", t, func() {
		st := WhichQuote('x', true)
		So(st.IsQuoteFound, ShouldBeFalse)
	})

	Convey("Quote.Wrap / SelfWrap / IsEqual / GetOther", t, func() {
		So(Double.Wrap(""), ShouldEqual, `""`)
		So(Double.Wrap("hi"), ShouldEqual, `"hi"`)
		So(Double.SelfWrap(), ShouldNotBeBlank)
		So(Double.IsEqual('"'), ShouldBeTrue)
		So(Double.IsEqual('x'), ShouldBeFalse)
		So(Single.GetOther(), ShouldEqual, Single)
		So(Double.GetOther(), ShouldEqual, Double)
	})

	Convey("Quote.WrapAny / WrapString / WrapSkipOnExist / WrapRegardless", t, func() {
		So(Double.WrapAny("hi"), ShouldEqual, `"hi"`)
		So(Double.WrapString("hi"), ShouldEqual, `"hi"`)
		So(Double.WrapSkipOnExist(`"hi"`), ShouldEqual, `"hi"`)
		So(Double.WrapRegardless("hi"), ShouldEqual, `"hi"`)
	})

	Convey("Quote.WrapAnySkipOnExist", t, func() {
		So(Double.WrapAnySkipOnExist(`"hi"`), ShouldEqual, `"hi"`)
		So(Double.WrapAnySkipOnExist("hi"), ShouldEqual, `"hi"`)
	})

	Convey("Quote.IsWrapped / UnWrap / WrapWithOptions", t, func() {
		So(Double.IsWrapped(`"hi"`), ShouldBeTrue)
		So(Double.IsWrapped("hi"), ShouldBeFalse)
		So(Double.UnWrap(`"hi"`), ShouldEqual, "hi")
		So(Double.WrapWithOptions(true, `"hi"`), ShouldEqual, `"hi"`)
	})

	Convey("Quote.WrapFmtString — replaces {wrapped} placeholder", t, func() {
		got := Double.WrapFmtString("prefix {wrapped} suffix", "hi")
		So(got, ShouldEqual, `prefix "hi" suffix`)
	})
}
