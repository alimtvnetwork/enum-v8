package creationtests

import (
	"fmt"

	"github.com/alimtvnetwork/core-v9/codestack"
	"github.com/alimtvnetwork/core-v9/constants"
)

func generateAllBasicEnumTestCases() {
	length := len(simpleEnumCollectionTestCases)
	tab := constants.Tab

	fmt.Println("var allBasicEnumsCollection = [...]enuminf.BasicEnumer{")

	for i := 0; i < length; i++ {
		item := simpleEnumCollectionTestCases[i]
		typeName := item.TypeName()
		name := item.Name()
		fullInvalidName := codestack.NameOf.JoinPackageNameWithRelative(
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
