package logtype

// Max returns the highest valid (non-Invalid) log type.
// Pattern-8 fix: Invalid is the trailing sentinel in the const block, so
// BasicEnumImpl.Max() would incorrectly return Invalid. Return the last
// real member (Pattern) explicitly.
func Max() Variant {
	return Pattern
}
