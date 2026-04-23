package promptclitype

import "github.com/alimtvnetwork/core-v8/errcore"

func ValidationError(
	rawString string,
	expected Variant,
) error {
	converted, err := New(rawString)

	if err != nil {
		return errcore.ExpectingErrorSimpleNoType(
			"Expecting "+expected.Name(),
			expected.String(),
			rawString+err.Error())
	}

	if converted == expected {
		return nil
	}

	return errcore.ExpectingErrorSimpleNoType(
		"Expecting "+expected.Name(),
		expected.String(),
		rawString)
}

func StringMustBe(
	rawString string,
	expected Variant,
) {
	err := ValidationError(rawString, expected)
	errcore.MustBeEmpty(err)
}
