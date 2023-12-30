package services

import (
	"net/http"

	"github.com/chanaseptiari/learn-auth-golang/apps/initializer"
	"github.com/chanaseptiari/learn-auth-golang/helper"
	"github.com/chanaseptiari/learn-auth-golang/models"
	"github.com/chanaseptiari/learn-auth-golang/repositories"
	"github.com/gin-gonic/gin"
)

func LoginService(Input models.LoginInput, cnt *gin.Context) (int, interface{}) {
	// Error First Validate
	if err := cnt.ShouldBindJSON(&Input); err != nil {
		return http.StatusBadRequest, helper.NewErrorResponse(http.StatusBadRequest, "Validate", err.Error())
	}

	// Check Username
	resRepo, errRepo := repositories.GetUsername(initializer.DB, (*models.LoginSearch)(&Input))
	if errRepo != nil {
		return http.StatusOK, helper.NewErrorResponse(http.StatusOK, "Username Not Found", errRepo.Error())
	}

	err := helper.CheckHash(resRepo["Password"].(string), Input.Password)
	if err == false {
		// Wrong Password
		return http.StatusOK, helper.NewErrorResponse(http.StatusOK, "Password", "Wrong password")
	}

	token := helper.GenerateToken(Input.Username)

	// // Success
	return http.StatusOK, helper.NewSuccessResponse(http.StatusOK, "Success", token)
}
