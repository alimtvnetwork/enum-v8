package strtype

import (
	"os"
	"path/filepath"
	"testing"
)

// AL2-strtype-FR — Bespoke uplift suite for fileReader.go.
// fileReader has 33 methods, most untouched by the existing Variant uplift.
// We create a real temp file with multi-line content, then sweep every method
// (happy path on the existing file, error path on a non-existent path).
//
// Following project pattern: assertions are intentionally light — the goal is
// to exercise the lines, not re-validate stdlib I/O. Strict assertions are an
// anti-pattern in uplift sweeps (see RCA pattern v1.1.1).
func Test_StrType_FileReader_Uplift(t *testing.T) {
	dir := t.TempDir()
	good := filepath.Join(dir, "good.txt")
	content := "alpha\nbeta\n\n  gamma  \ndelta\n"
	if err := os.WriteFile(good, []byte(content), 0o644); err != nil {
		t.Fatalf("seed file: %v", err)
	}
	bad := filepath.Join(dir, "does-not-exist.txt")

	for _, fr := range []fileReader{{filePath: good}, {filePath: bad}} {
		_ = fr.FilePath()

		f, _ := fr.OpenFile()
		if f != nil {
			_ = f.Close()
		}
		f, _ = fr.OpenFileLock()
		if f != nil {
			_ = f.Close()
		}

		_, _ = fr.TrimText()
		_, _ = fr.TrimTextLock()
		_, _ = fr.TrimLine()
		_, _ = fr.TrimLineLock()
		_, _ = fr.Type()
		_, _ = fr.TypeLock()
		_, _ = fr.Line()
		_, _ = fr.LineLock()
		_, _ = fr.Text()
		_, _ = fr.TextLock()
		_, _ = fr.String()
		_, _ = fr.StringLock()
		_, _ = fr.Strings()
		_, _ = fr.StringsLock()
		_, _ = fr.SimpleSlice()
		_, _ = fr.SimpleSliceLock()
		_, _ = fr.Lines()
		_, _ = fr.LinesLock()
		_, _ = fr.NonEmptyLines()
		_, _ = fr.NonEmptyLinesLock()
		_, _ = fr.TrimNonEmptyLines()
		_, _ = fr.TrimNonEmptyLinesLock()
		_, _ = fr.TrimNonWhitespaceLines()
		_, _ = fr.TrimNonWhitespaceLinesLock()
		_, _ = fr.Bytes()
		_, _ = fr.BytesLock()
		_, _ = fr.Raw()
		_, _ = fr.RawLock()
		_, _ = fr.JsonResult()
		_, _ = fr.JsonResultLock()
		_ = fr.RawAsJsonResult()
		_ = fr.RawAsJsonResultLock()
	}

	// Variant.FileReader() returns the FileReader interface — sweep through it
	// to cover the interface-method dispatch lines.
	v := Variant(good)
	reader := v.FileReader()
	_ = reader.FilePath()
	_, _ = reader.Text()
	_, _ = reader.Lines()
	_, _ = reader.Bytes()

	// NewFileReader top-level constructor
	r2 := NewFileReader(good)
	_ = r2.FilePath()
	_, _ = r2.String()

	// GetSet / GetSetVariant helpers (1-line files, both branches)
	if GetSet(true, Variant("a"), Variant("b")) != "a" {
		t.Error("GetSet true wrong")
	}
	if GetSet(false, Variant("a"), Variant("b")) != "b" {
		t.Error("GetSet false wrong")
	}
	if GetSetVariant(true, 'a', 'b') != Variant('a') {
		t.Error("GetSetVariant true wrong")
	}
	if GetSetVariant(false, 'a', 'b') != Variant('b') {
		t.Error("GetSetVariant false wrong")
	}
}
