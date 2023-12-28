package controllers

import (
	"github.com/chanaseptiari/learn-auth-golang/models"
	"github.com/chanaseptiari/learn-auth-golang/services"
	"github.com/gin-gonic/gin"
)

func Login(cnt *gin.Context) {
	var input models.LoginInput
	code, res := services.LoginService(input, cnt)
	cnt.SecureJSON(code, res)
}
