package creationtests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core-v8/errcore"
)

func GenerateTestCases(t *testing.T) {
	generateAllEnumGeneralTestCases(
		false)
}

// Test_AllEnums_ContractsTesting
//
//	Generates testcases by generateAllEnumGeneralTestCases()
//	Runs by allEnumGeneralTestCases
func Test_AllEnums_ContractsTesting(t *testing.T) {
	shouldBe := errcore.ShouldBe.StrEqMsg

	for _, testCase := range allEnumGeneralTestCases {
		// Arrange
		currentEnum := testCase.InitialBasicEnumer
		var actualEnumDynamicMap enumimpl.DynamicMap = currentEnum.RangesDynamicMap()
		typeName := currentEnum.TypeName()
		currentRangesCsv := currentEnum.RangeNamesCsv()
		customHeader := typeName

		Convey(customHeader, t, func() {
			So(typeName, ShouldEqual, testCase.TypeName)
			So(currentEnum.EnumType(), ShouldEqual, testCase.ExpectedEnumType)
			So(currentEnum.Name(), ShouldResemble, testCase.ExpectedInvalidName)
			So(currentEnum.ValueString(), ShouldResemble, testCase.ExpectedInvalidValueString)
			So(currentEnum.MinValueString(), ShouldResemble, testCase.StringMin)
			So(currentEnum.MaxValueString(), ShouldResemble, testCase.StringMax)
		})

		Convey(customHeader+shouldBe(testCase.ExpectedRangesNamesCsv, currentRangesCsv), t, func() {
			So(currentRangesCsv, ShouldResemble, testCase.ExpectedRangesNamesCsv)
		})

		Convey(typeName+" should match map[string]interface{}", t, func() {
			diffMessage := actualEnumDynamicMap.LogShouldDiffMessage(
				true,
				typeName+" - type mismatch",
				testCase.ExpectedMapValues)

			So(diffMessage, ShouldBeEmpty)
		})
	}
}
