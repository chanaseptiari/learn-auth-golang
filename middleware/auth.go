package middleware

import (
	"net/http"

	"github.com/chanaseptiari/learn-auth-golang/helper"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(cnt *gin.Context) {

		header := cnt.Request.Header["X-Api-Key"]
		if len(header) == 0 {
			cnt.JSON(http.StatusForbidden, helper.NewErrorResponse(http.StatusForbidden, "Header", "Token not found"))
			cnt.Abort()
			return
		}
		res, err := helper.VerifyToken(cnt.Request.Header["X-Api-Key"][0])
		if err != nil || !res {
			cnt.JSON(http.StatusForbidden, helper.NewErrorResponse(http.StatusForbidden, "Token", err.Error()))
			cnt.Abort()
			return
		}

		// Continue Request
		cnt.Next()
	}
}
