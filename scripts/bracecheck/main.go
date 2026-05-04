// Command bracecheck — Go syntax pre-check.
//
// Walks the repository (skipping vendor/, .git/, node_modules/, data/,
// cross-repo/, and any path containing /testdata/) and verifies every
// .go file:
//
//   1. Parses with go/parser.ParseFile (skips files with build-tag
//      issues that wouldn't compile on the host platform anyway).
//   2. Has balanced braces, brackets, and parens — a fast lexical
//      check that catches mismatches the parser sometimes reports
//      with confusing line numbers.
//
// Exit codes:
//   0 — all files clean.
//   1 — at least one file has a syntax or balance issue. Each issue
//       is printed on stdout in the form:
//         <relpath>:<line>:<col>: <message>
//
// Designed to be invoked by scripts/CoveragePreChecks.psm1 and run.sh
// before the slower coverage runner.
package main

import (
	"fmt"
	"go/parser"
	"go/token"
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
	root, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "bracecheck: cwd:", err)
		os.Exit(2)
	}

	var issues []string
	checked := 0

	walkErr := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil // tolerate transient FS errors
		}
		if d.IsDir() {
			if skipDirs[d.Name()] || strings.HasPrefix(d.Name(), ".") && d.Name() != "." {
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
			issues = append(issues, fmt.Sprintf("%s:0:0: cannot read: %v", rel, readErr))
			return nil
		}

		// Fast lexical balance check (skips strings, runes, comments).
		if msg, line, col := checkBalance(src); msg != "" {
			issues = append(issues, fmt.Sprintf("%s:%d:%d: %s", rel, line, col, msg))
			return nil
		}

		// Full parse (parser errors include precise positions).
		fset := token.NewFileSet()
		if _, parseErr := parser.ParseFile(fset, path, src, parser.AllErrors); parseErr != nil {
			// Surface each parser error on its own line.
			for _, line := range strings.Split(strings.TrimSpace(parseErr.Error()), "\n") {
				issues = append(issues, fmt.Sprintf("%s: %s", rel, strings.TrimSpace(line)))
			}
		}
		return nil
	})

	if walkErr != nil {
		fmt.Fprintln(os.Stderr, "bracecheck: walk:", walkErr)
		os.Exit(2)
	}

	if len(issues) > 0 {
		for _, i := range issues {
			fmt.Println(i)
		}
		fmt.Printf("✗ bracecheck: %d issue(s) across %d file(s)\n", len(issues), checked)
		os.Exit(1)
	}

	fmt.Printf("✓ bracecheck: %d files clean\n", checked)
}

// checkBalance scans src ignoring contents of strings, runes, and comments.
// Returns (msg, line, col) on the first imbalance, or ("", 0, 0) if balanced.
func checkBalance(src []byte) (string, int, int) {
	type pos struct {
		ch        byte
		line, col int
	}
	var stack []pos
	line, col := 1, 1

	advance := func(b byte) {
		if b == '\n' {
			line++
			col = 1
		} else {
			col++
		}
	}

	matches := map[byte]byte{')': '(', ']': '[', '}': '{'}

	for i := 0; i < len(src); i++ {
		b := src[i]

		// Line comment
		if b == '/' && i+1 < len(src) && src[i+1] == '/' {
			for i < len(src) && src[i] != '\n' {
				i++
			}
			line++
			col = 1
			continue
		}
		// Block comment
		if b == '/' && i+1 < len(src) && src[i+1] == '*' {
			i += 2
			for i+1 < len(src) && !(src[i] == '*' && src[i+1] == '/') {
				if src[i] == '\n' {
					line++
					col = 1
				} else {
					col++
				}
				i++
			}
			i++ // consume the trailing '/'
			col += 2
			continue
		}
		// String literals
		if b == '"' {
			start := pos{b, line, col}
			advance(b)
			i++
			for i < len(src) && src[i] != '"' {
				if src[i] == '\\' && i+1 < len(src) {
					advance(src[i])
					i++
				}
				if src[i] == '\n' {
					return "unterminated string literal", start.line, start.col
				}
				advance(src[i])
				i++
			}
			if i >= len(src) {
				return "unterminated string literal", start.line, start.col
			}
			advance(src[i])
			continue
		}
		// Raw strings
		if b == '`' {
			advance(b)
			i++
			for i < len(src) && src[i] != '`' {
				advance(src[i])
				i++
			}
			if i >= len(src) {
				return "unterminated raw string literal", line, col
			}
			advance(src[i])
			continue
		}
		// Rune literals
		if b == '\'' {
			advance(b)
			i++
			for i < len(src) && src[i] != '\'' {
				if src[i] == '\\' && i+1 < len(src) {
					advance(src[i])
					i++
				}
				advance(src[i])
				i++
			}
			if i >= len(src) {
				return "unterminated rune literal", line, col
			}
			advance(src[i])
			continue
		}

		switch b {
		case '(', '[', '{':
			stack = append(stack, pos{b, line, col})
		case ')', ']', '}':
			if len(stack) == 0 {
				return fmt.Sprintf("unmatched closing '%c'", b), line, col
			}
			top := stack[len(stack)-1]
			if top.ch != matches[b] {
				return fmt.Sprintf("mismatched bracket: '%c' opened at %d:%d, closed by '%c'",
					top.ch, top.line, top.col, b), line, col
			}
			stack = stack[:len(stack)-1]
		}
		advance(b)
	}

	if len(stack) > 0 {
		top := stack[len(stack)-1]
		return fmt.Sprintf("unclosed '%c'", top.ch), top.line, top.col
	}
	return "", 0, 0
}
