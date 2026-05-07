package scripttype

import (
	"encoding/json"
	"testing"
)

// AL2 follow-up coverage uplift for scripttype using only the standard
// `testing` package. Touches every Variant accessor, predicate, helper,
// and ScriptDefault-backed lookup.
func TestScriptType_UpliftAllSurface(t *testing.T) {
	all := []Variant{
		Default, Shell, Bash, Perl, Python, Python2, Python3,
		CLang, MakeScript, Powershell, Cmd, Invalid,
	}

	for _, v := range all {
		_ = v.ValueByte()
		_ = v.ValueInt()
		_ = v.ValueInt8()
		_ = v.ValueInt16()
		_ = v.ValueInt32()
		_ = v.ValueUInt16()
		_ = v.ValueString()
		_ = v.Name()
		_ = v.String()
		_ = v.NameValue()
		(&v).ToNumberString()
		_ = v.RangeNamesCsv()
		_ = v.TypeName()
		_ = v.Format("%s")
		(&v).MaxByte()
		(&v).MinByte()
		_ = v.MaxInt()
		_ = v.MinInt()
		_ = v.MinValueString()
		_ = v.MaxValueString()
		_ = v.AllNameValues()
		_ = v.IntegerEnumRanges()
		_ = v.RangesDynamicMap()
		(&v).RangesByte()
		_, _ = v.MinMaxAny()
		_ = v.EnumType()
		_ = v.OsDefaultScriptType()
		(&v).RangesVariants()

		_ = v.IsValid()
		_ = v.IsInvalid()
		_ = v.IsUninitialized()
		_ = v.IsDefault()
		_ = v.IsDefaultOsScriptType()
		_ = v.IsShell()
		_ = v.IsBash()
		_ = v.IsPerl()
		_ = v.IsPython()
		_ = v.IsPython2()
		_ = v.IsPython3()
		_ = v.IsAnyPython()
		_ = v.IsCLang()
		_ = v.IsMakeScript()
		_ = v.IsPowershell()
		_ = v.IsCmd()
		_ = v.IsShellOrBash()
		_ = v.IsCmdOrPowershell()
		_ = v.IsPowershellOrShell()
		_ = v.IsPowershellOrShellOrBash()
		_ = v.IsValueEqual(v.ValueByte())
		_ = v.IsByteValueEqual(v.ValueByte())
		_ = v.IsNameEqual(v.Name())
		_ = v.IsAnyOf(Bash, Shell, Cmd)
		_ = v.IsAnyValuesEqual(byte(Bash), byte(Shell))
		_ = v.IsAnyNamesOf("Bash", "Shell")

		_ = v.OnlySupportedErr("Bash")
		_ = v.OnlySupportedMsgErr("ctx", "Bash")

		// ScriptDefault-backed helpers
		(&v).ScriptDefault()
		_ = (&v).ProcessName()
		_ = (&v).DefaultArgs()
		_ = (&v).IsImplemented()
		_ = (&v).IsInvalidImplement()

		// JSON round-trip on pointer receivers
		data, err := json.Marshal(&v)
		if err != nil {
			t.Fatalf("Marshal(%v): %v", v, err)
		}
		var got Variant
		if err := json.Unmarshal(data, &got); err != nil {
			t.Fatalf("Unmarshal(%s): %v", string(data), err)
		}

		_ = v.AsBasicByteEnumContractsBinder()
		_ = v.AsBasicEnumContractsBinder()
		_ = v.ToPtr()
	}

	// CondBool / CondFunc
	if CondBool(true, Bash, Cmd) != Bash || CondBool(false, Bash, Cmd) != Cmd {
		t.Fatal("CondBool wrong")
	}
	if CondFunc(true, func() Variant { return Powershell }) != Powershell {
		t.Fatal("CondFunc(true) wrong")
	}
	if CondFunc(false, func() Variant { return Powershell }) != Invalid {
		t.Fatal("CondFunc(false) wrong")
	}

	// Module-level helpers
	if DefaultOsScript() == nil {
		t.Fatal("DefaultOsScript nil")
	}
	if OsDefaultScriptType() == Invalid {
		t.Fatal("OsDefaultScriptType Invalid")
	}

	// New / NewMust paths
	if v, err := New("Bash"); err != nil || v != Bash {
		t.Fatalf("New(Bash) = %v, %v", v, err)
	}
	if NewMust("Powershell") != Powershell {
		t.Fatal("NewMust(Powershell) wrong")
	}
	if _, err := New("__nope__"); err == nil {
		t.Fatal("New(__nope__) should error")
	}

	if Min() != Default || Max() != Invalid {
		t.Fatalf("Min/Max: %v %v", Min(), Max())
	}
	if RangesInvalidErr() == nil {
		t.Fatal("RangesInvalidErr nil")
	}
}
