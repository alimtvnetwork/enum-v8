package creationtests

import (
	"github.com/alimtvnetwork/core-v9/coreinterface/enuminf"
	"github.com/alimtvnetwork/core-v9/enums/stringcompareas"
	"github.com/alimtvnetwork/core-v9/issetter"
	"github.com/alimtvnetwork/core-v9/reqtype"
	"github.com/alimtvnetwork/enum-v8/accesstype"
	"github.com/alimtvnetwork/enum-v8/completionstate"
	"github.com/alimtvnetwork/enum-v8/dbaction"
	"github.com/alimtvnetwork/enum-v8/dbexposetype"
	"github.com/alimtvnetwork/enum-v8/eventtype"
	"github.com/alimtvnetwork/enum-v8/instructiontype"
	"github.com/alimtvnetwork/enum-v8/leveltype"
	"github.com/alimtvnetwork/enum-v8/licensetype"
	"github.com/alimtvnetwork/enum-v8/linuxservicestate"
	"github.com/alimtvnetwork/enum-v8/linuxtype"
	"github.com/alimtvnetwork/enum-v8/logtype"
	"github.com/alimtvnetwork/enum-v8/onofftype"
	"github.com/alimtvnetwork/enum-v8/osgroupexecution"
	"github.com/alimtvnetwork/enum-v8/overwritetype"
	"github.com/alimtvnetwork/enum-v8/pathpatterntype"
	"github.com/alimtvnetwork/enum-v8/resauthtype"
	"github.com/alimtvnetwork/enum-v8/revokereason"
	"github.com/alimtvnetwork/enum-v8/scripttype"
	"github.com/alimtvnetwork/enum-v8/servicestate"
	"github.com/alimtvnetwork/enum-v8/sqljointype"
	"github.com/alimtvnetwork/enum-v8/taskcategory"
	"github.com/alimtvnetwork/enum-v8/taskpriority"
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
