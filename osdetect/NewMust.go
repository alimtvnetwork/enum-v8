package osdetect

import "github.com/alimtvnetwork/core-v8/errcore"

// NewMust
//
//	Creates string to the type Variant
func NewMust(name string) Variant {
	newType, err := New(name)
	errcore.HandleErr(err)
	
	return newType
}
