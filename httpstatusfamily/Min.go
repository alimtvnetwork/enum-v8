package httpstatusfamily

// Min returns the lowest valid (non-Invalid) status family.
// Pattern-8 fix: BasicEnumImpl.Min() may surface the trailing Invalid
// sentinel; return the real first member directly.
func Min() Variant {
	return Informational
}
