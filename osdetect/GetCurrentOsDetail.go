package osdetect

import (
	"github.com/alimtvnetwork/core-v9/codestack"
	"github.com/alimtvnetwork/core-v9/errcore"
)

func GetCurrentOsDetail() (*OperatingSystemDetail, error) {
	osDetailWithErr := currentOsDetailGeneratorOnce.
		Value().(*OsDetailWithErr)
	
	if osDetailWithErr != nil {
		return osDetailWithErr.OperatingSystemDetail, errcore.ToError(osDetailWithErr.Error)
	}
	
	return nil, errcore.NotSupportedType.Error(
		"couldn't cast or generate os details!",
		codestack.StacksTo.StringDefault())
}
