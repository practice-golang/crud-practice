package db

import (
	"database/sql"
)

type Mysql struct{ dsn string }

func (d *Mysql) connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", d.dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *Mysql) CreateDB() error {
	sql := `CREATE DATABASE IF NOT EXISTS ` + Info.DatabaseName + `;`
	_, err := Con.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func (d *Mysql) CreateTable() error {
	sql := `
	CREATE TABLE IF NOT EXISTS ` + Info.DatabaseName + `.` + Info.TableName + ` (
		IDX INT(11) NOT NULL AUTO_INCREMENT,
		TITLE VARCHAR(255) NULL DEFAULT NULL COLLATE 'utf8_general_ci',
		AUTHOR VARCHAR(255) NULL DEFAULT NULL COLLATE 'utf8_general_ci',
		PRIMARY KEY (IDX) USING BTREE
	)
	COLLATE='utf8_general_ci'
	ENGINE=InnoDB;
	`

	_, err := Con.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func (d *Mysql) DropTable() error {
	sql := `DROP TABLE ` + Info.DatabaseName + `.` + Info.TableName + `;`
	_, err := Con.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func (d *Mysql) RenameTable() error {
	sql := `RENAME TABLE ` + Info.DatabaseName + `.` + Info.TableName + ` TO ` + Info.DatabaseName + `.` + Info.TableName + `_RENAMED;`
	_, err := Con.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func (d *Mysql) Exec(sql string, colValues []interface{}, options string) (int64, int64, error) {
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
