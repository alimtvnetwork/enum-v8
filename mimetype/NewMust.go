package mimetype

import "github.com/alimtvnetwork/core-v9/errcore"

func NewMust(name string) Variant {
	v, err := New(name)
	errcore.HandleErr(err)
	return v
}
