package pathpatterntype

import "strings"

// New
//
//	Variant gets created from a Variant JSON name, direct name,
//	curly name, path name, or the "Name(value)" form produced by
//	AllNameValues (e.g. "Root(1)").
func New(name string) (Variant, error) {
	v, err := BasicEnumImpl.GetValueByName(name)

	if err == nil {
		return Variant(v), nil
	}

	// Tolerate the "Name(value)" form returned by AllNameValues.
	if idx := strings.IndexByte(name, '('); idx > 0 && strings.HasSuffix(name, ")") {
		stripped := name[:idx]
		if v2, err2 := BasicEnumImpl.GetValueByName(stripped); err2 == nil {
			return Variant(v2), nil
		}
	}

	return findUsingInternalMapping(name, err)
}
