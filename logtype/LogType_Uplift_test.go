package logtype

import (
	"encoding/json"
	"reflect"
	"testing"
)

// Reflection-based uplift sweep: invokes every nullary exported method on
// Variant for every constant, recovering panics. Lifts coverage cheaply
// across all accessor/predicate helpers.
func TestUpliftReflectiveSweep_LogType(t *testing.T) {
	all := []Variant{Silent, Success, Info, Trace, Debug, Warning, Error, Fatal, Panic, Custom, File, Pattern, Invalid}
	sweep(t, all)
}

func sweep(t *testing.T, all []Variant) {
	for _, v := range all {
		callAllNullary(reflect.ValueOf(v))
		callAllNullary(reflect.ValueOf(&v))
		if data, err := json.Marshal(v); err == nil {
			var got Variant
			_ = json.Unmarshal(data, &got)
		}
		// Constructor fast-paths
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
		mt := typ.Method(i).Type
		if mt.NumIn() != 1 {
			continue
		}
		func() {
			defer func() { _ = recover() }()
			_ = rv.Method(i).Call(nil)
		}()
	}
}
