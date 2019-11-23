package config

import "fmt"

var (
	MySQLSource = "root:251700@tcp(127.0.0.1:3306)/fileserver?charset=utf8mb4"
)

func UpdateDBHost(username, userpwd, host string) {
	MySQLSource = fmt.Sprintf("%s:%s@tcp(%s)/fileserver?charset=utf8", username, userpwd, host)
}
