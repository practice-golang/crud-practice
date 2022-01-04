package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

func Test_dbConn(t *testing.T) {
	tests := []struct {
		name   string
		wantDb *sql.DB
	}{
		{
			name:   "dbConn",
			wantDb: &sql.DB{},
		},
	}

	for _, tt := range tests {
		tt.wantDb = dbConn()

		t.Run(tt.name, func(t *testing.T) {
			db := dbConn()

			require.Equal(t, reflect.TypeOf(tt.wantDb), reflect.TypeOf(db))
		})
	}
}

func TestInsertData(t *testing.T) {
	type args struct {
		book  Book
		table string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "InsertData",
			args: args{
				book: Book{
					Title:  "test_title_" + fmt.Sprint(rand.Intn(99)),
					Author: "test_author_" + fmt.Sprint(rand.Intn(99)),
				},
				table: "books",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := InsertData(tt.args.book, tt.args.table)
			if err != nil {
				t.Error("InsertData", err)
			}

			require.NotNil(t, r)

			affectedROW, err := r.RowsAffected()
			if err != nil {
				t.Error("InsertData", err)
			}
			require.Equal(t, int64(1), affectedROW, "not equal")
		})
	}
}
