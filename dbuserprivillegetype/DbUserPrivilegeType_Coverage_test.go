package dbuserprivilegetype

import (
	"encoding/json"
	"testing"
)

// AL2-02 Batch B coverage suite for dbuserprivilegetype.

func TestDbPriv_NewRoundTrip(t *testing.T) {
	names := []string{"Invalid", "All", "Select", "Insert", "Create", "Update", "Alter",
		"Delete", "Drop", "Execute", "Event", "CreateView", "Index", "LockTables",
		"References", "ShowView", "Trigger"}
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

func TestDbPriv_NewInvalid(t *testing.T) {
	if v, err := New("__bogus__"); err == nil || v != Invalid {
		t.Errorf("New bogus: v=%v err=%v", v, err)
	}
}

func TestDbPriv_NewMust(t *testing.T) {
	if NewMust("Select") != Select {
		t.Error("NewMust mismatch")
	}
}

func TestDbPriv_AllPredicates(t *testing.T) {
	checks := []struct {
		name string
		got  bool
	}{
		{"IsNone", Invalid.IsNone()},
		{"IsAll", All.IsAll()},
		{"IsSelect", Select.IsSelect()},
		{"IsRead", Select.IsRead()},
		{"IsReadOrSelect", Select.IsReadOrSelect()},
		{"IsInsert", Insert.IsInsert()},
		{"IsCreate", Create.IsCreate()},
		{"IsUpdate", Update.IsUpdate()},
		{"IsAlter", Alter.IsAlter()},
		{"IsRenameOrChange", Alter.IsRenameOrChange()},
		{"IsDelete", Delete.IsDelete()},
		{"IsDrop", Drop.IsDrop()},
		{"IsExecute", Execute.IsExecute()},
		{"IsEvent", Event.IsEvent()},
		{"IsCreateView", CreateView.IsCreateView()},
		{"IsIndex", Index.IsIndex()},
		{"IsLockTables", LockTables.IsLockTables()},
		{"IsReferences", References.IsReferences()},
		{"IsShowView", ShowView.IsShowView()},
		{"IsTrigger", Trigger.IsTrigger()},
		{"IsInvalid", Invalid.IsInvalid()},
	}
	for _, c := range checks {
		if !c.got {
			t.Errorf("%s should be true", c.name)
		}
	}
	if !All.IsValid() || Invalid.IsValid() {
		t.Error("IsValid mismatch")
	}
	if !All.IsAllOr(Select) || !Select.IsAllOr(Select) || Insert.IsAllOr(Select) {
		t.Error("IsAllOr mismatch")
	}
	if !All.IsAllOrValue(byte(Select)) || !Select.IsAllOrValue(byte(Select)) {
		t.Error("IsAllOrValue mismatch")
	}
	if !Create.IsCreateOrUpdate() || !Update.IsCreateOrUpdate() {
		t.Error("IsCreateOrUpdate mismatch")
	}
	if !Insert.IsInsertOrUpdate() || !Update.IsInsertOrUpdate() {
		t.Error("IsInsertOrUpdate mismatch")
	}
	if !Insert.IsCreateOrUpdateOrInsertLogically() ||
		!Create.IsCreateOrUpdateOrInsertLogically() ||
		!Update.IsCreateOrUpdateOrInsertLogically() {
		t.Error("IsCreateOrUpdateOrInsertLogically mismatch")
	}
	if !Delete.IsDropLogically() || !Drop.IsDropLogically() {
		t.Error("IsDropLogically mismatch")
	}
	if !Create.IsCrudOnlyLogically() || Select.IsCrudOnlyLogically() {
		t.Error("IsCrudOnlyLogically mismatch")
	}
	if Create.IsNotCrudOnlyLogically() || !Select.IsNotCrudOnlyLogically() {
		t.Error("IsNotCrudOnlyLogically mismatch")
	}
	if !Select.IsReadOrEditLogically() || !Update.IsReadOrEditLogically() {
		t.Error("IsReadOrEditLogically mismatch")
	}
	if !Select.IsReadOrUpdateLogically() {
		t.Error("IsReadOrUpdateLogically mismatch")
	}
	if !Update.IsEditOrUpdateLogically() || !Insert.IsEditOrUpdateLogically() {
		t.Error("IsEditOrUpdateLogically mismatch")
	}
	if !Update.IsUpdateOrRemoveLogically() || !Delete.IsUpdateOrRemoveLogically() {
		t.Error("IsUpdateOrRemoveLogically mismatch")
	}
	if !Select.IsAnyOf(Select, Update) || Select.IsAnyOf(Update, Insert) {
		t.Error("IsAnyOf mismatch")
	}
	if !Select.IsNameOf("Select") || Select.IsNameOf("zzz") {
		t.Error("IsNameOf mismatch")
	}
	if !Select.IsAnyNamesOf("Select", "Insert") {
		t.Error("IsAnyNamesOf mismatch")
	}
	if !Select.IsByteValueEqual(byte(Select)) || !Select.IsValueEqual(byte(Select)) ||
		!Select.IsNameEqual("Select") || !Select.IsAnyValuesEqual(byte(Select)) {
		t.Error("value-equal mismatch")
	}
}

func TestDbPriv_NumericAccessors(t *testing.T) {
	v := Select
	_ = v.ValueInt()
	_ = v.ValueInt8()
	_ = v.ValueInt16()
	_ = v.ValueInt32()
	_ = v.ValueUInt16()
	_ = v.ValueByte()
	_ = v.Value()
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
	_ = v.OnlySupportedErr("Select")
	_ = v.OnlySupportedMsgErr("ctx", "Select")
	_ = RangesInvalidErr()
}

func TestDbPriv_JsonRoundTrip(t *testing.T) {
	for _, x := range []Variant{All, Select, Insert, Create, Update, Delete, Drop, Trigger} {
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

func TestDbPriv_PtrAndBinders(t *testing.T) {
	v := Select
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
		v.AsBasicEnumContractsBinder() == nil || v.AsCrudTyper() == nil ||
		v.AsPrivilegeTyper() == nil {
		t.Error("As* binder nil")
	}
}

func TestDbPriv_NotImplementedPanics(t *testing.T) {
	mustPanic := func(label string, fn func()) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("%s should panic", label)
			}
		}()
		fn()
	}
	v := Select
	mustPanic("IsSkipOnExist", func() { v.IsSkipOnExist() })
	mustPanic("IsDropOnExist", func() { v.IsDropOnExist() })
	mustPanic("IsCreateLogically", func() { v.IsCreateLogically() })
	mustPanic("IsCreateOrUpdateLogically", func() { v.IsCreateOrUpdateLogically() })
	mustPanic("IsUpdateOnExist", func() { v.IsUpdateOnExist() })
	mustPanic("IsOnExistCheckLogically", func() { v.IsOnExistCheckLogically() })
	mustPanic("IsOnExistOrSkipOnNonExistLogically", func() { v.IsOnExistOrSkipOnNonExistLogically() })
}
