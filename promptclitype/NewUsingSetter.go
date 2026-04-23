package promptclitype

import (
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/errcore"
	"github.com/alimtvnetwork/core-v8/issetter"
)

func NewUsingSetter(value issetter.Value) (Variant, error) {
	mappedItem, has := isSetterWithVariantMap[value]

	if !has {
		return Invalid, errcore.
			KeyNotExistInMapType.
			Error(
				typeConvFailedPrefixMsg+coredynamic.TypeName(value),
				mapReferenceMessage)
	}

	return Variant(mappedItem.Value()), nil
}
