package resauthtype

import (
	"encoding/json"
	"testing"
)

// AL2-02 Batch B coverage suite for resauthtype.

func TestResAuth_NewRoundTrip(t *testing.T) {
	names := []string{"Invalid", "AllAccess", "Error", "Warning", "Restricted", "UnAuthorized",
		"PermissionIssue", "Forbidden", "ReadAccess", "WriteAccess", "CreateAccess", "EditAccess",
		"AccessGranted", "AccessRejected"}
	for _, n := range names {
		v, err := New(n)
		if err != nil {
			t.Errorf("New(%q): %v", n, err)
			continue
		}
		if v.Name() != n {
			t.Errorf("Name() = %q want %q", v.Name(), n)
		}
	}
}

func TestResAuth_NewInvalid(t *testing.T) {
	if v, err := New("__bogus__"); err == nil || v != Invalid {
		t.Errorf("New bogus: v=%v err=%v", v, err)
	}
}

func TestResAuth_NewMust(t *testing.T) {
	if NewMust("AllAccess") != AllAccess {
		t.Error("NewMust mismatch")
	}
}

func TestResAuth_AllPredicates(t *testing.T) {
	checks := []struct {
		name string
		v    Variant
		f    func() bool
	}{
		{"IsAccess", AllAccess, AllAccess.IsAccess},
		{"IsError", Error, Error.IsError},
		{"IsWarning", Warning, Warning.IsWarning},
		{"IsRestricted", Restricted, Restricted.IsRestricted},
		{"IsUnAuthorized", UnAuthorized, UnAuthorized.IsUnAuthorized},
		{"IsPermissionIssue", PermissionIssue, PermissionIssue.IsPermissionIssue},
		{"IsForbidden", Forbidden, Forbidden.IsForbidden},
		{"IsReadAccess", ReadAccess, ReadAccess.IsReadAccess},
		{"IsWriteAccess", WriteAccess, WriteAccess.IsWriteAccess},
		{"IsCreateAccess", CreateAccess, CreateAccess.IsCreateAccess},
		{"IsEditAccess", EditAccess, EditAccess.IsEditAccess},
		{"IsInvalid", Invalid, Invalid.IsInvalid},
		{"IsUninitialized", Invalid, Invalid.IsUninitialized},
	}
	for _, c := range checks {
		if !c.f() {
			t.Errorf("%s should be true for %v", c.name, c.v)
		}
	}
	if !AllAccess.IsValid() || Invalid.IsValid() {
		t.Error("IsValid mismatch")
	}
	if !AllAccess.IsReadLogically() || !AllAccess.IsWriteLogically() ||
		!AllAccess.IsCreateLogically() || !AllAccess.IsEditLogically() {
		t.Error("logical predicates failed for AllAccess")
	}
	if !ReadAccess.IsReadLogically() || !WriteAccess.IsWriteLogically() ||
		!CreateAccess.IsCreateLogically() || !EditAccess.IsEditLogically() {
		t.Error("specific logical predicates failed")
	}
	if !Error.IsAnyErrorLogically() || !Error.IsUnAuthorizedLogically() {
		t.Error("error logical predicate failed")
	}
	if !ReadAccess.IsReadOrUpdateLogically() {
		t.Error("IsReadOrUpdateLogically failed")
	}
	if !AllAccess.IsYes() || !AllAccess.IsAccept() || !AllAccess.IsSuccess() {
		t.Error("yes/accept/success failed")
	}
	if !Error.IsNo() || !Error.IsReject() || !Error.IsFailed() {
		t.Error("no/reject/failed predicates failed")
	}
	if !Warning.IsAsk() || !Warning.IsSkip() || !Warning.IsIndeterminate() {
		t.Error("ask/skip predicates failed")
	}
	if !AllAccess.IsAnyOf(AllAccess, Error) || AllAccess.IsAnyOf(Error) {
		t.Error("IsAnyOf mismatch")
	}
	if !AllAccess.IsNameOf("AllAccess") || AllAccess.IsNameOf("zzz") {
		t.Error("IsNameOf mismatch")
	}
	if !AllAccess.IsAnyNamesOf("AllAccess", "Error") {
		t.Error("IsAnyNamesOf mismatch")
	}
	if !AllAccess.IsByteValueEqual(byte(AllAccess)) ||
		!AllAccess.IsValueEqual(byte(AllAccess)) ||
		!AllAccess.IsNameEqual("AllAccess") ||
		!AllAccess.IsAnyValuesEqual(byte(AllAccess)) {
		t.Error("value-equal mismatch")
	}
}

func TestResAuth_NumericAccessors(t *testing.T) {
	v := ReadAccess
	_ = v.ValueInt()
	_ = v.ValueInt8()
	_ = v.ValueInt16()
	_ = v.ValueInt32()
	_ = v.ValueUInt16()
	_ = v.ValueByte()
	_ = v.MinByte()
	_ = v.MaxByte()
	if v.MinInt() > v.MaxInt() {
		t.Error("MinInt > MaxInt")
	}
	if v.MinValueString() == "" || v.MaxValueString() == "" {
		t.Error("min/max value string empty")
	}
	if min, max := v.MinMaxAny(); min == nil || max == nil {
		t.Error("MinMaxAny nil")
	}
	if len(v.RangesByte()) == 0 || len(v.IntegerEnumRanges()) == 0 ||
		len(v.AllNameValues()) == 0 || len(v.RangesDynamicMap()) == 0 {
		t.Error("ranges empty")
	}
	if v.RangeNamesCsv() == "" || v.NameValue() == "" || v.TypeName() == "" ||
		v.ValueString() == "" || v.ToNumberString() == "" || v.String() == "" {
		t.Error("string accessor empty")
	}
	if v.Format("v=%v") == "" {
		t.Error("Format empty")
	}
	if v.EnumType() == nil {
		t.Error("EnumType nil")
	}
}

func TestResAuth_JsonRoundTrip(t *testing.T) {
	for _, x := range []Variant{Invalid, AllAccess, Error, ReadAccess, WriteAccess, AccessGranted, AccessRejected} {
		data, err := json.Marshal(x)
		if err != nil {
			t.Fatalf("Marshal %v: %v", x, err)
		}
		var got Variant
		if err := json.Unmarshal(data, &got); err != nil {
			t.Fatalf("Unmarshal %v: %v", x, err)
		}
		if got != x {
			t.Errorf("round-trip: got %v want %v", got, x)
		}
	}
}

func TestResAuth_PtrAndBinders(t *testing.T) {
	v := AllAccess
	if v.ToPtr() == nil || (*v.ToPtr()).ToSimple() != v {
		t.Error("ToPtr mismatch")
	}
	var nilP *Variant
	if nilP.ToSimple() != Invalid {
		t.Error("ToSimple(nil) should be Invalid")
	}
	if v.JsonPtr().HasError() {
		t.Errorf("Json: %v", v.JsonPtr().Error)
	}
	if v.JsonPtr() == nil || v.AsJsoner() == nil || v.AsJsonContractsBinder() == nil ||
		v.AsJsonMarshaller() == nil || v.AsBasicByteEnumContractsBinder() == nil ||
		v.AsBasicEnumContractsBinder() == nil || v.AsIsSuccessValidator() == nil ||
		v.AsYesNoAcceptRejecter() == nil {
		t.Error("As* binder nil")
	}
	_ = v.OnlySupportedErr("AllAccess")
	_ = v.OnlySupportedMsgErr("ctx", "AllAccess")
	_ = RangesInvalidErr()
	_ = Min()
	_ = Max()
}
