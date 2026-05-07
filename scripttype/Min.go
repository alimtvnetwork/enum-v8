package scripttype

// Min returns the lowest-valued enum (Default == 0).
// Invalid is the highest sentinel for this enum (placed last in the const block),
// so the numeric minimum of the range is Default, not Invalid.
func Min() Variant {
	return Default
}
