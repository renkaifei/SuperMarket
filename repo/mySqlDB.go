package repo

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	mySqlDB *sql.DB
)

const (
	sqlDrive            = "mysql"
	sqlConnectionString = "root:19881220@/superMarket?charset=utf8"
)

func init() {
	db, err := sql.Open(sqlDrive, sqlConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	mySqlDB = db
}
