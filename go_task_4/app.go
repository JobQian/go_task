package main

import (
	"fmt"
	initconfig "go_task_4/init"
	"go_task_4/pkg/utils"
)

func main() {
	initconfig.InitViperConfig("setting", "yaml", "./configs")
	utils.InitLogger()
	res := initconfig.InitDataBase()
	if res != nil {
		utils.Logger.Debug("数据库连接失败: " + res.Error())
		panic(fmt.Errorf("数据库连接失败: %w", res))
	}

	sqlDB, err := initconfig.DB.DB()

	if err != nil {
		utils.Logger.Debug("数据库关闭失败: " + err.Error())
		panic(fmt.Errorf("数据库关闭失败: %w", err))
	}

	defer sqlDB.Close()
}
