package service

import (
	"database/sql"

	"github.com/astaxie/beego/logs"

	//mysqlドライバ
	_ "github.com/go-sql-driver/mysql"
)

//DataBaseShardMax dbシャード最大数
const DataBaseShardMax = 2

//UserEntryDataBaseName ...
const UserEntryDataBaseName = "user_data"

//GetMysqlConnection シャード分けされたユーザデータdb
func GetMysqlConnection(database string) *sql.DB {
	db, e := sql.Open("mysql", "root:password@/"+database)
	if e != nil {
		logs.Error("database access error!")
	}
	return db
}
