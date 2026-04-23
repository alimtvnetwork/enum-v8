package osdetect

import (
	"github.com/alimtvnetwork/core-v8/constants/regkeysconsts"
	"github.com/alimtvnetwork/core-v8/filemode"
)

const (
	powershell    = "powershell"
	bash          = "bash"
	windowsVendor = "microsoft"
	// macOsSysVersionCommand
	//
	// "sw_vers" returns command line output:
	//   ProductName:	Mac OS X
	//   ProductVersion:	10.15.7
	//   BuildVersion:	19H524
	macOsSysVersionCommand          = "sw_vers"
	macOsName                       = "macos"
	windowsRegistryKeyPathForOsInfo = regkeysconsts.WindowsOsInfo // `SOFTWARE\Microsoft\Windows NT\CurrentVersion`
	windowsClientInstallationType   = "Client"
	windowsServerInstallationType   = "Server"
	cacheFileMode                   = filemode.CacheFullAccess
	macOsProductName                = "ProductName"
	macOsProductVersion             = "ProductVersion"
	macOsBuildVersion               = "BuildVersion"
	dockerDetectPath                = "/.dockerenv"
	osDetailTempDirName             = "os-detail"
	osCachedTempFileName            = "cached-details.json"
	windows11BuildIdentifier        = 22000 // https://t.ly/Jsr1
)
