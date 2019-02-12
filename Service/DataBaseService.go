package service

import (
	"database/sql"
	"strconv"

	"github.com/astaxie/beego/logs"

	_ "github.com/go-sql-driver/mysql"
)

func GetMysqlConnection(shardId int) *sql.DB {
	No := strconv.Itoa(shardId)
	db, e := sql.Open("mysql", "root:password@/unix(/var/run/mysqld/mysqld.sock)/user_data"+No)
	if e != nil {
		logs.Debug("database access error!")
	}
	return db
}
