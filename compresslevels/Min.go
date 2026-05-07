package compresslevels

// Min returns the lowest valid (non-Invalid) compression level.
// Pattern-8 fix: Invalid is the trailing sentinel in the const block, so
// BasicEnumImpl.Min() would still return the first valid value (Default,
// index 0) — which is correct here, but we expose Min() explicitly for
// API symmetry with Max() and to keep the enum's contract obvious.
func Min() Variant {
	return Default
}
