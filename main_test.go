package main

import (
	"crud-practice/config"
	"crud-practice/db"
	"os"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "modernc.org/sqlite"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
		db   db.DBInfo
		args []string
	}{
		{
			name: "EMPTY ARGS",
			db:   config.DatabaseInfoSQLite,
			args: []string{"program"},
		},
		{
			name: "WRONG ARGS",
			db:   config.DatabaseInfoSQLite,
			args: []string{"program", "wrong", "params"},
		},
		{
			name: "INSERT",
			db:   config.DatabaseInfoSQLite,
			args: []string{"program", "insert", "book_name", "author_name"},
		},
		{
			name: "SELECT ALL",
			db:   config.DatabaseInfoSQLite,
			args: []string{"program", "select"},
		},
		{
			name: "SELECT 1",
			db:   config.DatabaseInfoSQLite,
			args: []string{"program", "select", "1"},
		},
		{
			name: "SELECT 0(=ALL)",
			db:   config.DatabaseInfoSQLite,
			args: []string{"program", "select", "wtf", "wtf"},
		},
		{
			name: "UPDATE",
			db:   config.DatabaseInfoSQLite,
			args: []string{"program", "update", "1", "book_name", "author_name"},
		},
		{
			name: "UPDATE NO ARGS",
			db:   config.DatabaseInfoSQLite,
			args: []string{"program", "update"},
		},
		{
			name: "DELETE",
			db:   config.DatabaseInfoSQLite,
			args: []string{"program", "delete", "1"},
		},
		{
			name: "DELETE 0(=NONE)",
			db:   config.DatabaseInfoSQLite,
			args: []string{"program", "delete", "wtf", "wtf"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args

			err := beginJob()

			if tt.name == "EMPTY ARGS" {
				if strings.Contains(err.Error(), "no args") {
					return
				}
			}
			if tt.name == "WRONG ARGS" {
				if strings.Contains(err.Error(), "check inputed parameters") {
					return
				}
			}

			defer os.Remove(db.Info.FilePath)
			defer db.Con.Close()

			if err != nil {
				if tt.name == "DELETE 0(=NONE)" {
					if strings.Contains(err.Error(), "idx value have to exist and to be larger than 0") {
						return
					}
				}
				if tt.name == "UPDATE NO ARGS" {
					if strings.Contains(err.Error(), "need more params") {
						return
					}
				}
				t.Error(err)
			}
		})
	}
}
