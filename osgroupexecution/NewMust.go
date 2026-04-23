package osgroupexecution

import "github.com/alimtvnetwork/core-v8/errcore"

func NewMust(name string) Precedence {
	newType, err := New(name)
	errcore.HandleErr(err)

	return newType
}
