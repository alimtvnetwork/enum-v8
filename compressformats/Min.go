package compressformats

// Min returns the lowest valid (non-Invalid) compress format.
// Pattern-8 fix: previously returned Invalid (the trailing sentinel); the
// real first member is Zip (iota 0).
func Min() Variant {
	return Zip
}
