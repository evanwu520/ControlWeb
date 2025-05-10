package routers

import (
	"net/http"

	"game.com/controlWeb/controllers"
	"game.com/controlWeb/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/css", "./css")
	r.Static("/js", "./js")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", nil)
	})

	userGroup := r.Group("/user")
	userController := controllers.NewUserController()
	userGroup.POST("/login", userController.Login)
	userGroup.POST("/online/:account", userController.OnlinePlayers)

	bonusGroup := r.Group("/bonus").Use(middlewares.TokenCheck)
	bonusController := controllers.NewBonusController()
	bonusGroup.POST("", bonusController.Reward)
	bonusGroup.GET("", bonusController.RewardRecords)

	return r

}
