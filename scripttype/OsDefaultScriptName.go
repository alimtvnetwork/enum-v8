package scripttype

import "github.com/alimtvnetwork/core-v8/osconsts"

func OsDefaultScriptType() Variant {
	if osconsts.IsWindows {
		return Powershell
	}

	return Bash
}
