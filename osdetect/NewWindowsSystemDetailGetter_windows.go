package osdetect

import (
	"github.com/alimtvnetwork/core-v9/codestack"
	"github.com/alimtvnetwork/core-v9/errcore"
	"golang.org/x/sys/windows/registry"
)

func NewWindowsSystemDetailGetter() (windowsSysDetailDefiner, error) {
	k, err := registry.OpenKey(
		registry.LOCAL_MACHINE,
		windowsRegistryKeyPathForOsInfo,
		registry.QUERY_VALUE)
	
	if err != nil {
		return nil, errcore.FailedToParseType.CombineWithAnother(
			"registry.LOCAL_MACHINE",
			"couldn't read registry key!"+err.Error(),
			windowsRegistryKeyPathForOsInfo).ErrorNoRefs(
			codestack.StacksTo.StringDefault())
	}
	
	generator := &windowsSystemDetailGenerator{
		rawErrCollection: errcore.RawErrCollection{},
		rootRegistryKey:  k,
	}
	
	return generator, nil
}
