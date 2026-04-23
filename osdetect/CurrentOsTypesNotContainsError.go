package osdetect

import (
	"github.com/alimtvnetwork/core-v8/corecsv"
	"github.com/alimtvnetwork/core-v8/errcore"
)

func CurrentOsTypesNotContainsError(items ...Variant) error {
	currentTypesMap := CurrentOsTypesMap()
	names := make([]string, len(currentTypesMap))
	
	for i, item := range items {
		_, has := currentTypesMap[item]
		
		if has {
			return nil
		}
		
		names[i] = item.Name()
	}
	
	expectingAnyOfMessage := corecsv.RangeNamesWithValuesIndexesCsvString(
		names...)
	title := "Current os type not found"
	
	return errcore.ExpectingErrorSimpleNoType(
		title,
		expectingAnyOfMessage,
		currentTypesMap)
}
