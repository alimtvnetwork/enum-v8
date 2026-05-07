package osgroupexecution

import (
	"encoding/json"
	"testing"
)

// AL2-01 Batch A coverage suite for osgroupexecution.
// Covers New/NewMust + every IsX predicate + JSON round-trip + Variant accessors.

func TestOsGroupExecution_NewRoundTrip(t *testing.T) {
	names := []string{"Create", "Delete", "Update", "ManageByUsers", "AddGroupsToSudoers", "GroupManage"}
	for _, n := range names {
		v, err := New(n)
		if err != nil {
			t.Errorf("New(%q) error: %v", n, err)
			continue
		}
		if v.Name() != n {
			t.Errorf("Name() = %q, want %q", v.Name(), n)
		}
	}
}

func TestOsGroupExecution_NewInvalid(t *testing.T) {
	v, err := New("__bogus__")
	if err == nil || v != Invalid {
		t.Errorf("New bogus: v=%v err=%v", v, err)
	}
}

func TestOsGroupExecution_NewMust(t *testing.T) {
	if NewMust("Create") != Create {
		t.Error("NewMust(Create) mismatch")
	}
}

func TestOsGroupExecution_Predicates(t *testing.T) {
	if !Create.IsCreate() || !Delete.IsDelete() || !Update.IsUpdate() || !ManageByUsers.IsManageByUsers() {
		t.Error("Is* predicate mismatch")
	}
	if !Invalid.IsInvalid() || Create.IsInvalid() {
		t.Error("IsInvalid mismatch")
	}
	if !Create.IsValid() || Invalid.IsValid() {
		t.Error("IsValid mismatch")
	}
	if !Create.IsAnyOf(Create, Delete) || Create.IsAnyOf(Update, Delete) {
		t.Error("IsAnyOf mismatch")
	}
	if !Create.IsAnyNamesOf("Create", "Delete") {
		t.Error("IsAnyNamesOf mismatch")
	}
	if !Create.IsAnyValuesEqual(byte(Create), byte(Delete)) {
		t.Error("IsAnyValuesEqual mismatch")
	}
	if !Create.IsByteValueEqual(byte(Create)) {
		t.Error("IsByteValueEqual mismatch")
	}
	if !Create.IsValueEqual(byte(Create)) || !Create.IsNameEqual("Create") {
		t.Error("IsValueEqual / IsNameEqual mismatch")
	}
}

func TestOsGroupExecution_NumericAccessors(t *testing.T) {
	v := Create
	if v.ValueInt() != int(Create) || v.ValueInt8() != int8(Create) ||
		v.ValueInt16() != int16(Create) || v.ValueInt32() != int32(Create) ||
		v.ValueByte() != byte(Create) || v.ValueUInt16() != uint16(Create) ||
		v.Value() != byte(Create) {
		t.Error("numeric accessor mismatch")
	}
	if v.MinByte() > v.MaxByte() {
		t.Errorf("Min %d > Max %d", v.MinByte(), v.MaxByte())
	}
	if v.MinInt() > v.MaxInt() {
		t.Error("MinInt > MaxInt")
	}
	if v.MinValueString() == "" || v.MaxValueString() == "" {
		t.Error("Min/Max value string empty")
	}
	min, max := v.MinMaxAny()
	if min == nil || max == nil {
		t.Error("MinMaxAny nil")
	}
	if len(v.RangesByte()) == 0 || len(v.IntegerEnumRanges()) == 0 || len(v.AllNameValues()) == 0 {
		t.Error("ranges/names empty")
	}
	if v.RangeNamesCsv() == "" || v.NameValue() == "" || v.TypeName() == "" {
		t.Error("string accessor empty")
	}
	if len(v.RangesDynamicMap()) == 0 {
		t.Error("RangesDynamicMap empty")
	}
}

func TestOsGroupExecution_JsonRoundTrip(t *testing.T) {
	for _, p := range []Precedence{Create, Delete, Update, ManageByUsers, AddGroupsToSudoers, GroupManage} {
		data, err := json.Marshal(p)
		if err != nil {
			t.Fatalf("Marshal %v: %v", p, err)
		}
		var got Precedence
		if err := json.Unmarshal(data, &got); err != nil {
			t.Fatalf("Unmarshal %v: %v", p, err)
		}
		if got != p {
			t.Errorf("round-trip mismatch: got %v want %v", got, p)
		}
	}
}

func TestOsGroupExecution_FormatAndPtr(t *testing.T) {
	v := Update
	if v.Format("v=%v") == "" {
		t.Error("Format empty")
	}
	if v.ToPtr() == nil || *v.ToPtr() != v {
		t.Error("ToPtr mismatch")
	}
	if v.String() == "" || v.ValueString() == "" || v.ToNumberString() == "" {
		t.Error("string method empty")
	}
	if v.JsonPtr().HasError() {
		t.Errorf("Json error: %v", v.JsonPtr().Error)
	}
	if v.JsonPtr() == nil {
		t.Error("JsonPtr nil")
	}
	if v.AsJsoner() == nil || v.AsJsonContractsBinder() == nil ||
		v.AsJsonMarshaller() == nil || v.AsBasicByteEnumContractsBinder() == nil ||
		v.AsBasicEnumContractsBinder() == nil {
		t.Error("As* binder nil")
	}
	if v.EnumType() == nil {
		t.Error("EnumType nil")
	}
}
