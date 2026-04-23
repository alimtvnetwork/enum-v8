package scripttype

import (
	"github.com/alimtvnetwork/core-v8/constants"
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	bashDefaultScript      = RangesMap[Bash]
	cmdDefaultScript       = RangesMap[Powershell]
	scriptTypeStringRanges = [...]string{
		Default:    "Default",
		Shell:      "Shell",
		Bash:       "Bash",
		Perl:       "Perl",
		Python:     "Python",
		Python2:    "Python2",
		Python3:    "Python3",
		CLang:      "CLang",
		MakeScript: "MakeScript",
		Powershell: "Powershell",
		Cmd:        "Cmd",
		Invalid:    "Invalid",
	}

	scriptTypeRanges = [...]Variant{
		Default:    Default,
		Shell:      Shell,
		Bash:       Bash,
		Perl:       Perl,
		Python:     Python,
		Python2:    Python2,
		Python3:    Python3,
		CLang:      CLang,
		MakeScript: MakeScript,
		Powershell: Powershell,
		Cmd:        Cmd,
		Invalid:    Invalid,
	}

	RangesMap = map[Variant]*ScriptDefault{
		Invalid: {
			ScriptType: Invalid,
		},
		Default: {
			ScriptType:  Default,
			ProcessName: "",
			DefaultArguments: []string{
				constants.NonInteractiveFlag,
			},
		},
		Shell: {
			ScriptType:  Shell,
			ProcessName: constants.BinShellCmd,
			DefaultArguments: []string{
				constants.NonInteractiveFlag,
			},
			IsImplemented: true,
		},
		Bash: {
			ScriptType:  Bash,
			ProcessName: constants.BashDefaultPath,
			DefaultArguments: []string{
				constants.NonInteractiveFlag,
			},
			IsImplemented: true,
		},
		Perl: {
			ScriptType:  Perl,
			ProcessName: "perl",
			DefaultArguments: []string{
				constants.NonInteractivePerlFlag,
			},
			IsImplemented: true,
		},
		Python: {
			ScriptType:  Python,
			ProcessName: "python",
			DefaultArguments: []string{
				constants.NonInteractiveFlag,
			},
			IsImplemented: false,
		},
		Python2: {
			ScriptType:  Python2,
			ProcessName: "python2",
			DefaultArguments: []string{
				constants.NonInteractiveFlag,
			},
			IsImplemented: false,
		},
		Python3: {
			ScriptType:  Python3,
			ProcessName: "python3",
			DefaultArguments: []string{
				constants.NonInteractiveFlag,
			},
			IsImplemented: true,
		},
		CLang: {
			ScriptType:  CLang,
			ProcessName: "gcc",
			DefaultArguments: []string{
				constants.NonInteractiveFlag,
			},
			IsImplemented: false,
		},
		MakeScript: {
			ScriptType:  MakeScript,
			ProcessName: "python3",
			DefaultArguments: []string{
				constants.NonInteractiveFlag,
			},
			IsImplemented: true,
		},
		Powershell: {
			ScriptType:  Powershell,
			ProcessName: "pwsh",
			DefaultArguments: []string{
				constants.NonInteractiveFlag,
			},
			IsImplemented: true,
		},
		Cmd: {
			ScriptType:  Cmd,
			ProcessName: "cmd",
			DefaultArguments: []string{
				constants.NonInteractiveCmdFlag,
			},
			IsImplemented: true,
		},
	}

	osDefaultScriptType = OsDefaultScriptType()
	defaultOsTypeVal    = osDefaultScriptType.ValueByte()

	aliasMap = map[string]byte{
		"":                defaultOsTypeVal,
		"def":             defaultOsTypeVal,
		"default":         defaultOsTypeVal,
		"Default":         defaultOsTypeVal,
		"s":               Shell.ValueByte(),
		"sh":              Shell.ValueByte(),
		"Sh":              Shell.ValueByte(),
		"shell":           Shell.ValueByte(),
		"/shell":          Shell.ValueByte(),
		"b":               Bash.ValueByte(),
		"bash":            Bash.ValueByte(),
		"bh":              Bash.ValueByte(),
		"/bh":             Bash.ValueByte(),
		"perl":            Perl.ValueByte(),
		"pl":              Perl.ValueByte(),
		"py":              Python.ValueByte(),
		"py2":             Python2.ValueByte(),
		"py3":             Python3.ValueByte(),
		"gcc":             CLang.ValueByte(),
		"gcc++":           CLang.ValueByte(),
		"c++":             CLang.ValueByte(),
		"clang":           CLang.ValueByte(),
		"c":               CLang.ValueByte(),
		"make":            MakeScript.ValueByte(),
		"Make":            MakeScript.ValueByte(),
		"m":               MakeScript.ValueByte(),
		"pw":              Powershell.ValueByte(),
		"pwsh":            Powershell.ValueByte(),
		"pwsh.exe":        Powershell.ValueByte(),
		"power":           Powershell.ValueByte(),
		"powershell":      Powershell.ValueByte(),
		"/powershell.exe": Powershell.ValueByte(),
		"Powershell":      Powershell.ValueByte(),
		"PowerShell":      Powershell.ValueByte(),
		"pshell":          Powershell.ValueByte(),
		"p-shell":         Powershell.ValueByte(),
		"cmd":             Cmd.ValueByte(),
		"dos":             Cmd.ValueByte(),
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingFirstItemSliceAliasMap(
		Invalid,
		scriptTypeStringRanges[:],
		aliasMap)
)
