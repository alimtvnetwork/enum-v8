package inputiptype

import (
	"encoding/json"
	"testing"
)

// AL2-03 Batch C — networking/IP coverage suite.

func TestInputIpType_NewAndPredicates(t *testing.T) {
	for _, name := range []string{"Ip", "IpWithSubnet", "SubnetMask", "Gateway", "IpWithPort"} {
		v, err := New(name)
		if err != nil {
			t.Errorf("New(%q): %v", name, err)
		}
		if v.Name() != name {
			t.Errorf("Name mismatch: got %q want %q", v.Name(), name)
		}
		_ = NewMust(name)
	}
	if _, err := New("__bogus__"); err == nil {
		t.Error("New bogus should fail")
	}
	if !Ip.IsIp() || !IpWithSubnet.IsIpWithSubnet() || !SubnetMask.IsSubnetMask() ||
		!Gateway.IsGateway() || !IpWithPort.IsIPWithPort() {
		t.Error("Is* predicate wrong")
	}
	if Min() != Invalid || Max() != IpWithPort {
		t.Errorf("Min/Max wrong: %v %v", Min(), Max())
	}
	if RangesInvalidErr() == nil {
		t.Error("RangesInvalidErr should be informational non-nil")
	}
}

func TestInputIpType_Accessors(t *testing.T) {
	v := Gateway
	if v.ValueByte() != byte(v) || v.ValueInt() != int(v) || v.ValueInt8() != int8(v) ||
		v.ValueInt16() != int16(v) || v.ValueInt32() != int32(v) || v.ValueUInt16() != uint16(v) ||
		v.Value() != byte(v) {
		t.Error("numeric accessors mismatch")
	}
	if v.Name() == "" || v.String() == "" || v.NameValue() == "" || v.ValueString() == "" ||
		v.ToNumberString() == "" || v.RangeNamesCsv() == "" || v.TypeName() == "" {
		t.Error("string accessors empty")
	}
	if v.MaxByte() != byte(Max()) || v.MinByte() != byte(Min()) {
		t.Error("Min/MaxByte mismatch")
	}
	if v.MaxInt() < v.MinInt() || v.MinValueString() == "" || v.MaxValueString() == "" {
		t.Error("Min/Max range wrong")
	}
	if min, max := v.MinMaxAny(); min == nil || max == nil {
		t.Error("MinMaxAny nil")
	}
	if len(v.AllNameValues()) == 0 || len(v.IntegerEnumRanges()) == 0 ||
		len(v.RangesByte()) == 0 || len(v.RangesDynamicMap()) == 0 {
		t.Error("ranges empty")
	}
	if v.Format("%s") == "" || v.EnumType() == nil {
		t.Error("Format/EnumType wrong")
	}
}

func TestInputIpType_EqualityAndJson(t *testing.T) {
	v := Ip
	if !v.IsByteValueEqual(byte(Ip)) || v.IsByteValueEqual(byte(Gateway)) ||
		!v.IsValueEqual(byte(Ip)) || !v.IsNameEqual("Ip") {
		t.Error("equality wrong")
	}
	if !v.IsAnyValuesEqual(byte(Ip), byte(Gateway)) || v.IsAnyValuesEqual(byte(Invalid)) {
		t.Error("IsAnyValuesEqual wrong")
	}
	if !v.IsAnyNamesOf("Ip", "Gateway") || v.IsAnyNamesOf("nope") {
		t.Error("IsAnyNamesOf wrong")
	}
	if !v.IsValid() || Invalid.IsValid() || !Invalid.IsInvalid() {
		t.Error("valid/invalid wrong")
	}
	if !v.IsAnyOf(Ip, Gateway) || v.IsAnyOf(Gateway) {
		t.Error("IsAnyOf wrong")
	}
	for _, x := range []Variant{Invalid, Ip, IpWithSubnet, SubnetMask, Gateway, IpWithPort} {
		data, err := json.Marshal(x)
		if err != nil || len(data) == 0 {
			t.Errorf("Marshal(%v) err=%v", x, err)
		}
		var got Variant
		if err := json.Unmarshal(data, &got); err != nil {
			t.Errorf("Unmarshal(%v): %v", x, err)
		}
	}
	p := v.ToPtr()
	if p == nil || v.JsonPtr() == nil {
		t.Error("ToPtr/JsonPtr nil")
	}
	if v.AsJsonMarshaller() == nil || v.AsBasicByteEnumContractsBinder() == nil ||
		v.AsBasicEnumContractsBinder() == nil {
		t.Error("binder nil")
	}
	_ = v.OnlySupportedErr("Ip")
	_ = v.OnlySupportedMsgErr("ctx", "Ip")
}
