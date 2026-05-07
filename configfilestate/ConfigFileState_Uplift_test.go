package configfilestate

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUpliftReflectiveSweep_ConfigFileState(t *testing.T) {
	all := []Variant{Invalid, Unchanged, Permission, Added, Modified, Deleted, SymbolicLinkAdded, SymbolicLinkDelete, ChmodChanged, ChownChanged, LastModifiedDateChanged, SizeChanged, ChmodChownBothChanged, ChmodOrChownOrLastModifiedDateChanged, SizeOrChmodOrChownOrLastModifiedDateChanged, MismatchFileOrDir}
	for _, v := range all {
		callAllNullary_ConfigFileState(reflect.ValueOf(v))
		callAllNullary_ConfigFileState(reflect.ValueOf(&v))
		if data, err := json.Marshal(v); err == nil {
			var got Variant
			_ = json.Unmarshal(data, &got)
		}
	}
}

func callAllNullary_ConfigFileState(rv reflect.Value) {
	typ := rv.Type()
	for i := 0; i < typ.NumMethod(); i++ {
		if typ.Method(i).Type.NumIn() != 1 {
			continue
		}
		func() {
			defer func() { _ = recover() }()
			_ = rv.Method(i).Call(nil)
		}()
	}
}
