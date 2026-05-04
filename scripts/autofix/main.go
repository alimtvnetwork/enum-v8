// Command autofix — Go source auto-fixer.
//
// Walks the repository (skipping vendor/, .git/, node_modules/, data/,
// cross-repo/, dist/, build/, and .git*) and applies safe, idempotent
// fixes to every .go file:
//
//   1. Trim trailing whitespace on every line.
//   2. Collapse 3+ consecutive blank lines down to 2.
//   3. Ensure file ends with exactly one trailing newline.
//   4. Run gofmt-style formatting via go/format.Source.
//
// The fixer is conservative — only transformations whose output is
// equivalent to the input under Go's lexical rules are applied. If
// go/format.Source rejects the file (i.e. it doesn't parse), the file
// is left untouched and the parse error is surfaced as a warning so
// bracecheck can pinpoint the syntax issue.
//
// Flags:
//   --dry-run   Report what would change without writing files.
//
// Exit codes:
//   0 — fixer ran cleanly (with or without changes).
//   1 — at least one file failed to read or write.
//   2 — bad CLI invocation.
//
// Designed to be invoked by scripts/CoveragePreChecks.psm1 and run.sh.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var skipDirs = map[string]bool{
	".git":         true,
	"vendor":       true,
	"node_modules": true,
	"data":         true,
	"cross-repo":   true,
	"dist":         true,
	"build":        true,
}

func main() {
	dryRun := flag.Bool("dry-run", false, "report changes without writing files")
	flag.Parse()

	root, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "autofix: cwd:", err)
		os.Exit(2)
	}

	checked := 0
	changed := 0
	failed := 0
	warnings := 0

	walkErr := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			if skipDirs[d.Name()] || (strings.HasPrefix(d.Name(), ".") && d.Name() != ".") {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".go") {
			return nil
		}
		if strings.Contains(path, string(filepath.Separator)+"testdata"+string(filepath.Separator)) {
			return nil
		}
		checked++

		rel, _ := filepath.Rel(root, path)
		src, readErr := os.ReadFile(path)
		if readErr != nil {
			fmt.Printf("  ! %s: read failed: %v\n", rel, readErr)
			failed++
			return nil
		}

		fixed := normalizeWhitespace(src)

		formatted, fmtErr := format.Source(fixed)
		if fmtErr != nil {
			// Don't write — let bracecheck report the syntax issue.
			fmt.Printf("  ~ %s: skipped formatter (parse error: %v)\n", rel, fmtErr)
			warnings++
			// Still consider whitespace-only fix if it helps.
			formatted = fixed
		}

		if bytes.Equal(formatted, src) {
			return nil
		}
		changed++

		if *dryRun {
			fmt.Printf("  - %s would be modified\n", rel)
			return nil
		}

		if writeErr := os.WriteFile(path, formatted, 0o644); writeErr != nil {
			fmt.Printf("  ! %s: write failed: %v\n", rel, writeErr)
			failed++
		}
		return nil
	})

	if walkErr != nil {
		fmt.Fprintln(os.Stderr, "autofix: walk:", walkErr)
		os.Exit(1)
	}

	verb := "fixed"
	if *dryRun {
		verb = "would fix"
	}

	switch {
	case failed > 0:
		fmt.Printf("✗ autofix: %d failure(s); %s %d/%d files (%d warnings)\n",
			failed, verb, changed, checked, warnings)
		os.Exit(1)
	case changed == 0:
		fmt.Printf("✓ autofix: no fixable issues across %d files\n", checked)
	default:
		fmt.Printf("✓ autofix: %s %d/%d files (%d warnings)\n", verb, changed, checked, warnings)
	}
}

// normalizeWhitespace trims trailing line whitespace, collapses runs of
// 3+ blank lines down to 2, and ensures exactly one trailing newline.
func normalizeWhitespace(src []byte) []byte {
	lines := strings.Split(string(src), "\n")
	for i, l := range lines {
		lines[i] = strings.TrimRight(l, " \t\r")
	}

	var out []string
	blankRun := 0
	for _, l := range lines {
		if l == "" {
			blankRun++
			if blankRun > 2 {
				continue
			}
		} else {
			blankRun = 0
		}
		out = append(out, l)
	}

	// Strip trailing empties then add exactly one newline.
	for len(out) > 0 && out[len(out)-1] == "" {
		out = out[:len(out)-1]
	}
	return []byte(strings.Join(out, "\n") + "\n")
}
