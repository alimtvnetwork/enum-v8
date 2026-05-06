package brackets

import (
	"github.com/alimtvnetwork/core-v9/constants"
)

// Assumption here, s has single bracket on one side and s it is not empty.
//
// PI-008 (2026-05-06): fixed off-by-one in both branches. Left-only previously
// returned `s[1:length-1]` (dropped trailing char too); right-only previously
// returned `s[0:length-2]` (dropped two trailing chars). Correct is to strip
// exactly the single bracket on the wrapped side.
func unWrapSingle(s string, isLeft bool) string {
	length := len(s)

	if length <= 1 {
		// it has brackets only
		return constants.EmptyString
	}

	if isLeft {
		return s[1:length]
	}

	// right
	return s[0 : length-1]
}
