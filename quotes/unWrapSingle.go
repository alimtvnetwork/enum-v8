package quotes

import (
	"github.com/alimtvnetwork/core-v9/constants"
)

// Assumption here, s has single quote on one side and s it is not empty.
//
// PI-008 (2026-05-06): fixed off-by-one in the right-only branch. Previously
// returned `s[0:length-2]` which dropped two trailing chars; correct is
// `s[0:length-1]` to strip exactly the single trailing quote.
func unWrapSingle(s string, isLeft bool) string {
	length := len(s)

	if length <= 1 {
		// it has quotes only
		return constants.EmptyString
	}

	if isLeft {
		return s[1:length]
	}

	// right
	return s[0 : length-1]
}
