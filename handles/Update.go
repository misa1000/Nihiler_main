package handles

import (
	"Termbin/DAO"
	"Termbin/models"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	var a models.Clipboard
	DAO.InitSQL()
	content := c.PostForm("c")
	id := c.Param("id")
	models.Db.Where("id=?", id).Find(&a)
	if a.ID != 0 {
		models.Db.Where("id=?", id).Updates(models.Clipboard{Context: content}) //注意此处进行updates的批量修改，传入的是原始模板
		c.String(200, "localhost:888/%s updated", a.Short)
	} else {
		c.String(200, "没有找到该剪切板记录！！！")
	}
}
