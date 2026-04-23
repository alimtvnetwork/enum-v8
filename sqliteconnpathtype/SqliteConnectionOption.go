package sqliteconnpathtype

import (
	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coreutils/stringutil"
)

type SqliteConnectionOption struct {
	Root, Prefix, DbName string
	TypeName, Dynamic    string
	Sequence             string
}

func (it SqliteConnectionOption) CreateMap() map[string]string {
	return map[string]string{
		"{root}":     it.Root,
		"{prefix}":   it.Prefix,
		"{db-name}":  it.DbName,
		"{type}":     it.TypeName,
		"{dynamic}":  it.Dynamic,
		"{sequence}": it.Sequence,
	}
}

func (it SqliteConnectionOption) Compile(pathFormat string) string {
	return stringutil.
		ReplaceTemplate.
		DirectKeyUsingMap(
			pathFormat,
			it.CreateMap())
}

func (it SqliteConnectionOption) String() string {
	createdMap := it.CreateMap()

	return corejson.
		Serialize.
		ToString(createdMap)
}
