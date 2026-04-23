package pathpatterntype

import (
	"github.com/alimtvnetwork/core-v8/constants"
	"github.com/alimtvnetwork/core-v8/errcore"
	"github.com/alimtvnetwork/core-v8/simplewrap"
)

// findUsingInternalMapping
//
//	Variant gets created from Variant JSON name direct name or
//	curly name or path name also returns the variant.
//
// Used by:
//
//	New
//
// Example:
//   - "{id}" or "id" : should return Id
func findUsingInternalMapping(
	name string, err error,
) (Variant, error) {
	if name == "" {
		return Invalid, err
	}

	if name[0] != constants.CurlyBraceStartChar {
		name = simplewrap.WithCurly(name)
	}

	found, has := singlePatternFormatToVariantMap[name]

	if has {
		return found.Clone(), nil
	}

	// error
	return Invalid, errcore.ErrorWithRefToError(
		err,
		singlePatternFormatToVariantMap)
}
