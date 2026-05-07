package dbdrivertype

import (
	"encoding/json"
	"strings"
	"testing"
)

// AL2-07 bespoke coverage suite for dbdrivertype connection-string surface.

func TestDbDriverType_NewAndAccessors(t *testing.T) {
	for _, name := range []string{"Sqlite", "Redis", "MySql", "MariaDb", "PostgreSql", "MongoDb", "Oracle", "Firebird", "HSqlDb", "Json", "Yaml"} {
		v, err := New(name)
		if err != nil {
			t.Errorf("New(%q): %v", name, err)
		}
		if v.Name() != name {
			t.Errorf("Name mismatch: %q != %q", v.Name(), name)
		}
		_ = NewMust(name)
	}
	if _, err := New("__bogus__"); err == nil {
		t.Error("New bogus should fail")
	}
	if Min() != Invalid || Max() != Protobuf {
		t.Errorf("Min/Max wrong: %v %v", Min(), Max())
	}
	if RangesInvalidErr() == nil {
		t.Error("RangesInvalidErr nil")
	}

	if !MySql.IsSqlDb() || !PostgreSql.IsSqlDb() || MongoDb.IsSqlDb() {
		t.Error("IsSqlDb wrong")
	}
	if !MongoDb.IsNoSql() || !Json.IsNoSql() || MySql.IsNoSql() {
		t.Error("IsNoSql wrong")
	}
	if !MongoDb.IsMongoDb() || !Oracle.IsOracle() {
		t.Error("predicate wrong")
	}

	if MySql.DefaultPort() != 3306 || PostgreSql.DefaultPort() != 5432 ||
		MongoDb.DefaultPort() != 27017 {
		t.Error("DefaultPort wrong")
	}
	if p, has := PostgreSql.DefaultPortStatus(); !has || p != 5432 {
		t.Error("DefaultPortStatus wrong")
	}
	if _, has := Sqlite.DefaultPortStatus(); has {
		t.Error("Sqlite should have no default port")
	}

	if !Is("MySql", MySql) || Is("__bogus__", MySql) {
		t.Error("Is wrong")
	}

	v := MySql
	if v.ValueUInt16() != uint16(v) || len(v.AllNameValues()) == 0 ||
		len(v.IntegerEnumRanges()) == 0 {
		t.Error("accessors broken")
	}
	if v.MinValueString() == "" || v.MaxValueString() == "" || v.TypeName() == "" {
		t.Error("string accessors empty")
	}
	if min, max := v.MinMaxAny(); min == nil || max == nil {
		t.Error("MinMaxAny nil")
	}

	data, err := json.Marshal(v)
	if err != nil || len(data) == 0 {
		t.Errorf("Marshal: %v", err)
	}
	var got Variant
	if err := json.Unmarshal(data, &got); err != nil {
		t.Errorf("Unmarshal: %v", err)
	}
	_ = v.OnlySupportedErr("MySql")
	_ = v.OnlySupportedMsgErr("ctx", "MySql")
}

func TestDbDriverType_ConnectionCompile(t *testing.T) {
	conn := Connection{
		DbType: MySql,
		ConnectionOptions: ConnectionOptions{
			Host:     "127.0.0.1",
			Port:     "3306",
			User:     "root",
			Password: "secret",
			DbName:   "appdb",
			Options:  "?parseTime=true",
		},
	}
	if !conn.IsValid() || conn.IsInvalidConnectionString() {
		t.Error("MySql connection should be valid")
	}
	if !conn.HasConnectionString() {
		t.Error("MySql connection should have format")
	}
	if conn.ConnectionStringFormat() == "" || conn.ConnectionStringAllDbFormat() == "" {
		t.Error("formats empty")
	}
	out, err := conn.Compile()
	if err != nil || !strings.Contains(out, "127.0.0.1") || !strings.Contains(out, "appdb") {
		t.Errorf("Compile: out=%q err=%v", out, err)
	}
	out2, err := conn.CompileUsingParams("h", "1", "d", "u", "p", "")
	if err != nil || !strings.Contains(out2, "h") {
		t.Errorf("CompileUsingParams: out=%q err=%v", out2, err)
	}
	out3, err := conn.CompileUsingParamsNoOptions("h", "1", "d", "u", "p")
	if err != nil || out3 == "" {
		t.Errorf("CompileUsingParamsNoOptions: %v", err)
	}
	if conn.CompileUsingConnectionFormat("{ip}:{port}") != "127.0.0.1:3306" {
		t.Error("CompileUsingConnectionFormat wrong")
	}

	// invalid driver — Oracle is not in connectionStringFormatMap
	bad := Connection{DbType: Oracle}
	if bad.IsValid() || !bad.IsInvalidConnectionString() {
		t.Error("Oracle should be invalid for connection string")
	}
	if _, err := bad.Compile(); err == nil {
		t.Error("Compile should error for Oracle")
	}
	if _, err := bad.CompileUsingParams("h", "1", "d", "u", "p", ""); err == nil {
		t.Error("CompileUsingParams should error for Oracle")
	}
}

func TestDbDriverType_ConnectionCompiler(t *testing.T) {
	c := MySql.Connection()
	if !c.HasConnectionString() || c.IsInvalidConnectionString() || !c.IsValidConnectionString() {
		t.Error("MySql.Connection() validity wrong")
	}
	if c.Format() == "" || c.AllDbFormat() == "" {
		t.Error("compiler formats empty")
	}
	out, err := c.CompileUsingConnection(ConnectionOptions{
		Host: "h", Port: "1", User: "u", Password: "p", DbName: "d",
	})
	if err != nil || out == "" {
		t.Errorf("CompileUsingConnection: %v", err)
	}

	bad := Oracle.Connection()
	if bad.HasConnectionString() || !bad.IsInvalidConnectionString() {
		t.Error("Oracle compiler should be invalid")
	}
}

func TestDbDriverType_ConnectionOptionsMaps(t *testing.T) {
	opts := ConnectionOptions{
		Host: "h", Port: "1", User: "u", Password: "p",
		DbName: "d", Options: "o",
	}
	m := opts.CreateMap()
	if m["{db}"] != "d" || m["{ip}"] != "h" || m["{?options}"] != "o" {
		t.Error("CreateMap wrong")
	}
	m2 := opts.CreateMapUsingParams("H", "2", "D", "U", "P", "O")
	if m2["{db}"] != "D" || m2["{ip}"] != "H" {
		t.Error("CreateMapUsingParams wrong")
	}
	m3 := opts.CreateMapUsingParamsNoOptions("H", "2", "D", "U", "P")
	if m3["{?options}"] != "" {
		t.Error("CreateMapUsingParamsNoOptions should clear options")
	}
	if opts.CompileUsingConnectionFormat("{ip}:{port}") != "h:1" {
		t.Error("CompileUsingConnectionFormat wrong")
	}
}
