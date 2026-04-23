package servicestate

import (
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core-v8/reqtype"
)

var (
	Ranges = [...]string{
		Invalid: "Invalid",
		Status:  "status",
		Start:   "start",
		Restart: "restart",
		Reload:  "reload",
		Enable:  "enable",
		Disable: "disable",
		Stop:    "stop",
	}

	capitalNameMap = [...]string{
		Invalid: "Invalid",
		Status:  "Status",
		Start:   "Start",
		Restart: "Restart",
		Reload:  "Reload",
		Enable:  "Enable",
		Disable: "Disable",
		Stop:    "Stop",
	}

	actionToRequestMap = map[Action]reqtype.Request{
		Invalid: reqtype.Invalid,
		Status:  reqtype.Invalid,
		Start:   reqtype.Start,
		Restart: reqtype.Restart,
		Reload:  reqtype.Reload,
		Enable:  reqtype.Invalid,
		Disable: reqtype.Invalid,
		Stop:    reqtype.Stop,
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Status),
		Ranges[:])
)
