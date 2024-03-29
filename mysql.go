package ginc

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Mysql *gorm.DB

func InitMysql() {
	if viper.GetString("mysql.host") == "" {
		return
	}

	// 日志级别
	disableLog := viper.GetBool("mysql.disable_log")
	logLevel := logger.Info
	if disableLog {
		logLevel = logger.Silent
	}

	dbLogger := logger.New(
		Logger, // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logLevel,    // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)

	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	user := viper.GetString("mysql.user")
	password := viper.GetString("mysql.password")
	database := viper.GetString("mysql.database")

	dsn := fmt.Sprint(user, ":", password, "@tcp(", host, ":", port, ")/", database, "?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{
		Logger: dbLogger,
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	Mysql = db
}
