package init

import (
	"fmt"
	"go_task_4/internal/model"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type Config struct {
	Driver   string
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

func InitDataBase() error {

	dbconfig := &Config{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
		Username: viper.GetString("database.user"),
		Password: viper.GetString("database.passowrd"),
		DBName:   viper.GetString("database.dbname"),
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbconfig.Username,
		dbconfig.Password,
		dbconfig.Host,
		dbconfig.Port,
		dbconfig.DBName,
	)
	fmt.Println(dsn)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // 输出到控制台
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值 eg:500 * time.Millisecond
			LogLevel:                  logger.Info, // 打印所有 SQL
			IgnoreRecordNotFoundError: true,        // 忽略 "record not found" 错误
			Colorful:                  true,        // 彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		return err
	}

	DB = db
	//初始化 检查 创建 所有表
	res := DB.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
	if res != nil {
		return res
	}
	return nil
}
