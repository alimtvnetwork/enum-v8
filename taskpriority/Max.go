package taskpriority

// Max returns the highest valid (non-Invalid) task priority.
// Pattern-8 fix: Invalid is the trailing sentinel in the const block, so
// BasicEnumImpl.Max() would incorrectly return Invalid. Return the last
// real member (LowerPriority) explicitly.
func Max() Variant {
	return LowerPriority
}
