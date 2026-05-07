package instructiontype

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// AL2: Broad coverage uplift for instructiontype (was 48.2%).
// Exercises every Is*() predicate plus value/range/format surface.
func Test_InstructionType_Coverage2(t *testing.T) {
	all := []Variant{
		Invalid, Scoping, DependsOn, InstallPackages, OsServices,
		EnvironmentVariables, EnvironmentPaths, Ssl, CronTabs, Ethernet,
		PowerDns, MySql, PostgreSql, PhpMyAdmin, PhpPgAdmin,
		Nginx, Apache, PureFtp, FtpUser, Compress, Decompress,
		DownloadDecompress, OsUsersManage, OsGroupsManage, OsServicesCreate,
		LinuxServiceCreate, Ssh, PathModifiers, FileSystem,
		ScriptsCollection, Firewall, FirewallIpTables, CliInstructions,
		InstructionsExecuteAfterReboot, Verify,
	}

	Convey("instructiontype — every predicate runs without panic", t, func() {
		for _, v := range all {
			_ = v.IsUninitialized()
			_ = v.IsScoping()
			_ = v.IsDependsOn()
			_ = v.IsInstallPackages()
			_ = v.IsOsServices()
			_ = v.IsEnvironmentVariables()
			_ = v.IsEnvironmentPaths()
			_ = v.IsSsl()
			_ = v.IsCronTabs()
			_ = v.IsEthernet()
			_ = v.IsPowerDns()
			_ = v.IsMySql()
			_ = v.IsPostgreSql()
			_ = v.IsPhpMyAdmin()
			_ = v.IsPhpPgAdmin()
			_ = v.IsNginx()
			_ = v.IsApache()
			_ = v.IsPureFtp()
			_ = v.IsCompress()
			_ = v.IsDecompress()
			_ = v.IsDownloadDecompress()
			_ = v.IsOsUsersManage()
			_ = v.IsOsGroupsManage()
			_ = v.IsOsServicesCreate()
			_ = v.IsLinuxServiceCreate()
			_ = v.IsSsh()
			_ = v.IsPathModifiers()
			_ = v.IsFileSystem()
			_ = v.IsScriptsCollection()
			_ = v.IsFirewall()
			_ = v.IsFirewallIpTables()
			_ = v.IsCliInstructions()
			_ = v.IsInstructionsExecuteAfterReboot()
			_ = v.IsVerify()
			_ = v.IsInvalid()
			_ = v.IsValid()

			So(v.ValueByte(), ShouldEqual, byte(v))
			So(v.ValueInt(), ShouldEqual, int(v))
			So(v.ValueInt8(), ShouldEqual, int8(v))
			So(v.ValueInt16(), ShouldEqual, int16(v))
			So(v.ValueInt32(), ShouldEqual, int32(v))
			So(v.ValueUInt16(), ShouldEqual, uint16(v))
			So(v.IsByteValueEqual(byte(v)), ShouldBeTrue)
			So(v.IsValueEqual(byte(v)), ShouldBeTrue)
			So(v.IsAnyValuesEqual(byte(v)), ShouldBeTrue)
			So(v.IsNameEqual(v.Name()), ShouldBeTrue)
		}

		// Targeted positive predicate hits
		So(Scoping.IsScoping(), ShouldBeTrue)
		So(Nginx.IsNginx(), ShouldBeTrue)
		So(Apache.IsApache(), ShouldBeTrue)
		So(MySql.IsMySql(), ShouldBeTrue)
		So(Verify.IsVerify(), ShouldBeTrue)
		So(Invalid.IsInvalid(), ShouldBeTrue)
		So(Verify.IsValid(), ShouldBeTrue)
	})

	Convey("instructiontype — surface", t, func() {
		So(len(Scoping.AllNameValues()), ShouldBeGreaterThan, 0)
		So(Scoping.MaxInt(), ShouldBeGreaterThanOrEqualTo, Scoping.MinInt())
		So(Scoping.MinValueString(), ShouldNotBeBlank)
		So(Scoping.MaxValueString(), ShouldNotBeBlank)
		So(Scoping.RangesDynamicMap(), ShouldNotBeNil)
		So(Scoping.IntegerEnumRanges(), ShouldNotBeNil)
		minA, maxA := Scoping.MinMaxAny()
		So(minA, ShouldNotBeNil)
		So(maxA, ShouldNotBeNil)
		So(Scoping.EnumType(), ShouldNotBeNil)
		So(Scoping.Format("{name}"), ShouldNotBeBlank)
		So(Scoping.ToNumberString(), ShouldNotBeBlank)
		So(Scoping.OnlySupportedErr("Scoping"), ShouldNotBeNil)
		So(Scoping.OnlySupportedMsgErr("ctx", "Scoping"), ShouldNotBeNil)
		So(Scoping.IsAnyNamesOf("Scoping"), ShouldBeTrue)
	})

	Convey("instructiontype — JSON round-trip", t, func() {
		original := Nginx
		bs, err := json.Marshal(&original)
		So(err, ShouldBeNil)
		var got Variant
		So(json.Unmarshal(bs, &got), ShouldBeNil)
		So(got, ShouldEqual, original)
	})
}
