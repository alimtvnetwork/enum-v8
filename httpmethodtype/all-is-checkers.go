package httpmethodtype

func Is(rawString string, expected Variant) bool {
	v, err := New(rawString)
	if err != nil {
		return false
	}
	return v == expected
}
