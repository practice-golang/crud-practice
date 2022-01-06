package main // import "practice-crud"

import (
	"fmt"
	"log"
	"os"
	"practice-crud/config"
	"practice-crud/db"
	"practice-crud/function"
	"practice-crud/model"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"gopkg.in/guregu/null.v4"
	_ "modernc.org/sqlite"
)

func main() {
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

	// db.Info = config.DatabaseInfoSQLite
	// db.Info = config.DatabaseInfoMySQL
	db.Info = config.DatabaseInfoPgPublic
	// db.Info = config.DatabaseInfoPgSchema
	// db.Info = config.DatabaseInfoPgOtherDatabase

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
			fmt.Println(command[0] + ": Need more params")
			os.Exit(1)
		}

		book := model.Book{
			Title:  null.StringFrom(os.Args[2:3][0]),
			Author: null.StringFrom(os.Args[3:4][0]),
		}

		idx, count, err := function.InsertData(book)
		if err != nil {
			log.Fatal("InsertData:", err)
		}

		log.Println("Insert idx, count:", idx, count)

		break

	case "select":
		var id int

		if argLen < 1 {
			log.Fatal("Need more params")
		} else if argLen == 2 {
			idStr := os.Args[2:3]
			id, _ = strconv.Atoi(idStr[0])
		} else {
			id = 0
		}

		books, err := function.SelectData(id)
		if err != nil {
			log.Fatal("SelectData:", err)
		}
		for _, book := range books {
			log.Println(book.ID, book.Title, book.Author)
		}

		break

	case "update":
		var id int

		if argLen < 4 {
			log.Println("Need more params")
		} else {
			id, _ = strconv.Atoi(os.Args[2:3][0])

			book := model.Book{
				ID:     null.IntFrom(int64(id)),
				Title:  null.StringFrom(os.Args[3:4][0]),
				Author: null.StringFrom(os.Args[4:5][0]),
			}

			count, err := function.UpdateData(book)
			if err != nil {
				log.Fatal("UpdateData:", err)
			}

			log.Println("Update count:", count)
		}

		break

	case "delete":
		var id int

		if argLen < 1 {
			log.Fatal("Need more params")
		} else if argLen == 2 {
			idStr := os.Args[2:3]
			id, _ = strconv.Atoi(idStr[0])
		} else {
			id = 0
		}

		count, err := function.DeleteData(id)
		if err != nil {
			log.Fatal("DeleteData:", err)
		}

		log.Println("Delete count:", count)

		break

	default:
		log.Println("check inputed parameters")
	}
}
