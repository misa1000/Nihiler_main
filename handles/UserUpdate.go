package handles

import (
	"Termbin/DAO"
	"Termbin/models"
	"github.com/gin-gonic/gin"
)

func UserUpdate(c *gin.Context) {
	var a models.Clipboard_VIP
	DAO.InitSQL_VIPs()
	content := c.PostForm("c")
	id := c.Param("id")
	username := c.Param("username")
	models.Db.Where("id=?", id).Find(&a)
	if a.ID != 0 {
		if username == a.Username {
			models.Db.Where("id=?", id).Updates(models.Clipboard_VIP{Context: content}) //注意此处进行updates的批量修改，传入的是原始模板
			c.String(200, "localhost:888/%s updated", a.Short)
		} else {
			c.String(200, "此栏并非为你所做，因此你没有权限进行访问")
		}
	} else {
		c.String(200, "没有找到该剪切板记录！！！")
	}
}
