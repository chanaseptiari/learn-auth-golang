package apps

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/chanaseptiari/learn-auth-golang/controllers"
	"github.com/chanaseptiari/learn-auth-golang/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(os.Getenv("MODE"))

	router := gin.New()

	router.POST("/register", controllers.Register)

	router.POST("/login", controllers.Login)

	router.HEAD("/renew-token", controllers.RenewToken)
	v1 := router.Group("/v1").Use(middleware.AuthRequired())
	{
		v1.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"data": "Golang Restful API"})
		})

		v1.GET("/long_async", func(c *gin.Context) {
			// create copy to be used inside the goroutine
			cCp := c.Copy()
			go func() {
				// simulate a long task with time.Sleep(). 5 seconds
				time.Sleep(5 * time.Second)

				// note that you are using the copied context "cCp", IMPORTANT
				log.Println("Done! in path " + cCp.Request.URL.Path)
			}()
		})

		v1.GET("/long_sync", func(c *gin.Context) {
			// simulate a long task with time.Sleep(). 5 seconds
			time.Sleep(5 * time.Second)

			// since we are NOT using a goroutine, we do not have to copy the context
			log.Println("Done! in path " + c.Request.URL.Path)
		})

	}

	return router
}
