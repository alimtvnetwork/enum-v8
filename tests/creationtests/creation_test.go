package creationtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
)

func Test_Creation(t *testing.T) {
	for i, invalidChecker := range simpleEnumCollectionTestCases {
		// Arrange
		name := coredynamic.TypeName(invalidChecker)
		invalidHeader := fmt.Sprintf(
			"%d - %s - IsInvalid",
			i,
			name,
		)
		validHeader := fmt.Sprintf(
			"%d - %s - IsValid",
			i,
			name,
		)

		// Act
		isInvalid := invalidChecker.IsInvalid()
		isValid := invalidChecker.IsValid()

		// Assert
		Convey(invalidHeader, t, func() {
			So(isInvalid, ShouldBeTrue)
		})
		Convey(validHeader, t, func() {
			So(isValid, ShouldBeFalse)
		})
	}
}
