package handles

import "github.com/gin-gonic/gin"

func User(c *gin.Context) {
	username := c.Param("username")
	if Sessions[username] == 0 {
		c.String(200, "该用户并未登陆")
		c.Abort() //通过map--Sessions检查是否存有登陆信息
	} else {
		c.Next()
	}
}
