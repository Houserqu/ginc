package ginc

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

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
