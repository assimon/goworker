package bootstrap

import (
	"goworker/core/config"
	"goworker/core/database"
	"goworker/core/gredis"
	"goworker/core/logging"
)

func init()  {
	// 核心配置加载 core config load.
	config.AppInit()
	// 加载数据库配置.
	database.DatabaseInit()
	// 加载日志配置
	logging.LogInit()
	// 加载redis
	gredis.RedisInit()
}