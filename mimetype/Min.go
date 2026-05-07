package mimetype

// Min returns the lowest valid (non-Invalid) MIME top-level category.
// Pattern-8 fix: skip the trailing Invalid sentinel.
func Min() Variant { return Application }
