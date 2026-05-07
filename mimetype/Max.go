package mimetype

// Max returns the highest valid (non-Invalid) MIME top-level category.
// Pattern-8 fix: skip the trailing Invalid sentinel.
func Max() Variant { return Video }
