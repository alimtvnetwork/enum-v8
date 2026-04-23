package scripttype

import (
	"fmt"

	"github.com/alimtvnetwork/core-v8/converters"
)

type ScriptDefault struct {
	ScriptType       Variant
	IsImplemented    bool
	ProcessName      string
	DefaultArguments []string
}

func (it *ScriptDefault) String() string {
	return fmt.Sprint(
		it.ScriptType.String(),
		converters.AnyToValueString(*it))
}
