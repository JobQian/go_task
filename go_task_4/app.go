package main

import (
	initconfig "go_task_4/init"
	"go_task_4/pkg/utils"
)

func main() {
	initconfig.InitViperConfig("setting", "yaml", "./configs")
	utils.InitLogger()
	initconfig.InitDataBase()

	sqlDB, err := initconfig.DB.DB()

	if err != nil {
		utils.Logger.Debug(err.Error())
	}
	defer sqlDB.Close()

}
