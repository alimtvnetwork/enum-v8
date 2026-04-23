package linuxservicestate

import "github.com/alimtvnetwork/core-v8/errcore"

func NewMust(codeOrName string) ExitCode {
	exitCode, err := New(codeOrName)
	errcore.HandleErr(err)

	return exitCode
}
