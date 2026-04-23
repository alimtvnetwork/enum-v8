package packageinstallmethod

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
//   - "yes": On,
//   - "1":   On,
//   - "y":   On,
//   - "n":   Off,
//   - "no":  Off,
//   - "Off":  Off,
//   - "0":   Off,
func NewMust(name string) Variant {
	newType, err := New(name)
	errcore.HandleErr(err)

	return newType
}
