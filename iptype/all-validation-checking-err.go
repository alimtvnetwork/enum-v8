package iptype

import "github.com/alimtvnetwork/core-v8/errcore"

func ValidationError(
	rawString string,
	expectedEnum Variant,
) error {
	return BasicEnumImpl.ExpectingEnumValueError(
		rawString,
		expectedEnum)
}

func V4ValidationError(
	rawIpVersion string,
) error {
	return ValidationError(
		rawIpVersion,
		V4)
}

func V6ValidationError(
	rawIpVersion string,
) error {
	return ValidationError(
		rawIpVersion,
		V6)
}

func StringMustBe(
	rawString string,
	expected Variant,
) {
	err := ValidationError(rawString, expected)
	errcore.MustBeEmpty(err)
}
