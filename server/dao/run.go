/**
* @Author: yexichen
* @Date:2023/7/11-19:23
* @Desc
**/

package dao

import (
	"database/sql"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

// Db 全局变量
var (
	DB, err = sql.Open("mysql", "root:277187@tcp(127.0.0.1:3306)/chess?charset=utf8mb4&loc=Local&parseTime=True")
	Redisdb *redis.Client
)

func RUNDB() {
	//启用数据库
	DB, err = sql.Open("mysql", "root:277187@tcp(127.0.0.1:3306)/chess?charset=utf8mb4&loc=Local&parseTime=True")
	//initRedis()
}
