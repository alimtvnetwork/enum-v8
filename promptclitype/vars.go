package promptclitype

import (
	"github.com/alimtvnetwork/core-v9/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v9/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core-v9/errcore"
	"github.com/alimtvnetwork/core-v9/issetter"
	"github.com/alimtvnetwork/enum-v2/onofftype"
)

var (
	Ranges = [...]string{
		Invalid: "Invalid",
		Accept:  "Accept",
		Reject:  "Reject",
		Later:   "Later",
		Review:  "Review",
	}

	undefinedItems = [...]bool{
		Invalid: true,
		Later:   true,
	}

	lowerCaseNames = map[Variant]string{
		Invalid: "invalid",
		Accept:  "accept",
		Reject:  "reject",
		Later:   "later",
		Review:  "review",
	}

	onOffNamesLowerMap = map[Variant]string{
		Invalid: "invalid",
		Accept:  "on",
		Reject:  "off",
		Later:   "skip",
		Review:  "review",
	}

	onOffNames = map[Variant]string{
		Invalid: "Invalid",
		Accept:  "On",
		Reject:  "Off",
		Later:   "Skip",
		Review:  "Review",
	}

	trueFalseNames = map[Variant]string{
		Invalid: "uninitialized",
		Accept:  "True",
		Reject:  "False",
		Later:   "Skip",
		Review:  "Review",
	}

	nameToVariant = map[string]Variant{
		"":       Invalid,
		"-1":     Invalid,
		"ask":    Later,
		"*":      Later,
		"yes":    Accept,
		"1":      Accept,
		"y":      Accept,
		"n":      Reject,
		"no":     Reject,
		"Reject": Reject,
		"0":      Reject,
		"3":      Review,
		"review": Review,
		"Review": Review,
		"A":      Accept,
		"R":      Reject,
		"L":      Later,
		"C":      Review,
		"a":      Accept,
		"r":      Reject,
		"l":      Later,
		"c":      Review,
	}

	undefinedMap = map[Variant]bool{
		Invalid: true,
		Later:   true,
	}

	isSetterWithVariantMap = map[issetter.Value]Variant{
		issetter.Uninitialized: Invalid,
		issetter.Wildcard:      Later,
		issetter.True:          Accept,
		issetter.Set:           Accept,
		issetter.False:         Reject,
		issetter.Unset:         Reject,
	}

	variantToIsSetterBooleanMap = map[Variant]issetter.Value{
		Invalid: issetter.Uninitialized,
		Later:   issetter.Wildcard,
		Accept:  issetter.True,
		Reject:  issetter.False,
		Review:  issetter.Wildcard,
	}

	variantToOnOffEnumMap = map[Variant]onofftype.Variant{
		Invalid: onofftype.Invalid,
		Later:   onofftype.Ask,
		Accept:  onofftype.On,
		Reject:  onofftype.Off,
		Review:  onofftype.Ask, // todo fix ambiguity, alim
	}

	mapReferenceMessage = errcore.MessageWithRef(
		"mapping list",
		isSetterWithVariantMap)

	typeConvFailedPrefixMsg = BasicEnumImpl.TypeName() +
		" cannot be converted from "

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Invalid),
		Ranges[:])
)
