package controllers

import (
	"github.com/chanaseptiari/learn-auth-golang/services"
	"github.com/gin-gonic/gin"
)

func RenewToken(cnt *gin.Context) {
	code, res := services.RenewToken(cnt)
	cnt.JSON(code, res)
}
