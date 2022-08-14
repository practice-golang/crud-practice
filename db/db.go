package db

import (
	"crud-practice/model"
	"database/sql"
	"errors"
	"net/url"
	"strconv"

	go_ora "github.com/sijms/go-ora/v2"
)

type (
	DBInfo struct {
		DatabaseType  int
		Protocol      string
		Addr          string
		Port          string
		DatabaseName  string
		SchemaName    string
		TableName     string
		GrantID       string
		GrantPassword string
		FilePath      string
	}

	DBObject interface {
		connect() (*sql.DB, error)
		CreateDB() error
		CreateTable() error
		// Exec - Almost Same as sql.Exec()
		// Because of PostgreSQL and MS SQL Server, INSERT query and RETURN id way is not enough to use sql.Exec()
		// Return affected rows, last insert id, error
		// Not return sql.Result
		Exec(string, []interface{}, string) (int64, int64, error)
	}
)

var (
	Info            DBInfo   // DB connection info
	InfoOracleAdmin DBInfo   // Oracle Admin connection info
	Obj             DBObject // Duck interface
	Con             *sql.DB  // DB connection
)

func SetupDB() error {
	var err error

	switch Info.DatabaseType {

	case model.SQLITE:
		dsn := Info.FilePath
		Obj = &SQLite{dsn: dsn}

		Con, err = Obj.connect()
		if err != nil {
			return err
		}

	case model.MYSQL:
		// dsn := Info.GrantID + ":" + Info.GrantPassword + "@" + Info.Protocol + "(" + Info.Addr + ":" + Info.Port + ")/" + Info.DatabaseName
		dsn := Info.GrantID + ":" + Info.GrantPassword + "@" + Info.Protocol + "(" + Info.Addr + ":" + Info.Port + ")/"
		Obj = &Mysql{dsn: dsn}

		Con, err = Obj.connect()
		if err != nil {
			return err
		}

	case model.POSTGRES:
		dsn := `host=` + Info.Addr + ` port=` + Info.Port + ` user=` + Info.GrantID + ` password=` + Info.GrantPassword + ` dbname=` + Info.DatabaseName + ` sslmode=disable`
		Obj = &Postgres{dsn: dsn}

		Con, err = Obj.connect()
		if err != nil {
			return err
		}

	case model.SQLSERVER:
		dsn := "sqlserver://" + Info.GrantID + ":" + Info.GrantPassword + "@" + Info.Addr + ":" + Info.Port + "?" + Info.DatabaseName + "&connction+timeout=30"
		Obj = &SqlServer{dsn: dsn}

		Con, err = Obj.connect()
		if err != nil {
			return err
		}

	case model.ORACLE:
		port, _ := strconv.Atoi(Info.Port)
		dsn := go_ora.BuildUrl(Info.Addr, port, Info.DatabaseName, Info.GrantID, Info.GrantPassword, nil)
		if Info.FilePath != "" {
			dsn += "?SSL=enable&SSL Verify=false&WALLET=" + url.QueryEscape(Info.FilePath)
		}
		Obj = &Oracle{dsn: dsn}

		Con, err = Obj.connect()
		if err != nil {
			return err
		}

	default:
		return errors.New("database type not supported")
	}

	return nil
}
