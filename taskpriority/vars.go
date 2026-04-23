package taskpriority

import (
	"github.com/alimtvnetwork/core-v8/converters"
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Default:       "Default",
		DefaultLock:   "DefaultLock",
		Reminder:      "Reminder",
		Notification:  "Notification",
		SystemUpdate:  "SystemUpdate",
		LowerPriority: "LowerPriority",
		Invalid:       "Invalid",
	}

	nameToVariantMap = map[string]Variant{
		"Default":       Default,
		"DefaultLock":   DefaultLock,
		"Reminder":      Reminder,
		"Notification":  Notification,
		"SystemUpdate":  SystemUpdate,
		"LowerPriority": LowerPriority,
		"Invalid":       Invalid,
	}

	priorityMap = map[string]int{
		"Default":       40,
		"DefaultLock":   20,
		"Reminder":      10,
		"Notification":  10,
		"SystemUpdate":  10,
		"LowerPriority": 10,
		"Invalid":       0,
	}

	lockEnforcedMap = [...]bool{
		DefaultLock:  true,
		SystemUpdate: true,
	}

	priorityMapString = converters.AnyToValueString(
		priorityMap)

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Default),
		Ranges[:])
)
