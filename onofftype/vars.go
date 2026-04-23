package onofftype

import (
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core-v8/errcore"
	"github.com/alimtvnetwork/core-v8/issetter"
)

var (
	Ranges = [...]string{
		Invalid: "Invalid",
		Ask:     "Ask",
		On:      "On",
		Off:     "Off",
	}

	undefinedItems = [...]bool{
		Invalid: true,
		Ask:     true,
	}

	lowerCaseNames = map[Variant]string{
		Invalid: "invalid",
		Ask:     "ask",
		On:      "yes",
		Off:     "no",
	}

	onOffNamesLowerMap = map[Variant]string{
		Invalid: "invalid",
		Ask:     "ask",
		On:      "on",
		Off:     "off",
	}

	onOffNames = map[Variant]string{
		Invalid: "Invalid",
		Ask:     "Ask",
		On:      "On",
		Off:     "Off",
	}

	trueFalseNames = map[Variant]string{
		Invalid: "uninitialized",
		Ask:     "ask",
		On:      "true",
		Off:     "false",
	}

	nameToVariant = map[string]Variant{
		"":    Invalid,
		"-1":  Invalid,
		"ask": Ask,
		"*":   Ask,
		"yes": On,
		"1":   On,
		"y":   On,
		"n":   Off,
		"no":  Off,
		"Off": Off,
		"0":   Off,
	}

	undefinedMap = map[Variant]bool{
		Invalid: true,
		Ask:     true,
	}

	isSetterWithVariantMap = map[issetter.Value]Variant{
		issetter.Uninitialized: Invalid,
		issetter.Wildcard:      Ask,
		issetter.True:          On,
		issetter.Set:           On,
		issetter.False:         Off,
		issetter.Unset:         Off,
	}

	variantToIsSetterBooleanMap = map[Variant]issetter.Value{
		Invalid: issetter.Uninitialized,
		Ask:     issetter.Wildcard,
		On:      issetter.True,
		Off:     issetter.False,
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
