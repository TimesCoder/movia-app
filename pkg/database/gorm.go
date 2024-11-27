package database

import (
	"fmt"

	"github.com/TimesCoder/movie-app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDatabase(MysqlConfig config.MySQLConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		MysqlConfig.User,
		MysqlConfig.Password,
		MysqlConfig.HOST,
		MysqlConfig.PORT,
		MysqlConfig.Database,
	)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
