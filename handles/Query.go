package handles

import (
	"Termbin/DAO"
	"Termbin/models"
	"github.com/gin-gonic/gin"
)

func Query(c *gin.Context) {
	var a models.Clipboard
	DAO.InitSQL()
	id := c.Param("id")
	models.Db.Where("id=?", id).Find(&a)
	if a.ID != 0 { //注意此处进行updates的批量修改，传入的是原始模板
		c.String(200, "%s", a.Context)
	} else {
		c.String(200, "没有找到该剪切板记录！！！")
	}
}
