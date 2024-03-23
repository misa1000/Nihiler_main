package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)
import (
	"github.com/gin-gonic/gin"
)

var db *gorm.DB

type User struct {
	gorm.Model        ///结构体里面的变量名必须大写，否则在数据库里面根本创建不出表
	Name       string `json:"name" gorm:"column:name" binding:"required"`
	Game       string `json:"game" gorm:"column:game" binding:"required"`
	Sex        string `json:"sex"  gorm:"column:sex" binding:"required"`
}

func Add1(c *gin.Context) {
	var user User
	err := c.ShouldBindJSON(&user)
	if err == nil {
		result := db.Create(&user) //此处是返回一个DB类的指针，因此用它的错误类别来判断是否会出错。。。
		if result.Error == nil {
			c.String(200, "增加非常成功，数据是%s", user)
		} else {
			c.String(200, "出错了！！！！！！！！！！！！！%s", user)
		}
	} else {
		c.String(200, "绑定有错误！！！！！")
	}
}
func test1(c *gin.Context) {
	c.String(200, "all is well!")
}
func delete(c *gin.Context) {
	var user []User
	id := c.Param("ID")
	fmt.Println(id)
	db.Where("ID=?", id).Find(&user)
	if len(user) == 0 {
		c.String(200, "没有查到对应的id,请确认id是否存在")
	} else {
		db.Where("id=?", id).Delete(&user)
		c.String(200, "删除成功了")
	}
}
func main() {
	var err error
	dsn := "root:Aa35754391@tcp(127.0.0.1:3306)/crudceshi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}, //复合文字的“}”前面需要尾随“,”
	})
	sqlDB, err := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	//07640
	fmt.Println(db)
	fmt.Println(err)
	engine := gin.Default()
	db.AutoMigrate(&User{})
	engine.DELETE("/user/DELETE/:ID", delete)
	engine.POST("/user/ADD", Add1)
	engine.GET("/test", test1)
	engine.Run(":3305")
}
