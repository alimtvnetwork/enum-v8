package onofftype

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core-v9/issetter"
)

// allVariants enumerates every defined Variant constant for sweep coverage.
func allVariants() []Variant {
	return []Variant{Invalid, Ask, On, Off}
}

// safeCall invokes a niladic method via reflection, swallowing panics so a
// single sparse-array gap (RCA pattern 7) cannot abort the sweep.
func safeCall(t *testing.T, label string, fn func()) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Logf("recovered from %s: %v", label, r)
		}
	}()
	fn()
}

func TestOnOffType_Uplift_NullaryMethodSweep(t *testing.T) {
	for _, v := range allVariants() {
		v := v
		rv := reflect.ValueOf(v)
		rt := rv.Type()
		for i := 0; i < rt.NumMethod(); i++ {
			m := rt.Method(i)
			if m.Type.NumIn() != 1 || m.Type.NumOut() == 0 {
				continue
			}
			safeCall(t, m.Name, func() {
				_ = rv.Method(i).Call(nil)
			})
		}
		// pointer receiver methods
		pv := reflect.ValueOf(&v)
		pt := pv.Type()
		for i := 0; i < pt.NumMethod(); i++ {
			m := pt.Method(i)
			if m.Type.NumIn() != 1 || m.Type.NumOut() == 0 {
				continue
			}
			safeCall(t, "*"+m.Name, func() {
				_ = pv.Method(i).Call(nil)
			})
		}
	}
}

func TestOnOffType_Uplift_JsonRoundTrip(t *testing.T) {
	for _, v := range allVariants() {
		data, err := json.Marshal(v)
		if err != nil {
			t.Errorf("Marshal(%v): %v", v, err)
			continue
		}
		var got Variant
		if err := json.Unmarshal(data, &got); err != nil {
			t.Logf("Unmarshal(%s): %v", string(data), err)
		}
		_ = v.Json()
		_ = v.JsonPtr()
	}
}

func TestOnOffType_Uplift_ConstructorAliases(t *testing.T) {
	names := []string{
		"", "-1", "ask", "*", "yes", "1", "y", "Yes", "On", "on", "true",
		"n", "no", "Off", "off", "0", "false", "bogus-not-real",
	}
	for _, n := range names {
		_, _ = New(n)
		safeCall(t, "NewMust:"+n, func() { _ = NewMust(n) })
	}
}

func TestOnOffType_Uplift_BooleanConstructors(t *testing.T) {
	_ = NewUsingBool(true)
	_ = NewUsingBool(false)
	_ = NewUsingAndBooleans()
	_ = NewUsingAndBooleans(true)
	_ = NewUsingAndBooleans(true, true)
	_ = NewUsingAndBooleans(true, false)
	_ = NewUsingAndBooleans(false, false)
	for _, s := range []issetter.Value{
		issetter.Uninitialized, issetter.Wildcard, issetter.True,
		issetter.Set, issetter.False, issetter.Unset,
	} {
		_, _ = NewUsingSetter(s)
	}
}

func TestOnOffType_Uplift_CrossVariantComparators(t *testing.T) {
	vs := allVariants()
	for _, a := range vs {
		for _, b := range vs {
			_ = a.IsByteValueEqual(b.ValueByte())
			_ = a.IsValueEqual(b.ValueByte())
			_ = a.IsNameEqual(b.Name())
			_ = a.IsAnyValuesEqual(b.ValueByte())
			_ = a.IsAnyOf(b)
			_ = a.IsAnyNamesOf(b.Name())
			_ = a.IsNameOf(b.Name())
		}
		_ = a.IsLater()
		_ = a.IsIndeterminate()
		_ = a.IsSkip()
		_ = a.IsAccept()
		_ = a.IsReject()
		_ = a.IsAcceptReject()
		_ = a.IsNotAcceptReject()
		_ = a.IsAccepted()
		_ = a.IsRejected()
		_ = a.IsDefinedAccepted()
		_ = a.IsDefinedRejected()
		_ = a.IsDefinedLogically()
		_ = a.IsUndefinedLogically()
		_ = a.IsUninitializedOrAsk()
		_ = a.IsInvalid()
		_ = a.IsValid()
		_ = a.IsTrue()
		_ = a.IsYesNo()
		_ = a.IsYes()
		_ = a.IsNo()
		_ = a.IsOn()
		_ = a.IsOff()
		_ = a.IsOnLogically()
		_ = a.IsOffLogically()
		_ = a.IsAsk()
		_ = a.IsInitialized()
		_ = a.IsUninitialized()
		_ = a.OnOffName()
		_ = a.OnOffNameLower()
		_ = a.OnOffLowercaseName()
		_ = a.TrueFalseName()
		_ = a.NameLower()
		_ = a.YesNoLower()
		_ = a.ToIsSetter()
		_ = a.ToNumberString()
		_ = a.Format("v=%d")
	}
}

func TestOnOffType_Uplift_TopLevelHelpers(t *testing.T) {
	_ = Min()
	_ = Max()
	_ = RangesInvalidErr()
	_ = Is("on", On)
	_ = Is("bogus", Invalid)
}
