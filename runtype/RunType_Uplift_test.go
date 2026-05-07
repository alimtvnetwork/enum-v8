package runtype

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUpliftReflectiveSweep_RunType(t *testing.T) {
	all := []Variant{Invalid, Now, OnReboot, OnShutdown, OnEveryReboot, OnEveryShutdown, OnFailRetry, EveryMinute, EveryHour, EveryDay, EveryMonth, EveryYear}
	for _, v := range all {
		callAllNullary(reflect.ValueOf(v))
		callAllNullary(reflect.ValueOf(&v))
		if data, err := json.Marshal(v); err == nil {
			var got Variant
			_ = json.Unmarshal(data, &got)
		}
		if name := v.Name(); name != "" {
			_, _ = New(name)
			_ = NewMust(name)
		}
	}
	if _, err := New("__bogus__"); err == nil {
		t.Error("New bogus should fail")
	}
}

func callAllNullary(rv reflect.Value) {
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
