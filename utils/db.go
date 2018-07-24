package utils

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func OpenDB(username, password, mysqlPath, dbName string) (initDb *sql.DB, err error) {
	initDb, err = sql.Open("mysql", username+":"+password+"@"+mysqlPath+"/"+dbName+"?parseTime=true")
	if err != nil {
		log.Println(err)
	}
	DB = initDb
	return
}
