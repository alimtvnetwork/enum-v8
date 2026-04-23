package quotes

import (
	"github.com/alimtvnetwork/core-v8/constants"
)

// Assumption here, s has single quotes and s it is not empty
func unWrapSingle(s string, isLeft bool) string {
	length := len(s)

	if length <= 1 {
		// it has quotes only
		return constants.EmptyString
	}

	if isLeft {
		return s[1 : length-1]
	}

	// right
	return s[0 : length-2]
}
