package taskpriority

// Min returns the lowest valid (non-Invalid) task priority.
// Pattern-8 fix: Invalid is the trailing sentinel; first valid is Default (0).
func Min() Variant {
	return Default
}
