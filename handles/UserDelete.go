package handles

import (
	"Termbin/DAO"
	"Termbin/models"
	"github.com/gin-gonic/gin"
)

func UserDelete(c *gin.Context) {
	var a models.Clipboard_VIP
	DAO.InitSQL_VIPs()
	id := c.Param("id")
	username := c.Param("username")
	models.Db.Where("id=?", id).Find(&a)
	if a.ID != 0 {
		if a.Username == username {
			models.Db.Where("id=?", id).Delete(&a)
			c.String(200, "deleted %d", a.ID)
		} else {
			c.String(200, "你并非该栏作者，没有修改权限")
		}
	} else {
		c.String(200, "无法删除，可能无此数据或者该数据已经被删除！！！！！")
	}
}
