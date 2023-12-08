package db

import (
	"withChat/internal/db/mysql"
	"withChat/internal/db/redis"
)

func Initialize() {
	mysql.InitializeMysql() // 启动mysql连接
	redis.InitializeRedis() // 启动redis连接
}
