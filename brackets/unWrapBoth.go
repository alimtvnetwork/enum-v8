package brackets

import (
	"github.com/alimtvnetwork/core-v8/constants"
)

// Assumption here, both brackets exists and s it is not empty
func unWrapBoth(s string) string {
	length := len(s)

	if length <= 2 {
		// both are brackets only
		return constants.EmptyString
	}

	return s[1 : length-2]
}
