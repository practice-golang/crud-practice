package main // import "crud-practice"

import (
	"crud-practice/config"
	"crud-practice/crud"
	"crud-practice/db"
	"crud-practice/model"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/sijms/go-ora/v2"
	"gopkg.in/guregu/null.v4"
	_ "modernc.org/sqlite"
)

var appname = "crud-practice"

func beginJob() error {
	var err error

	// 실행파일(#1) 명령(#2) 대상(#3...)
	argLen := len(os.Args) - 1
	if argLen < 1 {
		fmt.Println("Usage:")
		fmt.Println(appname + " insert book_name author_name")
		fmt.Println(appname + " select [id_number]")
		fmt.Println(appname + " update id_number book_name author_name")
		fmt.Println(appname + " delete id_number")
		return errors.New("no args")
	}

	command := os.Args[1:2]

	err = db.SetupDB()
	if err != nil {
		log.Fatal("SetupDB:", err)
	}

	err = db.Obj.CreateDB()
	if err != nil {
		log.Fatal("CreateDB:", err)
	}

	err = db.Obj.CreateTable()
	if err != nil {
		log.Fatal("CreateTable:", err)
	}
	// defer db.Con.Close()

	switch command[0] {
	case "insert":
		if argLen < 3 {
			return errors.New(command[0] + ": need more params")
		}

		book := model.Book{
			Title:  null.StringFrom(os.Args[2:3][0]),
			Author: null.StringFrom(os.Args[3:4][0]),
		}

		idx, count, err := crud.InsertData(book)
		if err != nil {
			return err
		}

		log.Println("Insert idx, count:", idx, count)

	case "select":
		var id int = 0

		if argLen == 2 {
			idStr := os.Args[2:3]
			id, _ = strconv.Atoi(idStr[0])
		}

		books, err := crud.SelectData(id)
		if err != nil {
			return errors.New("SelectData " + err.Error())
		}

		for _, book := range books {
			log.Println(book.Idx, book.Title, book.Author)
		}

	case "update":
		var id int

		if argLen < 4 {
			return errors.New(command[0] + ": need more params")
		} else {
			id, _ = strconv.Atoi(os.Args[2:3][0])

			book := model.Book{
				Idx:    null.IntFrom(int64(id)),
				Title:  null.StringFrom(os.Args[3:4][0]),
				Author: null.StringFrom(os.Args[4:5][0]),
			}

			count, err := crud.UpdateData(book)
			if err != nil {
				return err
			}

			log.Println("Update count:", count)
		}

	case "delete":
		var id int

		if argLen == 2 {
			idStr := os.Args[2:3]
			id, _ = strconv.Atoi(idStr[0])
		} else {
			id = 0
		}

		count, err := crud.DeleteData(id)
		if err != nil {
			return err
		}

		log.Println("Delete count:", count)

	case "drop-table":
		err := db.Obj.DropTable()
		if err != nil {
			log.Fatal("DropTable:", err)
		}

	case "rename-table":
		err := db.Obj.RenameTable()
		if err != nil {
			log.Fatal("RenameTable:", err)
		}

	default:
		return errors.New("check inputed parameters")
	}

	return nil
}

func init() {
	// db.Info = config.DatabaseInfoSQLite
	// db.Info = config.DatabaseInfoMySQL
	// db.Info = config.DatabaseInfoPgPublic
	// db.Info = config.DatabaseInfoPgSchema
	// db.Info = config.DatabaseInfoPgOtherDatabase
	// db.Info = config.DatabaseInfoSqlServer

	db.Info = config.DatabaseInfoOracle
	db.InfoOracleAdmin = config.DatabaseInfoOracleSystem
	// db.Info = config.DatabaseInfoOracleCloud
	// db.InfoOracleAdmin = config.DatabaseInfoOracleCloudAdmin
}

func main() {
	err := beginJob()

	if err != nil {
		log.Fatal(err)
	}
}
