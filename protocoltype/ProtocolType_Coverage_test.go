package protocoltype

import (
	"encoding/json"
	"testing"
)

// AL2-03 Batch C coverage suite for protocoltype.

func TestProtocolType_MinMax(t *testing.T) {
	if Min() != Invalid {
		t.Errorf("Min %v", Min())
	}
	if Max() != Custom {
		t.Errorf("Max %v", Max())
	}
	if RangesInvalidErr() == nil {
		t.Error("RangesInvalidErr should be informational")
	}
	if !Invalid.IsInvalid() || Invalid.IsValid() {
		t.Error("invalid predicate wrong")
	}
	if !Tcp.IsValid() || !Tcp.IsAnyValuesEqual(byte(Tcp), byte(Udp)) {
		t.Error("valid/IsAnyValuesEqual wrong")
	}
}

func TestProtocolType_PredicateGroups(t *testing.T) {
	// Exercise group predicates if exposed via Variant methods.
	for _, v := range []Variant{Tcp, Udp, Icmp, Grpc, OAuth, Smtp, Imap, Pop3, Ip, IpV6, Custom} {
		if v.Name() == "" || v.ValueString() == "" {
			t.Errorf("%v empty name/value", v)
		}
	}
}

func TestProtocolType_Accessors(t *testing.T) {
	v := Https
	if v.ValueByte() != byte(v) || v.ValueInt() != int(v) || v.ValueInt8() != int8(v) ||
		v.ValueInt16() != int16(v) || v.ValueInt32() != int32(v) || v.ValueUInt16() != uint16(v) {
		t.Error("numeric accessor mismatch")
	}
	if v.MaxByte() != byte(Max()) || v.MinByte() != byte(Min()) {
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
	if v.Name() == "" || v.NameValue() == "" || v.String() == "" || v.RangeNamesCsv() == "" ||
		v.TypeName() == "" || v.ToNumberString() == "" || v.Format("%s") == "" {
		t.Error("string accessors empty")
	}
	if v.EnumType() == nil {
		t.Error("EnumType nil")
	}
}

func TestProtocolType_EqualityAndJson(t *testing.T) {
	v := Tcp
	if !v.IsByteValueEqual(byte(Tcp)) || v.IsByteValueEqual(byte(Udp)) ||
		!v.IsValueEqual(byte(Tcp)) || !v.IsNameEqual("Tcp") {
		t.Error("equality wrong")
	}
	if !v.IsAnyValuesEqual(byte(Tcp), byte(Udp)) || v.IsAnyValuesEqual(byte(Invalid)) {
		t.Error("IsAnyValuesEqual wrong")
	}
	if !v.IsAnyNamesOf("Tcp", "Udp") || v.IsAnyNamesOf("nope") {
		t.Error("IsAnyNamesOf wrong")
	}
	for _, x := range []Variant{Invalid, Tcp, Https, Custom} {
		data, err := json.Marshal(x)
		if err != nil || len(data) == 0 {
			t.Errorf("Marshal(%v): %v", x, err)
		}
		var got Variant
		if err := json.Unmarshal(data, &got); err != nil {
			t.Errorf("Unmarshal(%v): %v", x, err)
		}
	}
	if v.ToPtr() == nil || v.JsonPtr() == nil {
		t.Error("ToPtr/JsonPtr nil")
	}
	if v.AsJsonMarshaller() == nil || v.AsBasicByteEnumContractsBinder() == nil ||
		v.AsBasicEnumContractsBinder() == nil {
		t.Error("binder nil")
	}
	_ = v.OnlySupportedErr("Tcp")
	_ = v.OnlySupportedMsgErr("ctx", "Tcp")
}
