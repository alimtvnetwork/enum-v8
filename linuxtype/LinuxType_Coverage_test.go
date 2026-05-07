package linuxtype

import (
	"encoding/json"
	"testing"
)

// AL2-04 Batch D coverage suite for linuxtype.

func TestLinuxType_NewRoundTrip(t *testing.T) {
	for _, name := range Ranges {
		if name == "" {
			continue
		}
		v, err := New(name)
		if err != nil {
			t.Errorf("New(%q): %v", name, err)
			continue
		}
		if v.Name() == "" {
			t.Errorf("New(%q) empty name", name)
		}
		_ = NewMust(name)
	}
	if _, err := New("__bogus__"); err == nil {
		t.Error("New bogus should fail")
	}
}

func TestLinuxType_GroupMaps(t *testing.T) {
	if len(UbuntuServer.UbuntuServerMap()) == 0 {
		t.Error("UbuntuServerMap empty")
	}
	if len(Ubuntu := UbuntuServer.UbuntuMap()); Ubuntu == 0 {
		t.Error("UbuntuMap empty")
	}
	if len(DebianServer.DebianMap()) == 0 {
		t.Error("DebianMap empty")
	}
	if len(Docker.DockerMap()) == 0 {
		t.Error("DockerMap empty")
	}
}

func TestLinuxType_Accessors(t *testing.T) {
	v := UbuntuServer
	if v.ValueByte() != byte(v) || v.ValueInt() != int(v) || v.ValueInt8() != int8(v) ||
		v.ValueInt16() != int16(v) || v.ValueInt32() != int32(v) || v.ValueUInt16() != uint16(v) {
		t.Error("numeric accessors mismatch")
	}
	if v.MaxByte() != byte(BasicEnumImpl.Max()) || v.MinByte() != byte(BasicEnumImpl.Min()) {
		t.Error("Min/MaxByte mismatch")
	}
	if v.MaxInt() < v.MinInt() || v.MinValueString() == "" || v.MaxValueString() == "" {
		t.Error("Min/Max wrong")
	}
	if min, max := v.MinMaxAny(); min == nil || max == nil {
		t.Error("MinMaxAny nil")
	}
	if len(v.AllNameValues()) == 0 || len(v.IntegerEnumRanges()) == 0 ||
		len(v.RangesByte()) == 0 || len(v.RangesDynamicMap()) == 0 {
		t.Error("ranges empty")
	}
	if v.Name() == "" || v.NameValue() == "" || v.ValueString() == "" ||
		v.RangeNamesCsv() == "" || v.TypeName() == "" || v.Format("%s") == "" {
		t.Error("string accessors empty")
	}
	if v.EnumType() == nil {
		t.Error("EnumType nil")
	}
	if !v.IsByteValueEqual(byte(UbuntuServer)) || v.IsByteValueEqual(byte(Centos)) ||
		!v.IsValueEqual(byte(UbuntuServer)) || !v.IsNameEqual("UbuntuServer") {
		t.Error("equality wrong")
	}
	if !v.IsAnyValuesEqual(byte(UbuntuServer), byte(Centos)) ||
		!v.IsAnyNamesOf("UbuntuServer", "Centos") {
		t.Error("IsAny wrong")
	}
	if !v.IsAnyOf(byte(UbuntuServer)) {
		t.Error("IsAnyOf wrong")
	}
	data, err := json.Marshal(v)
	if err != nil || len(data) == 0 {
		t.Errorf("Marshal: %v", err)
	}
	var got Variant
	if err := json.Unmarshal(data, &got); err != nil {
		t.Errorf("Unmarshal: %v", err)
	}
	if v.ToPtr() == nil {
		t.Error("ToPtr nil")
	}
	if v.AsBasicEnumContractsBinder() == nil || v.AsBasicByteEnumContractsBinder() == nil ||
		v.AsLinuxTyper() == nil {
		t.Error("binder nil")
	}
	_ = v.OnlySupportedErr("UbuntuServer")
	_ = v.OnlySupportedMsgErr("ctx", "UbuntuServer")
}
