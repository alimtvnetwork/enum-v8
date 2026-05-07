package sqljointype

import (
	"encoding/json"
	"testing"
)

// AL2-02 Batch B coverage suite for sqljointype.

func TestSqlJoin_PredicatesAndSqlSyntax(t *testing.T) {
	cases := []struct {
		v   Variant
		sql string
	}{
		{Default, "JOIN"},
		{Invalid, ""},
		{Join, "JOIN"},
		{Inner, "INNER JOIN"},
		{Left, "LEFT JOIN"},
		{Right, "RIGHT JOIN"},
		{FullOuter, "FULL OUTER JOIN"},
		{Cross, "CROSS JOIN"},
	}
	for _, c := range cases {
		if got := c.v.ToSqlName(); got != c.sql {
			t.Errorf("ToSqlName(%v) = %q want %q", c.v, got, c.sql)
		}
	}
	if !Inner.IsValid() || Invalid.IsValid() {
		t.Error("IsValid mismatch")
	}
	if !Inner.IsValidLogically() || Invalid.IsValidLogically() {
		t.Error("IsValidLogically mismatch")
	}
	if !Inner.IsInnerJoinLogically() || !Default.IsInnerJoinLogically() ||
		Left.IsInnerJoinLogically() {
		t.Error("IsInnerJoinLogically mismatch")
	}
	if !Left.IsOuterJoinLogically() || !Right.IsOuterJoinLogically() ||
		!FullOuter.IsOuterJoinLogically() || Inner.IsOuterJoinLogically() {
		t.Error("IsOuterJoinLogically mismatch")
	}
	if !Join.IsJoin() || !Right.IsRight() || !Right.IsRightOuter() ||
		!Left.IsLeft() || !Left.IsLeftOuter() || !Cross.IsCross() ||
		!FullOuter.IsFullOuter() || !Invalid.IsInvalid() {
		t.Error("IsX predicate mismatch")
	}
	if !Inner.IsAnyOf(Inner, Left) || Inner.IsAnyOf(Left, Right) {
		t.Error("IsAnyOf mismatch")
	}
	if !Inner.IsByteValueEqual(byte(Inner)) || !Inner.IsValueEqual(byte(Inner)) ||
		!Inner.IsNameEqual("Inner") || !Inner.IsAnyValuesEqual(byte(Inner)) ||
		!Inner.IsAnyNamesOf("Inner", "Left") {
		t.Error("value-equal mismatch")
	}
}

func TestSqlJoin_NumericAccessors(t *testing.T) {
	v := Inner
	_ = v.ValueInt()
	_ = v.ValueInt8()
	_ = v.ValueInt16()
	_ = v.ValueInt32()
	_ = v.ValueUInt16()
	_ = v.ValueByte()
	_ = v.Value()
	_ = v.MinByte()
	_ = v.MaxByte()
	if v.MinInt() > v.MaxInt() {
		t.Error("MinInt > MaxInt")
	}
	if v.MinValueString() == "" || v.MaxValueString() == "" {
		t.Error("min/max value string empty")
	}
	if min, max := v.MinMaxAny(); min == nil || max == nil {
		t.Error("MinMaxAny nil")
	}
	if len(v.RangesByte()) == 0 || len(v.IntegerEnumRanges()) == 0 ||
		len(v.AllNameValues()) == 0 || len(v.RangesDynamicMap()) == 0 {
		t.Error("ranges empty")
	}
	if v.RangeNamesCsv() == "" || v.NameValue() == "" || v.TypeName() == "" ||
		v.ValueString() == "" || v.ToNumberString() == "" || v.String() == "" {
		t.Error("string accessor empty")
	}
	if v.Format("v=%v") == "" {
		t.Error("Format empty")
	}
	if v.EnumType() == nil {
		t.Error("EnumType nil")
	}
	_ = v.OnlySupportedErr("Inner")
	_ = v.OnlySupportedMsgErr("ctx", "Inner")
	_ = RangesInvalidErr()
}

func TestSqlJoin_JsonRoundTrip(t *testing.T) {
	for _, x := range []Variant{Default, Inner, Left, Right, FullOuter, Cross} {
		data, err := json.Marshal(x)
		if err != nil {
			t.Fatalf("Marshal %v: %v", x, err)
		}
		var got Variant
		if err := json.Unmarshal(data, &got); err != nil {
			t.Fatalf("Unmarshal %v: %v", x, err)
		}
		if got != x {
			t.Errorf("round-trip: got %v want %v", got, x)
		}
	}
}

func TestSqlJoin_CompileAndJoinClause(t *testing.T) {
	left := TableWithColumn{TableName: "users", ColumnName: "id"}
	right := TableWithColumn{TableName: "orders", ColumnName: "user_id"}
	if left.TableWithField() == "" {
		t.Error("TableWithField empty")
	}
	if left.Clone() != left {
		t.Error("Clone mismatch")
	}
	jc := Inner.JoinClause(left, right)
	if jc.JoinType != Inner {
		t.Error("JoinClause.JoinType mismatch")
	}
	if jc.Compile() == "" || jc.String() == "" {
		t.Error("Compile/String empty")
	}
	jc2 := jc.SetLeftClone(TableWithColumn{TableName: "u2", ColumnName: "id"})
	if jc2.Left.TableName != "u2" {
		t.Error("SetLeftClone failed")
	}
	jc3 := jc.SetRightClone(TableWithColumn{TableName: "o2", ColumnName: "uid"})
	if jc3.Right.TableName != "o2" {
		t.Error("SetRightClone failed")
	}
	if jc.Clone() != jc {
		t.Error("JoinerClause Clone mismatch")
	}
	if Inner.Compile(left, right) == "" {
		t.Error("Variant.Compile empty")
	}
}

func TestSqlJoin_PtrAndBinders(t *testing.T) {
	v := Inner
	if v.ToPtr() == nil || (*v.ToPtr()).ToSimple() != v {
		t.Error("ToPtr mismatch")
	}
	var nilP *Variant
	if nilP.ToSimple() != Inner {
		t.Error("ToSimple(nil) should be Inner")
	}
	if v.Json().HasError() {
		t.Errorf("Json: %v", v.Json().Error)
	}
	if v.JsonPtr() == nil || v.AsJsoner() == nil || v.AsJsonContractsBinder() == nil ||
		v.AsJsonMarshaller() == nil || v.AsBasicByteEnumContractsBinder() == nil ||
		v.AsBasicEnumContractsBinder() == nil {
		t.Error("As* binder nil")
	}
	_ = Min()
	_ = Max()
}
