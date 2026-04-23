package creationtests

import (
	"github.com/alimtvnetwork/core-v8/coreinterface/enuminf"
	"github.com/alimtvnetwork/core-v8/enums/stringcompareas"
	"github.com/alimtvnetwork/core-v8/issetter"
	"github.com/alimtvnetwork/core-v8/reqtype"
	"https://github.com/alimtvnetwork/enum-v1/accesstype"
	"https://github.com/alimtvnetwork/enum-v1/completionstate"
	"https://github.com/alimtvnetwork/enum-v1/dbaction"
	"https://github.com/alimtvnetwork/enum-v1/dbexposetype"
	"https://github.com/alimtvnetwork/enum-v1/eventtype"
	"https://github.com/alimtvnetwork/enum-v1/instructiontype"
	"https://github.com/alimtvnetwork/enum-v1/leveltype"
	"https://github.com/alimtvnetwork/enum-v1/licensetype"
	"https://github.com/alimtvnetwork/enum-v1/linuxservicestate"
	"https://github.com/alimtvnetwork/enum-v1/linuxtype"
	"https://github.com/alimtvnetwork/enum-v1/logtype"
	"https://github.com/alimtvnetwork/enum-v1/onofftype"
	"https://github.com/alimtvnetwork/enum-v1/osgroupexecution"
	"https://github.com/alimtvnetwork/enum-v1/overwritetype"
	"https://github.com/alimtvnetwork/enum-v1/pathpatterntype"
	"https://github.com/alimtvnetwork/enum-v1/resauthtype"
	"https://github.com/alimtvnetwork/enum-v1/revokereason"
	"https://github.com/alimtvnetwork/enum-v1/scripttype"
	"https://github.com/alimtvnetwork/enum-v1/servicestate"
	"https://github.com/alimtvnetwork/enum-v1/sqljointype"
	"https://github.com/alimtvnetwork/enum-v1/taskcategory"
	"https://github.com/alimtvnetwork/enum-v1/taskpriority"
)

var (
	bytesEnumContractsCollection = []enuminf.BasicEnumContractsBinder{
		reqtype.Invalid.AsBasicEnumContractsBinder(),
		stringcompareas.Invalid.AsBasicEnumContractsBinder(),
		accesstype.Invalid.AsBasicEnumContractsBinder(),
		completionstate.Invalid.AsBasicEnumContractsBinder(),
		dbaction.Invalid.AsBasicEnumContractsBinder(),
		dbexposetype.Invalid.AsBasicEnumContractsBinder(),
		eventtype.Invalid.AsBasicEnumContractsBinder(),
		instructiontype.Invalid.AsBasicEnumContractsBinder(),
		leveltype.Invalid.AsBasicEnumContractsBinder(),
		licensetype.Invalid.AsBasicEnumContractsBinder(),
		linuxservicestate.Invalid.AsBasicEnumContractsBinder(),
		linuxtype.Invalid.AsBasicEnumContractsBinder(),
		logtype.Invalid.AsBasicEnumContractsBinder(),
		onofftype.Invalid.AsBasicEnumContractsBinder(),
		osgroupexecution.Invalid.AsBasicEnumContractsBinder(),
		overwritetype.Invalid.AsBasicEnumContractsBinder(),
		pathpatterntype.Invalid.AsBasicEnumContractsBinder(),
		resauthtype.Invalid.AsBasicEnumContractsBinder(),
		revokereason.Unspecified.AsBasicEnumContractsBinder(),
		scripttype.Invalid.AsBasicEnumContractsBinder(),
		servicestate.Invalid.AsBasicEnumContractsBinder(),
		sqljointype.Invalid.AsBasicEnumContractsBinder(),
		taskcategory.Invalid.AsBasicEnumContractsBinder(),
		taskpriority.Invalid.AsBasicEnumContractsBinder(),
	}

	defaultOsScriptType = scripttype.OsDefaultScriptType()
	shellScriptType     = scripttype.Shell
	bashScriptType      = scripttype.Bash

	allScriptCreationTestCases = map[string]scripttype.Variant{
		"":                defaultOsScriptType,
		"def":             defaultOsScriptType,
		"default":         defaultOsScriptType,
		"Default":         defaultOsScriptType,
		"s":               shellScriptType,
		"sh":              shellScriptType,
		"Sh":              shellScriptType,
		"shell":           shellScriptType,
		"Shell":           shellScriptType,
		"/shell":          shellScriptType,
		"b":               bashScriptType,
		"bash":            bashScriptType,
		"Bash":            bashScriptType,
		"bh":              bashScriptType,
		"/bh":             bashScriptType,
		"Perl":            scripttype.Perl,
		"perl":            scripttype.Perl,
		"pl":              scripttype.Perl,
		"py":              scripttype.Python,
		"py2":             scripttype.Python2,
		"py3":             scripttype.Python3,
		"gcc":             scripttype.CLang,
		"gcc++":           scripttype.CLang,
		"c++":             scripttype.CLang,
		"CLang":           scripttype.CLang,
		"c":               scripttype.CLang,
		"Make":            scripttype.MakeScript,
		"MakeScript":      scripttype.MakeScript,
		"make":            scripttype.MakeScript,
		"m":               scripttype.MakeScript,
		"pw":              scripttype.Powershell,
		"pwsh":            scripttype.Powershell,
		"pwsh.exe":        scripttype.Powershell,
		"power":           scripttype.Powershell,
		"powershell":      scripttype.Powershell,
		"/powershell.exe": scripttype.Powershell,
		"Powershell":      scripttype.Powershell,
		"PowerShell":      scripttype.Powershell,
		"pshell":          scripttype.Powershell,
		"cmd":             scripttype.Cmd,
		"Cmd":             scripttype.Cmd,
		"dos":             scripttype.Cmd,
	}

	setterInvalid = issetter.Uninitialized
)
