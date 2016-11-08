package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"middleware"
	"net/http"
)

var Router *gin.Engine

func main() {
	gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	Router = gin.Default()     //获得路由实例

	Router.Static("/public", "./public")
	Router.LoadHTMLGlob("views/*")

	// Router.Use(func(c *gin.Context) {
	// 	fmt.Println("middleware start")
	// 	c.Next()
	// 	fmt.Println("middleware end")
	// })

	auth := middleware.Auth{}

	//注册接口
	Router.GET("/auth/cookie", auth.GetCookie)
	Router.POST("/auth/cookie", auth.SetCookie)
	Router.GET("/auth/session", auth.GetSession)
	Router.POST("/auth/session", auth.SetSession)

	foo := middleware.Foo{}
	Router.GET("/foo/view", func(c *gin.Context) {
		fmt.Println("/foo/view")
		c.Next()
		fmt.Println("/foo/view end")
	}, foo.Render)

	Router.GET("/foo/json", foo.Json)
	Router.POST("/foo/json", foo.Post)
	Router.GET("/name/:name", foo.Get)

	Router.GET("/simple/server/get", GetHandler)
	Router.POST("/simple/server/post", PostHandler)
	Router.PUT("/simple/server/put", PutHandler)
	Router.DELETE("/simple/server/delete", DeleteHandler)
	//监听端口
	// http.ListenAndServe(":8005", Router)

	Router.Run(":8005")
}

func GetHandler(c *gin.Context) {
	value, exist := c.Get("key")
	if !exist {
		value = "the key is not exist!"
	}
	c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("get success! %s\n", value)))
	return
}

func PostHandler(c *gin.Context) {
	type JsonHolder struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	holder := JsonHolder{Id: 1, Name: "my name"}
	//若返回json数据，可以直接使用gin封装好的JSON方法
	c.JSON(http.StatusOK, holder)
	return
}
func PutHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain", []byte("put success!\n"))
	return
}
func DeleteHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain", []byte("delete success!\n"))
	return
}
