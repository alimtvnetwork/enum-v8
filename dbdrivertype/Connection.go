package dbdrivertype

import (
	"github.com/alimtvnetwork/core-v8/constants"
	"github.com/alimtvnetwork/core-v8/coreutils/stringutil"
	"github.com/alimtvnetwork/core-v8/errcore"
)

type Connection struct {
	DbType Variant
	ConnectionOptions
}

func (it Connection) HasConnectionString() bool {
	_, has := connectionStringFormatMap[it.DbType]

	return has
}

func (it Connection) IsInvalidConnectionString() bool {
	_, has := connectionStringFormatMap[it.DbType]

	return !has
}

func (it Connection) IsValid() bool {
	_, has := connectionStringFormatMap[it.DbType]

	return has
}

func (it Connection) ConnectionStringFormat() string {
	return connectionStringFormatMap[it.DbType]
}

func (it Connection) ConnectionStringAllDbFormat() string {
	return connectionStringAllDbFormatMap[it.DbType]
}

func (it Connection) CreateMap() map[string]string {
	return map[string]string{
		"{db}":       it.DbName,
		"{ip}":       it.Host,
		"{port}":     it.Port,
		"{user}":     it.User,
		"{password}": it.User,
		"{?options}": it.Options,
	}
}

func (it Connection) CreateMapUsingParams(
	host, port, dbName string,
	user, password, options string,
) map[string]string {
	return map[string]string{
		"{db}":       dbName,
		"{ip}":       host,
		"{port}":     port,
		"{user}":     user,
		"{password}": password,
		"{?options}": options,
	}
}

func (it Connection) CreateMapUsingParamsNoOptions(
	host, port, dbName string,
	user, password string,
) map[string]string {
	return map[string]string{
		"{db}":       dbName,
		"{ip}":       host,
		"{port}":     port,
		"{user}":     user,
		"{password}": password,
		"{?options}": constants.EmptyString,
	}
}

func (it Connection) CompileUsingConnectionFormat(
	format string,
) string {
	createdMap := it.CreateMap()

	return stringutil.ReplaceTemplate.DirectKeyUsingMap(
		format,
		createdMap)
}

func (it Connection) Compile() (string, error) {
	format, has := connectionStringFormatMap[it.DbType]

	if !has {
		return format, it.invalidConnectionStringErr()
	}

	createdMap := it.CreateMap()

	return stringutil.ReplaceTemplate.DirectKeyUsingMap(
		format,
		createdMap), nil
}

func (it Connection) CompileUsingParamsNoOptions(
	host, port, dbName string,
	user, password string,
) (string, error) {
	return it.CompileUsingParams(
		host,
		port,
		dbName,
		user,
		password,
		constants.EmptyString, // no options
	)
}

func (it Connection) CompileUsingParams(
	host, port, dbName string,
	user, password, options string,
) (string, error) {
	format, has := connectionStringFormatMap[it.DbType]

	if !has {
		return format, it.invalidConnectionStringErr()
	}

	createdMap := it.CreateMapUsingParams(
		host,
		port,
		dbName,
		user,
		password,
		options)

	return stringutil.ReplaceTemplate.DirectKeyUsingMap(
		format,
		createdMap), nil
}

func (it Connection) invalidConnectionStringErr() error {
	return errcore.
		InvalidStringType.
		Error("connection string not available for "+it.DbType.Name(), it.DbType)
}
