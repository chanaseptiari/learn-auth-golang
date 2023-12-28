package controllers

import (
	"github.com/chanaseptiari/learn-auth-golang/models"
	"github.com/chanaseptiari/learn-auth-golang/services"
	"github.com/gin-gonic/gin"
)

func Register(cnt *gin.Context) {
	var input models.MsAuth
	code, res := services.RegisterService(input, cnt)
	cnt.SecureJSON(code, res)
}
