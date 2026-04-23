package iptype

import (
	"strings"

	"github.com/alimtvnetwork/core-v8/constants"
)

// GetUsingIpString
//
//	returns specific Ip version based on ip string
//
// Examples:
//   - ""                 : Invalid
//   - "192..."           : V4
//   - "fe:00....."       : V6
func GetUsingIpString(
	rawIpVersion string,
) Variant {
	if rawIpVersion == "" {
		return Invalid
	}

	if strings.Contains(
		rawIpVersion,
		constants.Colon) {
		return V6
	}

	return V4
}
