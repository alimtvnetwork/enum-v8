package quotes

import (
	"github.com/alimtvnetwork/core-v9/constants"
)

// Assumption here, both quotations exist and s it is not empty.
//
// PI-008 (2026-05-06): fixed off-by-one. Previously returned `s[1:length-2]`
// which dropped two chars from the right and one from the left, e.g. `"hi"`
// → "h" instead of "hi". Symmetric strip is `s[1:length-1]`.
func unWrapBoth(s string) string {
	length := len(s)

	if length <= 2 {
		// both are quotes only
		return constants.EmptyString
	}

	return s[1 : length-1]
}
