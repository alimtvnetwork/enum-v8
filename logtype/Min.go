package logtype

// Min returns the lowest valid (non-Invalid) log type.
// Pattern-8 fix: Invalid is the trailing sentinel; first valid is Silent (0).
func Min() Variant {
	return Silent
}
