package middleware

/*
 *认证
 */

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Auth struct{}

//+++++++++++++++++++++++++Cookie Begin++++++++++++++++++
func (Auth) GetCookie(c *gin.Context) {
	if cookie, err := c.Request.Cookie("test"); err == nil {
		c.String(http.StatusOK, cookie.Value)
		return
	}
	c.String(http.StatusOK, "cookie Name is empty")
}

func (Auth) SetCookie(c *gin.Context) {
	//当前url才有的cookie（未指定path）
	cookie := &http.Cookie{
		Name:  "test",
		Value: "0",
	}
	//golang cookie跟url有直接关系
	cookie1 := &http.Cookie{
		Name:  "zl",
		Value: "1",
		Path:  "/", //全局cookie（当前目录即整个工程）
	}

	http.SetCookie(c.Writer, cookie)
	http.SetCookie(c.Writer, cookie1)
	c.String(http.StatusOK, "cookie set success")
}

//+++++++++++++++++++++++++Cookie End++++++++++++++++++

func (Auth) GetSession(c *gin.Context) {
	c.String(http.StatusOK, "hello %s", "world")
}

func (Auth) SetSession(c *gin.Context) {
	c.String(http.StatusOK, "hell %s", "world")
}
