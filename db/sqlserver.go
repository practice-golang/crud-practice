package db

import (
	"database/sql"
)

type SqlServer struct{ dsn string }

func (d *SqlServer) CreateDB() error {
	sql := `
	USE master
	IF NOT EXISTS(
		SELECT name
		FROM sys.databases
		WHERE name=N'` + Info.DatabaseName + `'
	) CREATE DATABASE [` + Info.DatabaseName + `]
	`
	_, err := Con.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func (d *SqlServer) CreateTable() error {
	sql := `
	USE ` + Info.DatabaseName + `
	IF OBJECT_ID(N'` + Info.TableName + `', N'U') IS NULL
	CREATE TABLE ` + Info.TableName + ` (
		IDX INT NOT NULL IDENTITY PRIMARY KEY,
		TITLE VARCHAR(256) NULL,
		AUTHOR VARCHAR(256) NULL,
	)
	-- GO`

	_, err := Con.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func (d *SqlServer) connect() (*sql.DB, error) {
	db, err := sql.Open("sqlserver", d.dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *SqlServer) Exec(sql string, colValues []interface{}, options string) (int64, int64, error) {
	var err error
	var count int64 = 0
	var idx int64 = 0

	sql += ` SELECT ID = CONVERT(bigint, ISNULL(SCOPE_IDENTITY(), -1)) `

	err = Con.QueryRow(sql, colValues...).Scan(&idx)
	if err != nil {
		return count, idx, err
	}

	if idx > 0 {
		count = 1
	}

	return count, idx, nil
}
