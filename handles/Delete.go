package handles

import (
	"Termbin/DAO"
	"Termbin/models"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	var a models.Clipboard
	DAO.InitSQL()
	id := c.Param("id")
	models.Db.Where("id=?", id).Find(&a)
	if a.ID != 0 {
		models.Db.Where("id=?", id).Delete(&a)
		c.String(200, "deleted %d", a.ID)
	} else {
		c.String(200, "无法删除，可能无此数据或者该数据已经被删除！！！！！")
	}
}
