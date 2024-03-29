package handles

import (
	"Termbin/DAO"
	"Termbin/functions"
	"Termbin/models"
	"github.com/gin-gonic/gin"
)

var Sessions = make(map[string]int) //声明一个Sessions用来记录登录的信息

func Login(c *gin.Context) {
	var a models.User
	username := c.Param("username")
	password := c.Param("password")
	DAO.InitSQL_Register()
	models.Db.Where("username=?", username).Find(&a)
	if a.ID != 0 {
		c.String(200, "确实存在用户\n")
		if a.Password == functions.Hex(password) {
			c.String(200, "登陆成功！")
			Sessions[username] = 1
		} else {
			c.String(200, "您的密码错误！")
		}
	} else {
		c.String(200, "查无此人！！！")
	}
}
