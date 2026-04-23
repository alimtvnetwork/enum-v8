package sslcmdnames

import (
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:               "Invalid",
		Help:                  "Help",
		AddDomain:             "AddDomain",
		RemoveDomain:          "RemoveDomain",
		AddSsl:                "AddSsl",
		DryRunSsl:             "DryRunSsl",
		RenewSsl:              "RenewSsl",
		RemoveSsl:             "RemoveSsl",
		MoveSsl:               "MoveSsl",
		BackupSsl:             "BackupSsl",
		RevokeSsl:             "RevokeSsl",
		AutoSsl:               "AutoSsl",
		AutoRenewSsl:          "AutoRenewSsl",
		RemoveAll:             "RemoveAll",
		ExpireSsl:             "ExpireSsl",
		SslChallenge:          "SslChallenge",
		HttpSslChallenge:      "HttpSslChallenge",
		DnsSslChallenge:       "DnsSslChallenge",
		LastStates:            "LastStates",
		Sync:                  "Sync",
		SyncNow:               "SyncNow",
		SyncForce:             "SyncForce",
		CurrentStateName:      "CurrentStateName",
		Histories:             "Histories",
		ExportConfigSpecific:  "ExportConfigSpecific",
		ExportConfigAll:       "ExportConfigAll",
		ImportConfigSpecific:  "ImportConfigSpecific",
		ImportConfigAll:       "ImportConfigAll",
		SearchDomains:         "SearchDomains",
		ListPorts:             "ListPorts",
		ListSslEnabledServers: "ListSslEnabledServers",
		SearchSslByDomain:     "SearchSslByDomain",
		SearchSslByUser:       "SearchSslByUser",
		SearchSslByPackage:    "SearchSslByPackage",
		ListAboutToExpire:     "ListAboutToExpire",
		ListExpiredSsl:        "ListExpiredSsl",
		Backup:                "Backup",
		Import:                "Import",
	}

	aliasMap = map[string]byte{
		"add-domain":           AddDomain.Value(),
		"domain":               AddDomain.Value(),
		"add-ssl":              AddSsl.Value(),
		"add":                  AddSsl.Value(),
		"dry-run-ssl":          DryRunSsl.Value(),
		"dry-run":              DryRunSsl.Value(),
		"renew-ssl":            RenewSsl.Value(),
		"renew":                RenewSsl.Value(),
		"remove-ssl":           RemoveSsl.Value(),
		"remove":               RemoveSsl.Value(),
		"delete":               RemoveSsl.Value(),
		"delete-ssl":           RemoveSsl.Value(),
		"move-ssl":             MoveSsl.Value(),
		"backup-ssl":           BackupSsl.Value(),
		"revoke-ssl":           RevokeSsl.Value(),
		"auto-ssl":             AutoSsl.Value(),
		"auto":                 AutoSsl.Value(),
		"auto-renew-ssl":       AutoRenewSsl.Value(),
		"remove-all":           RemoveAll.Value(),
		"expire-ssl":           ExpireSsl.Value(),
		"ssl-challenge":        SslChallenge.Value(),
		"http-challenge":       HttpSslChallenge.Value(),
		"dns-challenge":        DnsSslChallenge.Value(),
		"last-states":          LastStates.Value(),
		"search-domains":       SearchSslByDomain.Value(),
		"search-user":          SearchSslByUser.Value(),
		"search-package":       SearchSslByPackage.Value(),
		"list-about-to-expire": ListAboutToExpire.Value(),
	}

	BasicEnumImpl = enumimpl.New.BasicByte.DefaultWithAliasMapAllCases(
		Invalid,
		Ranges[:],
		aliasMap)
)
