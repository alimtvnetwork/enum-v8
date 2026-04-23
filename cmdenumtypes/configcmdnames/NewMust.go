package configcmdnames

import "github.com/alimtvnetwork/core-v8/errcore"

func NewMust(name string) Variant {
	newType, err := New(name)
	errcore.HandleErr(err)

	return newType
}
