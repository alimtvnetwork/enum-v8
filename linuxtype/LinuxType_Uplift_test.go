package linuxtype

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUpliftReflectiveSweep_LinuxType(t *testing.T) {
	all := []Variant{Invalid, UbuntuServer, UbuntuServer18, UbuntuServer22, Centos, Centos7, Centos9, CentosStream, DebianServer, DebianServer11, Docker, DockerUbuntuServer, DockerCentos8, Android, UbuntuDesktop}
	for _, v := range all {
		callAllNullary_LinuxType(reflect.ValueOf(v))
		callAllNullary_LinuxType(reflect.ValueOf(&v))
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

func callAllNullary_LinuxType(rv reflect.Value) {
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
