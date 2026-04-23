package linuxservicestate

import (
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]byte{
		Invalid:                  Invalid.ValueByte(),
		ActiveRunning:            ActiveRunning.ValueByte(),
		DeadButPidExists:         DeadButPidExists.ValueByte(),
		DeadButVarLockFileExists: DeadButVarLockFileExists.ValueByte(),
		NotRunning:               NotRunning.ValueByte(),
		UnknownService:           UnknownService.ValueByte(),
		InvalidService:           InvalidService.ValueByte(),
		InvalidCode:              InvalidCode.ValueByte(),
	}

	StringRanges = [...]string{
		Invalid:                  "Invalid",
		ActiveRunning:            "ActiveRunning",
		DeadButPidExists:         "DeadButPidExists",
		DeadButVarLockFileExists: "DeadButVarLockFileExists",
		NotRunning:               "NotRunning",
		UnknownService:           "UnknownService",
		InvalidService:           "InvalidService",
		InvalidCode:              "InvalidCode",
	}

	// RawMapping
	//
	// Reference :
	// https://gitlab.com/auk-go/os-manuals/uploads/a3fc906f4ea29a59ebf29490391d0f86/image.png
	// https://t.ly/3jkY
	RawMapping = [...]ExitCode{
		0: ActiveRunning,
		1: DeadButPidExists,
		2: DeadButVarLockFileExists,
		3: NotRunning,
		4: UnknownService,
		5: UnknownService,
	}

	rawMappingLength = byte(len(RawMapping))

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(ActiveRunning),
		StringRanges[:])
)
