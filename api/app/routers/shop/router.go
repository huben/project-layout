package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.GET("/shop", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"code": 1, "msg": "OK", "data": "shop"})
	})
}
