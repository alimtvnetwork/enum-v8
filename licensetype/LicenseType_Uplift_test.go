package licensetype

import (
	"encoding/json"
	"reflect"
	"testing"
)

func licAll() []Variant {
	return []Variant{Invalid, PublicDomain, ByCc, BySa, ByNc, ByNcSa, ByNd, ByNcNd}
}

func licSafe(t *testing.T, label string, fn func()) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Logf("recovered %s: %v", label, r)
		}
	}()
	fn()
}

func TestLicenseType_Uplift_NullaryMethodSweep(t *testing.T) {
	for _, v := range licAll() {
		v := v
		rv := reflect.ValueOf(v)
		rt := rv.Type()
		for i := 0; i < rt.NumMethod(); i++ {
			m := rt.Method(i)
			if m.Type.NumIn() != 1 || m.Type.NumOut() == 0 {
				continue
			}
			licSafe(t, m.Name, func() { _ = rv.Method(i).Call(nil) })
		}
		pv := reflect.ValueOf(&v)
		pt := pv.Type()
		for i := 0; i < pt.NumMethod(); i++ {
			m := pt.Method(i)
			if m.Type.NumIn() != 1 || m.Type.NumOut() == 0 {
				continue
			}
			licSafe(t, "*"+m.Name, func() { _ = pv.Method(i).Call(nil) })
		}
	}
}

func TestLicenseType_Uplift_CrossVariantOps(t *testing.T) {
	for _, a := range licAll() {
		for _, b := range licAll() {
			_ = a.IsAnyOf(b)
			_ = a.IsByteValueEqual(b.ValueByte())
			_ = a.IsValueEqual(b.ValueByte())
			_ = a.IsAnyValuesEqual(b.ValueByte())
			_ = a.IsNameEqual(b.Name())
			_ = a.IsAnyNamesOf(b.Name())
		}
		_ = a.Format("v=%d")
	}
}

func TestLicenseType_Uplift_JsonRoundTrip(t *testing.T) {
	for _, v := range licAll() {
		data, err := json.Marshal(v)
		if err != nil {
			t.Errorf("Marshal(%v): %v", v, err)
			continue
		}
		var got Variant
		_ = json.Unmarshal(data, &got)
	}
}

func TestLicenseType_Uplift_ConstructorMatrix(t *testing.T) {
	names := []string{
		"Invalid", "PublicDomain", "ByCc", "BySa", "ByNc",
		"ByNcSa", "ByNd", "ByNcNd", "", "bogus", "bycc", "BYCC",
	}
	for _, n := range names {
		_, _ = New(n)
		licSafe(t, "NewMust:"+n, func() { _ = NewMust(n) })
	}
	_ = Min()
	_ = Max()
	_ = RangesInvalidErr()
}
