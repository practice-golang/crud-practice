package function

import (
	"crud-practice/db"
	"crud-practice/model"
	"errors"
	"log"

	"github.com/blockloop/scan"
	"github.com/practice-golang/np"
)

func InsertData(book model.Book) (int64, int64, error) {
	var err error
	var succeedCount int64 = 0
	var succeedIDX int64 = 0

	colNameTest1 := np.MakeSlice(book)
	log.Println(colNameTest1)
	colNameTest2 := np.MakeMap(book)
	log.Println(colNameTest2)

	colNameSTR := "`TITLE`, `AUTHOR`"
	colValueCavitySTR := `?, ?`
	switch {
	case db.Info.DatabaseType == db.SQLITE:
		colNameSTR = `"TITLE", "AUTHOR"`
	case db.Info.DatabaseType == db.POSTGRES:
		colNameSTR = `"TITLE", "AUTHOR"`
		colValueCavitySTR = `$1, $2`
	}

	colValues := []interface{}{}
	colValues = append(colValues, book.Title)
	colValues = append(colValues, book.Author)

	tablename := getTableName()

	sql := `
		INSERT INTO ` + tablename + `
			(` + colNameSTR + `)
		VALUES
			(` + colValueCavitySTR + `)
		`

	succeedCount, succeedIDX, err = db.Obj.Exec(sql, colValues, "IDX")
	if err != nil {
		return succeedCount, succeedIDX, err
	}

	return succeedCount, succeedIDX, nil
}

func SelectData(id int) ([]model.Book, error) {
	result := []model.Book{}

	colNameSTR := "`TITLE`, `AUTHOR`"
	switch {
	case db.Info.DatabaseType == db.SQLITE:
		colNameSTR = `"TITLE", "AUTHOR"`
	case db.Info.DatabaseType == db.POSTGRES:
		colNameSTR = `"TITLE", "AUTHOR"`
	}

	whereSTR := []interface{}{}

	tablename := getTableName()

	sql := `SELECT ` + colNameSTR + ` FROM ` + tablename
	if id > 0 {
		sql += ` WHERE IDX=?`
		whereSTR = append(whereSTR, id)
	}

	r, err := db.Con.Query(sql, whereSTR...)
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
	tablename := getTableName()

	sql := `
		UPDATE ` + tablename + ` SET
			TITLE=?, AUTHOR=?
		WHERE IDX=?
		`
	if db.Info.DatabaseType == db.POSTGRES {
		sql = `
			UPDATE ` + tablename + ` SET
				"TITLE"=$1, "AUTHOR"=$2
			WHERE "IDX"=$3
		`
	}
	changeValues := []interface{}{book.Title, book.Author}
	whereValues := []interface{}{book.Idx}
	colValues := append(changeValues, whereValues...)

	r, err := db.Con.Exec(sql, colValues...)
	if err != nil {
		return 0, err
	}

	succeedCount, _ := r.RowsAffected()

	return succeedCount, nil
}

func DeleteData(id int) (int64, error) {
	var succeedCount int64 = 0

	tablename := getTableName()

	if id > 0 {
		sql := `
			DELETE FROM ` + tablename + `
			WHERE IDX=?
			`
		whereValues := []interface{}{id}

		r, err := db.Con.Exec(sql, whereValues...)
		if err != nil {
			return succeedCount, err
		}

		succeedCount, _ = r.RowsAffected()
	} else {
		return succeedCount, errors.New("id value have to be larger than 0")
	}

	return succeedCount, nil
}
