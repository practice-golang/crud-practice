package db

import (
	"database/sql"
	"net/url"
	"strconv"
	"strings"

	go_ora "github.com/sijms/go-ora/v2"
)

type Oracle struct {
	dsn     string
	Version int64
}

func (d *Oracle) createAccount() {
	var err error

	tableSpace := "USERS"
	port, _ := strconv.Atoi(InfoOracleAdmin.Port)
	dsn := go_ora.BuildUrl(InfoOracleAdmin.Addr, port, InfoOracleAdmin.DatabaseName, InfoOracleAdmin.GrantID, InfoOracleAdmin.GrantPassword, nil)
	if InfoOracleAdmin.FilePath != "" {
		tableSpace = "DATA"
		dsn += "?SSL=enable&SSL Verify=false&WALLET=" + url.QueryEscape(InfoOracleAdmin.FilePath)
	}

	conn, err := sql.Open("oracle", dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	versionSTR := ""
	sql := `SELECT version FROM V$INSTANCE`
	err = conn.QueryRow(sql).Scan(&versionSTR)
	if err != nil {
		panic(err)
	}

	d.Version, _ = strconv.ParseInt(strings.Split(versionSTR, ".")[0], 10, 64)

	sql = `
	SELECT COUNT(USERNAME) AS COUNT
	FROM ALL_USERS
	WHERE USERNAME = '` + strings.ToUpper(Info.GrantID) + `'`

	var count int64
	_ = conn.QueryRow(sql).Scan(&count)
	if count > 0 {
		return
	}

	sql = `CREATE USER ` + Info.GrantID + ` IDENTIFIED BY "` + Info.GrantPassword + `"`
	if InfoOracleAdmin.FilePath != "" {
		sql += `
		DEFAULT TABLESPACE ` + tableSpace + `
		TEMPORARY TABLESPACE TEMP`
	}
	_, err = conn.Exec(sql)
	if err != nil {
		panic(err)
	}

	sql = `GRANT CONNECT, RESOURCE TO ` + Info.GrantID
	_, err = conn.Exec(sql)
	if err != nil {
		panic(err)
	}

	sql = `ALTER USER ` + Info.GrantID + ` DEFAULT TABLESPACE ` + tableSpace + ` QUOTA UNLIMITED ON ` + tableSpace
	_, err = conn.Exec(sql)
	if err != nil {
		panic(err)
	}
}

func (d *Oracle) connect() (*sql.DB, error) {
	db, err := sql.Open("oracle", d.dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *Oracle) CreateDB() error {
	err := Con.Ping()
	if err != nil {
		if strings.Contains(err.Error(), "ORA-01017") {
			d.createAccount()
			return nil
		}
	}

	return err
}

func (d *Oracle) CreateTable() error {
	var err error
	var count int64

	tableName := strings.ToUpper(`"` + Info.GrantID + `"."` + Info.TableName + `"`)

	sql := `
	SELECT COUNT(TABLE_NAME) AS COUNT
	FROM user_tables
	WHERE TABLE_NAME = '` + strings.ToUpper(Info.TableName) + `'`

	err = Con.QueryRow(sql).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	sql = `
	CREATE TABLE ` + tableName + ` (
		"IDX"       NUMBER(11) GENERATED ALWAYS AS IDENTITY (START WITH 1 INCREMENT BY 1) NOT NULL,
		"TITLE"     VARCHAR2(128),
		"AUTHOR"    VARCHAR2(128),

		UNIQUE("IDX")
	)`

	_, err = Con.Exec(sql)
	if err != nil {
		panic(err)
	}

	return nil
}

func (d *Oracle) DropTable() error {
	var err error
	var count int64

	tableName := strings.ToUpper(`"` + Info.GrantID + `"."` + Info.TableName + `"`)

	sql := `
	SELECT COUNT(TABLE_NAME) AS COUNT
	FROM user_tables
	WHERE TABLE_NAME = '` + strings.ToUpper(Info.TableName) + `'`

	err = Con.QueryRow(sql).Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		return nil
	}

	sql = `DROP TABLE ` + tableName
	_, err = Con.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func (d *Oracle) RenameTable() error {
	var err error

	tableName := strings.ToUpper(`"` + Info.TableName + `"`)
	tableNameRename := strings.ToUpper(`"` + Info.TableName + `_RENAMED"`)

	sql := `ALTER TABLE ` + tableName + ` RENAME TO ` + tableNameRename
	_, err = Con.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func (d *Oracle) Exec(sql string, colValues []interface{}, options string) (int64, int64, error) {
	var err error
	var count int64 = 0
	var idx int64 = 0

	result, err := Con.Exec(sql, colValues...)
	if err != nil {
		return count, idx, err
	}

	count, _ = result.RowsAffected()
	idx, _ = result.LastInsertId()

	return count, idx, nil
}
