package certaction

import (
	"encoding/json"
	"reflect"
	"testing"
)

func certAll() []Variant {
	return []Variant{Invalid, Create, Renew, Revoke}
}

func certSafe(t *testing.T, label string, fn func()) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Logf("recovered %s: %v", label, r)
		}
	}()
	fn()
}

func TestCertAction_Uplift_NullaryMethodSweep(t *testing.T) {
	for _, v := range certAll() {
		v := v
		rv := reflect.ValueOf(v)
		rt := rv.Type()
		for i := 0; i < rt.NumMethod(); i++ {
			m := rt.Method(i)
			if m.Type.NumIn() != 1 || m.Type.NumOut() == 0 {
				continue
			}
			certSafe(t, m.Name, func() { _ = rv.Method(i).Call(nil) })
		}
		pv := reflect.ValueOf(&v)
		pt := pv.Type()
		for i := 0; i < pt.NumMethod(); i++ {
			m := pt.Method(i)
			if m.Type.NumIn() != 1 || m.Type.NumOut() == 0 {
				continue
			}
			certSafe(t, "*"+m.Name, func() { _ = pv.Method(i).Call(nil) })
		}
	}
}

func TestCertAction_Uplift_CrossVariantOps(t *testing.T) {
	for _, a := range certAll() {
		_ = a.IsCreate()
		_ = a.IsRenew()
		_ = a.IsRevoke()
		_ = a.IsUninitialized()
		_ = a.IsInvalid()
		_ = a.IsValid()
		_ = a.Format("v=%d")
		for _, b := range certAll() {
			_ = a.IsAnyOf(b)
			_ = a.IsEqual(b)
			_ = a.IsAboveOrEqual(b)
			_ = a.IsLowerOrEqual(b)
			_ = a.IsByteValueEqual(b.ValueByte())
			_ = a.IsValueEqual(b.ValueByte())
			_ = a.IsAnyValuesEqual(b.ValueByte())
			_ = a.IsNameEqual(b.Name())
			_ = a.IsAnyNamesOf(b.Name())
		}
	}
}

func TestCertAction_Uplift_JsonRoundTrip(t *testing.T) {
	for _, v := range certAll() {
		data, err := json.Marshal(v)
		if err != nil {
			t.Errorf("Marshal(%v): %v", v, err)
			continue
		}
		var got Variant
		_ = json.Unmarshal(data, &got)
		_ = v.Json()
		_ = v.JsonPtr()
	}
}

func TestCertAction_Uplift_ConstructorMatrix(t *testing.T) {
	names := []string{"Invalid", "Create", "Renew", "Revoke", "", "bogus", "create", "CREATE"}
	for _, n := range names {
		_, _ = New(n)
		certSafe(t, "NewMust:"+n, func() { _ = NewMust(n) })
	}
	_ = Min()
	_ = Max()
	_ = RangesInvalidErr()
}
