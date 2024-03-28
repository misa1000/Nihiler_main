package handles

import (
	"Termbin/DAO"
	"Termbin/models"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var a models.User
	username := c.Param("username")
	password := c.Param("password")
	a.Username = username
	a.Password = password
	DAO.InitSQL_Register()
	models.Db.Where("username=?", username).Find(&a)
	if a.ID != 0 {
		c.String(200, "该用户名已经存在！请重新注册")
	} else {
		models.Db.Create(&a)
		c.String(200, "注册成功！")
	}
}
