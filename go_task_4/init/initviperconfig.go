package init

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitViperConfig(namestr string, typestr string, pathstr string) {

	// 配置文件名，不带扩展名
	viper.SetConfigName(namestr)
	// 配置文件类型
	viper.SetConfigType(typestr)
	// 配置文件路径（当前目录） 可以配置多个目录-viper会按顺序查找
	viper.AddConfigPath(pathstr)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("配置读取失败: %w", err))
	}

	// 自动匹配环境变量
	viper.AutomaticEnv()

	//热更新，只要配置文件变动了就自动reload
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件被修改:", in.Name)
	})

}
