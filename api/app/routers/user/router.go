package shop

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine)  {
	e.GET("/user", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"code": 1, "msg": "OK", "data": "user"})
	})
}