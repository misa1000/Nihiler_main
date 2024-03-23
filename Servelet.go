package main

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	c.HTML(200, "denglujiemian.html", nil)
}
func WELCOME(c *gin.Context) {
	username := c.PostForm("username") /*PostForm和DefaultPostForm用于从POST请求中查询得到参数*/
	password := c.PostForm("password")
	c.HTML(200, "Welcome.html", gin.H{"username": username,
		"password": password})
}
func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("H:\\GO\\BeginnerGO/*") /*注意此处必须有/*,否则是不可以找到指引文件的*/
	/*如果想要设计restful的代码风格，那么可以试着使用如下的relativePath设置，写为“/login/:username/:password”
	然后在执行函数中使用Param方法进行取得参数。*/
	engine.GET("/login", Login)
	engine.POST("/login", WELCOME)
	engine.Run(":888")
}
