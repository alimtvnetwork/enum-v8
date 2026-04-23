package revokereason

import (
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Unspecified:          "Unspecified",          // 0
		KeyCompromise:        "KeyCompromise",        // 1
		CaCompromise:         "CaCompromise",         // 2
		AffiliationChanged:   "AffiliationChanged",   // 3
		Superseded:           "Superseded",           // 4
		CessationOfOperation: "CessationOfOperation", // 5
		CertificateHold:      "CertificateHold",      // 6
		_Unused:              "_Unused",              // 7
		RemoveFromCRL:        "RemoveFromCRL",        // 8
		PrivilegeWithdrawn:   "PrivilegeWithdrawn",   // 9
		AaCompromise:         "AaCompromise",         // 10
	}

	RangesMap = map[string]Variant{
		"":                     Unspecified,
		"Unspecified":          Unspecified,
		"KeyCompromise":        KeyCompromise,
		"CaCompromise":         CaCompromise,
		"AffiliationChanged":   AffiliationChanged,
		"Superseded":           Superseded,
		"CessationOfOperation": CessationOfOperation,
		"CertificateHold":      CertificateHold,
		"_Unused":              _Unused, // 7
		"RemoveFromCRL":        RemoveFromCRL,
		"PrivilegeWithdrawn":   PrivilegeWithdrawn,
		"AaCompromise":         AaCompromise,
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Unspecified),
		Ranges[:])
)
