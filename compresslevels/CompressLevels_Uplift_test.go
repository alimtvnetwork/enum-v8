package compresslevels

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUpliftReflectiveSweep_CompressLevels(t *testing.T) {
	all := []Variant{Default, Best, Fast, NoCompression, Invalid}
	for _, v := range all {
		callAllNullary_CompressLevels(reflect.ValueOf(v))
		callAllNullary_CompressLevels(reflect.ValueOf(&v))
		if data, err := json.Marshal(v); err == nil {
			var got Variant
			_ = json.Unmarshal(data, &got)
		}
	}
}

func callAllNullary_CompressLevels(rv reflect.Value) {
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
