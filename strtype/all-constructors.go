package strtype

import (
	"fmt"
	"strconv"

	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coreinterface/enuminf"
)

func New(val string) Variant {
	return Variant(val)
}

func NewFileReader(filePath string) FileReader {
	return Variant(filePath).FileReader()
}

func NewUsingInteger(valueInteger int) Variant {
	return Variant(strconv.Itoa(valueInteger))
}

func NewUsingEnum(valueEnum enuminf.BasicEnumer) Variant {
	return Variant(valueEnum.Name())
}

func NewUsingStringer(valueStringer fmt.Stringer) Variant {
	return Variant(valueStringer.String())
}

func NewUsingJsoner(jsoner corejson.Jsoner) Variant {
	json := jsoner.JsonPtr()

	return NewUsingJsonResult(json)
}

func NewUsingJsonResult(jsonResult *corejson.Result) Variant {
	if jsonResult.HasError() {
		return Variant(jsonResult.MeaningfulErrorMessage())
	}

	return Variant(jsonResult.JsonString())
}
