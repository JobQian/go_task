package main

import (
	"fmt"
	initconfig "go_task_4/init"
<<<<<<< HEAD
	"go_task_4/internal/handler"
	"go_task_4/internal/repository"
	"go_task_4/internal/router"
	"go_task_4/internal/service"
	"go_task_4/pkg/utils"
	"net/http"

	"github.com/spf13/viper"
=======
	"go_task_4/pkg/utils"
>>>>>>> d7ec24b276d863d05d3c77537f56a03b8a283f67
)

func main() {
	initconfig.InitViperConfig("setting", "yaml", "./configs")
	utils.InitLogger()
	res := initconfig.InitDataBase()
	if res != nil {
		utils.Logger.Debug("数据库连接失败: " + res.Error())
		panic(fmt.Errorf("数据库连接失败: %w", res))
	}
<<<<<<< HEAD
	userrepository := repository.NewUserRepository(initconfig.DB)
	userService := service.NewUserService(userrepository)
	userHandler := handler.NewUserHandler(userService)
	router := router.NewRouter(userHandler)

	// 使用 http.Server 包装
	srv := &http.Server{
		Addr:           viper.GetString("server.port"),
		Handler:        router,
		ReadTimeout:    viper.GetDuration("server.read_timeout"),
		WriteTimeout:   viper.GetDuration("server.write_timeout"),
		MaxHeaderBytes: viper.GetInt("server.max_header_bytes"),
	}
	// 启动
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
=======
>>>>>>> d7ec24b276d863d05d3c77537f56a03b8a283f67

	sqlDB, err := initconfig.DB.DB()

	if err != nil {
		utils.Logger.Debug("获取底层数据库实例失败: " + err.Error())
	}
<<<<<<< HEAD
=======

>>>>>>> d7ec24b276d863d05d3c77537f56a03b8a283f67
	defer sqlDB.Close()
}
