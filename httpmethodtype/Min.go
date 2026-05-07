package httpmethodtype

// Min returns the lowest valid (non-Invalid) HTTP method variant.
// Pattern-8 fix: BasicEnumImpl.Min() may surface the trailing Invalid
// sentinel; return the real first member (Get) directly.
func Min() Variant {
	return Get
}
