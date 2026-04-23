package quotes

import (
	"github.com/alimtvnetwork/core-v8/constants"
)

// Assumption here, both quotations exist and s it is not empty
func unWrapBoth(s string) string {
	length := len(s)

	if length <= 2 {
		// both are quotes only
		return constants.EmptyString
	}

	return s[1 : length-2]
}
