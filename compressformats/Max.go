package compressformats

// Max returns the highest valid (non-Invalid) compress format.
// Pattern-8 fix: BasicEnumImpl.Max() returns the trailing Invalid sentinel
// here; return the last real member (TarBz2) directly.
func Max() Variant {
	return TarBz2
}
