package strtype

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// AL2: Broad coverage uplift for strtype (was 16.2%).
func Test_StrType_Coverage(t *testing.T) {
	Convey("strtype — value & length accessors", t, func() {
		v := Variant("hello")
		So(v.Value(), ShouldEqual, "hello")
		So(v.StringValue(), ShouldEqual, "hello")
		So(v.String(), ShouldEqual, "hello")
		So(v.Name(), ShouldEqual, "hello")
		So(v.Length(), ShouldEqual, 5)
		So(v.Size(), ShouldEqual, 5)
		So(v.Count(), ShouldEqual, 5)
		So(v.AllChars(), ShouldNotBeNil)
		So(v.AllRunes(), ShouldNotBeNil)
		l, _ := v.RunesLength()
		So(l, ShouldEqual, 5)
	})

	Convey("strtype — empty / whitespace / defined", t, func() {
		So(Variant("").IsEmpty(), ShouldBeTrue)
		So(Variant("x").IsDefined(), ShouldBeTrue)
		So(Variant("   ").IsWhitespace(), ShouldBeTrue)
		So(Variant("  hi  ").Trim().Value(), ShouldEqual, "hi")
		So(Variant("a").IsEqualTrim("a"), ShouldBeTrue)
		So(Variant("Invalid").IsInvalid(), ShouldBeTrue)
		So(Variant("ok").IsValid(), ShouldBeTrue)
		So(Variant("a").HasAnyItem(), ShouldBeTrue)
	})

	Convey("strtype — string ops", t, func() {
		So(Variant("aaa").Replace("a", "b").Value(), ShouldEqual, "bbb")
		So(Variant("abcabc").Remove("a").Value(), ShouldEqual, "bcbc")
		So(Variant("abc").RemoveMany("a", "c").Value(), ShouldEqual, "b")
		So(Variant("a,b,c").SplitBy(","), ShouldHaveLength, 3)
		So(Variant("a b c").SplitByWhitespace(), ShouldHaveLength, 3)
		So(Variant("a, b , c").SplitTrimmedNonEmpty(","), ShouldHaveLength, 3)
		So(Variant("hello").AddSuffixOnMissing("-x"), ShouldEqual, "hello-x")
		So(Variant("hello-x").AddSuffixOnMissing("-x"), ShouldEqual, "hello-x")
		So(Variant("hello").AddPrefixOnMissing("p-"), ShouldEqual, "p-hello")
		So(Variant("p-hello").AddPrefixOnMissing("p-"), ShouldEqual, "p-hello")
		So(Variant("hello").Add(" world").Value(), ShouldEqual, "hello world")
		So(Variant("a").Append(Variant("b")).Value(), ShouldEqual, "ab")
		So(Variant("b").Prepend(Variant("a")).Value(), ShouldEqual, "ab")
		So(Variant("b").PrependString("a").Value(), ShouldEqual, "ab")
		So(Variant("a").AddAnother(Variant("b")).Value(), ShouldEqual, "ab")
		So(Variant("a").Join(Variant("b"), Variant("c")).Value(), ShouldEqual, "a-b-c")
		So(Variant("a").JoinStrings("b", "c").Value(), ShouldEqual, "a-b-c")
		So(Variant("a").AppendIf(true, Variant("b")).Value(), ShouldEqual, "ab")
		So(Variant("a").AppendIf(false, Variant("b")).Value(), ShouldEqual, "a")
		So(Variant("a").AppendStringIf(true, "b").Value(), ShouldEqual, "ab")
		So(Variant("b").PrependIf(true, Variant("a")).Value(), ShouldEqual, "ab")
		So(Variant("b").PrependStringIf(true, "a").Value(), ShouldEqual, "ab")
	})

	Convey("strtype — comparison & search", t, func() {
		v := Variant("hello world")
		So(v.IsContains("world"), ShouldBeTrue)
		So(v.IsStartsWith("hello"), ShouldBeTrue)
		So(v.IsEndsWith("world"), ShouldBeTrue)
		So(v.HasPrefix("hello"), ShouldBeTrue)
		So(v.HasSuffix("world"), ShouldBeTrue)
		So(v.Index("world"), ShouldEqual, 6)
		So(v.LastIndexOf("o"), ShouldBeGreaterThanOrEqualTo, 0)
		So(v.IsEqual("hello world"), ShouldBeTrue)
		So(v.IsEqualAnother(Variant("hello world")), ShouldBeTrue)
		So(v.Is(Variant("hello world")), ShouldBeTrue)
		So(Variant("b").IsGreater("a"), ShouldBeTrue)
		So(Variant("b").IsGreaterEqual("b"), ShouldBeTrue)
		So(Variant("a").IsLess("b"), ShouldBeTrue)
		So(Variant("a").IsLessEqual("a"), ShouldBeTrue)
		So(Variant("Hello").IsNameEqual("Hello"), ShouldBeTrue)
		So(Variant("a").IsAnyNamesOf("a", "b"), ShouldBeTrue)
	})

	Convey("strtype — wraps & quotation", t, func() {
		v := Variant("hi")
		So(v.QuotationWrap(), ShouldNotBeBlank)
		So(v.CurlyWrap(), ShouldEqual, "{hi}")
		So(v.SquareWrap(), ShouldEqual, "[hi]")
	})

	Convey("strtype — safe substr", t, func() {
		v := Variant("hello world")
		So(string(v.SafeSubStringStart(5)), ShouldEqual, "hello")
		So(string(v.SafeSubStringEnd(6)), ShouldEqual, "world")
		So(string(v.SafeSubString(0, 5)), ShouldEqual, "hello")
		So(string(v.SafeSubString(0, 999)), ShouldEqual, "hello world")
	})

	Convey("strtype — boolean combinators", t, func() {
		So(Variant("").OrEmpty(Variant("x")), ShouldBeTrue)
		So(Variant("x").OrHasElement(Variant("")), ShouldBeTrue)
		So(Variant("x").AndHasElement(Variant("y")), ShouldBeTrue)
		So(Variant("").AndIsEmpty(Variant("")), ShouldBeTrue)
	})

	Convey("strtype — int conversions", t, func() {
		So(Variant("42").Integer(), ShouldEqual, 42)
		i, err := Variant("99").ConvInteger()
		So(err, ShouldBeNil)
		So(i, ShouldEqual, 99)
		v, ok := Variant("7").IntegerDefaultVal(0)
		So(ok, ShouldBeTrue)
		So(v, ShouldEqual, 7)
		v2, ok2 := Variant("xx").IntegerDefaultVal(99)
		So(ok2, ShouldBeFalse)
		So(v2, ShouldEqual, 99)
		So(Variant("42").IntType().Value(), ShouldEqual, 42)
	})

	Convey("strtype — surface accessors", t, func() {
		v := Variant("x")
		So(v.TypeName(), ShouldNotBeBlank)
		So(v.NameValue(), ShouldNotBeBlank)
		So(v.ValueByte(), ShouldEqual, byte('x'))
		_ = v.ValueInt()
		_ = v.ValueInt8()
		_ = v.ValueInt16()
		_ = v.ValueInt32()
		_ = v.ValueUInt16()
		_ = v.ValueString()
		_ = v.ToNumberString()
		So(v.AsBasicEnumer(), ShouldNotBeNil)
		So(v.ToPtr(), ShouldNotBeNil)
		_ = v.RangesDynamicMap()
		_ = v.AllNameValues()
		_ = v.MinMaxAny()
		_ = v.MinInt()
		_ = v.MaxInt()
		_ = v.MinValueString()
		_ = v.MaxValueString()
		_ = v.IntegerEnumRanges()
		_ = v.RangeNamesCsv()
		_ = v.EnumType()
		So(v.Format("{name}"), ShouldNotBeBlank)
	})

	Convey("strtype — package helpers New / NewUsingInteger", t, func() {
		So(string(New("z")), ShouldEqual, "z")
		So(string(NewUsingInteger(7)), ShouldEqual, "7")
	})
}
