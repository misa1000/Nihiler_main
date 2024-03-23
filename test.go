package main

import (
	"github.com/gin-gonic/gin"
)

func test2(c *gin.Context) {
	c.String(200, "Hello,the world!!!!")
}
func main() {
	engine := gin.Default()
	engine.GET("/hello", test2)
	engine.Run(":2233")
}
