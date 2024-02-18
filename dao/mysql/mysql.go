package mysql

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() (err error) {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		zap.L().Error("MySQL数据库连接失败", zap.Error(err))
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		zap.L().Error("MySQL获取实例失败, err: %v\n", zap.Error(err))
		return
	}
	// 设置最大连接数
	sqlDB.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))

	return
}

// Close 关闭数据库连接
func Close() {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}
