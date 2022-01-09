package crud

import (
	"crud-practice/db"
	"crud-practice/model"
	"crud-practice/np"
	"errors"

	"github.com/blockloop/scan"
)

func InsertData(book model.Book) (int64, int64, error) {
	var (
		err   error
		count int64 = 0
		idx   int64 = 0
	)

	dbtype := db.GetDatabaseTypeString()
	tablename := db.GetTableName()

	columns := np.CreateString(book, dbtype, "insert").Names

	holders, err := np.CreateHolders(dbtype, columns)
	if err != nil {
		return count, idx, err
	}

	sql := ` INSERT INTO ` + tablename + ` (` + columns + `) VALUES (` + holders + `)`
	colSlice := np.CreateMapSlice(book, "insert")

	count, idx, err = db.Obj.Exec(sql, colSlice["values"], "IDX")
	if err != nil {
		return count, idx, err
	}

	return count, idx, nil
}

func SelectData(id int) ([]model.Book, error) {
	result := []model.Book{}

	dbtype := db.GetDatabaseTypeString()
	tablename := db.GetTableName()

	book := model.Book{}

	columns := np.CreateString(book, dbtype, "").Names

	sql := `SELECT ` + columns + ` FROM ` + tablename

	where := []interface{}{}
	if id > 0 {
		substitute, _, _ := np.CreateAssignHolders(dbtype, db.QuotesName("IDX"), 0)
		sql += ` WHERE ` + substitute
		where = append(where, id)
	}

	r, err := db.Con.Query(sql, where...)
	if err != nil {
		return nil, err
	}

	err = scan.Rows(&result, r)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateData(book model.Book) (int64, error) {
	dbtype := db.GetDatabaseTypeString()
	tablename := db.GetTableName()

	columns := np.CreateString(book, dbtype, "").Names
	directive, offset, _ := np.CreateAssignHolders(dbtype, columns, 0)
	where, _, _ := np.CreateAssignHolders(dbtype, db.QuotesName("IDX"), offset)

	sql := `UPDATE ` + tablename + ` SET ` + directive + ` WHERE ` + where

	updateValues := []interface{}{book.Title, book.Author}
	whereValues := []interface{}{book.Idx}
	values := append(updateValues, whereValues...)

	r, err := db.Con.Exec(sql, values...)
	if err != nil {
		return 0, err
	}

	count, _ := r.RowsAffected()

	return count, nil
}

func DeleteData(id int) (int64, error) {
	var count int64 = 0

	dbtype := db.GetDatabaseTypeString()
	tablename := db.GetTableName()

	if id > 0 {
		where, _, _ := np.CreateAssignHolders(dbtype, db.QuotesName("IDX"), 0)
		sql := `DELETE FROM ` + tablename + ` WHERE ` + where
		whereValues := []interface{}{id}

		r, err := db.Con.Exec(sql, whereValues...)
		if err != nil {
			return count, err
		}

		count, _ = r.RowsAffected()
	} else {
		return count, errors.New("idx value have to exist and to be larger than 0")
	}

	return count, nil
}
