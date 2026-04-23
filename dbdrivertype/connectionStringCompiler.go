package dbdrivertype

import (
	"github.com/alimtvnetwork/core-v8/constants"
	"github.com/alimtvnetwork/core-v8/coreutils/stringutil"
	"github.com/alimtvnetwork/core-v8/errcore"
)

type connectionStringCompiler struct {
	dbType Variant
}

func (it connectionStringCompiler) HasConnectionString() bool {
	_, has := connectionStringFormatMap[it.dbType]

	return has
}

func (it connectionStringCompiler) IsInvalidConnectionString() bool {
	_, has := connectionStringFormatMap[it.dbType]

	return !has
}

func (it connectionStringCompiler) IsValidConnectionString() bool {
	_, has := connectionStringFormatMap[it.dbType]

	return has
}

func (it connectionStringCompiler) Format() string {
	return connectionStringFormatMap[it.dbType]
}

func (it connectionStringCompiler) AllDbFormat() string {
	return connectionStringAllDbFormatMap[it.dbType]
}

func (it connectionStringCompiler) CompileUsingConnection(
	conn ConnectionOptions,
) (string, error) {
	return Connection{
		DbType:            it.dbType,
		ConnectionOptions: conn,
	}.Compile()
}

func (it connectionStringCompiler) CompileUsingParams(
	host, port, dbName string,
	user, password string,
) (string, error) {
	connection := Connection{
		DbType: it.dbType,
		ConnectionOptions: ConnectionOptions{
			Host:               host,
			Port:               port,
			User:               user,
			Password:           password,
			Options:            constants.EmptyString,
			DbName:             dbName,
			IsSpecificDatabase: false,
		},
	}

	return connection.Compile()
}

func (it connectionStringCompiler) CompileUsingAllDbConnectionFormat(
	conn ConnectionOptions,
) (string, error) {
	format, has := connectionStringAllDbFormatMap[it.dbType]

	if !has {
		return format, it.invalidConnectionStringErr()
	}

	return conn.CompileUsingConnectionFormat(format), nil
}

func (it connectionStringCompiler) CompileUsingConnectionFormat(
	format string,
	conn ConnectionOptions,
) string {
	return conn.CompileUsingConnectionFormat(format)
}

func (it connectionStringCompiler) CompileUsingParamsOptions(
	host, port, dbName string,
	user, password, options string,
) (string, error) {
	connection := Connection{
		DbType: it.dbType,
		ConnectionOptions: ConnectionOptions{
			Host:               host,
			Port:               port,
			User:               user,
			Password:           password,
			Options:            options,
			DbName:             dbName,
			IsSpecificDatabase: false,
		},
	}

	return connection.Compile()
}

func (it connectionStringCompiler) CompileUsingMap(
	isCurlyReplace bool,
	replacerMap map[string]string,
) (string, error) {
	format, has := connectionStringFormatMap[it.dbType]

	if !has {
		return format, it.invalidConnectionStringErr()
	}

	return it.FormatCompileUsingMap(
		format,
		isCurlyReplace,
		replacerMap,
	)
}

func (it connectionStringCompiler) FormatCompileUsingMap(
	format string,
	isCurlyReplace bool,
	replacerMap map[string]string,
) (
	string, error,
) {
	return stringutil.ReplaceTemplate.UsingMapOptions(
		isCurlyReplace,
		format,
		replacerMap), nil
}

func (it connectionStringCompiler) CompileUsingMapMust(
	isCurlyReplace bool,
	replacerMap map[string]string,
) string {
	connectionStringCompiled, err := it.CompileUsingMap(
		isCurlyReplace,
		replacerMap)
	errcore.MustBeEmpty(err)

	return connectionStringCompiled
}

func (it connectionStringCompiler) invalidConnectionStringErr() error {
	return errcore.
		InvalidStringType.
		Error("connection string not available for "+it.dbType.Name(), it.dbType)
}
