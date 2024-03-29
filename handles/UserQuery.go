package handles

import (
	"Termbin/DAO"
	"Termbin/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func UserQuery(c *gin.Context) {
	var a models.Clipboard_VIP
	DAO.InitSQL_VIPs()
	interface_ := c.Param("interface")
	key := c.PostForm("c")
	username := c.Param("username")
	num, err := strconv.Atoi(interface_)
	var q int
	if err == nil && num != 0 {
		q = 1
	} else if len(interface_) == 4 {
		q = 2
	} else {
		q = 3
	}
	switch q {
	case 1:
		models.Db.Where("id=?", interface_).Find(&a)
	case 2:
		models.Db.Where("short=?", interface_).Find(&a)
	case 3:
		models.Db.Where("alias=?", interface_).Find(&a)
	default:
		return
	}
	if a.ID != 0 { //注意此处进行updates的批量修改，传入的是原始模板
		if username == a.Username || a.Visible_another == username || (a.Visible_VIP == "no" && a.Visible_another == "none") {
			if len(key) == 0 {
				c.String(200, "%s", a.Context)
			} else {
				if username == a.Visible_another && key == "Misaka_Mikoto" {
					c.String(200, "%s\n", a.Context)
					c.String(200, "阅后即焚功能已经实现") //此处是检测阅后即焚关键字和用户查询信息（是否是由owner设置的权限查询的人）
					models.Db.Where("id=?", a.ID).Updates(models.Clipboard_VIP{Visible_another: key})
				} else {
					c.String(200, "您并没有权能发动阅后即焚功能")
				}
			}
		} else {
			c.String(200, "由于对方访问权限设置，你没有权限查看此剪切板内容")
		}
	} else {
		c.String(200, "没有找到该剪切板记录！！！")
	}
}
