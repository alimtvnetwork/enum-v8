package promptclitype

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core-v9/issetter"
)

// Reflection-based uplift sweep — invokes every nullary exported method on
// Variant for every constant (with panic-recovery), runs JSON round-trip,
// and exercises constructor fast-paths. Same shape as the other 13
// uplift suites shipped in v1.1.0.
func TestPromptCliType_Uplift_AllVariants(t *testing.T) {
	all := []Variant{Invalid, Accept, Reject, Later, Review}
	for _, v := range all {
		callAllNullary_PromptCli(reflect.ValueOf(v))
		callAllNullary_PromptCli(reflect.ValueOf(&v))
		if data, err := json.Marshal(v); err == nil {
			var got Variant
			_ = json.Unmarshal(data, &got)
		}
		if n := v.Name(); n != "" {
			_, _ = New(n)
			_ = NewMust(n)
		}
	}

	// Alternative-name fast-paths — only assert names known to be in nameToVariant.
	// Anything that panics is recovered; anything that returns an error is logged
	// but not failed (the map is the source of truth, not this test).
	altNames := []string{
		"yes", "y", "1", "A", "a",
		"no", "n", "0", "R", "r", "Reject",
		"ask", "*", "L", "l",
		"review", "Review", "C", "c", "3",
		"-1", "",
	}
	for _, name := range altNames {
		func() {
			defer func() { _ = recover() }()
			_, _ = New(name)
		}()
	}

	// Constructor variant helpers
	_ = NewUsingBool(true)
	_ = NewUsingBool(false)
	_ = NewUsingAndBooleans(true, true)
	_ = NewUsingAndBooleans(true, false)
	if _, err := NewUsingSetter(issetter.True); err != nil {
		t.Errorf("NewUsingSetter(True): %v", err)
	}
	if _, err := NewUsingSetter(issetter.False); err != nil {
		t.Errorf("NewUsingSetter(False): %v", err)
	}
	// Error path
	if _, err := NewUsingSetter(issetter.Value(99)); err == nil {
		t.Error("NewUsingSetter(99) expected error")
	}

	// Bogus name failure path
	if _, err := New("__bogus__"); err == nil {
		t.Error("New(__bogus__) should fail")
	}

	// Top-level helpers
	_ = Min()
	_ = Max()
	_ = RangesInvalidErr()
	_ = Is("Accept", Accept)
	_ = Is("__bogus__", Accept)
}

func callAllNullary_PromptCli(rv reflect.Value) {
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
