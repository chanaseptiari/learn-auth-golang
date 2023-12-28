package apps

import (
	"net/http"
	"os"

	"github.com/chanaseptiari/learn-auth-golang/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(os.Getenv("MODE"))
	router := gin.New()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "Golang Restful API"})
	})

	router.POST("/register", controllers.Register)

	router.POST("/login", controllers.Login)

	return router
}
