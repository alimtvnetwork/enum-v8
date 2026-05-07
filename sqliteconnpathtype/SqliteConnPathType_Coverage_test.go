package sqliteconnpathtype

import (
	"encoding/json"
	"strings"
	"testing"
)

// AL2-02 Batch B coverage suite for sqliteconnpathtype.
// Covers SqliteConnectionOption (CreateMap/Compile/String), the
// sqliteConnectionCompiler value-type helpers (formats + CompileUsing*),
// and Variant accessors / JSON round-trip with the PI-005/PI-006 local
// overrides.

var allVariants = []Variant{
	Invalid, AllSqlitePath, AllWithTypeSqlitePath,
	AllWithTypeAndDynamicSqlitePath, AllWithTypeAndSequenceSqlitePath,
	PrefixSqlitePath, PrefixTypeSqlitePath, SpecificSqlitePath,
	DynamicSpecificSqlitePath, SequenceSpecificSqlitePath,
	DynamicSequenceSpecificSqlitePath,
}

func TestSqliteConn_OptionCompileAndMap(t *testing.T) {
	opt := SqliteConnectionOption{
		Root: "/srv", Prefix: "data", DbName: "users",
		TypeName: "user", Dynamic: "shard1", Sequence: "001",
	}
	m := opt.CreateMap()
	for _, k := range []string{"{root}", "{prefix}", "{db-name}", "{type}", "{dynamic}", "{sequence}"} {
		if _, ok := m[k]; !ok {
			t.Errorf("CreateMap missing %s", k)
		}
	}
	got := opt.Compile("{root}/{prefix}/{type}/{dynamic}-{sequence}-{db-name}.db")
	want := "/srv/data/user/shard1-001-users.db"
	if got != want {
		t.Errorf("Compile = %q want %q", got, want)
	}
	if s := opt.String(); !strings.Contains(s, "users") {
		t.Errorf("String missing db name: %s", s)
	}
}

func TestSqliteConn_CompilerFormats(t *testing.T) {
	c := sqliteConnectionCompiler{}
	formats := []string{
		c.AllFormat(), c.AllWithTypeFormat(), c.AllWithDynamicTypeFormat(),
		c.AllWithSequenceTypeFormat(), c.PrefixFormat(), c.PrefixTypeFormat(),
		c.DynamicSpecificFormat(), c.SequenceSpecificFormat(),
		c.DynamicSequenceSpecificFormat(), c.SpecificFormat(),
	}
	for i, f := range formats {
		if f == "" {
			t.Errorf("format[%d] empty", i)
		}
	}
	for _, v := range allVariants {
		if v == Invalid {
			continue
		}
		if c.PathFormat(v) == "" {
			t.Errorf("PathFormat(%v) empty", v)
		}
	}
}

func TestSqliteConn_CompilerCompileUsingHelpers(t *testing.T) {
	c := sqliteConnectionCompiler{}
	results := []string{
		c.CompileUsingAllWithParams("/r", "p", "db"),
		c.CompileUsingAllTypeWithParams("/r", "p", "t", "db"),
		c.CompileUsingAllTypeDynamicWithParams("/r", "p", "t", "d", "db"),
		c.CompileUsingAllTypeSequenceWithParams("/r", "p", "t", "s", "db"),
		c.CompileUsingPrefixWithParams("p", "t", "db"),
		c.CompileUsingPrefixTypeWithParams("p", "t", "db"),
		c.CompileUsingDynamicSpecificWithParams("d", "db"),
		c.CompileUsingSpecific("db"),
		c.CompileUsingDb("db"),
		c.CompileUsingSequenceSpecificWithParams("s", "db"),
		c.CompileUsingDynamicSequenceSpecificWithParams("d", "s", "db"),
		c.CompileUsingFormat("{db-name}.db", SqliteConnectionOption{DbName: "db"}),
		c.CompileUsingType(SpecificSqlitePath, SqliteConnectionOption{DbName: "x"}),
	}
	for i, r := range results {
		if r == "" || !strings.Contains(r, "db") {
			t.Errorf("compile result[%d] = %q", i, r)
		}
	}
}

func TestSqliteConn_VariantBasics(t *testing.T) {
	v := AllSqlitePath
	if v.String() != string(v) || v.Name() != string(v) {
		t.Error("String/Name mismatch")
	}
	if !v.IsValid() || Invalid.IsValid() {
		t.Error("IsValid wrong")
	}
	if !Invalid.IsInvalid() || v.IsInvalid() {
		t.Error("IsInvalid wrong")
	}
	if v.PathFormat() == "" {
		t.Error("PathFormat empty")
	}
	if v.ValueByte() != 0 || v.ValueUInt16() != 0 || v.ValueInt() != 0 ||
		v.ValueInt8() != 0 || v.ValueInt16() != 0 || v.ValueInt32() != 0 ||
		v.ToNumberString() != "0" {
		t.Error("numeric accessors should be zero for string-backed variant")
	}
	if v.ValueString() != string(v) {
		t.Error("ValueString mismatch")
	}
	if v.TypeName() == "" || v.RangeNamesCsv() == "" {
		t.Error("typename/csv empty")
	}
	if v.MinValueString() == "" || v.MaxValueString() == "" {
		t.Error("min/max value string empty")
	}
	if min, max := v.MinMaxAny(); min == nil || max == nil {
		t.Error("MinMaxAny nil")
	}
	if len(v.AllNameValues()) == 0 || len(v.StringRanges()) == 0 ||
		len(v.StringRangesPtr()) == 0 || len(v.RangesDynamicMap()) == 0 {
		t.Error("range collections empty")
	}
	_ = v.IntegerEnumRanges() // may be empty for string enum, just exercise
	if v.MaxInt() < v.MinInt() {
		t.Error("MaxInt < MinInt")
	}
	if v.EnumType() == nil {
		t.Error("EnumType nil")
	}
	if v.Format("%s") == "" {
		t.Error("Format empty")
	}
	if v.RangesInvalidMessage() == "" || v.RangesInvalidErr() == nil {
		t.Error("RangesInvalid* should be informational and non-empty")
	}
	if !v.IsValidRange() || v.IsInvalidRange() {
		t.Error("IsValidRange wrong")
	}
	if !v.IsAnyNamesOf(string(v), "other") {
		t.Error("IsAnyNamesOf should match")
	}
	if v.IsAnyNamesOf("nope") {
		t.Error("IsAnyNamesOf should not match")
	}
	if v.OnlySupportedErr(string(v)) == nil || v.OnlySupportedMsgErr("ctx", string(v)) == nil {
		t.Error("OnlySupportedErr should return informational error")
	}
}

func TestSqliteConn_VariantJsonRoundTrip(t *testing.T) {
	for _, v := range allVariants {
		data, err := json.Marshal(v)
		if err != nil {
			t.Fatalf("Marshal %v: %v", v, err)
		}
		var got Variant
		if err := json.Unmarshal(data, &got); err != nil {
			t.Errorf("Unmarshal %v: %v", v, err)
			continue
		}
		if got != v {
			t.Errorf("round-trip: got %v want %v", got, v)
		}
	}
	// Empty / "" / nil bytes fallback to MinValueString().
	var v Variant
	if err := v.UnmarshalJSON([]byte(`""`)); err != nil {
		t.Errorf("UnmarshalJSON empty: %v", err)
	}
	if string(v) == "" {
		t.Error("UnmarshalJSON empty: should fallback to MinValueString")
	}
	var v2 Variant
	if err := v2.UnmarshalJSON(nil); err != nil {
		t.Errorf("UnmarshalJSON nil: %v", err)
	}
	// Malformed quoted bytes path
	var v3 Variant
	_ = v3.UnmarshalJSON([]byte(`"__not_a_real_variant__"`))
}

func TestSqliteConn_VariantPtrAndBinders(t *testing.T) {
	v := AllSqlitePath
	if v.ToPtr() == nil {
		t.Error("ToPtr nil")
	}
	if v.AsBasicEnumContractsBinder() == nil ||
		v.AsStandardEnumerContractsBinder() == nil {
		t.Error("binders nil")
	}
	if v.AsJsonContractsBinder() == nil {
		t.Error("AsJsonContractsBinder nil")
	}
	jr := v.Json()
	if v.JsonPtr() == nil {
		t.Error("JsonPtr nil")
	}
	var dst Variant
	if err := dst.JsonParseSelfInject(&jr); err != nil {
		t.Errorf("JsonParseSelfInject: %v", err)
	}
	if _, err := v.UnmarshallEnumToValue([]byte(`"All"`)); err != nil {
		t.Errorf("UnmarshallEnumToValue: %v", err)
	}
}
