package main

import (
	"fmt"
	initconfig "go_task_4/init"
	"go_task_4/internal/handler"
	"go_task_4/internal/repository"
	"go_task_4/internal/router"
	"go_task_4/internal/service"
	"go_task_4/pkg/utils"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	initconfig.InitViperConfig("setting", "yaml", "./configs")
	utils.InitLogger()
	res := initconfig.InitDataBase()
	if res != nil {
		utils.Logger.Debug("数据库连接失败: " + res.Error())
		panic(fmt.Errorf("数据库连接失败: %w", res))
	}

	userrepository := repository.NewUserRepository(initconfig.DB)
	userservice := service.NewUserService(userrepository)
	userhandler := handler.NewUserHandler(userservice)
	postrepository := repository.NewPostRepository(initconfig.DB)
	postservice := service.NewPostService(postrepository)
	posthandler := handler.NewPostHandler(postservice)
	commentrepository := repository.NewCommentRepository(initconfig.DB)
	commentservcie := service.NewCommentService(commentrepository)
	commenthandler := handler.NewCommentHandler(commentservcie)

	router := router.NewRouter(userhandler, posthandler, commenthandler)

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

	sqlDB, err := initconfig.DB.DB()

	if err != nil {
		utils.Logger.Debug("获取底层数据库实例失败: " + err.Error())
	}
	defer sqlDB.Close()
}
