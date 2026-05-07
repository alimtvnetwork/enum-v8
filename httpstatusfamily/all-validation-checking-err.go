package httpstatusfamily

import "github.com/alimtvnetwork/core-v9/errcore"

func ValidationError(rawString string, expected Variant) error {
	return BasicEnumImpl.ExpectingEnumValueError(rawString, expected)
}

func StringMustBe(rawString string, expected Variant) {
	errcore.MustBeEmpty(ValidationError(rawString, expected))
}
