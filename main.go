package main

import (
	"Back/internal/repo/mysql"
	"Back/logger"
	"Back/pkg/snowflake"
	"Back/pkg/jwt"
	"Back/routers"
	"Back/settings"
	"fmt"
)

func main() {
	// if err := settings.Init(); err != nil {
	// 	fmt.Printf("load config failed, err:%v\n", err)
	// 	return
	// }

	jwt.Init(settings.Conf.JwtConfig)
	
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	// if err := redis.Init(settings.Conf.RedisConfig); err != nil {
	// 	fmt.Printf("init redis failed, err:%v\n", err)
	// 	return
	// }
	// defer redis.Close()
	if err := snowflake.Init(1); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	// 注册路由
	r := routers.SetupRouter()
	err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
