package config

import "practice-crud/db"

var DatabaseInfoSQLite = db.DBInfo{
	DatabaseType: db.SQLITE,
	DatabaseName: "books",
	TableName:    "books",
	FilePath:     "./books.db",
}

var DatabaseInfoMySQL = db.DBInfo{
	DatabaseType:  db.MYSQL,
	Protocol:      "tcp",
	Addr:          "localhost",
	Port:          "13306",
	DatabaseName:  "myslimsite",
	SchemaName:    "",
	TableName:     "books",
	GrantID:       "root",
	GrantPassword: "",
}

var DatabaseInfoPgPublic = db.DBInfo{
	DatabaseType:  db.POSTGRES,
	Protocol:      "tcp",
	Addr:          "localhost",
	Port:          "5432",
	DatabaseName:  "postgres",
	SchemaName:    "public",
	TableName:     "books",
	GrantID:       "root",
	GrantPassword: "pgsql",
}

var DatabaseInfoPgSchema = db.DBInfo{
	DatabaseType:  db.POSTGRES,
	Protocol:      "tcp",
	Addr:          "localhost",
	Port:          "5432",
	DatabaseName:  "postgres",
	SchemaName:    "myslimsite",
	TableName:     "books",
	GrantID:       "root",
	GrantPassword: "pgsql",
}

// For not using database name 'postgres', you should create database yourself
var DatabaseInfoPgOtherDatabase = db.DBInfo{
	DatabaseType:  db.POSTGRES,
	Protocol:      "tcp",
	Addr:          "localhost",
	Port:          "5432",
	DatabaseName:  "mysitedb",
	SchemaName:    "myslimsite",
	TableName:     "books",
	GrantID:       "root",
	GrantPassword: "pgsql",
}
