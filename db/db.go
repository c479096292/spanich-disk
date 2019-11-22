package db

import (
	"database/sql"
	"fmt"
	"github.com/c479096292/spinach-disk/config"
	_ "github.com/go-sql-driver/mysql"
	"log"
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

// ExecResult: sql函数执行的结果
type ExecResult struct {
	Suc  bool        `json:"suc"`
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ParseRows(rows *sql.Rows) []map[string]interface{} {
	columns, _ := rows.Columns() // 返回全部列名
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		// 将values的地址保存到scanArgs中
		// 下面rows.Scan时,实际上时将值保存到了values中
		scanArgs[j] = &values[j]
	}

	record := make(map[string]interface{})
	// 只考虑解析一行的场景,如果初始化大于0,则后面取值会取到nil
	records := make([]map[string]interface{}, 0)

	for rows.Next() { // 下一行数据存在则继续循环
		//将行数据保存到record字典
		err := rows.Scan(scanArgs...)
		checkErr(err)

		for i, col := range values {
			if col != nil {
				record[columns[i]] = col
			}
		}
		records = append(records, record)
	}
	return records
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}