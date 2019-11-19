package db

import (
	"database/sql"
	"fmt"
	"github.com/c479096292/spinach-disk/config"
	"os"
)

var db *sql.DB

func init()  {
	db, _ = sql.Open("mysql",config.DB)
	db.SetMaxOpenConns(600)
	db.SetMaxIdleConns(100)
	err := db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to mysql, err:" + err.Error())
		os.Exit(1)
	}
}


func NewDB() *sql.DB {
	return db
}

