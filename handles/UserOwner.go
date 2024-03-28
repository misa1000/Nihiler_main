package handles

import (
	"Termbin/DAO"
	"Termbin/models"
	"github.com/gin-gonic/gin"
)

func UserOwner(c *gin.Context) {
	var a models.Clipboard_VIP
	username := c.Param("username")
	id := c.Param("id")
	allowance := c.Param("visible_VIP")
	DAO.InitSQL_VIPs()
	models.Db.Where("id=?", id).Find(&a)
	if a.Username == username {
		models.Db.Where("id=?", id).Updates(models.Clipboard_VIP{Visible_VIP: allowance})
		c.String(200, "设置成功！")
		if a.Visible_another == "none" {
			models.Db.Where("id=?", id).Updates(models.Clipboard_VIP{Visible_another: "none"})
		} else {
			models.Db.Where("id=?", id).Updates(models.Clipboard_VIP{Visible_another: "none"}) //如果此时该剪切板没有执行过访问权限控制，那么加以设置
		}
	} else {
		c.String(200, "设置失败，您没有权限或者没有您并没有创建这一条记录。")
	}
}
