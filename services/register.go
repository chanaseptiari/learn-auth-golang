package services

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/chanaseptiari/learn-auth-golang/apps/initializer"
	"github.com/chanaseptiari/learn-auth-golang/helper"
	"github.com/chanaseptiari/learn-auth-golang/models"
	"github.com/chanaseptiari/learn-auth-golang/repositories"
	"github.com/gin-gonic/gin"
)

func RegisterService(Input models.MsAuth, cnt *gin.Context) (int, interface{}) {
	Input.ActiveAt = helper.GetEpochTime

	// Error First Validate
	if err := cnt.ShouldBindJSON(&Input); err != nil {
		return http.StatusBadRequest, helper.NewErrorResponse(http.StatusBadRequest, "Validate", err.Error())
	}

	// Checking Username
	pattern := `^[a-z]+[a-z0-9_\\.-]+$`
	r, err := regexp.Compile(pattern)
	if err != nil {
		return http.StatusBadRequest, helper.NewErrorResponse(http.StatusBadRequest, "Validate", err.Error())
	}

	// Check if the username matches the pattern.
	if !r.MatchString(Input.Username) {
		err := errors.New("Username is NOT valid.")
		return http.StatusBadRequest, helper.NewErrorResponse(http.StatusBadRequest, "Invalid", err.Error())
	}

	// HashPassword
	hash, _ := helper.HashPassword(Input.Password)
	Input.Password = hash

	// Error Input
	if err := repositories.ProsesRegister(initializer.DB, (*models.MsAuth)(&Input)); err != nil {
		return http.StatusBadRequest, helper.NewErrorResponse(http.StatusBadRequest, "Create Account", err.Error())
	}

	// Success
	return http.StatusOK, helper.NewSuccessResponse(http.StatusOK, "Success", "Register Success")
}
