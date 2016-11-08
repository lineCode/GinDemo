package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Foo struct{}

func (Foo) Get(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "hello %s", name)
}

func (Foo) Render(c *gin.Context) {
	c.HTML(http.StatusOK, "foo.html", gin.H{
		"key": "Main website",
	})
}

func (Foo) Json(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user": struct {
			Name   string
			Gender int
		}{Name: "a", Gender: 1}})
}

func (Foo) Post(c *gin.Context) {
	fmt.Println(c.Request.Body)
	_ = "breakpoint"

	result, _ := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	fmt.Println(result)
	var user model.User
	json.Unmarshal(result, &user)
	fmt.Println(user)

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func Printfx(args ...interface{}) {
	for _, val := range args {
		fmt.Println(val)
	}
}
