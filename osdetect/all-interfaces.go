package osdetect

import (
	"github.com/alimtvnetwork/enum-v1/inttype"
	"github.com/alimtvnetwork/enum-v1/strtype"
)

type windowsSysDetailDefiner interface {
	Value(
		name string,
	) strtype.Variant
	ValueInt(
		name string,
	) inttype.Variant
	CloseRegKeyRead()
	CompiledErrorWithStackTraces() error
	Finalize() error
	windowsSystemDetailGetter
}

type windowsSystemDetailGetter interface {
	SystemDetail() (*OperatingSystemDetail, error)
}
