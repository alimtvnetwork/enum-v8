package servicestate

import "testing"

// AL2-01 Batch A coverage suite for servicestate.
// Covers New/NewMust/Min/Max/RangesInvalidErr + all action predicates.

func TestServiceState_NewRoundTrip(t *testing.T) {
	names := []string{"status", "start", "restart", "reload", "enable", "disable", "stop"}
	for _, n := range names {
		v, err := New(n)
		if err != nil {
			t.Errorf("New(%q) error: %v", n, err)
			continue
		}
		if v == Invalid {
			t.Errorf("New(%q) returned Invalid", n)
		}
	}
}

func TestServiceState_NewInvalid(t *testing.T) {
	v, err := New("__bogus__")
	if err == nil || v != Invalid {
		t.Errorf("New bogus: v=%v err=%v", v, err)
	}
}

func TestServiceState_NewMust(t *testing.T) {
	if NewMust("start") != Start {
		t.Error("NewMust(start) mismatch")
	}
}

func TestServiceState_MinMax(t *testing.T) {
	if Min() != Status {
		t.Errorf("Min() = %v, want Status", Min())
	}
	if Max() == Invalid {
		t.Error("Max() should not be Invalid")
	}
}

func TestServiceState_RangesInvalidErr(t *testing.T) {
	if err := RangesInvalidErr(); err != nil {
		t.Errorf("RangesInvalidErr unexpected: %v", err)
	}
}

func TestServiceState_ActionPredicates(t *testing.T) {
	if !Invalid.IsUndefined() || Start.IsUndefined() {
		t.Error("IsUndefined mismatch")
	}
	if !StopEnableStart.IsStopEnableStart() || Start.IsStopEnableStart() {
		t.Error("IsStopEnableStart mismatch")
	}
	if !Stop.IsSuspend() || !Stop.IsPause() || !Start.IsResumed() {
		t.Error("Suspend/Pause/Resumed mismatch")
	}
	if Start.IsStopSleepStart() {
		t.Error("IsStopSleepStart should be false")
	}
	if !Start.IsAnyAction() || Start.IsNotAnyAction() {
		t.Error("IsAnyAction mismatch")
	}
	if !Invalid.IsStopDisable() {
		t.Error("Invalid.IsStopDisable should be true (alias for IsUndefined)")
	}
}

func TestServiceState_ActionAccessors(t *testing.T) {
	v := Start
	if v.ValueUInt16() != uint16(Start) || v.ValueInt() != int(Start) {
		t.Error("numeric accessor mismatch")
	}
	if len(v.AllNameValues()) == 0 || len(v.IntegerEnumRanges()) == 0 {
		t.Error("ranges empty")
	}
	if v.MinValueString() == "" || v.MaxValueString() == "" {
		t.Error("Min/Max value string empty")
	}
	min, max := v.MinMaxAny()
	if min == nil || max == nil {
		t.Error("MinMaxAny nil")
	}
	if len(v.RangesDynamicMap()) == 0 {
		t.Error("RangesDynamicMap empty")
	}
	if !v.IsAnyNamesOf("start", "stop") {
		t.Error("IsAnyNamesOf mismatch")
	}
	if !v.IsAnyValuesEqual(byte(Start)) {
		t.Error("IsAnyValuesEqual mismatch")
	}
	if !v.IsByteValueEqual(byte(Start)) || !v.IsNameEqual("start") {
		t.Error("IsByteValueEqual / IsNameEqual mismatch")
	}
	if err := v.OnlySupportedErr("start"); err != nil {
		t.Errorf("OnlySupportedErr unexpected: %v", err)
	}
	if err := v.OnlySupportedMsgErr("ctx", "start"); err != nil {
		t.Errorf("OnlySupportedMsgErr unexpected: %v", err)
	}
}

func TestServiceState_RangesAndCapitalNameMap(t *testing.T) {
	if Ranges[Start] != "start" || Ranges[Restart] != "restart" {
		t.Error("Ranges mapping mismatch")
	}
	if capitalNameMap[Start] != "Start" || capitalNameMap[Restart] != "Restart" {
		t.Error("capitalNameMap mismatch")
	}
	// actionToRequestMap covered indirectly; spot-check.
	if _, ok := actionToRequestMap[Start]; !ok {
		t.Error("actionToRequestMap missing Start")
	}
}
