package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err = viper.ReadInConfig()    // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		fmt.Errorf("fatal error config file: %w", err)
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件发生变化, 热重载中...")
	})
	// fmt.Println("get config: ", viper.Get("app.name"))

	return
}
