package strtype

import (
	"encoding/json"
	"testing"
)

// AL2-strtype: dedicated coverage uplift, separate from the bespoke
// Coverage suite to keep the Convey output organized.
func Test_StrType_Uplift(t *testing.T) {
	v := Variant("hello")

	// String / Name / Value / StringValue
	if v.String() != "hello" || v.Name() != "hello" || v.Value() != "hello" || v.StringValue() != "hello" {
		t.Error("string accessor mismatch")
	}
	if v.Length() != 5 || v.Size() != 5 || v.Count() != 5 {
		t.Error("length accessor mismatch")
	}
	if l, _ := v.RunesLength(); l != 5 {
		t.Error("RunesLength wrong")
	}
	_ = v.AllChars()
	_ = v.AllRunes()

	// Title* family
	_ = v.TitleQuotation("t")
	_ = v.TitleCurly("t")
	_ = v.TitleSquare("t")
	_ = v.TitleQuotationReferenceStrings("t", "a", "b")
	_ = v.TitleQuotationRefs("t", 1, "a")

	// Wraps
	_ = v.QuotationWrap()
	_ = v.CurlyWrap()
	_ = v.SquareWrap()

	// Empty / defined / whitespace / trim / equal-trim
	if !Variant("").IsEmpty() || Variant("a").IsEmpty() {
		t.Error("IsEmpty wrong")
	}
	if !Variant("x").IsDefined() || Variant("").IsDefined() {
		t.Error("IsDefined wrong")
	}
	_ = Variant("   ").IsWhitespace()
	_ = Variant("xxx").IsWhitespace()
	if Variant("  hi  ").Trim().Value() != "hi" {
		t.Error("Trim wrong")
	}
	_ = Variant("a").IsEqualTrim("a")
	_ = Variant("a").IsEqualTrim("b")

	// Replace family
	if Variant("aaa").Replace("a", "b").Value() != "bbb" {
		t.Error("Replace wrong")
	}
	_ = Variant("{x}").ReplaceUsingMapCurly(map[string]string{"x": "y"})
	_ = Variant("x").ReplaceUsingMapDirect(map[string]string{"x": "y"})
	_ = Variant("{x}").ReplaceUsingMapOption(true, map[string]string{"x": "y"})
	_ = Variant("x").ReplaceUsingMapOption(false, map[string]string{"x": "y"})

	// Remove / RemoveMany / RemoveManyBySplitting
	if Variant("abcabc").Remove("a").Value() != "bcbc" {
		t.Error("Remove wrong")
	}
	_ = Variant("abc").RemoveMany("a", "c")
	_ = Variant("a,b,c").RemoveManyBySplitting(",", "b")

	// Splits
	if len(Variant("a,b,c").SplitBy(",")) != 3 {
		t.Error("SplitBy wrong")
	}
	_, _ = Variant("k=v").SplitKeyVal("=")
	_, _ = Variant(" k = v ").SplitKeyValTrim("=")
	_, _ = Variant("k=v").SplitKeyValue("=")
	_, _ = Variant(" k = v ").SplitKeyValueTrim("=")
	_, _ = Variant("k=v").SplitKeyValueAsType("=")
	if len(Variant(" a , b , , c ").SplitTrimmedNonEmpty(",")) != 3 {
		t.Error("SplitTrimmedNonEmpty wrong")
	}
	if len(Variant("a b c").SplitByWhitespace()) != 3 {
		t.Error("SplitByWhitespace wrong")
	}

	// Suffix / Prefix
	if Variant("hello").AddSuffixOnMissing("-x") != "hello-x" ||
		Variant("hello-x").AddSuffixOnMissing("-x") != "hello-x" {
		t.Error("AddSuffixOnMissing wrong")
	}
	if Variant("hello").AddPrefixOnMissing("p-") != "p-hello" ||
		Variant("p-hello").AddPrefixOnMissing("p-") != "p-hello" {
		t.Error("AddPrefixOnMissing wrong")
	}

	// Append / Prepend / Add / AddAnother / Join
	if Variant("a").Append(Variant("b")).Value() != "ab" {
		t.Error("Append wrong")
	}
	if Variant("b").Prepend(Variant("a")).Value() != "ab" {
		t.Error("Prepend wrong")
	}
	if Variant("b").PrependString("a").Value() != "ab" {
		t.Error("PrependString wrong")
	}
	if Variant("a").Add(" b").Value() != "a b" {
		t.Error("Add wrong")
	}
	if Variant("a").AddAnother(Variant("b")).Value() != "ab" {
		t.Error("AddAnother wrong")
	}
	_ = Variant("a").Join(Variant("b"), Variant("c"))
	_ = Variant("a").JoinStrings("b", "c")
	if Variant("a").AppendIf(true, Variant("b")).Value() != "ab" ||
		Variant("a").AppendIf(false, Variant("b")).Value() != "a" {
		t.Error("AppendIf wrong")
	}
	if Variant("a").AppendStringIf(true, "b").Value() != "ab" ||
		Variant("a").AppendStringIf(false, "b").Value() != "a" {
		t.Error("AppendStringIf wrong")
	}
	if Variant("b").PrependIf(true, Variant("a")).Value() != "ab" ||
		Variant("b").PrependIf(false, Variant("a")).Value() != "b" {
		t.Error("PrependIf wrong")
	}
	if Variant("b").PrependStringIf(true, "a").Value() != "ab" ||
		Variant("b").PrependStringIf(false, "a").Value() != "b" {
		t.Error("PrependStringIf wrong")
	}

	// Comparison & search
	if !v.IsContains("ell") || !v.IsStartsWith("he") || !v.IsEndsWith("lo") {
		t.Error("contains/starts/ends wrong")
	}
	if !v.HasPrefix("he") || !v.HasSuffix("lo") {
		t.Error("HasPrefix/Suffix wrong")
	}
	if v.Index("ll") != 2 || v.LastIndexOf("l") != 3 {
		t.Error("Index/LastIndexOf wrong")
	}
	if !v.IsEqual("hello") || !v.IsEqualAnother(Variant("hello")) || !v.Is(Variant("hello")) {
		t.Error("equal wrong")
	}
	if !Variant("b").IsGreater("a") || !Variant("b").IsGreaterEqual("b") ||
		!Variant("a").IsLess("b") || !Variant("a").IsLessEqual("a") {
		t.Error("compare wrong")
	}
	if !v.IsNameEqual("hello") || !v.IsAnyNamesOf("a", "hello") || v.IsAnyNamesOf("a", "b") {
		t.Error("name eq wrong")
	}

	// Booleans
	if Variant("").HasAnyItem() || !Variant("a").HasAnyItem() {
		t.Error("HasAnyItem wrong")
	}
	_ = Variant("").OrEmpty(Variant("x"))
	_ = Variant("x").OrHasElement(Variant(""))
	_ = Variant("x").AndHasElement(Variant("y"))
	_ = Variant("").AndIsEmpty(Variant(""))

	// Substring / Split helpers
	if v.SafeSubString(0, 5).Value() != "hello" || v.SafeSubString(0, 999).Value() != "hello" {
		t.Error("SafeSubString wrong")
	}
	if v.SafeSubString(-1, 3).Value() != "hel" || v.SafeSubString(99, 100).Value() != "" {
		t.Error("SafeSubString edge wrong")
	}
	_ = v.SafeSubStringStart(2)
	_ = v.SafeSubStringEnd(3)
	_, _ = v.SafeSplit(2)
	_ = v.SimpleStringOnce(true)
	_ = v.SimpleStringOnce(false)

	// Int / Version
	if Variant("42").Integer() != 42 {
		t.Error("Integer wrong")
	}
	if i, err := Variant("42").ConvInteger(); err != nil || i != 42 {
		t.Error("ConvInteger wrong")
	}
	if val, ok := Variant("7").IntegerDefaultVal(0); !ok || val != 7 {
		t.Error("IntegerDefaultVal wrong")
	}
	if val, ok := Variant("xx").IntegerDefaultVal(99); ok || val != 99 {
		t.Error("IntegerDefaultVal default wrong")
	}
	if Variant("42").IntType().Value() != 42 {
		t.Error("IntType wrong")
	}
	_ = Variant("1.0.0").Version()

	// JSON round-trip
	data, err := json.Marshal(v)
	if err != nil {
		t.Errorf("Marshal: %v", err)
	}
	var got Variant
	if err := json.Unmarshal(data, &got); err != nil || got != v {
		t.Errorf("Unmarshal: %v got=%q", err, got)
	}

	// AsBasicEnumer / ToPtr / NameUsingMap / HasInAliasMap / ToErr
	_ = v.AsBasicEnumer()
	_ = v.ToPtr()
	_ = v.NameUsingMap(map[Variant]string{v: "h"})
	_ = v.HasInAliasMap(map[string]Variant{"hello": v}, v)
	if Variant("").ToErr() != nil {
		t.Error("ToErr empty should be nil")
	}
	if Variant("boom").ToErr() == nil {
		t.Error("ToErr non-empty should be err")
	}

	// Stub / boilerplate
	if v.ValueUInt16() != 0 {
		t.Error("ValueUInt16 stub")
	}
	_ = v.AllNameValues()
	_ = v.IntegerEnumRanges()
	_, _ = v.MinMaxAny()
	_ = v.MinValueString()
	_ = v.MaxValueString()
	_ = v.MinInt()
	_ = v.MaxInt()
	_ = v.RangesDynamicMap()
	_ = v.NameValue()
	_ = v.RangeNamesCsv()
	_ = v.EnumType()
	_ = v.TypeName()
	_ = v.Format("{name}")
	_ = v.ValueString()
	_ = v.ToNumberString()

	// IsValid / IsInvalid (string predicate path)
	if !Variant("Invalid").IsInvalid() || Variant("ok").IsInvalid() {
		t.Error("IsInvalid wrong")
	}
	if !Variant("ok").IsValid() || Variant("Invalid").IsValid() {
		t.Error("IsValid wrong")
	}

	// Numeric value helpers (must be parseable for Int8/16/32 paths)
	num := Variant("42")
	if num.ValueInt() != 42 || num.ValueInt8() != 42 || num.ValueInt16() != 42 || num.ValueInt32() != 42 {
		t.Error("Value Int* wrong")
	}
	if Variant("400").ValueInt8() == 0 { // out-of-range guard returns InvalidIndex
		// just exercise the branch — value doesn't matter
		_ = Variant("400").ValueInt8()
	}
	_ = Variant("99999999").ValueInt16()
	if b, ok := Variant("42").ByteType(); !ok || b.Value() != 42 {
		t.Error("ByteType wrong")
	}
	if _, ok := Variant("xx").ByteType(); ok {
		t.Error("ByteType invalid should fail")
	}
	if Variant("42").ValueByte() != 42 {
		t.Error("ValueByte wrong")
	}

	// ToByteUsingMap / ToByteUsingMapValidationErr / ToIntUsingMapValidationErr
	bm := map[string]byte{"x": 7}
	if val, ok := Variant("x").ToByteUsingMap(bm); !ok || val != 7 {
		t.Error("ToByteUsingMap hit wrong")
	}
	if _, ok := Variant("y").ToByteUsingMap(bm); ok {
		t.Error("ToByteUsingMap miss wrong")
	}
	if _, ok := Variant("x").ToByteUsingMap(nil); ok {
		t.Error("ToByteUsingMap empty wrong")
	}
	if _, err := Variant("x").ToByteUsingMapValidationErr(bm); err != nil {
		t.Errorf("ToByteUsingMapValidationErr hit: %v", err)
	}
	if _, err := Variant("y").ToByteUsingMapValidationErr(bm); err == nil {
		t.Error("ToByteUsingMapValidationErr miss should err")
	}
	if _, err := Variant("x").ToByteUsingMapValidationErr(nil); err == nil {
		t.Error("ToByteUsingMapValidationErr empty should err")
	}
	im := map[string]int{"x": 9}
	if _, err := Variant("x").ToIntUsingMapValidationErr(im); err != nil {
		t.Errorf("ToIntUsingMapValidationErr hit: %v", err)
	}
	if _, err := Variant("y").ToIntUsingMapValidationErr(im); err == nil {
		t.Error("ToIntUsingMapValidationErr miss should err")
	}
	if _, err := Variant("x").ToIntUsingMapValidationErr(nil); err == nil {
		t.Error("ToIntUsingMapValidationErr empty should err")
	}

	// Constructors
	if New("z") != "z" {
		t.Error("New wrong")
	}
	if NewUsingInteger(7) != "7" {
		t.Error("NewUsingInteger wrong")
	}
	_ = NewFileReader("/tmp/x")
	_ = v.FileReader()
}
