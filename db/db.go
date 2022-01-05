package db

import (
	"database/sql"
	"errors"
)

const (
	_        = iota
	SQLITE   // SQLite
	MYSQL    // MySQL
	POSTGRES // PostgreSQL
	MOCKDB   // mockdb
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
		CreateDB() error
		CreateTable() error
		connect() (*sql.DB, error)
	}
)

var (
	Info DBInfo   // DB connection info
	Obj  DBObject // Duck interface
	Con  *sql.DB  // DB connection
)

func SetupDB() error {
	var err error

	switch Info.DatabaseType {

	case SQLITE:
		dsn := Info.FilePath
		Obj = &SQLite{dsn: dsn}

		Con, err = Obj.connect()
		if err != nil {
			return err
		}

		break

	case MYSQL:
		// dsn := Info.GrantID + ":" + Info.GrantPassword + "@" + Info.Protocol + "(" + Info.Addr + ":" + Info.Port + ")/" + Info.DatabaseName
		dsn := Info.GrantID + ":" + Info.GrantPassword + "@" + Info.Protocol + "(" + Info.Addr + ":" + Info.Port + ")/"
		Obj = &Mysql{dsn: dsn}

		Con, err = Obj.connect()
		if err != nil {
			return err
		}

		break

	case POSTGRES:
		dsn := `host=` + Info.Addr + ` port=` + Info.Port + ` user=` + Info.GrantID + ` password=` + Info.GrantPassword + ` dbname=` + Info.DatabaseName + ` sslmode=disable`
		Obj = &Postgres{dsn: dsn}

		Con, err = Obj.connect()
		if err != nil {
			return err
		}

		break

	default:
		return errors.New("nothing to support DB")
	}

	return nil
}
