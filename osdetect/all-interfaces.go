package osdetect

import (
	"github.com/alimtvnetwork/enum-v7/inttype"
	"github.com/alimtvnetwork/enum-v7/strtype"
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
