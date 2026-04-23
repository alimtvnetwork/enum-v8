package dbdrivertype

import (
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:             "Invalid",
		Sqlite:              "Sqlite",
		Redis:               "Redis",
		MySql:               "MySql",
		MariaDb:             "MariaDb",
		PostgreSql:          "PostgreSql",
		MicrosoftSqlExpress: "MicrosoftSqlExpress",
		MicrosoftSqlServer:  "MicrosoftSqlServer",
		MicrosoftSqlCompact: "MicrosoftSqlCompact",
		MicrosoftAccess:     "MicrosoftAccess",
		Oracle:              "Oracle",
		Firebird:            "Firebird",
		MongoDb:             "MongoDb",
		CouchDb:             "CouchDb",
		AmazonDynamoDb:      "AmazonDynamoDb",
		HSqlDb:              "HSqlDb",
		Text:                "Text",
		Json:                "Json",
		Yaml:                "Yaml",
		Protobuf:            "Protobuf",
	}

	sqlDbs = map[Variant]bool{
		Sqlite:              true,
		MySql:               true,
		MariaDb:             true,
		PostgreSql:          true,
		MicrosoftSqlExpress: true,
		MicrosoftSqlServer:  true,
		MicrosoftSqlCompact: true,
		MicrosoftAccess:     true,
		Oracle:              true,
		Firebird:            true,
		HSqlDb:              true,
	}

	noSqlDbs = map[Variant]bool{
		Redis:          true,
		MongoDb:        true,
		CouchDb:        true,
		AmazonDynamoDb: true,
		Text:           true,
		Json:           true,
		Yaml:           true,
		Protobuf:       true,
	}

	connectionStringFormatMap = map[Variant]string{
		Sqlite:             "{db}",
		Redis:              "redis://{ip}:{port}{?options}",
		MySql:              "{user}:{password}@tcp({ip}:{port})/{db}{?options}",
		PostgreSql:         "host={ip} port={port} user={user} password={password} dbname={db}{?options}",
		MicrosoftSqlServer: "sqlserver://{user}:{password}@{ip}:{port}?database={db}{?options}",
		MongoDb:            "mongodb://[{user}:{password}]{ip}:{port}/{db}{?options}", // https://t.ly/yavi
	}

	connectionStringAllDbFormatMap = map[Variant]string{
		Sqlite:             "{db}",
		Redis:              "redis://{ip}:{port}{?options}",
		MySql:              "{user}:{password}@tcp({ip}:{port}){?options}",
		PostgreSql:         "host={ip} port={port} user={user} password={password}{?options}",
		MicrosoftSqlServer: "sqlserver://{user}:{password}@{ip}:{port}{?options}",
		MongoDb:            "mongodb://[{user}:{password}]{ip}:{port}{?options}", // https://t.ly/yavi
	}

	defaultDbPortsMap = map[Variant]uint16{
		MySql:              3306,
		MariaDb:            3306,
		PostgreSql:         5432,
		MicrosoftSqlServer: 1433,
		Oracle:             1521,
		Redis:              6379,
		MongoDb:            27017,
		Firebird:           3050,
		CouchDb:            5984,
		HSqlDb:             9001,
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Invalid),
		Ranges[:])
)
