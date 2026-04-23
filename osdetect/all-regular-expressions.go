package osdetect

import (
	"github.com/alimtvnetwork/core-v8/regexnew"
)

var (
	prettyNameLazyRegex           = regexnew.PrettyNameRegex
	exactIdFieldMatchLazyRegex    = regexnew.ExactIdFieldMatchingRegex
	versionIdLazyRegex            = regexnew.ExactVersionIdFieldMatchingRegex
	ubuntuLazyRegex               = regexnew.UbuntuNameCheckerRegex
	centOSLazyRegex               = regexnew.CentOsNameCheckerRegex
	redHatLazyRegex               = regexnew.RedHatNameCheckerRegex
	windowsVersionNumberLazyRegex = regexnew.WindowsVersionNumberCheckerRegex
)
