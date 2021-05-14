package routers

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

var options = []Option{}

func Include(opts ...Option) {
	options = append(options, opts...)
}

func init() {
	log.Print("init")
}

func Init() *gin.Engine  {
	log.Print("router init")
	r := gin.Default()
	for _, opt := range options {
		opt(r)
	}
	return r
}