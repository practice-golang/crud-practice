package main

import (
	"crud-practice/config"
	"crud-practice/db"
	"os"
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
			name: "SQLITE",
			db:   config.DatabaseInfoSQLite,
			args: []string{"program", "insert", "book_name", "author_name"},
		},
		{
			name: "SQLITE",
			db:   config.DatabaseInfoSQLite,
			args: []string{"program", "select"},
		},
		{
			name: "SQLITE",
			db:   config.DatabaseInfoSQLite,
			args: []string{"program", "update", "1", "book_name", "author_name"},
		},
		{
			name: "SQLITE",
			db:   config.DatabaseInfoSQLite,
			args: []string{"program", "delete", "1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args
			err := doJob()
			if err != nil {
				t.Error(err)
			}
		})
	}
}
