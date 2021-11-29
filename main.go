package main

import (
	// "Back/internal/repo/mysql"
	// "Back/logger"
	// "Back/pkg/snowflake"
	// "Back/pkg/jwt"
	// "Back/routers"
	_ "Back/cmd"
	"Back/initialize"
	_ "Back/initialize"

	"fmt"

	"github.com/spf13/pflag"
)

var envName = pflag.StringP("env", "n", "dev", "Input dev's Env: dev or prod")

func main() {
	fmt.Println("call main")
	pflag.Parse()
	fmt.Println("name=", *envName)
	if err := initialize.InitWithEnv(*envName); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}

	// jwt.Init(initialize.Conf.JwtConfig)
	
	// if err := logger.Init(initialize.Conf.LogConfig, initialize.Conf.Mode); err != nil {
	// 	fmt.Printf("init logger failed, err:%v\n", err)
	// 	return
	// }
	// if err := mysql.Init(initialize.Conf.MySQLConfig); err != nil {
	// 	fmt.Printf("init mysql failed, err:%v\n", err)
	// 	return
	// }
	// defer mysql.Close() // 程序退出关闭数据库连接
	// if err := redis.Init(settings.Conf.RedisConfig); err != nil {
	// 	fmt.Printf("init redis failed, err:%v\n", err)
	// 	return
	// }
	// defer redis.Close()
	// if err := snowflake.Init(1); err != nil {
	// 	fmt.Printf("init snowflake failed, err:%v\n", err)
	// 	return
	// }
	// 注册路由
	// r := routers.SetupRouter()
	// err := r.Run(fmt.Sprintf(":%d", initialize.Conf.Port))
	// if err != nil {
	// 	fmt.Printf("run server failed, err:%v\n", err)
	// 	return
	// }
}
