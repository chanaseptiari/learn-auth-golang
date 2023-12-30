package services

import (
	"net/http"

	"github.com/chanaseptiari/learn-auth-golang/helper"
	"github.com/gin-gonic/gin"
)

func RenewToken(cnt *gin.Context) (int, interface{}) {
	token := cnt.Request.Header["X-Api-Key"][0]

	res, err := helper.ParseToken(token)
	username, _ := res.Claims.GetAudience()
	if err != nil {
		token := helper.GenerateToken(username[0])
		// return 200, token
		return http.StatusOK, helper.NewSuccessResponse(http.StatusOK, "Renew Token", token)
	}
	return http.StatusOK, helper.NewSuccessResponse(http.StatusOK, "Token", token)
}
