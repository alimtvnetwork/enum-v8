package creationtests

import (
	"github.com/alimtvnetwork/core-v8/coredata/corerange"
	"github.com/alimtvnetwork/core-v8/coreinterface/enuminf"
)

type EnumTestWrapper struct {
	Header                     string
	InitialBasicEnumer         enuminf.BasicEnumer
	TypeName                   string
	ExpectedEnumType           enuminf.EnumTyper
	ExpectedMapValues          map[string]interface{} // key - name, value - can be any
	ExpectedInvalidName        string
	ExpectedInvalidValueString string
	ExpectedRangesNamesCsv     string
	IntegerMinMax              corerange.MinMaxInt
	StringMin                  string
	StringMax                  string
}
