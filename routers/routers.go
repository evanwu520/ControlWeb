package routers

import (
	"net/http"

	"game.com/controlWeb/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("css", "css")
	r.Static("js", "js")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", nil)
	})

	userGroup := r.Group("/user")

	userController := controllers.NewUserController()
	userGroup.POST("/login", userController.Login)

	return r

}
