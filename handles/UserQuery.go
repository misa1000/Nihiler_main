package handles

import (
	"Termbin/DAO"
	"Termbin/models"
	"github.com/gin-gonic/gin"
)

func UserQuery(c *gin.Context) {
	var a models.Clipboard_VIP
	DAO.InitSQL_VIPs()
	id := c.Param("id")
	username := c.Param("username")
	models.Db.Where("id=?", id).Find(&a)
	if a.ID != 0 { //注意此处进行updates的批量修改，传入的是原始模板
		if username == a.Username || a.Visible_another == username || (a.Visible_VIP == "no" && a.Visible_another == "none") {
			c.String(200, "%s", a.Context)
		} else {
			c.String(200, "由于对方访问权限设置，你没有权限查看此剪切板内容")
		}
	} else {
		c.String(200, "没有找到该剪切板记录！！！")
	}
}
