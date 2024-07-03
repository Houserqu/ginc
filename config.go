package ginc

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConfig() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath != "" {
		log.Println("Load config file from: ", configPath)
		viper.SetConfigFile(configPath)
	} else {
		log.Println("Load config file from: ./config.yaml")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
	}

	// 设置默认值
	viper.SetDefault("server.addr", "0.0.0.0:8080")

	// 加载配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// 如果开启 nacos 配置中心，则初始化 nacos 配置，否则监听配置文件变化（不能同时使用 nacos 和监听本地文件）
	if viper.GetBool("nacos.enable") {
		// 初始化 nacos 配置
		InitNacosConfig()
	} else {
		// 监听配置文件变化
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("Config file changed:", e.Name)
		})
	}
}
