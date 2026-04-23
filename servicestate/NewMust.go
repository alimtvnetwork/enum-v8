package servicestate

import "github.com/alimtvnetwork/core-v8/errcore"

func NewMust(name string) Action {
	exitCode, err := New(name)
	errcore.HandleErr(err)

	return exitCode
}
