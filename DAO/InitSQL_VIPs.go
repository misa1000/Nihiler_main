package DAO

import (
	"Termbin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func InitSQL_VIPs() {
	var err error
	var a models.Clipboard_VIP
	dsn := "root:Aa35754391@tcp(127.0.0.1:3306)/Termbin?charset=utf8mb4&parseTime=True&loc=Local"
	models.Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}, //复合文字的“}”前面需要尾随“,”
	})

	if err == nil {
		sqlDB, err := models.Db.DB()
		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		if err == nil {
			sqlDB.SetMaxIdleConns(10)

			// SetMaxOpenConns 设置打开数据库连接的最大数量。
			sqlDB.SetMaxOpenConns(100)

			// SetConnMaxLifetime 设置了连接可复用的最大时间。
			sqlDB.SetConnMaxLifetime(time.Hour)
			//07640
			models.Db.AutoMigrate(&a)
		}
	}
}
