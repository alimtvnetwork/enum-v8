package instructiontype

import (
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:                        "Invalid",
		Scoping:                        "Scoping",
		DependsOn:                      "DependsOn",
		InstallPackages:                "InstallPackages",
		OsServices:                     "OsServices",
		EnvironmentVariables:           "EnvironmentVariables",
		EnvironmentPaths:               "EnvironmentPaths",
		Ssl:                            "Ssl",
		CronTabs:                       "CronTabs",
		Ethernet:                       "Ethernet",
		PowerDns:                       "PowerDns",
		MySql:                          "MySql",
		PostgreSql:                     "PostgreSql",
		PhpMyAdmin:                     "PhpMyAdmin",
		PhpPgAdmin:                     "PhpPgAdmin",
		Nginx:                          "Nginx",
		Apache:                         "Apache",
		PureFtp:                        "PureFtp",
		FtpUser:                        "FtpUser",
		Compress:                       "Compress",
		Decompress:                     "Decompress",
		DownloadDecompress:             "DownloadDecompress",
		OsUsersManage:                  "OsUsersManage",
		OsGroupsManage:                 "OsGroupsManage",
		OsServicesCreate:               "OsServicesCreate",
		LinuxServiceCreate:             "LinuxServiceCreate",
		Ssh:                            "Ssh",
		PathModifiers:                  "PathModifiers",
		FileSystem:                     "FileSystem",
		ScriptsCollection:              "ScriptsCollection",
		Firewall:                       "Firewall",
		FirewallIpTables:               "FirewallIpTables",
		CliInstructions:                "CliInstructions",
		InstructionsExecuteAfterReboot: "InstructionsExecuteAfterReboot",
		Verify:                         "Verify",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.DefaultAllCases(
		Invalid,
		Ranges[:])
)
