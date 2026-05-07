package osarchs

// Max returns the highest valid (non-Invalid) architecture.
// Invalid is the trailing sentinel in the const block, so we explicitly
// return the last real architecture (X64) instead of BasicEnumImpl.Max()
// which would return Invalid.
func Max() Architecture {
	return X64
}
