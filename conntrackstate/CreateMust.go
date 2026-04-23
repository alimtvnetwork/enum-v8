package conntrackstate

import "github.com/alimtvnetwork/core-v8/errcore"

func CreateMust(name string) Variant {
	newType, err := Create(name)
	errcore.HandleErr(err)

	return newType
}
