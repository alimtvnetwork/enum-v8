package completionstate

import (
	"github.com/alimtvnetwork/core-v9/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v9/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:               "Invalid",
		Initiate:              "Initiate",
		Running:               "Running",
		Success:               "Success",
		SuccessWithWarning:    "SuccessWithWarning",
		FailedMiddleWithError: "FailedMiddleWithError",
		CompleteWithError:     "CompleteWithError",
	}

	CompletionMap = map[Variant]bool{
		Success:               true,
		SuccessWithWarning:    true,
		FailedMiddleWithError: true,
		CompleteWithError:     true,
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.SafeTypeName(Invalid),
		Ranges[:])
)
