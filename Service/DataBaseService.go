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

//MasterDataBaseName ...
const MasterDataBaseName = "master_data"

//ShardUserDataBaseName ...
const ShardUserDataBaseName = "user_data_"

//GetMysqlConnection 指定したdbへのコネクトを取得(プールはドライバがやってくれる)
func GetMysqlConnection(database string) *sql.DB {
	db, e := sql.Open("mysql", "root:password@/"+database) // <-この辺、appconfに移すべきなんだろうけどさ...
	if e != nil {
		logs.Error("database access error!")
	}
	return db
}
