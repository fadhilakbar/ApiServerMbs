package routers

import (
	"ApiServerMbs/controllers"
	"ApiServerMbs/functions"
	"github.com/gin-gonic/gin"
	"os"
)

func SetupRouter(engine *gin.Engine) {

	engine.GET("/", func(c *gin.Context) {
		c.JSON(200, functions.GetResponse("00", os.Getenv("APP_NAME")+" "+os.Getenv("APP_VERSION")))
	})

	engine.POST("/InsertOutbox", controllers.InsertOutbox)
	engine.POST("/InsertOutboxMail", controllers.InsertOutboxMail)
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	engine.NoRoute(func(c *gin.Context) {
		c.JSON(404, functions.GetResponse("00", "Method Not Allowed"))
	})
}
