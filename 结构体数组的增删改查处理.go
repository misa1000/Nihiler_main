package main

import "github.com/gin-gonic/gin"

/*type User struct {
	Name string
	Sex  string
	Game string
}*/

var users []User

func Add(c *gin.Context) {
	var a User
	a.Name = c.Param("Name")
	a.Sex = c.Param("Sex")
	a.Game = c.Param("Game")
	users = append(users, a)
	c.String(200, "结构体数组当前状态%s", users)
}
func Delete(c *gin.Context) {
	Name := c.Param("Name")
	for i := 0; i < len(users); i++ {
		if Name == users[i].Name {
			users = append(users[:i], users[1+i:]...)
			break
		} else {
			continue
		}
	}
	c.String(200, "结构体数组当前状态%s", users)
}
func Find(c *gin.Context) {
	Name := c.Param("Name")
	var a User
	for i := 0; i < len(users); i++ {
		if Name == users[i].Name {
			a = users[i]
			break
		} else {
			continue
		}
	}
	c.String(200, "%s对应的相关信息有%s", Name, a)
}
func Update(c *gin.Context) {
	Name := c.Param("Name")
	for i := 0; i < len(users); i++ {
		if Name == users[i].Name {
			users[i].Sex = c.Param("Sex")
			users[i].Game = c.Param("Game")
			break
		} else {
			continue
		}
	}
	c.String(200, "更新已经完成了,当前结构体数组为%s", users)
}
func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("H:\\GO\\BeginnerGO/*")
	engine.GET("/find/:Name", Find)
	engine.GET("/add/:Name/:Sex/:Game", Add)
	engine.GET("/delete/:Name", Delete)
	engine.GET("/update/:Name/:Sex/:Game", Update)
	engine.Run(":888")
}

//路由可以进行分组，语法如下

/*engine:=gin.Default()
 v1:=engine.Group("/video")
v1.GET()*/
/*小结：写程序还是有小八嘎（bug），
1.不同的请求类型应该对应使用不同的取参方式
2.循环的逻辑实现还是没有搞清楚！！！所以还是得小心谨慎*/
