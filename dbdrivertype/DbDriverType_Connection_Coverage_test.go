package dbdrivertype

import (
	"strings"
	"testing"
)

// AL2-07 — Bespoke Connection / ConnectionOptions / connectionStringCompiler suite.
// Exercises the three business-logic source files that DbDriverType_Coverage_test.go
// (Variant-only) does not reach. For each Variant with a known connection-string
// format, build a representative Connection, compile, and assert the rendered
// substrings match the expected DSN shape.
func Test_DbDriverType_ConnectionCompile(t *testing.T) {
	cases := []struct {
		dbType    Variant
		mustHave  []string
	}{
		{Sqlite, []string{"/tmp/test.db"}},
		{Redis, []string{"redis://", "10.0.0.1", "6379"}},
		{MySql, []string{"alice", "secret", "tcp(", "10.0.0.1:3306)", "/payments"}},
		{PostgreSql, []string{"host=10.0.0.1", "port=5432", "user=alice", "password=secret", "dbname=payments"}},
		{MicrosoftSqlServer, []string{"sqlserver://alice:secret@10.0.0.1:1433", "database=payments"}},
		{MongoDb, []string{"mongodb://", "alice", "secret", "10.0.0.1", "27017", "/payments"}},
	}

	for _, tc := range cases {
		t.Run(tc.dbType.Name(), func(t *testing.T) {
			conn := Connection{
				DbType: tc.dbType,
				ConnectionOptions: ConnectionOptions{
					Host:     "10.0.0.1",
					Port:     portFor(tc.dbType),
					DbName:   dbNameFor(tc.dbType),
					User:     "alice",
					Password: "secret",
					Options:  "?ssl=true",
				},
			}

			if !conn.HasConnectionString() {
				t.Fatalf("%s should have connection string", tc.dbType.Name())
			}
			if conn.IsInvalidConnectionString() {
				t.Fatalf("%s wrongly reported invalid", tc.dbType.Name())
			}
			if !conn.IsValid() {
				t.Fatalf("%s should be valid", tc.dbType.Name())
			}

			_ = conn.ConnectionStringFormat()
			_ = conn.ConnectionStringAllDbFormat()
			_ = conn.CreateMap()
			_ = conn.CreateMapUsingParams("h", "p", "d", "u", "pw", "o")
			_ = conn.CreateMapUsingParamsNoOptions("h", "p", "d", "u", "pw")

			compiled, err := conn.Compile()
			if err != nil {
				t.Fatalf("compile %s: %v", tc.dbType.Name(), err)
			}
			for _, sub := range tc.mustHave {
				if !strings.Contains(compiled, sub) {
					t.Errorf("%s compiled %q missing %q", tc.dbType.Name(), compiled, sub)
				}
			}

			withParams, err := conn.CompileUsingParams("h", "p", "d", "u", "pw", "")
			if err != nil {
				t.Fatalf("compile-params %s: %v", tc.dbType.Name(), err)
			}
			_ = withParams

			noOpts, err := conn.CompileUsingParamsNoOptions("h", "p", "d", "u", "pw")
			if err != nil {
				t.Fatalf("compile-no-opts %s: %v", tc.dbType.Name(), err)
			}
			_ = noOpts

			_ = conn.CompileUsingConnectionFormat(conn.ConnectionStringFormat())

			// Bare ConnectionOptions surface
			opts := conn.ConnectionOptions
			_ = opts.CreateMap()
			_ = opts.CreateMapUsingParams("h", "p", "d", "u", "pw", "o")
			_ = opts.CreateMapUsingParamsNoOptions("h", "p", "d", "u", "pw")
			_ = opts.CompileUsingConnectionFormat(conn.ConnectionStringFormat())
		})
	}
}

func Test_DbDriverType_Connection_InvalidDbType(t *testing.T) {
	conn := Connection{DbType: Invalid}
	if conn.HasConnectionString() {
		t.Fatal("Invalid should not have connection string")
	}
	if !conn.IsInvalidConnectionString() {
		t.Fatal("Invalid should report invalid")
	}
	if conn.IsValid() {
		t.Fatal("Invalid should not be valid")
	}
	_, err := conn.Compile()
	if err == nil {
		t.Fatal("Invalid Compile should error")
	}
	_, err = conn.CompileUsingParams("h", "p", "d", "u", "pw", "o")
	if err == nil {
		t.Fatal("Invalid CompileUsingParams should error")
	}
	_, err = conn.CompileUsingParamsNoOptions("h", "p", "d", "u", "pw")
	if err == nil {
		t.Fatal("Invalid CompileUsingParamsNoOptions should error")
	}
}

func Test_DbDriverType_ConnectionStringCompiler(t *testing.T) {
	for _, dbType := range []Variant{Sqlite, Redis, MySql, PostgreSql, MicrosoftSqlServer, MongoDb} {
		c := connectionStringCompiler{dbType: dbType}
		if !c.HasConnectionString() {
			t.Errorf("%s compiler missing format", dbType.Name())
		}
		if c.IsInvalidConnectionString() {
			t.Errorf("%s compiler reported invalid", dbType.Name())
		}
		if !c.IsValidConnectionString() {
			t.Errorf("%s compiler not valid", dbType.Name())
		}
		_ = c.Format()
		_ = c.AllDbFormat()

		conn := Connection{
			DbType: dbType,
			ConnectionOptions: ConnectionOptions{
				Host: "h", Port: "1", User: "u", Password: "pw", DbName: "d", Options: "o",
			},
		}
		_, _ = c.CompileUsingConnection(conn)
		_, _ = c.CompileUsingParams("h", "1", "d", "u", "pw", "o")
		_, _ = c.CompileUsingAllDbConnectionFormat(conn)
		_ = c.CompileUsingConnectionFormat(c.Format(), conn)
		_, _ = c.CompileUsingParamsOptions(ConnectionOptions{Host: "h", Port: "1", DbName: "d", User: "u", Password: "pw"})
		_, _ = c.CompileUsingMap(map[string]string{"{db}": "d"})
		_ = c.FormatCompileUsingMap(c.Format(), map[string]string{"{db}": "d"})
		_ = c.CompileUsingMapMust(map[string]string{"{db}": "d"})
	}

	bad := connectionStringCompiler{dbType: Invalid}
	if bad.HasConnectionString() {
		t.Fatal("Invalid compiler should not have format")
	}
	if !bad.IsInvalidConnectionString() {
		t.Fatal("Invalid compiler should be invalid")
	}
	if bad.IsValidConnectionString() {
		t.Fatal("Invalid compiler should not be valid")
	}
	_, err := bad.CompileUsingMap(map[string]string{})
	if err == nil {
		t.Fatal("Invalid compiler CompileUsingMap should error")
	}
}

func portFor(v Variant) string {
	if p, ok := defaultDbPortsMap[v]; ok {
		return itoa(p)
	}
	return ""
}

func dbNameFor(v Variant) string {
	if v == Sqlite {
		return "/tmp/test.db"
	}
	return "payments"
}

func itoa(p uint16) string {
	if p == 0 {
		return ""
	}
	digits := []byte{}
	for p > 0 {
		digits = append([]byte{byte('0' + p%10)}, digits...)
		p /= 10
	}
	return string(digits)
}
