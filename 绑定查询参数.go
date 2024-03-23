package main

import "github.com/gin-gonic/gin"

type user struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func Form(c *gin.Context) {
	c.HTML(200, "form.html", nil)
}
func Formchuli(c *gin.Context) {
	var User user
	c.ShouldBind(&User)
	c.String(200, "from data %s", User)
}

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("H:\\GO\\BeginnerGO/*")
	engine.POST("/formchuli", Formchuli)
	engine.GET("/form", Form)
	engine.Run(":885")
}
