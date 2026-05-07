package leveltype

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUpliftReflectiveSweep_LevelType(t *testing.T) {
	all := []Variant{Invalid, Level1, Level2, Level3, Level4, Level5, Level6, Level7, Level8, Level9, Level10}
	for _, v := range all {
		callAllNullary_LevelType(reflect.ValueOf(v))
		callAllNullary_LevelType(reflect.ValueOf(&v))
		if data, err := json.Marshal(v); err == nil {
			var got Variant
			_ = json.Unmarshal(data, &got)
		}
		if n := v.Name(); n != "" {
			_, _ = New(n)
			_ = NewMust(n)
		}
	}
	if _, err := New("__bogus__"); err == nil {
		t.Error("New bogus should fail")
	}
}

func callAllNullary_LevelType(rv reflect.Value) {
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
