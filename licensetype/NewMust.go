package licensetype

import "github.com/alimtvnetwork/core-v8/errcore"

func NewMust(name string) Variant {
	exitCode, err := New(name)
	errcore.HandleErr(err)

	return exitCode
}
