package sshcmdnames

import (
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:              "Invalid",
		Help:                 "Help",
		Install:              "InstallInstall",
		Fix:                  "Fix",
		InstallFix:           "InstallFix",
		Enable:               "Enable",
		Disable:              "Disable",
		SetPort:              "SetPort",
		EnableAtPort:         "EnableAtPort",
		DisablePort:          "DisablePort",
		AllowIps:             "AllowIps",
		DisableIps:           "DisableIps",
		AddSshKey:            "AddSshKey",
		AddAuthorizedKey:     "AddAuthorizedKey",
		RemoveAuthKey:        "RemoveAuthKey",
		GenerateKey:          "GenerateKey",
		DisablePasswordLogin: "DisablePasswordLogin",
		EnablePasswordLogin:  "EnablePasswordLogin",
		WhichPort:            "WhichPort",
		PortStatus:           "PortStatus",
		Config:               "Config",
		SetBannerText:        "SetBannerText",
		RemoveBanner:         "RemoveBanner",
		ApplyDefaultConfig:   "ApplyDefaultConfig",
		Status:               "Status",
		ListAuthKeys:         "ListAuthKeys",
		ListAuthKeysJson:     "ListAuthKeysJson",
		ListJson:             "ListJson",
		List:                 "List",
		ConfigPaths:          "ConfigPaths",
		Histories:            "Histories",
		Backup:               "Backup",
		Import:               "Import",
	}

	aliasMap = map[string]byte{
		"set-port":               SetPort.Value(),
		"apply-default-config":   ApplyDefaultConfig.Value(),
		"fix":                    Fix.Value(),
		"install-fix":            InstallFix.Value(),
		"default-config":         ApplyDefaultConfig.Value(),
		"allow-ips":              AllowIps.Value(),
		"deny-ips":               DisableIps.Value(),
		"set-banner":             SetBannerText.Value(),
		"remove-banner":          RemoveBanner.Value(),
		"add-authorized-key":     AddAuthorizedKey.Value(),
		"add-auth-key":           AddAuthorizedKey.Value(),
		"remove-auth-key":        RemoveAuthKey.Value(),
		"remove-authorized-key":  RemoveAuthKey.Value(),
		"gen-key":                GenerateKey.Value(),
		"which-port":             WhichPort.Value(),
		"generate-key":           GenerateKey.Value(),
		"list-auth-keys":         ListAuthKeys.Value(),
		"list-auth-keys-json":    ListAuthKeysJson.Value(),
		"config-paths":           ConfigPaths.Value(),
		"disable-password-login": DisablePasswordLogin.Value(),
		"configure":              Config.Value(),
	}

	BasicEnumImpl = enumimpl.New.BasicByte.DefaultWithAliasMapAllCases(
		Invalid,
		Ranges[:],
		aliasMap)
)
