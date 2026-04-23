package dbdrivertype

import "github.com/alimtvnetwork/core-v8/errcore"

func ValidationError(
	rawString string,
	expectedEnum Variant,
) error {
	return BasicEnumImpl.ExpectingEnumValueError(
		rawString,
		expectedEnum)
}

func StringMustBe(
	rawString string,
	expected Variant,
) {
	err := ValidationError(rawString, expected)
	errcore.MustBeEmpty(err)
}
