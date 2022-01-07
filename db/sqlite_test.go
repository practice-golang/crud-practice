package db

import (
	"os"
	"testing"

	_ "modernc.org/sqlite"
)

func TestSQLite_connect(t *testing.T) {
	tests := []struct {
		name string
		dsn  string
		d    *SQLite
	}{
		{
			name: "SQLITE",
			dsn:  "../test.db",
			d:    &SQLite{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.dsn = tt.dsn
			Con, err := tt.d.connect()
			if err != nil {
				os.Remove(tt.dsn)
				t.Error(err)
				return
			}
			defer os.Remove(tt.dsn)
			defer Con.Close()
		})
	}
}

func TestSQLite_CreateTable(t *testing.T) {
	tests := []struct {
		name string
		dsn  string
		d    *SQLite
	}{
		{
			name: "SQLITE",
			dsn:  "../test.db",
			d:    &SQLite{},
		},
	}
	var err error
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.dsn = tt.dsn
			Con, err = tt.d.connect()
			if err != nil {
				os.Remove(tt.dsn)
				t.Error(err)
				return
			}
			defer os.Remove(tt.dsn)
			defer Con.Close()

			err = tt.d.CreateTable()
			if err != nil {
				os.Remove(tt.dsn)
				t.Error(err)
				return
			}
		})
	}
}

// func TestSQLite_Exec(t *testing.T) {
// 	type args struct {
// 		sql       string
// 		colValues []interface{}
// 		options   string
// 	}
// 	tests := []struct {
// 		name    string
// 		dsn     string
// 		d       *SQLite
// 		args    args
// 		want    int64
// 		want1   int64
// 		wantErr bool
// 	}{
// 		{
// 			name: "SQLITE",
// 			dsn:  "../test.db",
// 			d:    &SQLite{},
// 		},
// 	}
// 	var err error
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.d.dsn = tt.dsn
// 			Con, err = tt.d.connect()
// 			if err != nil {
// 				t.Error(err)
// 				return
// 			}

// 			err = tt.d.CreateTable()
// 			if err != nil {
// 				t.Error(err)
// 				return
// 			}

// 			got, got1, err := tt.d.Exec(tt.args.sql, tt.args.colValues, tt.args.options)
// 			if err != nil {
// 				t.Errorf("SQLite.Exec() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 		})
// 	}
// }
