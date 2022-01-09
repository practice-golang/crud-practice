package db

import "strings"

func GetTableName() string {
	tablename := ""
	switch Info.DatabaseType {
	case SQLITE:
		tablename = `"` + Info.TableName + `"`
	case MYSQL:
		tablename = Info.DatabaseName + `.` + Info.TableName
	case POSTGRES:
		tablename = `"` + Info.SchemaName + `"."` + Info.TableName + `"`
	case SQLSERVER:
		tablename = `"` + Info.DatabaseName + `"."` + Info.SchemaName + `"."` + Info.TableName + `"`
	}

	return tablename
}

func GetDatabaseTypeString() string {
	dbtype := ""

	switch Info.DatabaseType {
	case SQLITE:
		dbtype = "sqlite"
	case MYSQL:
		dbtype = "mysql"
	case POSTGRES:
		dbtype = "postgres"
	case SQLSERVER:
		dbtype = "sqlserver"
	}

	return dbtype
}

func QuotesName(data string) string {
	result := ""

	switch Info.DatabaseType {
	case SQLITE:
		data = strings.ReplaceAll(data, `"`, `""`)
		result = `"` + data + `"`
	case MYSQL:
		data = strings.ReplaceAll(data, "`", "``")
		result = "'" + data + "'"
	case POSTGRES:
		data = strings.ReplaceAll(data, `"`, `""`)
		result = `"` + data + `"`
	case SQLSERVER:
		data = strings.ReplaceAll(data, `"`, `""`)
		result = `"` + data + `"`
	}

	return result
}

func QuotesValue(data string) string {
	result := ""

	switch Info.DatabaseType {
	case SQLITE:
		data = strings.ReplaceAll(data, "'", "''")
		result = "'" + data + "'"
	case MYSQL:
		data = strings.ReplaceAll(data, "'", "\\'")
		result = "'" + data + "'"
	case POSTGRES:
		data = strings.ReplaceAll(data, "'", "''")
		result = "'" + data + "'"
	case SQLSERVER:
		data = strings.ReplaceAll(data, "'", "''")
		result = "'" + data + "'"
	}

	return result
}
