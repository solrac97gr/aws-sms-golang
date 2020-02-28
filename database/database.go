package database

import (
	"database/sql"

	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	err   error
	DBCon *sql.DB
)

func CreateConection() *sql.DB {
	DBCon, err = sql.Open("mysql", os.Getenv("CONNECTION"))
	if err != nil {
		panic(err.Error())
	}
	return DBCon
}
