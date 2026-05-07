package rootcmdnames

// RangesInvalidErr returns the standard "value not in supported ranges"
// error for this enum.
func RangesInvalidErr() error {
	return BasicEnumImpl.RangesInvalidErr()
}
