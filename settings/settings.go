package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Mode         string `mapstructure:"mode"`
	Port         int    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
	*JwtConfig `mapstructure:"jwt"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"db"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}


type JwtConfig struct {
	SecretKey string `mapstructure:"secretKey"`
}


func init() {


	fmt.Println("call setting init")
	viper.SetConfigFile("./configs/config.yaml")

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		viper.Unmarshal(&Conf)
	})

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("ReadInConfig failed, err: %v", err))
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unmarshal to Conf failed, err:%v", err))
	}
	
	// return err
}
