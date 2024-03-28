package handles

import (
	"Termbin/DAO"
	"Termbin/functions"
	"Termbin/models"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var a models.Clipboard
	var b models.Clipboard //声明一个结构体变量，方便对数据库进行写入操作
	url := "localhost:888/create/"
	content := c.PostForm("c")
	a.Context = content
	a.Digest = functions.Hex(content)
	a.Short = functions.Short()
	a.Size = len(content)
	a.Url = url + a.Short
	DAO.InitSQL()                  //链接数据库
	result := models.Db.Create(&a) //创建表
	models.Db.Where("context=?", content).Find(&b)
	if result.Error == nil {
		c.String(200, "date:  %s\ndigest: %s\nshort: %s\nsize: %d\nurl: %s\nstatus: %s\nuuid: %d\n",
			b.CreatedAt, b.Digest, b.Short, b.Size, b.Url, "created", b.ID)
	} else {
		c.String(200, "失败了！！！")
	}
}
