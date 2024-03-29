package main

import (
	"Termbin/handles"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	user := engine.Group("/user")
	user.POST("/register/:username/:password", handles.Register)
	user.POST("/login/:username/:password", handles.Login)
	protected := engine.Group("/:username") //设置有权限限制的访问路由组
	protected.Use(handles.User)             //使用中间件检查用户的登陆信息，是实际开发中常见的操作！
	{
		protected.POST("/create", handles.UserCreate)
		protected.POST("/create/:alias", handles.UserCreate)
		protected.POST("/delete/:id", handles.UserDelete)
		protected.POST("/update/:id", handles.UserUpdate)
		protected.POST("/query/:interface", handles.UserQuery)
		protected.POST("/owner/:id/:visible_VIP", handles.UserOwner)
		protected.POST("/another/:id/:visible_another", handles.UserAnother)
		protected.POST("/create_Delay", handles.Create_CelayDelete)
	}
	engine.GET("/:id", handles.Query)
	engine.POST("/create", handles.Create)
	engine.PUT("/:id", handles.Update)
	engine.DELETE("/:id", handles.Delete)
	engine.Run(":888")
}
