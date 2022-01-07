package function

import (
	"crud-practice/db"
	"strings"
)

func getTableName() string {
	tablename := ""
	switch db.Info.DatabaseType {
	case db.SQLITE:
		tablename = `"` + db.Info.TableName + `"`
	case db.MYSQL:
		tablename = db.Info.DatabaseName + `.` + db.Info.TableName
	case db.POSTGRES:
		tablename = `"` + db.Info.SchemaName + `"."` + db.Info.TableName + `"`
	}

	return tablename
}

func getDatabaseTypeString() string {
	dbtype := ""
	switch db.Info.DatabaseType {
	case db.SQLITE:
		dbtype = "sqlite"
	case db.MYSQL:
		dbtype = "mysql"
	case db.POSTGRES:
		dbtype = "postgres"
	}

	return dbtype
}

func quotesName(data string) string {
	result := ""

	switch db.Info.DatabaseType {
	case db.SQLITE:
		data = strings.ReplaceAll(data, `"`, `""`)
		result = `"` + data + `"`
	case db.MYSQL:
		data = strings.ReplaceAll(data, "`", "``")
		result = "'" + data + "'"
	case db.POSTGRES:
		data = strings.ReplaceAll(data, `"`, `""`)
		result = `"` + data + `"`
	}

	return result
}

// func quotesValue(data string) string {
// 	result := ""

// 	switch db.Info.DatabaseType {
// 	case db.SQLITE:
// 		data = strings.ReplaceAll(data, "'", "''")
// 		result = "'" + data + "'"
// 	case db.MYSQL:
// 		data = strings.ReplaceAll(data, "'", "\\'")
// 		result = "'" + data + "'"
// 	case db.POSTGRES:
// 		data = strings.ReplaceAll(data, "'", "''")
// 		result = "'" + data + "'"
// 	}

// 	return result
// }
