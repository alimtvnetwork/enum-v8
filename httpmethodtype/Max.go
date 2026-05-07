package httpmethodtype

// Max returns the highest valid (non-Invalid) HTTP method variant.
// Pattern-8 fix: BasicEnumImpl.Max() returns the trailing Invalid sentinel
// here; return the last real member (Options) directly.
func Max() Variant {
	return Options
}
