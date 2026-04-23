package promptclitype

import "github.com/alimtvnetwork/core-v8/errcore"

// NewMust
//
//	Creates string to the type Variant
//
// Mapping (using @nameToVariant):
//   - "":    Invalid,
//   - "-1":  Invalid,
//   - "ask": Ask,
//   - "*":   Ask,
//   - "yes": Accept,
//   - "1":   Accept,
//   - "y":   Accept,
//   - "n":   Reject,
//   - "no":  Reject,
//   - "Reject":  Reject,
//   - "0":   Reject,
func NewMust(name string) Variant {
	newType, err := New(name)
	errcore.HandleErr(err)

	return newType
}
