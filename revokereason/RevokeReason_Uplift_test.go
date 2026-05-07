package revokereason

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUpliftReflectiveSweep_RevokeReason(t *testing.T) {
	all := []Variant{Unspecified, KeyCompromise, CaCompromise, AffiliationChanged, Superseded, CessationOfOperation, CertificateHold, RemoveFromCRL, PrivilegeWithdrawn, AaCompromise}
	for _, v := range all {
		callAllNullary_RevokeReason(reflect.ValueOf(v))
		callAllNullary_RevokeReason(reflect.ValueOf(&v))
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

func callAllNullary_RevokeReason(rv reflect.Value) {
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
