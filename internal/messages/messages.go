package messages

import (
	"github.com/alimtvnetwork/core-v8/corecomparator"
	"github.com/alimtvnetwork/core-v8/errcore"
)

var (
	ComparatorOutOfRangeMessage = errcore.RangeNotMeet(
		errcore.ComparatorShouldBeWithinRangeType.String(),
		corecomparator.Min(),
		corecomparator.Max(),
		corecomparator.Ranges())
)
