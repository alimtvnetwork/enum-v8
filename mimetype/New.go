package mimetype

func New(name string) (Variant, error) {
	val, err := BasicEnumImpl.GetValueByName(name)
	if err != nil {
		return Invalid, err
	}
	return Variant(val), nil
}
