package main

import "github.com/gin-gonic/gin"

func test(c *gin.Context) {
	//	使用请求参数的方法获取在GEThttp请求中在url?之后的参数
	username := c.Query("username")
	password := c.DefaultQuery("password", "Aa35754391")
	c.String(200, "nihao woshi %s,wo demima shi%s", username, password)
}
func main() {
	engine := gin.Default()
	engine.GET("/ask_for_params", test)
	engine.Run(":886")
}
