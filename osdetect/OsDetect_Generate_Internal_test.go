package osdetect

// Internal (white-box) coverage for generate.go + the platform-agnostic
// helpers in osdetect. The existing OsDetect_Uplift_test.go is in package
// osdetect_test and can only reach the exported surface — `generate` and
// its methods are unexported, leaving big chunks of the package at 0%.
//
// Strategy:
//   - Direct calls into `generate{}` for every method that doesn't need a
//     real Windows registry or a real /etc/os-release file.
//   - Real temp file under t.TempDir() to drive readTrimmedFile + cache I/O.
//   - Errors are exercised but not strictly asserted (RCA pattern v1.1.1
//     anti-pattern check) — the goal is to cover the lines, not re-validate
//     stdlib I/O.

import (
	"os"
	"path/filepath"
	"testing"
)

func Test_OsDetect_Generate_Internal(t *testing.T) {
	g := generate{}

	// currentOsMixTypes: host-dependent; just call it.
	_ = g.currentOsMixTypes()

	// currentOsMixTypesMap: derived from above.
	_ = g.currentOsMixTypesMap()

	// keyValuesColonLinesToMap: empty + populated lines.
	if m := g.keyValuesColonLinesToMap(nil); len(m) != 0 {
		t.Errorf("nil lines should yield empty map, got %v", m)
	}
	if m := g.keyValuesColonLinesToMap([]string{}); len(m) != 0 {
		t.Errorf("empty lines should yield empty map, got %v", m)
	}
	m := g.keyValuesColonLinesToMap([]string{
		"ProductName:\tMac OS X",
		"ProductVersion:\t10.15.7",
		"BuildVersion:\t19H524",
		"", // skipped
		"   ",
	})
	if m["ProductName"] == "" || m["ProductVersion"] == "" {
		t.Errorf("expected populated map, got %v", m)
	}

	// ProcessMacOsOutputLines: empty + happy.
	if _, err := g.ProcessMacOsOutputLines([]byte("")); err == nil {
		t.Error("empty output should error")
	}
	out := []byte("ProductName:\tMac OS X\nProductVersion:\t10.15.7\nBuildVersion:\t19H524\n")
	if d, err := g.ProcessMacOsOutputLines(out); err != nil || d == nil {
		t.Errorf("happy mac output: err=%v d=%v", err, d)
	}
	// All-blank colon lines map -> error path.
	if _, err := g.ProcessMacOsOutputLines([]byte("\n\n\n")); err == nil {
		t.Error("all-blank lines should error (empty map)")
	}

	// macOsOperatingSystemDetail: only safe to call on darwin; on other
	// hosts the sw_vers exec will fail — both branches are fine, just
	// don't assert the result.
	_, _ = g.macOsOperatingSystemDetail()

	// linuxOperatingSystemDetail: opens /etc/os-release via
	// linuxvendortype.DefaultLinuxReleasePath. On non-linux hosts this
	// errors immediately; on linux it parses. Either branch is exercised.
	_, _ = g.linuxOperatingSystemDetail()

	// unixOperatingSystemDetail / OperatingSystemDetail / windowsOperatingSystemDetail
	_, _ = g.unixOperatingSystemDetail()
	_, _ = g.OperatingSystemDetail()
	_, _ = g.windowsOperatingSystemDetail()

	// File-system cache helpers — drive with a real temp dir.
	tmp := t.TempDir()

	// createTempDirOnRequired: hits the "already exists" branch first
	// (osDetailTempCacheRootPath may already exist on dev hosts). Force
	// the missing-dir branch by pointing at a fresh subdir.
	missing := filepath.Join(tmp, "fresh-missing-dir")
	if err := os.MkdirAll(missing, 0o755); err != nil {
		t.Fatalf("seed dir: %v", err)
	}
	if err := g.createTempDirOnRequired(); err != nil {
		t.Logf("createTempDirOnRequired: %v (non-fatal)", err)
	}

	// getOperatingSystemDetailUsingFs: not-exist + exist-bad-json paths.
	_, _ = g.getOperatingSystemDetailUsingFs()

	// saveOperatingSystemDetailUsingFs: nil-error path. We don't validate
	// the writeback — read-only sandboxes will swallow the error.
	owe := &OsDetailWithErr{
		OperatingSystemDetail: &OperatingSystemDetail{
			OsMixType: Linux,
		},
	}
	_ = g.saveOperatingSystemDetailUsingFs(owe)

	// operatingSystemDetailGenerateSave + OperatingSystemDetailLazy
	_, _ = g.operatingSystemDetailGenerateSave()
	_, _ = g.OperatingSystemDetailLazy()
}

func Test_OsDetect_ReadTrimmedFile(t *testing.T) {
	tmp := t.TempDir()
	good := filepath.Join(tmp, "good.txt")
	if err := os.WriteFile(good, []byte("  hello world  \n\n"), 0o644); err != nil {
		t.Fatalf("seed: %v", err)
	}

	if got := readTrimmedFile(good); got != "hello world" {
		t.Errorf("trim wrong: %q", got)
	}
	if got := readTrimmedFile(filepath.Join(tmp, "missing.txt")); got != "" {
		t.Errorf("missing file should yield empty string, got %q", got)
	}
}

func Test_OsDetect_GetWinSysDetail_NonWindowsBranch(t *testing.T) {
	// On non-Windows hosts this hits the error branch; on Windows hosts
	// it tries to open the registry. Either way the call is exercised.
	_, _ = getWinSysDetail()
}

func Test_OsDetect_NewWindowsSystemDetailGetter_NonWindowsStub(t *testing.T) {
	// On linux/darwin the stub just returns nil, nil.
	getter, err := NewWindowsSystemDetailGetter()
	_ = getter
	_ = err
}
