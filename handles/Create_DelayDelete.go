package handles

import (
	"Termbin/DAO"
	"Termbin/functions"
	"Termbin/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func Create_CelayDelete(c *gin.Context) {
	var a models.Clipboard_VIP
	var b models.Clipboard_VIP //声明一个结构体变量，方便对数据库进行写入操作
	url := "localhost:888/create/"
	time1 := c.PostForm("sunset")     //接受设置的秒数大小字符串
	time2, err := strconv.Atoi(time1) //字符串转化为int
	if err == nil {
	} else {
		return
	}
	time3 := fmt.Sprintf("%ds", time2)              //int转化为时间字符
	delayDuration, err := time.ParseDuration(time3) //转Duration
	content := c.PostForm("c")
	username := c.Param("username")
	a.Context = content
	a.Digest = functions.Hex(content)
	a.Short = functions.Short()
	a.Size = len(content)
	a.Url = url + a.Short
	a.Username = username
	a.Visible_VIP = "no"
	a.Visible_another = "none"
	DAO.InitSQL_VIPs()             //链接数据库
	result := models.Db.Create(&a) //创建表
	models.Db.Where("context=?", content).Find(&b)
	if result.Error == nil {
		c.String(200, "date:  %s\ndigest: %s\nshort: %s\nsize: %d\nurl: %s\nstatus: %s\nuuid: %d\nusername:%s",
			b.CreatedAt, b.Digest, b.Short, b.Size, b.Url, "created", b.ID, b.Username)
		go func() {
			time.Sleep(delayDuration)
			models.Db.Where("id=?", b.ID).Delete(&b)
		}() //在此处开辟一个微量级线程
	} else {
		c.String(200, "失败了！！！")
	}
}
