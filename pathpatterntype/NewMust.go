package pathpatterntype

import "github.com/alimtvnetwork/core-v8/errcore"

// NewMust
//
//	Variant gets created from Variant JSON name direct name or
//	curly name or path name also returns the variant.
//
// Example:
//   - "Id" or "\"Id\"" or {id}
//     or id or idValue as string("5") : should return Id
func NewMust(name string) Variant {
	exitCode, err := New(name)
	errcore.HandleErr(err)

	return exitCode
}
