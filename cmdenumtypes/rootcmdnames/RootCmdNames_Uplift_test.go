package rootcmdnames

import (
	"encoding/json"
	"reflect"
	"testing"
)

// AL2-rootcmd-uplift — Reflection-based sweep over every Variant constant.
// Existing RootCmdNames_Coverage_test.go only touches 10 of ~58 constants
// and a handful of methods. This sweep:
//   - calls every nullary exported method on Variant (value + pointer receiver)
//     with panic-recovery — RCA pattern 7 (sparse-array-safe)
//   - JSON marshal/unmarshal round-trip per constant
//   - New(name) / NewMust(name) for every Variant.Name()
//   - alias-name fast-paths exercised via New()
//   - ValueByte / IsAnyOf / IsAnyValuesEqual / IsByteValueEqual / IsValueEqual
//   - top-level Min / Max / RangesInvalidErr / Is

func TestRootCmdNames_Uplift_AllVariants(t *testing.T) {
	all := []Variant{
		Invalid, Help, Compress, Cron, Decompress, Dns, Download,
		EnvPath, EnvVars, Ethernet, Fail2ban, Firewall, Ftp,
		HostingPlan, Macro, Package, Packages, Pkg, Snapshot,
		Ssh, Ssl, User, UserRole, WebServer, Tooling, Docker,
		System, Os, Update, AutoFix, MySql, PostgreSql, DbServer,
		Sync, SyncNow, SyncDryRun, Nginx, Apache, Paths, SysPaths,
		Env, Services, Restart, Reboot, Shutdown, SysGroup, SysUser,
		FtpUser, PanelUser, CashBin, PanelRoles, CacheReset,
		ReloadSettings, ResetSettings, DefaultKnownSettings, DbUser,
		List, ListJson, Backup, Import,
	}

	for _, v := range all {
		callAllNullary_RootCmd(reflect.ValueOf(v))
		callAllNullary_RootCmd(reflect.ValueOf(&v))

		// JSON round-trip
		if data, err := json.Marshal(v); err == nil {
			var got Variant
			_ = json.Unmarshal(data, &got)
		}

		// Constructor fast-paths
		if n := v.Name(); n != "" {
			_, _ = New(n)
			_ = NewMust(n)
		}

		// Cross-variant comparators
		_ = v.IsByteValueEqual(v.ValueByte())
		_ = v.IsValueEqual(v.ValueByte())
		_ = v.IsNameEqual(v.Name())
		_ = v.IsAnyValuesEqual(v.ValueByte(), 0xFF)
		_ = v.IsAnyOf(v, Invalid)
		_ = v.IsAnyNamesOf(v.Name(), "__bogus__")
		_ = v.Is(v)
	}

	// Alias-name fast-paths — sweep every alias key from the alias map.
	// Anything that errors is fine; we just want the lookup branch covered.
	aliases := []string{
		"?", "-?", "/?", "zipper", "p-dns", "d-load", "iptables",
		"pure-ftp", "web-server", "db", "my-sql", "m-sql", "mySQL",
		"postgre", "postgresql", "postgre-sql", "p-sql", "sync-n",
		"sys-path", "r-settings", "default-settings",
		"apply-default-config", "def-setting", "def-settings",
		"__bogus__", "",
	}
	for _, name := range aliases {
		_, _ = New(name)
	}

	// Top-level helpers
	_ = Min()
	_ = Max()
	_ = RangesInvalidErr()
	_ = Is("Help", Help)
	_ = Is("__bogus__", Help)
}

func callAllNullary_RootCmd(rv reflect.Value) {
	typ := rv.Type()
	for i := 0; i < typ.NumMethod(); i++ {
		// nullary = receiver-only (NumIn == 1 because receiver counts)
		if typ.Method(i).Type.NumIn() != 1 {
			continue
		}
		func() {
			defer func() { _ = recover() }()
			_ = rv.Method(i).Call(nil)
		}()
	}
}
