package promptclitype

import (
	"github.com/alimtvnetwork/core-v9/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v9/errcore"
	"github.com/alimtvnetwork/core-v9/issetter"
)

func NewUsingSetter(value issetter.Value) (Variant, error) {
	mappedItem, has := isSetterWithVariantMap[value]

	if !has {
		return Invalid, errcore.
			KeyNotExistInMapType.
			Error(
				typeConvFailedPrefixMsg+coredynamic.SafeTypeName(value),
				mapReferenceMessage)
	}

	return Variant(mappedItem.Value()), nil
}
