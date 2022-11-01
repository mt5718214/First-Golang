package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var SqlDB *sql.DB

func init() {
	// db connection
	var err error
	SqlDB, err = sql.Open("mysql", "demo:demo123@tcp(localhost:3306)/demo")
	if err != nil {
		fmt.Println("DB連線資訊有誤請再次確認")
	}
	if err := SqlDB.Ping(); err != nil {
		fmt.Println("開啟 MySQL 連線發生錯誤，原因為：", err.Error())
		log.Fatal("db connect error: ", err.Error())
	}
}
