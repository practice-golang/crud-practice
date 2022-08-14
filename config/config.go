package config

import (
	"crud-practice/db"
	"crud-practice/model"
)

var DatabaseInfoSQLite = db.DBInfo{
	DatabaseType: model.SQLITE,
	DatabaseName: "books",
	TableName:    "books",
	FilePath:     "./books.db",
}

var DatabaseInfoMySQL = db.DBInfo{
	DatabaseType:  model.MYSQL,
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
	DatabaseType:  model.POSTGRES,
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
	DatabaseType:  model.POSTGRES,
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
	DatabaseType:  model.POSTGRES,
	Protocol:      "tcp",
	Addr:          "localhost",
	Port:          "5432",
	DatabaseName:  "mysitedb",
	SchemaName:    "myslimsite",
	TableName:     "books",
	GrantID:       "root",
	GrantPassword: "pgsql",
}

var DatabaseInfoSqlServer = db.DBInfo{
	DatabaseType:  model.SQLSERVER,
	Protocol:      "tcp",
	Addr:          "localhost",
	Port:          "1433",
	DatabaseName:  "mysitedb",
	SchemaName:    "dbo",
	TableName:     "books",
	GrantID:       "sa",
	GrantPassword: "SQLServer1433",
}

// ID = system
var DatabaseInfoOracleSystem = db.DBInfo{
	DatabaseType:  model.ORACLE,
	Protocol:      "tcp",
	Addr:          "localhost",
	Port:          "1521",
	DatabaseName:  "XE",
	SchemaName:    "",
	TableName:     "books",
	GrantID:       "system",
	GrantPassword: "oracle",
	FilePath:      "",
}

// GrantID = DatabaseName
var DatabaseInfoOracle = db.DBInfo{
	DatabaseType:  model.ORACLE,
	Protocol:      "tcp",
	Addr:          "localhost",
	Port:          "1521",
	DatabaseName:  "myaccount",
	SchemaName:    "",
	TableName:     "books",
	GrantID:       "myaccount",
	GrantPassword: "mypassword",
	FilePath:      "",
}

// ID = ADMIN
var DatabaseInfoOracleCloudAdmin = db.DBInfo{
	DatabaseType:  model.ORACLE,
	Protocol:      "tcp",
	Addr:          "adb.ap-chuncheon-1.oraclecloud.com",
	Port:          "1522",
	DatabaseName:  "g8364568bcaeec3_myora_low.adb.oraclecloud.com",
	SchemaName:    "",
	TableName:     "books",
	GrantID:       "admin",
	GrantPassword: "dkssudglrPtpdy0!",
	FilePath:      "./wallet_admin",
}

// GrantID = myaccount
var DatabaseInfoOracleCloud = db.DBInfo{
	DatabaseType:  model.ORACLE,
	Protocol:      "tcp",
	Addr:          "adb.ap-chuncheon-1.oraclecloud.com",
	Port:          "1522",
	DatabaseName:  "g8364568bcaeec3_myora_low.adb.oraclecloud.com",
	SchemaName:    "",
	TableName:     "books",
	GrantID:       "edp1096",
	GrantPassword: "Dkstkswjdqh@08gh",
	FilePath:      "./wallet_edp1096",
}
