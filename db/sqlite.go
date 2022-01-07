package db

import (
	"database/sql"
)

type SQLite struct{ dsn string }

func (d *SQLite) connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite", d.dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// CreateDB - enough to use connect so, not use
func (d *SQLite) CreateDB() error { return nil }

func (d *SQLite) CreateTable() error {
	sql := `
	CREATE TABLE IF NOT EXISTS "` + Info.TableName + `" (
		"IDX"			INTEGER,
		"TITLE"			TEXT,
		"AUTHOR"		TEXT,
		PRIMARY KEY("IDX" AUTOINCREMENT)
	);`

	_, err := Con.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func (d *SQLite) Exec(sql string, colValues []interface{}, options string) (int64, int64, error) {
	var err error
	var count int64 = 0
	var idx int64 = 0

	result, err := Con.Exec(sql, colValues...)
	if err != nil {
		return count, idx, err
	}

	count, _ = result.RowsAffected()
	idx, _ = result.LastInsertId()

	return count, idx, nil
}
