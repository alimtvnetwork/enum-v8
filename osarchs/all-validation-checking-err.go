package osarchs

import "github.com/alimtvnetwork/core-v8/errcore"

func ValidationError(
	rawString string,
	expectedEnum Architecture,
) error {
	return BasicEnumImpl.ExpectingEnumValueError(
		rawString,
		expectedEnum)
}

func StringMustBe(
	rawString string,
	expected Architecture,
) {
	err := ValidationError(rawString, expected)
	errcore.MustBeEmpty(err)
}
