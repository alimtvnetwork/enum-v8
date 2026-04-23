package nginxlogtype

import (
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:                     "Invalid",
		Notice:                      "Notice",         // Generic
		Warning:                     "Warning",        // Generic
		Error:                       "Error",          // Generic
		AlertError:                  "AlertError",     // Generic
		FileIssueError:              "FileIssueError", // Specific
		SyntaxIssueError:            "SyntaxIssueError",
		DuplicateDomainWarningError: "DuplicateDomainWarningError",
		DuplicateDefaultError:       "DuplicateDefaultError",
	}

	RangesMap = map[string]Variant{
		"Invalid":                     Invalid,
		"Notice":                      Notice,
		"Warning":                     Warning,
		"Error":                       Error,
		"AlertError":                  AlertError,
		"FileIssueError":              FileIssueError,
		"SyntaxIssueError":            SyntaxIssueError,
		"DuplicateDomainWarningError": DuplicateDomainWarningError,
		"DuplicateDefaultError":       DuplicateDefaultError,
		"unknown":                     Invalid,
		"notice":                      Notice,
		"warn":                        Warning,
		"emerg":                       Error,
		"alert":                       AlertError,
	}

	errorMap = map[Variant]bool{
		Error:                       true,
		AlertError:                  true,
		FileIssueError:              true,
		SyntaxIssueError:            true,
		DuplicateDomainWarningError: true,
		DuplicateDefaultError:       true,
	}

	BasicEnumImpl = enumimpl.New.BasicByte.DefaultAllCases(
		Invalid,
		Ranges[:])
)
