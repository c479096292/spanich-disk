package config

import "fmt"

var (
	MySQLSource = "test:test@tcp(127.0.0.1:3306)/fileserver?charset=utf8mb4"
)

func UpdateDBHost(host string) {
	MySQLSource = fmt.Sprintf("root:251700@tcp(%s)/fileserver?charset=utf8mb4", host)
}
