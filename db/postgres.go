package db

import (
	"database/sql"
	"errors"
	"strings"
)

type Postgres struct{ dsn string }

// CreateDB - Create DB and Schema
func (d *Postgres) CreateDB() error {
	var err error

	sql := ""

	// Check if DB exists - not use
	// sql = `
	// SELECT datname
	// FROM pg_database
	// WHERE datistemplate=FALSE
	// 	AND datname='` + Info.DatabaseName + `';
	// `

	// Database creation - not use
	// sql = `CREATE DATABASE ` + Info.DatabaseName + `;`
	// _, err = Con.Exec(sql)
	// if err != nil {
	// 	return err
	// }

	sql = `CREATE SCHEMA IF NOT EXISTS ` + Info.SchemaName + `;`
	_, err = Con.Exec(sql)
	if err != nil {
		if strings.Contains(err.Error(), "does not exist") {
			msg := "Database does not exist\n" +
				"With Postgres, create database yourself like below SQL query:" +
				"\nCREATE DATABASE " + Info.DatabaseName + ";"
			return errors.New(msg)
		}
		return err
	}

	return nil
}

func (d *Postgres) CreateTable() error {
	sql := `
	CREATE TABLE IF NOT EXISTS ` + Info.SchemaName + `.` + Info.TableName + ` (
		"IDX" SERIAL PRIMARY KEY,
		"TITLE" VARCHAR(255) NULL DEFAULT NULL,
		"AUTHOR" VARCHAR(255) NULL DEFAULT NULL
	)`
	// "TITLE" VARCHAR(255) UNIQUE NULL DEFAULT NULL

	_, err := Con.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func (d *Postgres) connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", d.dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}