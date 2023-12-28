package services

import (
	"net/http"

	"github.com/chanaseptiari/learn-auth-golang/apps/initializer"
	"github.com/chanaseptiari/learn-auth-golang/helper"
	"github.com/chanaseptiari/learn-auth-golang/models"
	"github.com/gin-gonic/gin"
)

func LoginService(Input models.LoginInput, cnt *gin.Context) (int, interface{}) {
	var check []models.LoginSearch

	// Error First Validate
	if err := cnt.ShouldBindJSON(&Input); err != nil {
		return http.StatusBadRequest, helper.NewErrorResponse(http.StatusBadRequest, "Error", err.Error())
	}

	initializer.DB.Table("ms_auth").Where("Username = ? ", Input.Username).Take(&check)

	err := helper.CheckHash(check[len(check)-1].Password, Input.Password)
	if err == false {
		// Wrong Password
		return http.StatusOK, helper.NewErrorResponse(http.StatusOK, "Error", "Wrong password")
	}

	token := helper.GenerateToken(Input.Username)

	// Success
	return http.StatusOK, helper.NewSuccessResponse(http.StatusOK, "Success", token)
}
