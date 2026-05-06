package inttype

import (
	"github.com/alimtvnetwork/core-v9/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v9/coredata/corejson"
)

var (
	typeName = coredynamic.SafeTypeName(Variant(-1))

	bytesToDeserializer = corejson.
				Deserialize.
				BytesTo
)
