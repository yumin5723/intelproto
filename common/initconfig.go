package common

import (
	"fmt"
	"intel-demo-inventory/global"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化配置文件
func InitConfig(env string) {
	//读取配置文件
	ReadConfig(env)
	//生效mysql配置文件
	initMysql()
	//生效redis配置文件
}

var (
	redisOnce   sync.Once
	err         error
	redisHelper *RedisHelper
)

func initMysql() {
	var buffer strings.Builder
	buffer.WriteString(global.GVA_CONFIG.Mysql.Username)
	buffer.WriteString(":")
	buffer.WriteString(global.GVA_CONFIG.Mysql.Password)
	buffer.WriteString("@tcp(")
	buffer.WriteString(global.GVA_CONFIG.Mysql.Path)
	buffer.WriteString(")/")
	buffer.WriteString(global.GVA_CONFIG.Mysql.Dbname)
	buffer.WriteString("?charset=utf8mb4")
	buffer.WriteString("&parseTime=True&loc=Local")

	dsn := buffer.String()
	global.GVA_DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("mysql connect fail...")
		return
	}
	sqlDB, _ := global.GVA_DB.DB()
	sqlDB.SetMaxIdleConns(global.GVA_CONFIG.Mysql.MaxIdleConns)
	sqlDB.SetMaxOpenConns(global.GVA_CONFIG.Mysql.MaxOpenConns)
}

type RedisHelper struct {
	*redis.Client
}

func initRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:         global.GVA_CONFIG.Redis.Addr,
		Password:     global.GVA_CONFIG.Redis.Password,
		DB:           global.GVA_CONFIG.Redis.DB,
		DialTimeout:  time.Duration(global.GVA_CONFIG.Redis.DiaTimeout) * time.Second,
		ReadTimeout:  time.Duration(global.GVA_CONFIG.Redis.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(global.GVA_CONFIG.Redis.WriteTimeout) * time.Second,
		PoolSize:     global.GVA_CONFIG.Redis.PoolSize,
		PoolTimeout:  time.Duration(global.GVA_CONFIG.Redis.PoolTimeout) * time.Second,
	})
	redisOnce.Do(func() {
		rdh := new(RedisHelper)
		rdh.Client = rdb
		redisHelper = rdh
		global.GVA_REDIS = redisHelper.Client
	})
}

// 初始化配置文件
func ReadConfig(env string) {
	// 配置文件路径 (包名 + 配置文件名 )
	defaultConfigFile := "config/config-" + env + ".yaml"

	v := viper.New()
	v.SetConfigFile(defaultConfigFile)

	// 读取配置文件中的配置信息，并将信息保存 到 v中
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: #{err}\n"))
	}
	// 监控配置文件
	v.WatchConfig()

	// 配置文件改变，则将 v中的配置信息，刷新到 global.GVA_CONFIG
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	// 将 v 中的配置信息 反序列化成 结构体 (将v 中配置信息 刷新到 global.GVA_CONFIG)
	if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}

	// 保存 viper 实例 v
	global.GVA_VP = v
	initMysql()
	initRedis()
}
