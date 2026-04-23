package creationtests

import (
	"fmt"

	"github.com/alimtvnetwork/core-v8/codestack"
	"github.com/alimtvnetwork/core-v8/constants"
)

func generateAllBasicEnumTestCases() {
	length := len(simpleEnumCollectionTestCases)
	tab := constants.Tab

	fmt.Println("var allBasicEnumsCollection = [...]enuminf.BasicEnumer{")

	for i := 0; i < length; i++ {
		item := simpleEnumCollectionTestCases[i]
		typeName := item.TypeName()
		name := item.Name()
		fullInvalidName := codestack.JoinPackageNameWithRelative(
			typeName,
			name)

		fmt.Println(
			tab,
			fullInvalidName+".AsBasicEnumContractsBinder(),",
		)
	}

	fmt.Println(
		"}",
	)
}
