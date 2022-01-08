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

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"gopkg.in/guregu/null.v4"
	_ "modernc.org/sqlite"
)

func doJob() error {
	var err error

	// 실행파일(#1) 명령(#2) 대상(#3...)
	argLen := len(os.Args) - 1
	if argLen < 1 {
		fmt.Println("Usage:")
		fmt.Println("appmain.exe insert book_name author_name")
		fmt.Println("appmain.exe select [id_number]")
		fmt.Println("appmain.exe update id_number book_name author_name")
		fmt.Println("appmain.exe delete id_number")
		os.Exit(1)
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
	defer db.Con.Close()

	switch command[0] {
	case "insert":
		if argLen < 3 {
			return errors.New(command[0] + ": Need more params")
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

		if argLen < 1 {
			return errors.New(command[0] + ": Need more params")
		} else if argLen == 2 {
			idStr := os.Args[2:3]
			id, _ = strconv.Atoi(idStr[0])
		}

		books, err := crud.SelectData(id)
		if err != nil {
			return err
		}

		for _, book := range books {
			log.Println(book.Idx, book.Title, book.Author)
		}

	case "update":
		var id int

		if argLen < 4 {
			return errors.New(command[0] + ": Need more params")
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

		if argLen < 1 {
			return errors.New(command[0] + ": Need more params")
		} else if argLen == 2 {
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

	default:
		return errors.New("check inputed parameters")
	}

	return nil
}

func init() {
	db.Info = config.DatabaseInfoSQLite
	// db.Info = config.DatabaseInfoMySQL
	// db.Info = config.DatabaseInfoPgPublic
	// db.Info = config.DatabaseInfoPgSchema
	// db.Info = config.DatabaseInfoPgOtherDatabase
}

func main() {
	err := doJob()

	if err != nil {
		log.Fatal(err)
	}
}
