package function

import "crud-practice/db"

func getTableName() string {
	tablename := ""
	switch db.Info.DatabaseType {
	case db.SQLITE:
		tablename = `"` + db.Info.TableName + `"`
	case db.MYSQL:
		tablename = db.Info.DatabaseName + `.` + db.Info.TableName
	case db.POSTGRES:
		tablename = db.Info.SchemaName + `.` + db.Info.TableName
	}

	return tablename
}
