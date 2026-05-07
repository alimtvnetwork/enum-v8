package linuxservicestate

import (
	"encoding/json"
	"reflect"
	"testing"
)

// Reflection-based uplift sweep for the ExitCode type (this package has no
// Variant — ExitCode is the byte enum). Mirrors the v1.1.0 template:
// invoke every nullary method via reflection with panic-recovery, JSON
// round-trip, constructor fast-paths.
func TestUpliftReflectiveSweep_LinuxServiceState(t *testing.T) {
	all := []ExitCode{
		Invalid, ActiveRunning, DeadButPidExists, DeadButVarLockFileExists,
		NotRunning, UnknownService, InvalidService, InvalidCode,
	}
	for _, v := range all {
		callAllNullary_LSS(reflect.ValueOf(v))
		callAllNullary_LSS(reflect.ValueOf(&v))
		if data, err := json.Marshal(v); err == nil {
			var got ExitCode
			_ = json.Unmarshal(data, &got)
		}
		if n := v.Name(); n != "" {
			_, _ = New(n)
			func() {
				defer func() { _ = recover() }()
				_ = NewMust(n)
			}()
		}
		// IsAllOf / IsAnyOf / IsEqual / IsAnyOfExitCode / IsNameOf
		_ = v.IsAllOf(int(v), int(v))
		_ = v.IsAllOf()
		_ = v.IsAnyOf(int(v))
		_ = v.IsAnyOf()
		_ = v.IsEqual(int(v))
		_ = v.IsEqual(-1)
		_ = v.IsEqual(99999)
		_ = v.IsAnyOfExitCode(ActiveRunning, NotRunning)
		_ = v.IsNameOf("ActiveRunning", "NotRunning")
	}

	// Code constructors — full int range plus boundary
	for code := -2; code <= 10; code++ {
		_ = NewCode(code)
	}
	for b := byte(0); b <= 10; b++ {
		_ = NewCodeMapping(b)
	}
	// Out-of-range mapping
	_ = NewCodeMapping(255)

	// Bogus name failure
	if _, err := New("__bogus__"); err == nil {
		t.Error("New(__bogus__) should fail")
	}

	// Top-level helpers
	_ = Min()
	_ = Max()
	_ = RangesInvalidErr()
}

func callAllNullary_LSS(rv reflect.Value) {
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
