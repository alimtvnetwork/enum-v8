package compresslevels

// Max returns the highest valid (non-Invalid) compression level.
// Pattern-8 fix: Invalid is the trailing sentinel in the const block, so
// BasicEnumImpl.Max() would incorrectly return Invalid. Return the last
// real member (NoCompression) explicitly.
func Max() Variant {
	return NoCompression
}
