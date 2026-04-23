package sqliteconnpathtype

import (
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	sqliteConnectionFormats = map[Variant]string{
		AllSqlitePath:                     "{root}/{prefix}/{db-name}.db",
		AllWithTypeSqlitePath:             "{root}/{prefix}/{type}/{db-name}.db",
		AllWithTypeAndDynamicSqlitePath:   "{root}/{prefix}/{type}/{dynamic}-{db-name}.db",
		AllWithTypeAndSequenceSqlitePath:  "{root}/{prefix}/{type}/{sequence}-{db-name}.db",
		PrefixSqlitePath:                  "{prefix}/{db-name}.db",
		PrefixTypeSqlitePath:              "{prefix}/{type}/{db-name}.db",
		DynamicSpecificSqlitePath:         "{dynamic}-{db-name}.db",
		SequenceSpecificSqlitePath:        "{sequence}-{db-name}.db",
		DynamicSequenceSpecificSqlitePath: "{dynamic}/{sequence}-{db-name}.db",
		SpecificSqlitePath:                "{db-name}.db",
	}

	rangesMap = [...]string{}

	BasicEnumImpl = enumimpl.New.BasicString.CreateUsingStringersSpread(
		coredynamic.TypeName(Invalid),
		Invalid,
		AllSqlitePath,
		AllWithTypeSqlitePath,
		AllWithTypeAndDynamicSqlitePath,
		AllWithTypeAndSequenceSqlitePath,
		PrefixSqlitePath,
		PrefixTypeSqlitePath,
		SpecificSqlitePath,
		DynamicSpecificSqlitePath,
		SequenceSpecificSqlitePath,
		DynamicSequenceSpecificSqlitePath)
)
