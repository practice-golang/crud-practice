package db

import (
	"database/sql"
)

type SQLite struct{ dsn string }

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

func (d *SQLite) connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite", d.dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
