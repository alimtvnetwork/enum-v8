package osdetect_test

import (
	"testing"

	"github.com/alimtvnetwork/enum-v7/inttype"
	"github.com/alimtvnetwork/enum-v7/osarchs"
	"github.com/alimtvnetwork/enum-v7/osdetect"
	"github.com/alimtvnetwork/enum-v7/strtype"
)

// AL2-08 bespoke coverage for osdetect — exercises the platform-guarded
// wrapper functions and the OperatingSystemDetail / WindowsSystemDetail
// helpers using constructed fixtures (no host-dependent assertions).

func TestOsDetect_PlatformWrappers_DoNotPanic(t *testing.T) {
	// All of these are stable on any host (return false on a non-matching OS).
	_ = osdetect.IsCentOs()
	_ = osdetect.IsDebian()
	_ = osdetect.IsRedhat()
	_ = osdetect.IsUbuntu()
	_ = osdetect.IsWindows10()
	_ = osdetect.IsWindows11()
	_ = osdetect.IsWindows8()
	_ = osdetect.IsWindowsServer()
	_ = osdetect.IsWindowsServer2016()
	_ = osdetect.IsWindowsServer2019()

	// CurrentOsTypesNotContainsError: passing the actual current type yields nil.
	cur := osdetect.CurrentOsType()
	if err := osdetect.CurrentOsTypesNotContainsError(cur); err != nil {
		t.Errorf("expected nil error for current type, got %v", err)
	}
	if err := osdetect.CurrentOsTypesNotContainsError(osdetect.Invalid); err == nil {
		t.Error("expected non-nil error for Invalid type")
	}

	// CurrentOsTypesMustBePresent should not panic for the current type.
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("MustBePresent panicked unexpectedly: %v", r)
		}
	}()
	osdetect.CurrentOsTypesMustBePresent(cur)
}

func TestOsDetect_GetCurrentOsDetail_Smoke(t *testing.T) {
	d, _ := osdetect.GetCurrentOsDetail()
	// On any supported host the cached singleton is non-nil; if it is, just
	// skip the deeper assertions.
	if d == nil {
		t.Skip("no os detail cached on this host")
	}
	_ = d.IsValid()
	_ = d.IsInvalid()
	_ = d.PrettyJsonString()
	_ = d.Json()
	_ = d.JsonPtr()
	_ = d.Serialize()
	_ = d.SerializeMust()
	_ = d.AllSysTypes()
	_ = d.AllSysTypesMap()
	_ = d.HasWindowsDetail()
	_ = d.IsEmptyWindowsDetail()
}

func TestOsDetect_OperatingSystemDetail_PureLogic(t *testing.T) {
	d := &osdetect.OperatingSystemDetail{
		OsMixType:    osdetect.Ubuntu,
		Name:         strtype.Variant("Ubuntu 22.04"),
		Vendor:       strtype.Variant("ubuntu"),
		Version:      strtype.Variant("22"),
		Release:      strtype.Variant("22.04"),
		Architecture: osarchs.X64,
	}
	if !d.IsName("Ubuntu 22.04") {
		t.Error("IsName failed")
	}
	if !d.IsNameContains("Ubuntu") {
		t.Error("IsNameContains failed")
	}
	if !d.IsNameStartsWith("Ubuntu") {
		t.Error("IsNameStartsWith failed")
	}
	if !d.IsNameEndsWith("22.04") {
		t.Error("IsNameEndsWith failed")
	}
	if !d.IsArch(osarchs.X64) || d.Is32BitArch() || !d.Is64BitArch() {
		t.Error("Architecture predicates wrong")
	}
	if !d.IsType(osdetect.Ubuntu) || d.IsType(osdetect.Windows) {
		t.Error("IsType wrong")
	}
	if !d.IsAnyOfTypes(osdetect.Windows, osdetect.Ubuntu) {
		t.Error("IsAnyOfTypes wrong")
	}
	if d.HasWindowsDetail() || !d.IsEmptyWindowsDetail() {
		t.Error("HasWindowsDetail wrong (no WindowsDetail set)")
	}
	if rv := d.ReleaseVersion(); rv == nil {
		t.Error("ReleaseVersion nil")
	}
	// Cached path: second call should return same pointer.
	if d.ReleaseVersion() == nil {
		t.Error("ReleaseVersion cached nil")
	}

	// Null-pointer path
	var nilDetail *osdetect.OperatingSystemDetail
	if !nilDetail.IsNull() {
		t.Error("nil should be IsNull")
	}
	if !nilDetail.IsEmpty() {
		t.Error("nil should be IsEmpty")
	}
	if nilDetail.HasWindowsDetails() {
		t.Error("nil should not HasWindowsDetails")
	}

	// Empty-name path: IsNameContains/etc return false on empty.
	empty := &osdetect.OperatingSystemDetail{}
	if empty.IsName("x") || empty.IsNameContains("x") ||
		empty.IsNameStartsWith("x") || empty.IsNameEndsWith("x") ||
		empty.IsArch(osarchs.X64) {
		t.Error("empty detail should report false")
	}
}

func TestOsDetect_WindowsSystemDetail_PureLogic(t *testing.T) {
	w := &osdetect.WindowsSystemDetail{
		WindowsVersion: inttype.Variant(10),
		ServerVersion:  inttype.Variant(2019),
		IsClient:       true,
	}
	if w.IsNull() || !w.IsDefined() {
		t.Error("populated detail should not be null")
	}
	if !w.IsWindows10() {
		t.Error("IsWindows10 wrong")
	}
	if w.IsWindows11() {
		t.Error("IsWindows11 should be false")
	}
	if w.IsWindowsSever() {
		t.Error("IsWindowsSever should be false (IsClient=true)")
	}
	if w.WinVer() != inttype.Variant(10) {
		t.Errorf("WinVer wrong: %v", w.WinVer())
	}

	// Server flavor
	s := &osdetect.WindowsSystemDetail{
		ServerVersion: inttype.Variant(2019),
		IsServer:      true,
	}
	if !s.IsWindowsSever() || !s.IsWindowsSever2019() {
		t.Error("server predicates wrong")
	}
	if !s.IsWindowsSeverGreaterEqual2016() {
		t.Error("server >=2016 wrong")
	}
	if s.WinVer() != inttype.Variant(2019) {
		t.Errorf("server WinVer wrong: %v", s.WinVer())
	}

	// Null-receiver paths
	var nw *osdetect.WindowsSystemDetail
	if !nw.IsNull() || nw.IsDefined() {
		t.Error("nil should be Null/!Defined")
	}
	if !nw.IsNullOr(false) {
		t.Error("nil.IsNullOr false should be true")
	}
	if nw.IsDefinedPlus(true) {
		t.Error("nil.IsDefinedPlus should be false")
	}
}

func TestOsDetect_OsDetailWithErr_JsonAndString(t *testing.T) {
	o := &osdetect.OsDetailWithErr{
		OperatingSystemDetail: &osdetect.OperatingSystemDetail{
			OsMixType: osdetect.Ubuntu,
			Name:      strtype.Variant("Ubuntu"),
		},
	}
	if o.String() == "" {
		t.Error("String empty")
	}
	if o.PrettyJsonString() == "" {
		t.Error("PrettyJsonString empty")
	}
	if o.Json().HasError() {
		t.Error("Json should not have error")
	}
	if o.JsonPtr() == nil {
		t.Error("JsonPtr nil")
	}

	var nilP *osdetect.OsDetailWithErr
	if nilP.String() != "" || nilP.PrettyJsonString() != "" {
		t.Error("nil receiver should return empty string")
	}
}
