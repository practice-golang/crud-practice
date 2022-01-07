package function

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

	dbtype := getDatabaseTypeString()
	tablename := getTableName()

	colNames := np.CreateString(book, dbtype, "insert").Names

	colValueBinds, err := np.CreateHolders(dbtype, colNames)
	if err != nil {
		return count, idx, err
	}

	sql := ` INSERT INTO ` + tablename + ` (` + colNames + `) VALUES (` + colValueBinds + `)`
	colSlice := np.CreateMapSlice(book, "insert")

	count, idx, err = db.Obj.Exec(sql, colSlice["values"], "IDX")
	if err != nil {
		return count, idx, err
	}

	return count, idx, nil
}

func SelectData(id int) ([]model.Book, error) {
	result := []model.Book{}

	dbtype := getDatabaseTypeString()
	tablename := getTableName()

	book := model.Book{}

	colNames := np.CreateString(book, dbtype, "").Names

	sql := `SELECT ` + colNames + ` FROM ` + tablename

	where := []interface{}{}
	if id > 0 {
		sql += ` WHERE IDX=?`
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
	dbtype := getDatabaseTypeString()
	tablename := getTableName()

	columns := np.CreateString(book, dbtype, "").Names
	directive, offset, _ := np.CreateUpdateHolders(dbtype, columns, 0)
	where, _, _ := np.CreateUpdateHolders(dbtype, quotesName("IDX"), offset)

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

	tablename := getTableName()

	if id > 0 {
		sql := `DELETE FROM ` + tablename + ` WHERE IDX=?`
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
