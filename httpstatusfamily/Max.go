package httpstatusfamily

// Max returns the highest valid (non-Invalid) status family.
// Pattern-8 fix: BasicEnumImpl.Max() returns the trailing Invalid sentinel
// here; return the last real member directly.
func Max() Variant {
	return ServerError
}
