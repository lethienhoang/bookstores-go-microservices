package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_USER   = "DB_USER"
	DB_PASSWD = "DB_PASSWD"
	DB_HOST   = "DB_HOST"
	DB_SCHEME = "DB_SCHEME"
)

var (
	DB *sql.DB
)

func NewDB(dbUser string, dbPasswd string, dbHost string, dbScheme string) {
	fmt.Println("Starting database...")
	//[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", dbUser, dbPasswd, dbHost, dbScheme)

	var err error
	DB, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
