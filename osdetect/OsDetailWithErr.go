package osdetect

import (
	"github.com/alimtvnetwork/core-v8/coredata/corejson"
)

type OsDetailWithErr struct {
	OperatingSystemDetail *OperatingSystemDetail
	Error                 string
}

func (it *OsDetailWithErr) String() string {
	if it == nil {
		return ""
	}
	
	json := it.JsonPtr()
	
	if json.HasError() {
		return json.MeaningfulErrorMessage()
	}
	
	return json.String()
}

func (it *OsDetailWithErr) PrettyJsonString() string {
	if it == nil {
		return ""
	}
	
	json := it.JsonPtr()
	
	if json.HasError() {
		return json.MeaningfulErrorMessage()
	}
	
	return json.PrettyJsonString()
}

func (it *OsDetailWithErr) Json() corejson.Result {
	return corejson.New(it)
}

func (it *OsDetailWithErr) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *OsDetailWithErr) JsonParseSelfInject(jsonResult *corejson.Result) error {
	return jsonResult.Deserialize(it)
}

func (it OsDetailWithErr) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return &it
}
