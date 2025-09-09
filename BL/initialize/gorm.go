package initialize

import (
	"blog-go/global"
	"os"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() *gorm.DB {
	mysqlConfig := global.Config.Mysql
	dsn := mysqlConfig.Dsn()
	// global.Log.Info(fmt.Sprintf("Connect: Mysql [%v]", dsn))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(mysqlConfig.LogLevel()),
	})
	if err != nil {
		global.Log.Error("Failed to connect to MySQL:", zap.Error(err))
		os.Exit(-1)
	}

	sqlDB, _ := db.DB()

	sqlDB.SetMaxIdleConns(mysqlConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(mysqlConfig.MaxOpenConns)

	return db

}
