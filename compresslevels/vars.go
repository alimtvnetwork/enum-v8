package compresslevels

import (
	"compress/flate"

	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	ranges = [...]int8{
		Default:       -1,
		Best:          9,
		Fast:          1,
		NoCompression: 0,
	}

	rangesMap = map[int8]Variant{
		-1: Default,
		9:  Best,
		1:  Fast,
		0:  NoCompression,
	}

	stringRanges = [...]string{
		Default:       "Default",
		Best:          "Best",
		Fast:          "Fast",
		NoCompression: "NoCompression",
		Invalid:       "Invalid",
	}

	stringRangesMap = map[string]Variant{
		"Default":       Default,
		"Best":          Best,
		"Fast":          Fast,
		"NoCompression": NoCompression,
	}

	flateRanges = [...]int8{
		Default:       flate.DefaultCompression,
		Best:          flate.BestCompression,
		Fast:          flate.BestSpeed,
		NoCompression: flate.NoCompression,
	}

	flateRangesMap = map[int8]Variant{
		flate.DefaultCompression: Default,
		flate.BestCompression:    Best,
		flate.BestSpeed:          Fast,
		flate.NoCompression:      NoCompression,
	}

	BasicEnumImpl = enumimpl.New.BasicInt8.UsingTypeSlice(
		coredynamic.TypeName(Default),
		stringRanges[:])
)
