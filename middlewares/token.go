package middlewares

import (
	"fmt"
	"net/http"

	"game.com/controlWeb/cache"
	"github.com/gin-gonic/gin"
)

func TokenCheck(c *gin.Context) {

	fmt.Println(c.Request.URL)
	token := c.GetHeader("token")
	info, exist := cache.GetToken(token)

	if !exist {
		c.HTML(http.StatusForbidden, "login.tmpl", nil)
		return
	}
	c.Set("adminName", info.Account)
	fmt.Println(c.Get("adminName"))
	c.Next()
}
